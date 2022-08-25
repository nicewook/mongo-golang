package service

import (
	"github.com/nicewook/mg/internal/product/dto"
	"github.com/nicewook/mg/internal/product/repository"
)

type ProductSvc struct {
	repository repository.ProductRepository
}

var _ ProductService = (*ProductSvc)(nil)

func NewProductSvc(repo repository.ProductRepository) ProductService {
	return &ProductSvc{
		repository: repo,
	}
}

func (svc *ProductSvc) InsertOne(r dto.ProductInsertOneReq) (dto.ProductInsertOneResp, error) {
	entReq := r.ToEntity()
	entResp, err := svc.repository.InsertOne(entReq)

	var dtoResp dto.ProductInsertOneResp
	dtoResp.ToDTO(entResp)
	return dtoResp, err
}

func (svc *ProductSvc) FindOne(r dto.ProductFindOneReq) (dtoResp dto.ProductFindOneResp, err error) {
	entReq := r.ToEntity()
	entResp, err := svc.repository.FindOne(entReq)
	dtoResp.ToDTO(entResp)

	return dtoResp, err
}

func (svc *ProductSvc) FindMany(r dto.ProductFindManyReq) (dtoResp dto.ProductFindManyResp, err error) {
	entReq := r.ToEntity()
	entResp, err := svc.repository.FindMany(entReq)
	dtoResp.ToDTO(entResp)

	return dtoResp, err
}
