package repository

import (
	"checkout-api/domain/model"
	"checkout-api/exception"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

type IUserRepo interface {
	Create(user *model.User) error
	Update(update *model.User) error
	GetByID(id string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetList() ([]model.User, error)
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (slf *UserRepo) Create(user *model.User) error {
	err := slf.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (slf *UserRepo) Update(update *model.User) error {
	query := slf.DB.Model(&model.User{}).Where("id = ?", update.ID).Updates(update)
	affected := query.RowsAffected
	if affected == 0 {
		return exception.DbObjNotFound
	}
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (slf *UserRepo) GetByID(id string) (*model.User, error) {
	var user model.User
	err := slf.DB.First(&user, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (slf *UserRepo) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := slf.DB.First(&user, "username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (slf *UserRepo) GetList() ([]model.User, error) {
	var users []model.User
	err := slf.DB.Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return users, nil
}

func (slf *UserRepo) Delete(id string) error {
	err := slf.DB.Where("id = ?", id).Delete(&model.User{}).Error
	if err == gorm.ErrRecordNotFound {
		return exception.DbObjNotFound
	} else if err != nil {
		return err
	}
	return nil
}
