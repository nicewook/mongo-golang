package entity

type ProductFindOneReq struct {
	Database   string
	Collection string
	Type       string
}

type ProductFindOneResp struct {
	Product
}
