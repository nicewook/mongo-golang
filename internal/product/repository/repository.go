package repository

import (
	"github.com/nicewok/mg/internal/product/entity"
)

type ProductRepository interface {
	InsertOne(r entity.ProductInsertOneReq) entity.ProductInsertOneResp
	FindOne(r entity.ProductFindOneReq) (entity.ProductFindOneResp, error)
}
