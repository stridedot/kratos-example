package data

import (
	"context"
	"github.com/stridedot/kratos-example/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type houseRepo struct {
	data *Data
	log  *log.Helper
}

// NewHouseRepo .
func NewHouseRepo(data *Data, logger log.Logger) biz.HouseRepo {
	return &houseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (hr *houseRepo) PaginatedListHouse(ctx context.Context, req *biz.Request) ([]*biz.House, error) {
	// var table string 
	// var offset int32
	// offset = (req.Page - 1) * req.PageSize

	// if req.Type == 1 {
	// 	table = "hk591.rent"
	// } else {
	// 	table = "hk591.sale"
	// }
	sqlStr := "SELECT id, title, price, cover FROM hk591.rent WHERE `status` = 2 AND closed = 2 ORDER BY id DESC LIMIT 0, 20"
	
	houses := []*biz.House{} 
	err := hr.data.db.Select(&houses, sqlStr)
	
	return houses, err
}