package postgres

import (
	"app/helpers"
	"app/model"
	"os"

	"github.com/pkg/errors"
)

func CreateProduct(newProduct model.Products) error {

	db := DB

	var existingProduct model.Products

	// Check if the admin ID exists in the admins table
	var adminCount int64
	if err := db.Table("admins").Where("admin_id = ?", newProduct.AdminID).Count(&adminCount).Error; err != nil {
		return err
	}

	if adminCount == 0 {
		return errors.New("Error: Admin ID not found in the database")
	}

	// Check if a product with the same name already exists
	if err := db.Table("products").Where("product_name = ?", newProduct.ProductName).First(&existingProduct).Error; err == nil {
		return errors.New("Product with the same name already exists")
	}

	if err := db.Create(&newProduct).Error; err != nil {
		return err
	}

	return nil

}

func SearchProduct(request model.SearchRequest) (interface{}, error) {

	var product []model.Products
	db := DB

	queryBuilder := db.Table("products").Where("1=1")

	if request.SortBy != "" {
		// SortBy should contain both column name and sorting order
		sortBy := request.SortBy
		queryBuilder = queryBuilder.Order(sortBy)
	}

	offset := (request.Page - 1) * request.Limit
	queryBuilder = queryBuilder.Offset(offset).Limit(request.Limit)

	if err := queryBuilder.Find(&product).Error; err != nil {
		return err, nil
	}

	if request.Output == "CSV" {
		csvData, err := helpers.ConvertToCSVProduct(product)
		if err != nil {
			return nil, err
		}

		folderPath := "file/CSV/"
		fileName := folderPath + "productOutput." + request.Output
		if err := os.WriteFile(fileName, []byte(csvData), 0644); err != nil {
			return nil, err
		}

		return csvData, nil
	} else if request.Output == "" || request.Output == "JSON" {
		return product, nil
	}

	return request.Output, nil
}

func UpdateProduct(productID string, updatedProduct *model.Products) error {
	var existingProduct model.Products
	db := DB

	// Retrieve the existing product by its ID
	if err := db.Table("products").First(&existingProduct, "product_id = ?", productID).Error; err != nil {
		return err
	}

	existingProduct.ProductName = updatedProduct.ProductName
	existingProduct.ProductPrice = updatedProduct.ProductPrice

	if err := db.Table("products").Where("product_id = ?", productID).Updates(updatedProduct).Error; err != nil {
		return err
	}

	return nil

}

func ViewProduct(ProductID string) (*model.Products, error) {
	db := DB

	var products model.Products
	if err := db.Table("products").First(&products, "product_id = ?", ProductID).Error; err != nil {
		return nil, err
	}

	return &products, nil

}
