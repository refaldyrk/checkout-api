package main

import (
	"checkout-api/config"
	"checkout-api/repository"
	"checkout-api/utils/data"
	"fmt"
	"os"
)

func CliHandler(args []string) {
	args = args[1:]

	db, err := config.NewDb(config.Envs.DB_URI)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)
	productRepo := repository.NewProductRepo(db)

	switch args[0] {
	case "seed-superuser":
		fmt.Println("running seed superuser...")
		data.SeedSuperuser(userRepo, args[1:]...)
	case "seed-data":
		fmt.Println("running seed data...")
		data.SeedData(productRepo)
	case "dev":
		fmt.Println("running playground...")
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
	fmt.Println("done")
}
