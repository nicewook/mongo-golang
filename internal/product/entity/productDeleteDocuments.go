package entity

import "net/url"

type ProductDeleteDocumentsReq struct {
	Database    string
	Collection  string
	QueryParams url.Values
}

type ProductDeleteDocumentsResp struct {
	Count int64
}
