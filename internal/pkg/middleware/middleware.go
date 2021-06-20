package middleware

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"test_task_advertising/internal/models"
)

func CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type,X-CSRF-Token")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
		if c.Request().Method == http.MethodOptions {
			return nil
		}
		return next(c)
	}
}

func PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				errJson := ctx.JSON(http.StatusOK, &models.Answer{Code: http.StatusInternalServerError,
					Msg: "Sorry, service cannot serve this request =("})
				if errJson != nil {
					log.Println("Error in recover middleware in ctx.JSON")
				}
				return
			}
		}()
		return next(ctx)
	}
}
