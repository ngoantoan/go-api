package service

import (
	"log"

	"seoulspa_api/dto"
	"seoulspa_api/entity"
	"seoulspa_api/repository"

	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.Admin_users
	Profile(userID string) entity.Admin_users
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.Admin_users {
	userToUpdate := entity.Admin_users{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.Admin_users {
	return service.userRepository.ProfileUser(userID)
}
