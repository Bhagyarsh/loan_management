package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhagyarsh/loan_management/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	const serverAdrress = "0.0.0.0:8000"
	log.Println("stating server at port " + serverAdrress)
	// r.GET("/ping", func(c *gin.Context) {
	// 	log.Println(c.Query("t"))
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/loan", controllers.GetAlloan)
	r.GET("/loan/:id", controllers.GetloanByID)
	r.PATCH("/loan/status/update/:id", controllers.ApproveloanByID)

	r.POST("/loan/create", controllers.Createloan)
	r.DELETE("/loan/:id", controllers.Cancelloan)
	fmt.Println(http.StatusOK)
	r.Run(serverAdrress)

}
