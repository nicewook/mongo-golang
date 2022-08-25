package entity

type ProductInsertOneReq struct {
	Database   string
	Collection string
	Product
}

type ProductInsertOneResp struct {
	InsertedID string
}
