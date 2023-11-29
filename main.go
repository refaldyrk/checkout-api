package main

import (
	"checkout-api/config"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	_ "checkout-api/docs"
)

// @title Backend checkout-api API
// @securitydefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme (add 'Bearer ' prefix).
func main() {
	args := os.Args
	if len(args) > 1 {
		CliHandler(args)
	} else {
		log.Printf("Envs: %v", config.Envs)
		log.Println("starting rest api app...")

		router := gin.Default()
		SetupServer(router)
		router.Run(config.Envs.ADDR)

		fmt.Println("starting rest api app...")
	}
}
