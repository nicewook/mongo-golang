package entity

import "net/url"

type ProductCountDocumentsReq struct {
	Database    string
	Collection  string
	QueryParams url.Values
}

type ProductCountDocumentsResp struct {
	Count int64
}
