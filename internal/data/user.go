package data

import (
	"context"
	"github.com/stridedot/kratos-example/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) GetUserByUserId(ctx context.Context, req *biz.UserRequest) (*biz.User, error) {
	user := &biz.User{}
	err := ur.data.db.Get(user, "SELECT user_id, realname, mobile FROM hk591.user WHERE user_id = ?", req.UserId)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (ur *userRepo) GetUserByMobile(ctx context.Context, req *biz.UserRequest) (*biz.User, error) {
	user := &biz.User{}
	err := ur.data.db.Get(user, "SELECT user_id, realname, mobile FROM hk591.user WHERE mobile = ?", req.Mobile)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (ur *userRepo) CountUsers(ctx context.Context, req *biz.UserRequest) (int64, error) {
	var count int64
	err := ur.data.db.Get(&count, "SELECT COUNT(*) FROM hk591.user")
	if err != nil {
		return 0, err
	}
	return count, err
}

func (ur *userRepo) GetPaginatedUsers(ctx context.Context, req *biz.UserRequest) ([]*biz.User, error) {

	users := []*biz.User{}
	err := ur.data.db.Select(
		&users, 
		"SELECT user_id, realname, mobile FROM hk591.user ORDER BY user_id DESC LIMIT ?, ?",
		(req.Page - 1) * req.PageSize,
		req.PageSize,
	)
	return users, err
}
