package postgres

import (
	"app/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var JwtSKey = []byte("your_secret_key")

// Jwt
func GenerateJWT(admin model.Admin, expiry time.Duration) (string, error) {

	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	// Create a new token object with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": admin.Admin_id,
		"exp":      exp, // Token expiration time

	})

	// Sign the token with a secret key and get the complete encoded token as a string
	tokenString, err := token.SignedString(JwtSKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
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
