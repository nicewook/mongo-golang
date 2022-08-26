package dto

import (
	"fmt"

	"github.com/nicewook/mg/internal/product/entity"
)

type ProductAddReviewReq struct {
	DatabaseName   string
	CollectionName string
	ProductName    string
	Review
}

func (r ProductAddReviewReq) ToEntity() entity.ProductAddReviewReq {
	entReq := entity.ProductAddReviewReq{
		Database:    r.DatabaseName,
		Collection:  r.CollectionName,
		ProductName: r.ProductName,
	}
	DeepCopy(&entReq.Review, &r.Review)
	return entReq
}

type ProductAddReviewResp struct {
	Message string
}

func (resp *ProductAddReviewResp) ToDTO(entResp entity.ProductAddReviewResp) {
	resp.Message = fmt.Sprintf("%d document matched, %d document modified", entResp.MatchedCount, entResp.ModifiedCount)
}
