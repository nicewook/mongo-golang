package repository

import (
	"github.com/nicewook/mg/internal/product/entity"
)

type ProductRepository interface {
	InsertOne(r entity.ProductInsertOneReq) (entity.ProductInsertOneResp, error)
	FindOne(r entity.ProductFindOneReq) (entity.ProductFindOneResp, error)
}
