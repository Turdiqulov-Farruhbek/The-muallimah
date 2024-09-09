package storage

import (
	"context"

	pb "auth-service/internal/pkg/genproto"

	_ "github.com/lib/pq"
)

type StorageI interface {
	User() UserI
}

type UserI interface {
	CreateUser(context.Context, *pb.UserCreateReq) (*pb.Void, error)
	GetUserByID(context.Context, *pb.ById) (*pb.UserGetRes, error)
	GetUserByEmail(context.Context, *pb.ByEmail) (*pb.UserGetRes, error)
	UpdateUser(context.Context, *pb.UserUpdateReq) (*pb.Void, error)
	ChangeUserPassword(context.Context, *pb.UserRecoverPasswordReq) (*pb.Void, error)
	DeleteUser(context.Context, *pb.ById) (*pb.Void, error)
	ListUsers(context.Context, *pb.UsersGetAllReq) (*pb.UsersGetAllRes, error)
	IsEmailExists(context.Context, *pb.UserEmailCheckReq) (*pb.UserEmailCheckRes, error)
	GetUserSecurityByEmail(context.Context, *pb.ByEmail) (*pb.UserSecurityRes, error)
	ConfirmUser(context.Context, *pb.ByEmail) (*pb.Void, error)
	ChangeUserPFP(context.Context, *pb.UserChangePFPReq) (*pb.Void, error)
}
