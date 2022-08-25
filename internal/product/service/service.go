package service

import "github.com/nicewok/mg/internal/product/dto"

type ProductService interface {
	InsertOne(r dto.ProductInsertOneReq) (dto.ProductInsertOneResp, error)
	FindOne(r dto.ProductFindOneReq) (dto.ProductFindOneResp, error)
}
