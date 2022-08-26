package dto

import (
	"net/url"

	"github.com/nicewook/mg/internal/product/entity"
)

type ProductCountDocumentsReq struct {
	DatabaseName   string
	CollectionName string
	QueryParams    url.Values
}

func (r ProductCountDocumentsReq) ToEntity() entity.ProductCountDocumentsReq {
	return entity.ProductCountDocumentsReq{
		Database:    r.DatabaseName,
		Collection:  r.CollectionName,
		QueryParams: r.QueryParams,
	}
}

type ProductCountDocumentsResp struct {
	Count int64
}

func (resp *ProductCountDocumentsResp) ToDTO(entResp entity.ProductCountDocumentsResp) {
	resp.Count = entResp.Count
}
