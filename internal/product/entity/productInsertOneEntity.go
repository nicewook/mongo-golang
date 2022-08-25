package entity

type ProductInsertReq struct {
	Database   string
	Collection string
	Products   []Product
}

type ProductInsertResp struct {
	InsertedIDs []string
}
