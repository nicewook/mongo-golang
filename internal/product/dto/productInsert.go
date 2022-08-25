package dto

import (
	"github.com/nicewook/mg/internal/product/entity"
)

type ProductInsertReq struct {
	DatabaseName   string
	CollectionName string
	Products       []Product
}

func (r ProductInsertReq) ToEntity() entity.ProductInsertReq {
	ent := entity.ProductInsertReq{
		Database:   r.DatabaseName,
		Collection: r.CollectionName,
	}
	DeepCopy(&ent.Products, &r.Products)
	return ent
}

type ProductInsertResp struct {
	InsertedIDs []string
}

func (resp *ProductInsertResp) ToDTO(entResp entity.ProductInsertResp) {
	resp.InsertedIDs = make([]string, len(entResp.InsertedIDs))
	copy(resp.InsertedIDs, entResp.InsertedIDs)
}
