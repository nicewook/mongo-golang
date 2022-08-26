package entity

import "net/url"

type ProductFindManyReq struct {
	Database    string
	Collection  string
	QueryParams url.Values
}

type ProductFindManyResp struct {
	Products []Product
}
