package jwt

import (
	"go-admin/pkg/e"
	userservice "go-admin/service/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		bearerToken := c.Request.Header.Get("Authorization")
		space := strings.Index(bearerToken, " ") + 1
		token := bearerToken[space:]
		log.Printf("%v", token)
		if token == "" {
			code = e.ERROR_AUTH_NEED_TOKEN
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.INTERNAL_ERROR
				}
			}
			userService := userservice.User{}
			jwtUID, err := userService.GetIDByMD5(claims.Username)
			log.Println(jwtUID)
			if err != nil {
				code = e.ERROR_AUTH_GETUID_BY_TOKEN
			}
			if jwtUID == 0 {
				code = e.ERROR_AUTH_GETUID_BY_TOKEN
			}
			c.Set("jwtuid", jwtUID)
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
