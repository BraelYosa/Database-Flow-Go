package connection

import (
	"app/helpers"
	"app/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JwtSKey = []byte("your_secret_key")

// Jwt
func GenerateJWT(admin model.Admin, expiry time.Duration) (string, error) {

	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	// Create a new token object with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_mail": admin.AdminEmail,
		"admin_name": admin.AdminSurname,
		"exp":        exp, // Token expiration time

	})

	// Sign the token with a secret key and get the complete encoded token as a string
	tokenString, err := token.SignedString(JwtSKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Open Database
func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

func CreateUser(newUser model.Users) error {
	db := DB

	var existingUser model.Users
	if err := db.Table("users").Where("name=?", newUser.Name).First(&existingUser).Error; err == nil {
		return errors.New("User with the same name is already exists")
	}

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func CreateAdmin(newAuthor model.Admin) error {
	db := DB

	var existingAdmin model.Admin
	if err := db.Table("admins").Where("admin_name=?", newAuthor.AdminSurname).First(&existingAdmin).Error; err == nil {
		return errors.New("Author with the same name is already exists")
	}

	if err := db.Create(&newAuthor).Error; err != nil {
		return err
	}

	return nil
}

func GetAdminByEmail(email string) (*model.Admin, error) {

	var admin model.Admin
	if err := DB.Table("admins").Where("admin_mail = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

// Search Query
func SearchUsers(request model.SearchRequest) (interface{}, error) {
	var users []model.Users
	db := DB

	queryBuilder := db.Table("users").Where("1=1") //True / Found

	// Apply filters based on search request
	helpers.FilterLocName(request, queryBuilder)

	if request.SortBy != "" {
		queryBuilder = queryBuilder.Order(request.SortBy)
	}

	// Apply pagination
	offset := (request.Page - 1) * request.Limit
	queryBuilder = queryBuilder.Offset(offset).Limit(request.Limit)

	if err := queryBuilder.Find(&users).Error; err != nil {
		return nil, err
	}

	// Convert the search results to the requested output format
	if request.Output == "CSV" {
		csvData, err := helpers.ConvertToCSV(users)
		if err != nil {
			return nil, err
		}

		folderPath := "file/CSV/"
		fileName := folderPath + "output." + request.Output
		if err := os.WriteFile(fileName, []byte(csvData), 0644); err != nil {
			return nil, err
		}

		return csvData, nil
	} else if request.Output == "" || request.Output == "JSON" {
		return users, nil
	}

	return request.Output, nil

}

// updates a user with the provided ID
func UpdateUser(userID string, updateUser *model.Users) error {
	db := DB

	// Find the user by ID / Primary key
	var existingUser model.Users
	if err := db.Table("users").First(&existingUser, userID).Error; err != nil {
		return err
	}

	// Update the user's fields
	existingUser.Name = updateUser.Name
	existingUser.Age = updateUser.Age
	existingUser.Location = updateUser.Location
	existingUser.Hobby = updateUser.Hobby

	// Save the updated user
	if err := db.Save(&existingUser).Error; err != nil {
		return err
	}

	return nil
}

func ViewUser(userID string) (*model.Users, error) {
	db := DB

	// Find the user by ID
	var user model.Users
	if err := db.Table("users").First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
