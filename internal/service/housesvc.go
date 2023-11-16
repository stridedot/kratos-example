package service

import (
	"context"

	v1 "hk591/api/property/v1"
	"hk591/internal/biz"
)

type HouseSvcService struct {
	v1.UnimplementedHouseSvcServer

	uc *biz.HouseUsecase
}

func NewHouseSvcService(uc *biz.HouseUsecase) *HouseSvcService {
	return &HouseSvcService{uc: uc}
}

func (s *HouseSvcService) PaginatedListHouse(ctx context.Context, req *v1.PagintatedListHouseRequest) (*v1.PaginatedListListHouseReply, error) {
	houses, _ := s.uc.PaginatedListHouse(ctx, req)

	data := make([]*v1.House, 0, len(houses))
	for _, v := range houses {
		data = append(data, &v1.House{
			Id:    v.ID,
			Title: v.Title,
			Price: v.Price,
			CoverUrl: v.Cover,
		})
	}

	reply := &v1.PaginatedListListHouseReply{
		Code: 20000,
		Msg: "success",
		Data: data,
	}

	return reply, nil
}
