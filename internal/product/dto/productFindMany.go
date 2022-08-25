package dto

import (
	"github.com/nicewook/mg/internal/product/entity"
)

type ProductFindManyReq struct {
	DatabaseName   string
	CollectionName string
	Type           string
}

func (r ProductFindManyReq) ToEntity() entity.ProductFindManyReq {
	return entity.ProductFindManyReq{
		Database:   r.DatabaseName,
		Collection: r.CollectionName,
		Type:       r.Type,
	}
}

type ProductFindManyResp struct {
	Products []Product
}

func (resp *ProductFindManyResp) ToDTO(entResp entity.ProductFindManyResp) {
	DeepCopy(&resp.Products, &entResp.Products)
}
