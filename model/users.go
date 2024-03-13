package model

import (
	"gorm.io/gorm"
)

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
	SortBy   string                 `json:"sort_By"`
	SortDesc bool                   `json:"sort_Desc"`
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
	AddressName  string `gorm:"column:address_name" json:"address_name"`
	LocationArea string `gorm:"column:location_area" json:"location_area"`
}

type Admin struct {
	gorm.Model
	AdminEmail   string `gorm:"column:admin_mail" json:"admin_mail"`
	AdminSurname string `gorm:"column:admin_name" json:"admin_name"`
	AdminPass    string `gorm:"column:admin_pass" json:"admin_pass"` //hash password
}
