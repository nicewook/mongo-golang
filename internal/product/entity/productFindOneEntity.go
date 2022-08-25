package entity

import "net/url"

type ProductFindOneReq struct {
	Database    string
	Collection  string
	QueryParams url.Values
}

type ProductFindOneResp struct {
	Product
}
