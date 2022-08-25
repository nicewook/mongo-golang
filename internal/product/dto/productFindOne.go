package dto

import (
	"net/url"

	"github.com/nicewook/mg/internal/product/entity"
)

type ProductFindOneReq struct {
	DatabaseName   string
	CollectionName string
	QueryParams    url.Values
}

func (r ProductFindOneReq) ToEntity() entity.ProductFindOneReq {
	return entity.ProductFindOneReq{
		Database:    r.DatabaseName,
		Collection:  r.CollectionName,
		QueryParams: r.QueryParams,
	}
}

type ProductFindOneResp struct {
	Product
}

func (resp *ProductFindOneResp) ToDTO(entResp entity.ProductFindOneResp) {
	DeepCopy(&resp.Product, &entResp.Product)
}
