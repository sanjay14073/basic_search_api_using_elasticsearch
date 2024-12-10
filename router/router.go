package router

import (
	"github.com/gin-gonic/gin"
	serach_controllers "search_api.com/controllers"
)

func Router(router *gin.Engine){
	//Run it once to ensure data is persent in elastic search
	router.POST("/",serach_controllers.AddData())
	//Pass ids to get it
	router.GET("/",serach_controllers.GetData())

}