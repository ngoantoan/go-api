package entity

type ListUser struct {
	ID         int `json:"id"`
	Fullname   string `json:"fullname"`
	Store_name string `json:"store_name"`
}