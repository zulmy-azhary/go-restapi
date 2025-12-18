package utils

import "github.com/gofiber/fiber/v2"

func SuccessResponse(message string, data ...interface{}) fiber.Map {
	response := fiber.Map{
		"success": true,
		"message": message,
	}

	// Only add data field if data is provided
	if len(data) > 0 && data[0] != nil {
		response["data"] = data[0]
	}

	return response
}

type ErrorDetail struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}

func ErrorResponse(message string, details ...interface{}) fiber.Map {
	response := fiber.Map{
		"success": false,
		"message": message,
	}

	// If error details are provided
	if len(details) > 0 {
		if code, ok := details[0].(int); ok {
			errorType := getErrorType(code)

			// Check if custom type is provided (override auto type)
			if len(details) >= 2 {
				if customType, ok := details[1].(string); ok {
					errorType = customType
				}
			}

			response["error"] = ErrorDetail{
				Code: code,
				Type: errorType,
			}
		}
	}

	return response
}

func PaginatedResponse(message string, data interface{}, page, perPage, total int) fiber.Map {
	totalPages := (total + perPage - 1) / perPage

	return fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
		"pagination": fiber.Map{
			"page":        page,
			"per_page":    perPage,
			"total":       total,
			"total_pages": totalPages,
		},
	}
}
