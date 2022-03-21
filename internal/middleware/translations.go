package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

//validator做入参校验时，错误信息是英文
//编写中间件将其转换成中文
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(),zh.New(),zh_Hant.New())
		//获取上下文的语言环境
		locale := c.GetHeader("local")
		trans,_:= uni.GetTranslator(locale)
		v,ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v,trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v,trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v,trans)
				break
			}
			c.Set("trans",trans)
		}
		c.Next()
	}
}


