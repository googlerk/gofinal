package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/googlerk/gofinal/middleware"
	"github.com/googlerk/gofinal/task"
	_ "github.com/lib/pq"
)

type Customers struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

var db *sql.DB

func init() {
	/*
		$env:DATABASE_URL='postgres://xfbpudtl:hV_LerX70gbhx4mRspt8Ajg1Y_ighZe2@john.db.elephantsql.com:5432/xfbpudtl'
		$env:DATABASE_URL='postgres://lthjkddt:sBLjG23bpGqPQrUuzoY44GjY1fPv9GxQ@john.db.elephantsql.com:5432/lthjkddt'
	*/
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS customers (
	    id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
	    status TEXT
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	callAPI := r.Group("")

	callAPI.Use(middleware.AuthMiddleware)

	callAPI.POST("/customers", task.PostCreateCustomersHandler)
	callAPI.GET("/customers/:id", task.GetCustomersByIdHandler)
	callAPI.GET("/customers", task.GetCustomersHandler)
	callAPI.PUT("/customers/:id", task.PutUpdateCustomersHandler)
	callAPI.DELETE("/customers/:id", task.DeleteCustomersHandler)

	return r
}

func main() {
	fmt.Println("customer service")
	gofinal := setupRouter()
	gofinal.Run(":2009")
	//run port ":2009"
}
