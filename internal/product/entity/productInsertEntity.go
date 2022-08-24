package entity

type ProductInsertReq struct {
	Database   string
	Collection string
	Product
}

type ProductInsertResp struct {
	InsertedID string
	Message    string
}
