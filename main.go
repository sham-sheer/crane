package main

import (
	"log"

	"net/http"
	"github.com/gin-gonic/gin"


	server "github.com/crane/server"
)



func main() {
	db := server.NewDB()

	router := gin.Default()

	server := server.NewServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":9000", router))
}
