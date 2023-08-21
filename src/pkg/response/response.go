package response

import "github.com/labstack/echo/v4"

func ErrResponse(c echo.Context, status int, message string) error {
	type errResponse struct {
		Message string `json:"message"`
	}

	return c.JSON(status, &errResponse{
		Message: message,
	})
}

func SuccessResponse(c echo.Context, status int, data any) error {
	return c.JSON(status, data)
}
