package controllers

import (
	"app/helpers"
	"app/helpers/connection"
	"app/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {
	var admin model.Admin

	if err := c.Bind(&admin); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error Bind admin")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(admin.AdminPass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin.AdminPass = string(hashedPass)

	if err := connection.CreateAdmin(admin); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error creating admin")
	}

	token, err := connection.GenerateJWT(admin, 24)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, map[string]any{
		"Token": token,
	})

}

func Login(c echo.Context) error {

	var adminLogin model.Admin

	if err := c.Bind(&adminLogin); err != nil {
		return err
	}

	admin, err := connection.GetAdminByEmail(adminLogin.AdminEmail)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error get admin Email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.AdminPass), []byte(adminLogin.AdminPass)); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error compare")
	}

	token, err := connection.GenerateJWT(*admin, 24*time.Hour)
	if err != nil {
		return err
	}

	expTime := time.Now().Add(24 * time.Hour)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":  token,
		"Exp at": expTime.Format(time.RFC3339),
	})

}
