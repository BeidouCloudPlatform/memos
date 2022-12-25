package server

import "github.com/labstack/echo/v4"

func composeResponse(data interface{}) interface{} {
	type R struct {
		Data interface{} `json:"data"`
	}

	return R{
		Data: data,
	}
}

func ua(c echo.Context) string {
	ua := c.Request().Header.Get("Memos-Type")
	switch ua {
	case "wx", "wxapp":
		return "微信小程序"
	default:
		return "default"
	}
}
