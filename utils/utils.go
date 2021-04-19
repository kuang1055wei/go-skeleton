package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

//获取翻译错误信息
func GetValidateError(err error) interface{} {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}
	return RemoveTopStruct(errs.Translate(Trans))
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
