package repositories

import (
	"github.com/google/uuid"
	"github.com/lovelyrrg51/go_backend/app/common"
	"github.com/lovelyrrg51/go_backend/app/logger"
	"github.com/lovelyrrg51/go_backend/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) (*models.User, *common.AppError)
	FindById(id uuid.UUID) (*models.User, *common.AppError)
	FindByUsername(username string) (*models.User, *common.AppError)
	FindByEmail(email string) (*models.User, *common.AppError)
	FindByUsernameOrEmail(key string) (*models.User, *common.AppError)
}

type DefaultUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *DefaultUserRepository {
	return &DefaultUserRepository{db}
}

func (d *DefaultUserRepository) Save(user models.User) (*models.User, *common.AppError) {
	if err := d.db.Save(&user).Error; err != nil {
		logger.Error("Error when update user " + err.Error())
		return nil, common.NewUnexpectedError("Unexpected error when update user " + err.Error())
	}
	return &user, nil
}

func (d *DefaultUserRepository) FindById(id uuid.UUID) (*models.User, *common.AppError) {
	var user models.User

	if err := d.db.First(&user, id).Error; err != nil {
		logger.Error("Error when find user by id " + err.Error())
		return nil, common.NewUnexpectedError("Unexpected error when find user by id " + err.Error())
	}
	return &user, nil
}

func (d *DefaultUserRepository) FindByUsername(username string) (*models.User, *common.AppError) {
	var user models.User

	err := d.db.First(&user, models.User{Username: username}).Error
	if err != nil {
		return nil, common.NewUnexpectedError(err.Error())
	}
	return &user, nil
}

func (d *DefaultUserRepository) FindByEmail(email string) (*models.User, *common.AppError) {
	var user models.User

	err := d.db.First(&user, models.User{Email: email}).Error
	if err != nil {
		return nil, common.NewUnexpectedError(err.Error())
	}
	return &user, nil
}

func (d *DefaultUserRepository) FindByUsernameOrEmail(key string) (*models.User, *common.AppError) {
	var user models.User

	err := d.db.First(&user, "username = ? OR email = ?", key, key).Error
	if err != nil {
		return nil, common.NewUnexpectedError(err.Error())
	}
	return &user, nil
}
