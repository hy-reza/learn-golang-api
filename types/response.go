package types

type BookResponse struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
}
