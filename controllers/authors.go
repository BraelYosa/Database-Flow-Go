package controllers

import (
	"app/helpers"
	"app/helpers/connection/postgres"
	"app/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {
	var admin model.Admin

	if err := c.Bind(&admin); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_Bind_admin")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(admin.AdminPass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin.AdminPass = string(hashedPass)

	if err := postgres.CreateAdmin(admin); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_creating_admin")
	}

	token, err := postgres.GenerateJWT(admin, 24)
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

	admin, err := postgres.GetAdminByEmail(adminLogin.AdminEmail)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_get_admin_Email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.AdminPass), []byte(adminLogin.AdminPass)); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_compare")
	}

	token, err := postgres.GenerateJWT(*admin, 24*time.Hour)
	if err != nil {
		return err
	}

	expTime := time.Now().Add(24 * time.Hour)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":  token,
		"Exp_at": expTime.Format(time.RFC3339),
	})

}
