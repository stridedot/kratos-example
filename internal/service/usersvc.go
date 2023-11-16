package service

import (
	"context"
	"time"

	pb "hk591/api/property/v1"
	"hk591/internal/biz"
)

type UserSvcService struct {
	pb.UnimplementedUserSvcServer

	uc *biz.UserUsecase
}

func NewUserSvcService(uc *biz.UserUsecase) *UserSvcService {
	return &UserSvcService{uc: uc}
}

func (s *UserSvcService) GetUserByUserId(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.uc.GetUserByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return &pb.GetUserReply{
		Code:    20000,
		Msg: "success",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data: &pb.User{
			UserId:   user.UserId,
			Username: user.Username,
			Mobile:   user.Mobile,
		},
	}, nil
}
func (s *UserSvcService) GetUserByMobile(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.uc.GetUserByMobile(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Code:    20000,
		Msg: "success",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data: &pb.User{
			UserId:   user.UserId,
			Username: user.Username,
			Mobile:   user.Mobile,
		},
	}, nil
}

func (s *UserSvcService) GetPaginatedUsers(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	data, err := s.uc.GetPaginatedUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	
	meta := &pb.Meta{
		Total: data.Meta.Total,
		CurrentPage: data.Meta.CurrentPage,
		PerPage: data.Meta.PerPage,
		HasNext: data.Meta.HasNext,
	}

	results, ok := data.Items.([]*biz.User)
	if !ok {
		return nil, err
	}

	items := []*pb.User{}

	for _, item := range results {
		items = append(items, &pb.User{
			UserId:   item.UserId,
			Username: item.Username,
			Mobile:   item.Mobile,
		})
	}

	return &pb.ListUserReply{
		Code:    20000,
		Msg: "success",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data: &pb.Data{
			Data: items,
			Meta: meta,	
		},
	}, nil
}
