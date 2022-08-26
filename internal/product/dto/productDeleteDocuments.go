package dto

import (
	"net/url"

	"github.com/nicewook/mg/internal/product/entity"
)

type ProductDeleteDocumentsReq struct {
	DatabaseName   string
	CollectionName string
	QueryParams    url.Values
}

func (r ProductDeleteDocumentsReq) ToEntity() entity.ProductDeleteDocumentsReq {
	return entity.ProductDeleteDocumentsReq{
		Database:    r.DatabaseName,
		Collection:  r.CollectionName,
		QueryParams: r.QueryParams,
	}
}

type ProductDeleteDocumentsResp struct {
	Count int64
}

func (resp *ProductDeleteDocumentsResp) ToDTO(entResp entity.ProductDeleteDocumentsResp) {
	resp.Count = entResp.Count
}
