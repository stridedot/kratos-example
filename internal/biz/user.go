package biz

import (
	"context"

	v1 "github.com/stridedot/kratos-example/api/property/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRequest struct {
	UserId    int64
	Page      int64
	PageSize  int64
	Mobile    string
}

type User struct {
	UserId   int64  `db:"user_id"`
	Username string `db:"realname"`
	Mobile   string `db:"mobile"`
}

type UserRepo interface {
	GetUserByUserId(ctx context.Context, req *UserRequest) (*User, error)
	GetUserByMobile(ctx context.Context, req *UserRequest) (*User, error)
	GetPaginatedUsers(ctx context.Context, req *UserRequest) ([]*User, error)
	CountUsers(ctx context.Context, req *UserRequest) (int64, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetUserByUserId(ctx context.Context, req *v1.GetUserRequest) (*User, error) {
	request := &UserRequest{	
		UserId: req.UserId,
	}
	return uc.repo.GetUserByUserId(ctx, request)
}

func (uc *UserUsecase) GetUserByMobile(ctx context.Context, req *v1.GetUserRequest) (*User, error) {
	request := &UserRequest{
		Mobile: req.Mobile,
	}
	return uc.repo.GetUserByMobile(ctx, request)
}

func (uc *UserUsecase) GetPaginatedUsers(ctx context.Context, req *v1.ListUserRequest) (*Data, error) {
	request := &UserRequest{
		Page: req.Page,
		PageSize: req.PageSize,
	}
	
	users, err := uc.repo.GetPaginatedUsers(ctx, request)
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.CountUsers(ctx, request)
	if err != nil {
		return nil, err
	}

	var hasNext int64
	if (req.Page * req.PageSize < count) {
		hasNext = 1
	}

	return &Data{
		Items: users,
		Meta: Meta{
			Total: count,
			CurrentPage: req.Page,
			PerPage: req.PageSize,
			HasNext: hasNext,
		},
	}, nil
}


