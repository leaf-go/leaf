package mounts

import (
	"github.com/go-playground/validator/v10"
	"leaf-go/utils"
	"regexp"
	"strconv"
	"x"
)

var (
	Validator *x.Validator
	rules     = x.Rules{
		"boolean": func(fl validator.FieldLevel) bool {
			switch fl.Field().Interface().(type) {
			case bool:
				return true
			}

			return false
		},
		"mobile": func(fl validator.FieldLevel) bool {
			if mobile := fl.Field().String(); mobile != "" {
				mobile := utils.Regexp.ReplaceAllString(`[-|\s]`, "", mobile)
				return utils.Validate.Mobile(mobile)
			}
			return false
		},
		"mobileWithTel": func(fl validator.FieldLevel) bool {
			if mobile := fl.Field().String(); mobile != "" {
				mobile := utils.Regexp.ReplaceAllString(`[-|\s]`, "", mobile)
				return utils.Validate.MobileWithTel(mobile)
			}
			return false
		},
		"bankCard": func(fl validator.FieldLevel) bool {
			if bankCard := fl.Field().String(); bankCard != "" {
				return utils.Validate.BankCard(bankCard)
			}
			return false
		},
		"version": func(fl validator.FieldLevel) bool {
			if version := fl.Field().String(); version != "" {
				return utils.Validate.Version(version)
			}
			return false
		},
		"password": func(fl validator.FieldLevel) bool {
			if password := fl.Field().String(); password != "" {
				ok, _ := regexp.MatchString(`[0-9a-zA-Z]{6,20}`, password)
				if !ok {
					return false
				}

				ok, _ = regexp.MatchString(`^(\d+)|([a-zA-Z])$`, password)
				return !ok
			}
			return false
		},
		"idCard": func(fl validator.FieldLevel) bool {
			if idCard := fl.Field().String(); idCard != "" {
				ok, _ := regexp.MatchString(`^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`, idCard)
				return ok
			}
			return false
		},
		"region_code": func(fl validator.FieldLevel) bool {
			if code := fl.Field().Int(); code != 0 {
				return code >= 100000 && code <= 900000
			}
			return false
		},
		"order_no": func(fl validator.FieldLevel) bool {
			if orderNo := fl.Field().Int(); orderNo != 0 {
				s := strconv.FormatInt(orderNo, 10)
				ok, _ := regexp.MatchString(`^[1-9]\d{17}$`, s)
				return ok
			}
			return false
		},
	}
)

func init() {
	Validator = x.NewValidator(
		validator.New(),
		rules,
		"failed",
	)
}
