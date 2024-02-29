package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Age  int    `gorm:"column:age" json:"age"`
	Location
	Hobby string `gorm:"column:hobby" json:"hobby"`
}

// SearchRequest represents the search parameters / pagination
type SearchRequest struct {
	Query    string                 `json:"query"`
	Filters  map[string]interface{} `json:"filters"`
	Page     int                    `json:"page"`
	Limit    int                    `json:"limit"`
	SortBy   string                 `json:"sortBy"`
	SortDesc bool                   `json:"sortDesc"`
	Output   string                 `json:"output"`
}

type UpdateReq struct {
	UserID     string `json:"user_id"`
	UpdateUser Users  `json:"updated_user"`
}

type ViewRequest struct {
	UserID string `json:"user_id"`
}

type Location struct {
	AddressName  string `gorm:"column:addressname" json:"addressname"`
	LocationArea string `gorm:"column:locationarea" json:"locationarea"`
}
