package tools

import (
	"bytes"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"log/slog"
)

func ValidateData(data interface{}) error {
	if data == nil {
		return errors.New("验证参数为空！")
	}
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		var e *validator.InvalidValidationError
		if errors.As(err, &e) {
			slog.Error("输入参数错误！", "reason", err)
			return errors.New("输入参数错误！" + err.Error())
		} else {
			var buf bytes.Buffer
			buf.WriteString("参数验证错误！")
			for s, e := range err.(validator.ValidationErrors) {
				slog.Error(e.Error())
				buf.WriteString(e.Field())
				if s < len(err.(validator.ValidationErrors))-1 {
					buf.WriteString(",")
				}
			}
			slog.Error("参数验证错误！")
			return errors.New(buf.String())
		}
	}
	return nil
}
