package helpers

import (
	"app/model"
	"bytes"
	"encoding/csv"
	"strconv"
)

func ConvertToCSVProduct(input []model.Products) (string, error) {
	var buf bytes.Buffer

	writer := csv.NewWriter(&buf)

	writer.Comma = ';'

	header := []string{"Product Name", "Price"}
	if err := writer.Write(header); err != nil {
		return "", err
	}

	for _, product := range input {
		row := []string{product.ProductName, strconv.Itoa(product.ProductPrice)}
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ConvertToCSV(input []model.Users) (string, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Excel ;
	// Golang ,

	writer.Comma = ';'

	// Write CSV header
	header := []string{"Name", "Age", "Address Name", "Location Area", "Hobby"}
	if err := writer.Write(header); err != nil {
		return "", err
	}

	// Write CSV rows
	for _, user := range input {
		row := []string{user.Name, strconv.Itoa(user.Age), user.AddressName, user.LocationArea, user.Hobby}
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	return buf.String(), nil
}
