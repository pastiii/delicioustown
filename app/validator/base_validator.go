package validator

import (
	"DeliciousTown/app/util/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate = validator.New()
	chinese  = zh.New()
	uni      = ut.New(chinese, chinese)
	trans, _ = uni.GetTranslator("zh")
	_        = validate.RegisterValidation("companyIdValidation", CompanyIdValidation) // 注册自定义验证规则
	_        = validate.RegisterValidation("gsv", GenericStringValidation)
	_        = validate.RegisterValidation("adminUserStatus", AdminUserStatusValidation)
)

func BaseValidate(context *gin.Context, s interface{}) bool {
	_ = zhs.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(s)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, v := range errors.Translate(trans) {
				response.ValidateError(context, v)
				break
			}
		} else {
			response.ValidateError(context, err.Error())
		}

		return false
	}

	return true
}
