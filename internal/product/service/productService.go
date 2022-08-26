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

func (svc *ProductSvc) Insert(r dto.ProductInsertReq) (dto.ProductInsertResp, error) {
	entReq := r.ToEntity()
	entResp, err := svc.repository.Insert(entReq)

	var dtoResp dto.ProductInsertResp
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

func (svc *ProductSvc) CountDocuments(r dto.ProductCountDocumentsReq) (dtoResp dto.ProductCountDocumentsResp, err error) {
	entReq := r.ToEntity()
	entResp, err := svc.repository.CountDocuments(entReq)
	dtoResp.ToDTO(entResp)

	return dtoResp, err
}

func (svc *ProductSvc) AddReview(r dto.ProductAddReviewReq) (dtoResp dto.ProductAddReviewResp, err error) {
	entReq := r.ToEntity()
	entResp, err := svc.repository.AddReview(entReq)
	dtoResp.ToDTO(entResp)

	return dtoResp, err
}

func (svc *ProductSvc) DeleteDocuments(r dto.ProductDeleteDocumentsReq) (dtoResp dto.ProductDeleteDocumentsResp, err error) {
	entReq := r.ToEntity()
	entResp, err := svc.repository.DeleteDocuments(entReq)
	dtoResp.ToDTO(entResp)

	return dtoResp, err
}
