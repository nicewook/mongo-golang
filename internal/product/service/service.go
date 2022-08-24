package service

import "github.com/nicewok/mg/internal/product/dto"

type ProductService interface {
	InsertOne(r dto.ProductInsertReq) dto.ProductInsertResp
}
