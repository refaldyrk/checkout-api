package data

import (
	"checkout-api/config"
	"checkout-api/domain/model"
	"checkout-api/exception"
	"checkout-api/repository"
	"checkout-api/utils/helper"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// args should be empty for default seed (superuser;superuser)
// or args must be containing 2 strings for custom username, and passwords
func SeedSuperuser(userRepo repository.IUserRepo, args ...string) {
	// validate args
	if len(args) != 2 && len(args) != 0 {
		fmt.Println("invalid args, should be empty for default seed (superuser;superuser) or 2 strings for custom username, and passwords")
		return
	}

	username := "superuser"
	password := "superuser"
	if len(args) == 2 {
		username = args[0]
		password = args[1]
	}

	// check for existing user
	existingUser, err := userRepo.GetByUsername(username)
	if err != nil && err != exception.DbObjNotFound {
		panic(err)
	}
	if existingUser != nil {
		fmt.Printf("%s already exists\n", username)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := &model.User{
		ID:       helper.GenerateUUID(),
		Username: username,
		Password: string(hashedPass),
	}
	userRepo.Create(user)

	// generate access token
	token, err := helper.GenerateJwtToken(username, user.ID, user.SessionID, config.Envs.JWT_SECRET, config.Envs.JWT_EXP)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}

func SeedData(productRepo repository.IProductRepo) {
	// products
	products := []model.Product{
		{
			ID:       helper.GenerateUUID(),
			Name:     "susu",
			Price:    100,
			Category: "minuman",
		},
		{
			ID:       helper.GenerateUUID(),
			Name:     "kopi",
			Price:    300,
			Category: "minuman",
		},
		{
			ID:       helper.GenerateUUID(),
			Name:     "jeruk",
			Price:    400,
			Category: "makanan",
		},
	}

	for _, product := range products {
		_, err := productRepo.GetByNameAndPrice(product.Name, product.Price)
		if err == exception.DbObjNotFound {
			err = productRepo.Create(&product)
			if err != nil {
				panic(err)
			}
			continue
		} else if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
	}

}
