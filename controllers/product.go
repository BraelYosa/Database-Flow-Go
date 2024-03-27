package controllers

import (
	"app/helpers"
	"app/helpers/connection/postgres"
	"app/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateProduct(c echo.Context) error {
	newResponse := helpers.Response()

	var getRequest []model.Products

	if err := c.Bind(&getRequest); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_Bind_Product")
	}

	for _, product := range getRequest {
		if err := postgres.CreateProduct(product); err != nil {
			if strings.Contains(err.Error(), "Admin ID not found") {
				return helpers.ErrorResponse(c, http.StatusNotFound, "Admin ID not found")
			}

			if strings.Contains(err.Error(), "Product with the same name already exists") {
				return helpers.ErrorResponse(c, http.StatusConflict, "Product with the same name already exists")
			}

			// any other errors
			return helpers.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		}
	}

	newResponse["Status"] = "Created"
	newResponse["Status_Type"] = http.StatusAccepted
	newResponse["Created"] = getRequest

	return helpers.SuccessResponse(c, newResponse, http.StatusAccepted)

}

func SearchProduct(c echo.Context) error {
	newResponse := helpers.Response()

	// Parse the request body to get the search parameters
	var searchRequest model.SearchRequest
	if err := c.Bind(&searchRequest); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid_request_body")
	}

	products, err := postgres.SearchProduct(searchRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newResponse["Status"] = "No_records_found"
			newResponse["Status_Type"] = http.StatusNotFound
			return helpers.ErrorResponse(c, http.StatusNotFound, newResponse)
		}
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_searching_products")
	}

	newResponse["Status"] = "Success"
	newResponse["Status_Type"] = http.StatusOK
	newResponse["Found"] = products

	if strings.ToUpper(searchRequest.Output) == "CSV" {
		return c.JSONPretty(200, map[string]any{
			"code":    200,
			"message": "",
			"data":    c.File("./file/CSV/productOutput.CSV"),
		}, " ")
	}
	return helpers.SuccessResponse(c, newResponse, http.StatusOK)
}

func UpdateProduct(c echo.Context) error {

	newResponse := helpers.Response()

	var updateReq model.UpdateReq

	if err := c.Bind(&updateReq); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_Bind")
	}

	err := postgres.UpdateProduct(updateReq.ProductID, &updateReq.UpdateProduct)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_Update_Product")
	}

	newResponse["Status"] = "Updated"
	newResponse["Status_Type"] = http.StatusOK
	newResponse["Created"] = updateReq.UpdateProduct

	return helpers.SuccessResponse(c, newResponse, http.StatusOK)

}

func ViewProduct(c echo.Context) error {
	newResponse := helpers.Response()

	ProductID := c.QueryParam("product_id")
	if ProductID == "" {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_ProductID")
	}

	product, err := postgres.ViewProduct(ProductID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_View_products")
		}
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_View_Product")
	}

	newResponse["Status"] = "Success"
	newResponse["Status_Type"] = http.StatusOK
	newResponse["Found"] = product

	return helpers.SuccessResponse(c, newResponse, http.StatusOK)

}
