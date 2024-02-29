package controllers

import (
	"app/helpers"
	"app/helpers/connection"
	"app/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Took me just a day
func CreateUser(c echo.Context) error {
	newResponse := helpers.Response()

	var dataResponse []model.Users
	if err := c.Bind(&dataResponse); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error Bind Data")
	}

	for _, user := range dataResponse {
		if err := connection.CreateUser(user); err != nil {
			if strings.Contains(err.Error(), "User with the same name already exits") {
				return helpers.ErrorResponse(c, http.StatusInternalServerError, "User with the same name already exits")
			}
			return helpers.ErrorResponse(c, http.StatusInternalServerError, "User with the same name already exits")
		}
	}

	newResponse["Status"] = "Created"
	newResponse["Status Type"] = http.StatusAccepted
	newResponse["Created"] = dataResponse

	return helpers.SuccessResponse(c, newResponse, http.StatusAccepted)
}

// Took me 4 days for this func
func SearchUsers(c echo.Context) error {
	newResponse := helpers.Response()

	// Parse the request body to get the search parameters
	var searchRequest model.SearchRequest
	if err := c.Bind(&searchRequest); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	users, err := connection.SearchUsers(searchRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newResponse["Status"] = "No records found"
			newResponse["Status Type"] = http.StatusNotFound
			return helpers.ErrorResponse(c, http.StatusNotFound, newResponse)
		}
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error searching users")
	}

	newResponse["Status"] = "Success"
	newResponse["Status Type"] = http.StatusOK
	newResponse["Found"] = users

	if strings.ToUpper(searchRequest.Output) == "CSV" {
		return c.JSONPretty(200, map[string]any{
			"code":    200,
			"message": "",
			"data":    c.File("./file/CSV/output.CSV"),
		}, " ")
	}
	return helpers.SuccessResponse(c, newResponse, http.StatusOK)

	// Response to struct
}

// Took me 2 days
func UpdateUser(c echo.Context) error {

	newResponse := helpers.Response()

	var updateReq model.UpdateReq
	if err := c.Bind(&updateReq); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Invalid req body")
	}

	err := connection.UpdateUser(updateReq.UserID, &updateReq.UpdateUser)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error updating user")
	}

	newResponse["Status"] = "Updated"
	newResponse["Status Type"] = http.StatusOK
	newResponse["Created"] = updateReq.UpdateUser

	return helpers.SuccessResponse(c, newResponse, http.StatusOK)

}

// Took me 2
func ViewUser(c echo.Context) error {
	newResponse := helpers.Response()

	userID := c.QueryParam("userID")
	if userID == "" {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "User ID is required")
	}

	// Retrieve user by ID
	user, err := connection.ViewUser(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.ErrorResponse(c, http.StatusNotFound, "User not found")
		}
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error retrieving user")
	}

	newResponse["Status"] = "Success"
	newResponse["Status Type"] = http.StatusOK
	newResponse["Found"] = user

	return helpers.SuccessResponse(c, newResponse, http.StatusOK)
}
