package util

import (
	"reflect"
)

func NewRespMap(statusCode int32, statusMsg string) *map[string] interface{} {
	return &map[string] interface{}{
		"StatusCode": statusCode,
		"StatusMsg":  statusMsg,
	}
}

func NewRespStruct(statusCode int32, statusMsg string, src interface{}) interface{} {
	var v_statusCode int32 = statusCode
	var v_statusMsg *string = nil
	if statusMsg != "" {
		v_statusMsg = &statusMsg
	}
	reflect.ValueOf(src).FieldByName("statusCode").Set(reflect.ValueOf(v_statusCode))
	reflect.ValueOf(src).FieldByName("statusMsg").Set(reflect.ValueOf(v_statusMsg))
	return src
}