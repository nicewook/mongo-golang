package service

import (
	"log"

	"github.com/nicewok/mg/internal/product/dto"
	"github.com/nicewok/mg/internal/product/repository"
)

type ProductSvc struct {
	repository repository.ProductRepository
}

var _ ProductService = (*ProductSvc)(nil)

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

	log.Println("entReq", entReq)
	log.Println("entResp", entResp)
	log.Println("dtoResp", dtoResp)
	return dtoResp, err
}

func NewProductSvc(repo repository.ProductRepository) ProductService {
	return &ProductSvc{
		repository: repo,
	}
}
