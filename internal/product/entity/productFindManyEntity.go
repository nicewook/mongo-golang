package entity

type ProductFindManyReq struct {
	Database   string
	Collection string
	Type       string
}

type ProductFindManyResp struct {
	Products []Product
}
