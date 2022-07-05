package service

import (
	"log"

	entity "seoulspa-api/entity/admin_users"
	"seoulspa_api/repository"

	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(username string, password string) interface{}
	// CreateUser(user dto.RegisterDTO) entity.Admin_users
	FindByEmail(email string) entity.Admin_users
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username, password)
	if v, ok := res.(entity.Admin_users); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) FindByEmail(email string) entity.Admin_users {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
