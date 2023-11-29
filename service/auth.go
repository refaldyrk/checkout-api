package service

import (
	"checkout-api/config"
	"checkout-api/domain/model"
	"checkout-api/repository"
	error_utils "checkout-api/utils/error"
	"checkout-api/utils/helper"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.IUserRepo
}

type IAuthService interface {
	Login(username string, password string) (string, error)
	CheckToken(token string) (*model.User, error)
	Register(username string, password string) error
}

func NewAuthService(userRepo repository.IUserRepo) IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (slf *AuthService) Login(username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", &error_utils.CustomErr{
			Code:    400,
			Message: "username and password are required",
		}
	}

	// check existance by username
	oldUser, err := slf.userRepo.GetByUsername(username)
	if err != nil {
		return "", &error_utils.CustomErr{
			Code:    400,
			Message: "username not found",
		}
	}

	// check password
	isPwMatch := helper.ComparePasswordHash(password, oldUser.Password)
	if !isPwMatch {
		return "", &error_utils.CustomErr{
			Code:    401,
			Message: "password incorrect",
		}
	}

	// generate token
	newSessionID := helper.GenerateUUID()
	token, err := helper.GenerateJwtToken(username, oldUser.ID, newSessionID, config.Envs.JWT_SECRET, config.Envs.JWT_EXP)
	if err != nil {
		return "", &error_utils.CustomErr{
			Code:    500,
			Message: "error when generating jwt token",
		}
	}

	// update session id from user
	oldUser.SessionID = newSessionID
	err = slf.userRepo.Update(oldUser)
	if err != nil {
		return "", &error_utils.CustomErr{
			Code:    404,
			Message: fmt.Sprintf("user with id %s not found", oldUser.ID),
		}
	}

	return token, nil
}

func (slf *AuthService) CheckToken(token string) (*model.User, error) {
	claims, err := helper.ValidateJWT(token)
	if err != nil {
		return nil, &error_utils.CustomErr{
			Code:    401,
			Message: "invalid token",
		}
	}

	userID := claims["user_id"].(string)
	user, err := slf.userRepo.GetByID(userID)
	if err != nil {
		return nil, &error_utils.CustomErr{
			Code:    401,
			Message: "user not found",
		}
	}

	return user, nil
}

func (slf *AuthService) Register(username string, password string) error {
	// user & email existance validation
	userByUsername, _ := slf.userRepo.GetByUsername(username)
	if userByUsername != nil {
		return &error_utils.CustomErr{
			Code:    400,
			Message: "username already exist",
		}
	}

	// hash password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// new object creation
	uuid := helper.GenerateUUID()
	// timeNow := time.Now().Unix()
	newUser := &model.User{
		ID:       uuid,
		Username: username,
		Password: string(hashedPass),
	}
	err = slf.userRepo.Create(newUser)
	if err != nil {
		return err
	}

	return nil
}
