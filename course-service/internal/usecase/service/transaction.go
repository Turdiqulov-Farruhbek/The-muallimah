package service

import (
	"context"
	"fmt"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
)


type TransactionService struct {
	stg storage.StorageI
	pb.UnimplementedTransactionServiceServer
}

func NewTransactionService(stg storage.StorageI) *TransactionService {
    return &TransactionService{stg: stg}
}

func (s *TransactionService) CreateTransaction(c context.Context, req *pb.TransactionCreateReq) (*pb.Void, error) {
	if req.UserCourseId != "" && req.UserCourseId != "string"{
		user_C,err := s.stg.UserCourse().GetUserCourse(&pb.ById{Id: req.UserCourseId})
		if err!= nil {
            return nil, err
        }
		expense := user_C.Course.Price
		// get balance 
		b_res,err := s.stg.Transaction().GetBalance(&pb.ById{Id: user_C.User.Id})
		if err!= nil {
            return nil, err
        }
		if expense > b_res.Balance {
			dif := expense - b_res.Balance
			return nil, fmt.Errorf("invalid balance you are short by %v", dif)
		}
		return s.stg.Transaction().CreateTransaction(req)
	}

	// get balance 
	b_res,err := s.stg.Transaction().GetBalance(&pb.ById{Id: req.UserId})
	if err!= nil {
		return nil, err
	}
	if req.Amount > b_res.Balance {
        dif := req.Amount - b_res.Balance
        return nil, fmt.Errorf("invalid balance you are short by %v", dif)
    }

    return s.stg.Transaction().CreateTransaction(req)
}

func (s *TransactionService) GetTransaction(c context.Context, req *pb.ById) (*pb.TransactionGetRes, error) {
    return s.stg.Transaction().GetTransaction(req)
}


func (s *TransactionService) ListTransactions(c context.Context, req *pb.TransactionListReq) (*pb.TransactionListRes, error) {
    return s.stg.Transaction().ListTransactions(req)
}

func (s *TransactionService) GetBalance(c context.Context, req *pb.ById) (*pb.BalanceGetRes, error) {
    return s.stg.Transaction().GetBalance(req)
}
