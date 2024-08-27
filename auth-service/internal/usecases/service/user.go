package service

import (
	pb "auth-service/internal/pkg/genproto"
	"auth-service/internal/storage"
	"context"
	"errors"
	"time"
)

type UserService struct {
	stg storage.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg storage.StorageI) *UserService {
	return &UserService{stg: stg}
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.UserCreateReq) (*pb.Void, error) {
	layout := "2006-01-02"
	_, err := time.Parse(layout, req.Dob)
	if err != nil {
		return nil, errors.New("invalid date of birth format. details: " + err.Error())
	}
	return u.stg.User().CreateUser(ctx, req)
}

func (u *UserService) GetUserByID(ctx context.Context, req *pb.ById) (*pb.UserGetRes, error) {
	return u.stg.User().GetUserByID(ctx, req)
}

func (u *UserService) GetUserByEmail(ctx context.Context, req *pb.ByEmail) (*pb.UserGetRes, error) {
	return u.stg.User().GetUserByEmail(ctx, req)
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.UserUpdateReq) (*pb.Void, error) {
	layout := "2006-01-02"
	_, err := time.Parse(layout, req.Body.Dob)
	if err != nil {
		return nil, errors.New("invalid date of birth format. details: " + err.Error())
	}
	return u.stg.User().UpdateUser(ctx, req)
}

func (u *UserService) ChangeUserPassword(ctx context.Context, req *pb.UserChangePasswordReq) (*pb.Void, error) {
	return u.stg.User().ChangeUserPassword(ctx, req)
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	return u.stg.User().DeleteUser(ctx, req)
}

func (u *UserService) ListUsers(ctx context.Context, req *pb.UsersGetAllReq) (*pb.UsersGetAllRes, error) {
	return u.stg.User().ListUsers(ctx, req)
}

func (u *UserService) IsEmailExists(ctx context.Context, req *pb.UserEmailCheckReq) (*pb.UserEmailCheckRes, error) {
	return u.stg.User().IsEmailExists(ctx, req)
}

func (u *UserService) GetUserSecurityByEmail(ctx context.Context, req *pb.ByEmail) (*pb.UserSecurityRes, error) {
	return u.stg.User().GetUserSecurityByEmail(ctx, req)
}

func (u *UserService) ConfirmUser(ctx context.Context, req *pb.ByEmail) (*pb.Void, error) {
	return u.stg.User().ConfirmUser(ctx, req)
}

func (u *UserService) ChangeUserPFP(ctx context.Context, req *pb.UserChangePFPReq) (*pb.Void, error) {
	return u.stg.User().ChangeUserPFP(ctx, req)
}
