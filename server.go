package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"search_api.com/router"
)

func main() {
	server:=gin.Default()
	
	router.Router(server)

	err:=server.Run("127.0.0.1:3000")
	if err!=nil{
		log.Print("There is a error starting the server")
	}
}