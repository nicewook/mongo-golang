package dto

import (
	"bytes"
	"encoding/gob"

	"github.com/nicewok/mg/internal/product/entity"
)

func DeepCopy(src, dist interface{}) (err error) {
	buf := bytes.Buffer{}
	if err = gob.NewEncoder(&buf).Encode(src); err != nil {
		return
	}
	return gob.NewDecoder(&buf).Decode(dist)
}

type ProductInsertReq struct {
	DatabaseName   string
	CollectionName string
	Product
}

func (r ProductInsertReq) ToEntity() entity.ProductInsertReq {
	ent := entity.ProductInsertReq{
		Database:   r.DatabaseName,
		Collection: r.CollectionName,
	}
	DeepCopy(r.Product, ent.Product)
	return ent
}

type ProductInsertResp struct {
	InsertedID string
	Message    string
}

func (resp *ProductInsertResp) ToDTO(entResp entity.ProductInsertResp) {
	resp.InsertedID = entResp.InsertedID
	resp.Message = entResp.Message
}
