package dto

import (
	"github.com/nicewok/mg/internal/product/entity"
)

type ProductInsertOneReq struct {
	DatabaseName   string
	CollectionName string
	Product
}

func (r ProductInsertOneReq) ToEntity() entity.ProductInsertOneReq {
	ent := entity.ProductInsertOneReq{
		Database:   r.DatabaseName,
		Collection: r.CollectionName,
	}
	DeepCopy(r.Product, ent.Product)
	return ent
}

type ProductInsertOneResp struct {
	InsertedID string
	Message    string
}

func (resp *ProductInsertOneResp) ToDTO(entResp entity.ProductInsertOneResp) {
	resp.InsertedID = entResp.InsertedID
	resp.Message = entResp.Message
}
