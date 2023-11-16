package biz

import (
	"context"

	v1 "github.com/stridedot/kratos-example/api/property/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type Request struct {
	Type int32
	Page int32
	PageSize int32
}

type House struct {
	ID     int32  `db:"id"`
	Title  string `db:"title"`
	Price  string `db:"price"`
	Cover  string `db:"cover"`
}

type HouseRepo interface {
	PaginatedListHouse(ctx context.Context, req *Request) ([]*House, error)
}

type HouseUsecase struct {
	repo HouseRepo
	log  *log.Helper
}

func NewHouseUsecase(repo HouseRepo, logger log.Logger) *HouseUsecase {
	return &HouseUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *HouseUsecase) PaginatedListHouse(ctx context.Context, req *v1.PagintatedListHouseRequest) ([]*House, error) {
	request := &Request{
		Type:     req.Type,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	return uc.repo.PaginatedListHouse(ctx, request)
}
