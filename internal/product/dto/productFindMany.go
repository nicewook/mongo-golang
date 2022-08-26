package dto

import (
	"net/url"

	"github.com/nicewook/mg/internal/product/entity"
)

type ProductFindManyReq struct {
	DatabaseName   string
	CollectionName string
	QueryParams    url.Values
}

func (r ProductFindManyReq) ToEntity() entity.ProductFindManyReq {
	return entity.ProductFindManyReq{
		Database:    r.DatabaseName,
		Collection:  r.CollectionName,
		QueryParams: r.QueryParams,
	}
}

type ProductFindManyResp []Product

func (resp *ProductFindManyResp) ToDTO(entResp entity.ProductFindManyResp) {
	DeepCopy(&resp, &entResp.Products)
}
