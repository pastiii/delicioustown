package validator

import (
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"regexp"
)

func CompanyIdValidation(fl validator.FieldLevel) bool {
	verificationRole := `^[a-z0-9\-]*$`
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		re, err := regexp.Compile(verificationRole)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return re.MatchString(field.String())
	default:
		return false
	}
}

func GenericStringValidation(fl validator.FieldLevel) bool {
	verificationRole := `(?:")|(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		re, err := regexp.Compile(verificationRole)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return !re.MatchString(field.String())
	default:
		return false
	}
}

func AdminUserStatusValidation(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.String() != "on" && field.String() != "off" {
		return false
	}

	return true
}
