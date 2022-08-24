package service

import (
	"github.com/nicewok/mg/internal/product/dto"
	"github.com/nicewok/mg/internal/product/repository"
)

type ProductSvc struct {
	repository repository.ProductRepository
}

var _ ProductService = (*ProductSvc)(nil)

func (svc *ProductSvc) InsertOne(r dto.ProductInsertReq) dto.ProductInsertResp {
	entReq := r.ToEntity()
	entResp := svc.repository.InsertOne(entReq)

	var dtoResp dto.ProductInsertResp
	dtoResp.ToDTO(entResp)
	return dtoResp
}

func NewProductSvc(repo repository.ProductRepository) ProductService {
	return &ProductSvc{
		repository: repo,
	}
}
