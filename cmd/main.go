package main

import (
	"github.com/MarceloBxD/bookstore/internal/interface/http"
	book "github.com/MarceloBxD/goandk8s/internal/domain/book"
	"github.com/MarceloBxD/goandk8s/internal/infra/db"
	"github.com/MarceloBxD/goandk8s/internal/interface/http"
	bookUseCase "github.com/MarceloBxD/goandk8s/internal/usecase/book"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	dsn := "host=localhost user=postgres password=postgres dbname=bookstore port=5432 sslmode=disable"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	dbConn.AutoMigrate(
		&book.Book{},
	)

	bookRepo := db.NewGormBookRepository(dbConn)
	createBookUC := bookUseCase.NewCreateBookUseCase(bookRepo)

	r := gin.Default()
	http.NewBookHandler(r, createBookUC)

	r.Run(":8080")
}