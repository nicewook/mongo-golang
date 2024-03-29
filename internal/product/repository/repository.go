package repository

import (
	"github.com/nicewook/mg/internal/product/entity"
)

type ProductRepository interface {
	Insert(r entity.ProductInsertReq) (entity.ProductInsertResp, error)

	FindOne(r entity.ProductFindOneReq) (entity.ProductFindOneResp, error)
	FindMany(r entity.ProductFindManyReq) (entity.ProductFindManyResp, error)
	CountDocuments(r entity.ProductCountDocumentsReq) (entity.ProductCountDocumentsResp, error)

	AddReview(r entity.ProductAddReviewReq) (entity.ProductAddReviewResp, error)
	AddTag(r entity.ProductAddTagReq) (entity.ProductAddTagResp, error)

	DeleteDocuments(r entity.ProductDeleteDocumentsReq) (entity.ProductDeleteDocumentsResp, error)
}
