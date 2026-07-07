package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Vincent088/btg-backend-test/go-backend/internal/config"
	deliveryhttp "github.com/Vincent088/btg-backend-test/go-backend/internal/delivery/http"
	"github.com/Vincent088/btg-backend-test/go-backend/internal/repository"
	"github.com/Vincent088/btg-backend-test/go-backend/internal/usecase"
)

func main() {
	_ = godotenv.Load()

	db := config.NewPostgresConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	defer db.Close()

	customerRepo := repository.NewCustomerRepository(db)
	familyRepo := repository.NewFamilyListRepository(db)
	nationalityRepo := repository.NewNationalityRepository(db)

	customerUsecase := usecase.NewCustomerUsecase(customerRepo, familyRepo, nationalityRepo)
	nationalityUsecase := usecase.NewNationalityUsecase(nationalityRepo)

	customerHandler := deliveryhttp.NewCustomerHandler(customerUsecase)
	nationalityHandler := deliveryhttp.NewNationalityHandler(nationalityUsecase)

	router := deliveryhttp.NewRouter(customerHandler, nationalityHandler)

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
