package main

import (
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	"test_task_advertising/internal/adsPost/delivery/http"
	"test_task_advertising/internal/adsPost/repository/postgres"
	"test_task_advertising/internal/adsPost/usecase"
	"test_task_advertising/internal/pkg/config"
	loggerZap "test_task_advertising/internal/pkg/logger"
	"test_task_advertising/internal/pkg/middleware"
	"time"
)

// todo add server config file like DB config
const (
	goServerPortStr = ":8080"
	apiStr          = "/api/v1"
	logLvl          = "INFO"
)

func main() {

	confDB, err := config.GetDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	adsPostRepoPostgres, errConnect := postgres.NewAdsPostRepository(confDB)
	if errConnect != nil {
		log.Fatal(errConnect)
	}
	defer func() {
		errClose := adsPostRepoPostgres.CloseAdsPost()
		if errClose != nil {
			log.Fatal(errClose)
		}
	}()
	log.Println("PostgreSQL connect successfully")

	logger := loggerZap.NewLogger(logLvl)

	adsPostUseCase := usecase.NewUseCase(adsPostRepoPostgres, logger)

	adsPostHandler := http.AdsPostHandler{AdsPostUseCase: adsPostUseCase}

	e := echo.New()

	middlewareWithLogger := middleware.MiddlewareWithLogger{Logger: logger}

	// middleware
	e.Use(middleware.PanicMiddleware)
	e.Use(middlewareWithLogger.LogMiddleware)
	e.Use(middleware.CORS)

	// ads posts
	e.POST(apiStr+"/adsPost", adsPostHandler.CreateAdsPost)
	e.GET(apiStr+"/adsPost", adsPostHandler.GetAdsPost)
	e.GET(apiStr+"/adsPosts", adsPostHandler.GetAdsPostArr)

	// better set timeout and CORS in nginx
	e.Server.ReadTimeout = time.Second * 10
	e.Server.WriteTimeout = time.Second * 10

	log.Fatal(e.Start(goServerPortStr))
}
