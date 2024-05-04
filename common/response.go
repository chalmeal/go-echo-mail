package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Res(c echo.Context, s int, r interface{}) error {
	var res map[string]interface{}
	if s != http.StatusOK {
		res = map[string]interface{}{
			"Status":  s,
			"Message": r,
		}
	} else {
		res = map[string]interface{}{
			"Status": s,
			"Result": r,
		}
	}
	return c.JSON(s, res)
}
