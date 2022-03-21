package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"myproject/pkg/app"
	"myproject/pkg/errorcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var(
			token string
			ecode = errorcode.Success
		)
		s,exists := c.GetQuery("token")
		if exists{
			token = s
		}else{
			token = c.GetHeader("token")
		}
		if token == ""{
			ecode = errorcode.InvalidParams
		}else{
			_,err := app.ParseToken(token)
			if err != nil{
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errorcode.UnauthorizedTokenTimeout
				default:
					ecode = errorcode.UnauthorizedTokenError
				}
			}
		}
		if ecode != errorcode.Success{
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
