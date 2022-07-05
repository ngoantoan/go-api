package controller

import (
	"net/http"
	entity "seoulspa-api/entity/admin_users"
	"strconv"

	dto "seoulspa_api/dto/admin-users"
	"seoulspa_api/helper"
	"seoulspa_api/service"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	// Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Username, loginDTO.Password)
	if v, ok := authResult.(entity.Admin_users); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Vui lòng kiểm tra lại thông tin đăng nhập", "Tài khoản hoặc mật khẩu không đúng", authResult)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
