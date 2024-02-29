package helpers

import (
	"app/model"
	"strings"

	"gorm.io/gorm"
)

// Old Filter
// FilterBySimilarity filters users based on similarity criteria
func FilterBySimilarity(users []model.Users, asFilter map[string]interface{}) []model.Users {
	var filterResult []model.Users

	for _, user := range users {
		// Convert the model.User to a map for filtering
		value := map[string]interface{}{
			"name":     user.Name,
			"age":      float64(user.Age),
			"location": user.Location,
		}

		isSame := true

		// Perform the similarity filter
		for key, filterValue := range asFilter {
			// Handle numeric fields differently
			switch key {
			case "age":
				userAge, ok := value[key].(float64)
				filterAge, okFilter := filterValue.(float64)
				if !ok || !okFilter || userAge != filterAge {
					isSame = false
				}
			case "name", "location": // Add cases for "name" and "location"
				userValue, ok := value[key].(string)
				if !ok || !strings.Contains(strings.ToLower(userValue), strings.ToLower(filterValue.(string))) {
					isSame = false
				}
			default:
				userValue, ok := value[key].(string)
				if !ok || userValue != filterValue {
					isSame = false
				}
			}

			if !isSame {
				break
			}
		}

		if isSame {
			filterResult = append(filterResult, user)
		}
	}

	return filterResult
}

// New Filter
func FilterLocName(request model.SearchRequest, query *gorm.DB) {
	if request.Query != "" {
		textSearch := strings.ToLower(request.Query)
		textSearchArray := strings.Split(textSearch, " ")

		for _, str := range textSearchArray {
			query.Where(`(
                LOWER(name) LIKE ? OR
                LOWER(locationarea) LIKE ?
            )`, "%"+str+"%", "%"+str+"%")
		}
	}
}
