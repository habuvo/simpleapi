package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"simpleapi/handlers"
	"simpleapi/lib"
)

const (
	sqlConn = "demo:demo@tcp(db)/demo"
)

func init() {
	var err error
	if lib.Env.DB, err = sql.Open("mysql", sqlConn); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	defer lib.Close(lib.Env.DB)

	if err := lib.InitDB(lib.Env.DB); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

}

func main() {

	g := gin.Default()

	// I'm fan of gRPC so i architect one route per entity instead four ones)
	// action must be create, read, update , delete
	g.POST("/company/:action", handlers.HandleCompany)
	g.POST("/contract/:action", handlers.HandleContract)
	g.POST("/purchase/create", handlers.MakePurchase)
	log.Println("server listen 8081")
	log.Fatal(http.ListenAndServe(":8081", g))

}
