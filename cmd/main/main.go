package main

import (
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	adsPostDelivery "test_task_advertising/internal/pkg/adsPost/delivery/http"
	"test_task_advertising/internal/pkg/adsPost/repository/postgres"
	"test_task_advertising/internal/pkg/adsPost/usecase"
	"test_task_advertising/internal/pkg/middleware"
)

// todo read conf file
const (
	postgresConnStr         = "user=docker password=docker dbname=myService sslmode=disable port=5432 host=pg"
	postgresConnectionCount = 10
	goServerPortStr         = ":8080"
	apiStr                  = "/api/v1"
)

func main() {

	// connect postgreSQL
	adsPostRepoPostgres, errConnect := postgres.NewAdsPostRepository(postgresConnStr, postgresConnectionCount)
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

	adsPostUseCase := usecase.NewUseCase(adsPostRepoPostgres)

	adsPostHandler := adsPostDelivery.AdsPostHandler{AdsPostUseCase: adsPostUseCase}

	e := echo.New()

	// middleware
	e.Use(middleware.PanicMiddleware)
	e.Use(middleware.CORS)

	// ads posts
	e.POST(apiStr+"/adsPost", adsPostHandler.CreateAdsPost)
	e.GET(apiStr+"/adsPost", adsPostHandler.GetAdsPost)
	e.GET(apiStr+"/adsPosts", adsPostHandler.GetAdsPostArr)

	log.Fatal(e.Start(goServerPortStr))
}
