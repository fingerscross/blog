package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

//GetValidMsg  返回结构体中的msg参数

func GetValidMsg(err error, obj any) string {

	//使用的时候需要传obj的指针
	getObj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	errors, ok := err.(validator.ValidationErrors)
	if ok {
		//断言成功
		for _, e := range errors {
			//循环每一个错误信息
			//根据报错字段名 获取结构体具体字段
			f, exits := getObj.Elem().FieldByName(e.Field())
			if exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}