package dto

import (
	"github.com/nicewook/mg/internal/product/entity"
)

type ProductFindOneReq struct {
	DatabaseName   string
	CollectionName string
	Type           string
}

func (r ProductFindOneReq) ToEntity() entity.ProductFindOneReq {
	return entity.ProductFindOneReq{
		Database:   r.DatabaseName,
		Collection: r.CollectionName,
		Type:       r.Type,
	}
}

type ProductFindOneResp struct {
	Product
}

func (resp *ProductFindOneResp) ToDTO(entResp entity.ProductFindOneResp) {
	DeepCopy(&entResp.Product, &resp.Product)
}
