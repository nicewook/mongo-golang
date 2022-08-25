package service

import "github.com/nicewok/mg/internal/product/dto"

type ProductService interface {
	InsertOne(r dto.ProductInsertOneReq) dto.ProductInsertOneResp
	FindOne(r dto.ProductFindOneReq) (dto.ProductFindOneResp, error)
}
