package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lovelyrrg51/go_backend/app/common"
	"github.com/lovelyrrg51/go_backend/app/dtos"
	"github.com/lovelyrrg51/go_backend/app/models"
	"github.com/lovelyrrg51/go_backend/app/repositories"
	"github.com/lovelyrrg51/go_backend/app/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(r dtos.AuthRegisterRequest) (*dtos.AuthRegisterResponse, *common.AppError)
	Login(r dtos.AuthLoginRequest) (*dtos.AuthLoginResponse, *common.AppError)
	GetProfile(userId uuid.UUID) (*dtos.UserProfileResponse, *common.AppError)
}

type DefaultUserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) DefaultUserService {
	return DefaultUserService{
		userRepo: repo,
	}
}

func (uService DefaultUserService) GetNormalUserResponse(u models.User) *dtos.NormalUserResponse {
	normalResponse := dtos.NormalUserResponse{
		Username:       u.Username,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		IsEmailConfirm: u.IsEmailConfirm,
	}

	return &normalResponse
}

func (uService DefaultUserService) Register(r dtos.AuthRegisterRequest) (*dtos.AuthRegisterResponse, *common.AppError) {
	// Check User Existed with Email
	user, _ := uService.userRepo.FindByEmail(r.Email)
	if user != nil {
		return nil, common.NewBadRequestError("Email Already Exist.")
	}

	// Check User Existed with Username
	user, _ = uService.userRepo.FindByUsername(r.Username)
	if user != nil {
		return nil, common.NewBadRequestError("Username Already Exist.")
	}

	// Conver Password to Hash
	hashedPassword, bcryptErr := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if bcryptErr != nil {
		fmt.Printf("UserService: Bcrypt Hash Password Error " + fmt.Sprintf("%s", bcryptErr))
		return nil, common.NewUnexpectedError("Somethig went wrong.")
	}

	// Get New UUID
	uuid := uuid.New()
	// Create New User && Save
	var newUser = models.User{
		ID:             uuid,
		Username:       r.Username,
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		Email:          r.Email,
		Password:       string(hashedPassword),
		IsDelete:       false,
		IsEmailConfirm: false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	_, saveErr := uService.userRepo.Save(newUser)
	if saveErr != nil {
		return nil, saveErr
	}

	// Response
	normalUser := uService.GetNormalUserResponse(newUser)
	response := dtos.AuthRegisterResponse{
		Status: "success",
		Data:   *normalUser,
	}

	return &response, nil
}

func (uService DefaultUserService) Login(r dtos.AuthLoginRequest) (*dtos.AuthLoginResponse, *common.AppError) {
	// Check User with Email or Username
	existedUser, err := uService.userRepo.FindByUsernameOrEmail(r.Username)
	if err != nil {
		return nil, common.NewBadRequestError("Invalid username or password.")
	}

	// Check User Password
	checkPassErr := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(r.Password))
	if checkPassErr != nil {
		return nil, common.NewBadRequestError("Invalid username or password.")
	}

	// Generate Standard JWT
	_, authToken, err := utils.GenerateStandardJWT(existedUser.ID)
	if err != nil {
		return nil, err
	}

	// Response
	normalUser := uService.GetNormalUserResponse(*existedUser)
	response := dtos.AuthLoginResponse{
		Status:    "success",
		Data:      *normalUser,
		AuthToken: *authToken,
	}

	return &response, nil
}

func (uService DefaultUserService) GetProfile(userId uuid.UUID) (*dtos.UserProfileResponse, *common.AppError) {
	// Get User
	existedUser, err := uService.userRepo.FindById(userId)
	if err != nil {
		return nil, common.NewBadRequestError("No existed User")
	}

	// Generate Standard JWT
	_, authToken, err := utils.GenerateStandardJWT(existedUser.ID)
	if err != nil {
		return nil, err
	}

	// Response
	normalUser := uService.GetNormalUserResponse(*existedUser)
	response := dtos.UserProfileResponse{
		Status:    "success",
		Data:      *normalUser,
		AuthToken: *authToken,
	}

	return &response, nil
}
