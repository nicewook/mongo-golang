package service

import "github.com/nicewook/mg/internal/product/dto"

type ProductService interface {
	Insert(r dto.ProductInsertReq) (dto.ProductInsertResp, error)
	FindOne(r dto.ProductFindOneReq) (dto.ProductFindOneResp, error)
	FindMany(r dto.ProductFindManyReq) (dto.ProductFindManyResp, error)
	CountDocuments(r dto.ProductCountDocumentsReq) (dto.ProductCountDocumentsResp, error)
	DeleteDocuments(r dto.ProductDeleteDocumentsReq) (dto.ProductDeleteDocumentsResp, error)
}
