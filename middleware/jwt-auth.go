package middleware

import (
	"log"
	"net/http"

	"seoulspa_api/helper"
	"seoulspa_api/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Không thể xử lý đưọc yêu cầu", "Không tìm thấy token", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token không hợp lệ", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
