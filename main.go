package main

import (
	"log"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	server "github.com/crane/server"
)

func main() {
	db := server.NewDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT"},
		AllowHeaders:     []string{"Origin","Content-Type","Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	  }))

	server := server.NewServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":9000", router))
}
