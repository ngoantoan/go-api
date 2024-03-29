package repository

import (
	"log"
	entity "seoulspa-api/entity/admin_users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user entity.Admin_users) entity.Admin_users
	UpdateUser(user entity.Admin_users) entity.Admin_users
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.Admin_users
	ProfileUser(userID string) entity.Admin_users
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.Admin_users) entity.Admin_users {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.Admin_users) entity.Admin_users {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.Admin_users
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user entity.Admin_users
	res := db.connection.Where("username = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.Admin_users
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.Admin_users {
	var user entity.Admin_users
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.Admin_users {
	var user entity.Admin_users
	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)

	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
