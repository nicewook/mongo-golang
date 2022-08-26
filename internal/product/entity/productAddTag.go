package entity

type ProductAddTagReq struct {
	Database    string
	Collection  string
	ProductName string
	Tag         string
}

type ProductAddTagResp struct { // the same struct as *mongo.UpdateResult
	MatchedCount  int64       // The number of documents matched by the filter.
	ModifiedCount int64       // The number of documents modified by the operation.
	UpsertedCount int64       // The number of documents upserted by the operation.
	UpsertedID    interface{} // The _id field of the upserted document, or nil if no upsert was done.
}
