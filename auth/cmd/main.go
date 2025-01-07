package main

import (
	"fmt"
	"githubmbarkhanBusTracker/auth/database"
	"githubmbarkhanBusTracker/auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start Service Auth ............")
	err := database.Connect()
	if err != nil {
		fmt.Println(err)

	} else {
		database.AutoMigrate()
		fmt.Println("Connected")
		r := gin.Default()
		routes.Api(r)
		r.Run(":8000")
	}

}