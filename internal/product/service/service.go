package service

import "github.com/nicewook/mg/internal/product/dto"

type ProductService interface {
	InsertOne(r dto.ProductInsertOneReq) (dto.ProductInsertOneResp, error)
	FindOne(r dto.ProductFindOneReq) (dto.ProductFindOneResp, error)
	FindMany(r dto.ProductFindManyReq) (dto.ProductFindManyResp, error)
}
