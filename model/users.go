package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	User_ID uint   `gorm:"primarykey"`
	Name    string `gorm:"column:name" json:"name"`
	Age     int    `gorm:"column:age" json:"age"`
	Location
	Hobby     string `gorm:"column:hobby" json:"hobby"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
	UserID        string   `json:"user_id"`
	ProductID     string   `json:"product_id"`
	UpdateUser    Users    `json:"updated_user"`
	UpdateProduct Products `json:"update_product"`
}

type ViewRequest struct {
	UserID string `json:"user_id"`
}

type Location struct {
	AddressName  string `gorm:"column:address_name" json:"address_name"`
	LocationArea string `gorm:"column:location_area" json:"location_area"`
}

type Admin struct {
	Admin_id     uint   `gorm:"primaryKey"`
	AdminEmail   string `gorm:"column:admin_mail" json:"admin_mail"`
	AdminSurname string `gorm:"column:admin_name" json:"admin_name"`
	AdminPass    string `gorm:"column:admin_pass" json:"admin_pass"` //hash password
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Products struct {
	ProductID    uint64 `gorm:"primaryKey" column:"product_id"`
	AdminID      uint   `gorm:"column:admin_id" json:"admin_id"`
	ProductName  string `gorm:"column:product_name" json:"product_name"`
	ProductPrice int    `gorm:"column:product_price" json:"product_price"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
