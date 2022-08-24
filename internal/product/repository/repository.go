package repository

import (
	"github.com/nicewok/mg/internal/product/entity"
)

type ProductRepository interface {
	InsertOne(r entity.ProductInsertReq) entity.ProductInsertResp
}
