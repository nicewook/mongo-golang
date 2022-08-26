package dto

import (
	"fmt"

	"github.com/nicewook/mg/internal/product/entity"
)

type ProductAddTagReq struct {
	DatabaseName   string
	CollectionName string
	ProductName    string
	Tag            string
}

func (r ProductAddTagReq) ToEntity() entity.ProductAddTagReq {
	entReq := entity.ProductAddTagReq{
		Database:    r.DatabaseName,
		Collection:  r.CollectionName,
		ProductName: r.ProductName,
		Tag:         r.Tag,
	}
	return entReq
}

type ProductAddTagResp struct {
	Message string
}

func (resp *ProductAddTagResp) ToDTO(entResp entity.ProductAddTagResp) {
	resp.Message = fmt.Sprintf("%d document matched, %d document modified", entResp.MatchedCount, entResp.ModifiedCount)
}
