package types

type BookRequest struct {
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	// Price  json.Number `json:"harga" binding:"required,number"`
}
