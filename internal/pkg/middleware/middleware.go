package middleware

import (
	"fmt"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"net/http"
	"test_task_advertising/internal/constants"
	"test_task_advertising/internal/models"
	"time"
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

type MiddlewareWithLogger struct {
	Logger *zap.Logger
}

func (ml *MiddlewareWithLogger) LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		start := time.Now()
		requestId := fmt.Sprintf("%016x", rand.Int())[:10]
		ctx.Set(constants.REQUEST_ID_KEY, requestId)

		defer func() {
			ml.Logger.Info(ctx.Path(),
				zap.String("RequestId:", requestId),
				zap.String("Method:", ctx.Request().Method),
				zap.String("RemoteAddr:", ctx.Request().RemoteAddr),
				zap.Time("StartTime:", start),
				zap.Duration("DurationTime:", time.Since(start)),
			)
		}()
		return next(ctx)
	}
}
