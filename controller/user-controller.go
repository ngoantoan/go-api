package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"golang_api/dto"
	"golang_api/helper"
	"golang_api/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"golang_api/config"

	"gorm.io/gorm"

	"golang_api/entity"

	"github.com/gin-gonic/gin/binding"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

//UserController is a ....
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	// authHeader := context.GetHeader("Authorization")
	// token, err := c.jwtService.ValidateToken(authHeader)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// id := fmt.Sprintf("%v", claims["user_id"])
	// user := c.userService.Profile(id)
	// res := helper.BuildResponse(true, "OK", user)
	// context.JSON(http.StatusOK, res)

	type ParamsOne struct {
		Username string `json:"username"`
		Phone    string `json:"phone"`
	}

	var f ParamsOne
	// Read ones
	context.ShouldBindBodyWith(&f, binding.JSON)

	query := `
		SELECT
			us.id,
			CONCAT(us.last_name, " ", us.first_name) AS fullname,
			st.description as store_name
		FROM
			admin_users us
		JOIN admin_stores st ON st.id = us.main_store_id
		WHERE us.main_store_id = 8`
	rows, err := db.Raw(query).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var t entity.ListUser
	var ts []entity.ListUser
	for rows.Next() {
		rows.Scan(
			&t.ID,
			&t.Fullname,
			&t.Store_name,
		)

		ts = append(ts, t)
	}

	res := helper.BuildResponse(true, "OK", f)
	context.JSON(http.StatusOK, res)
}
