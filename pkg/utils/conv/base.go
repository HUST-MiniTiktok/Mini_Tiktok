package conv

import (
	"errors"
	"reflect"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
)

type HertzBaseResponse struct {
	StatusCode int32
	StatusMsg  string
}

func ToHertzBaseResponse(err error) *HertzBaseResponse {
	if err == nil {
		return &HertzBaseResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
		}
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return &HertzBaseResponse{
			StatusCode: e.ErrCode,
			StatusMsg:  e.ErrMsg,
		}
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return &HertzBaseResponse{
		StatusCode: s.ErrCode,
		StatusMsg:  s.ErrMsg,
	}
}

func ToKitexBaseResponse(err error, src interface{}) interface{} {
	var v_statusCode int32
	var v_statusMsg *string
	e := errno.ErrNo{}
	if err == nil {
		v_statusCode = errno.Success.ErrCode
		v_statusMsg = &errno.Success.ErrMsg
	} else if errors.As(err, &e) {
		v_statusCode = e.ErrCode
		v_statusMsg = &e.ErrMsg
	} else {
		s := errno.ServiceErr.WithMessage(err.Error())
		v_statusCode = s.ErrCode
		v_statusMsg = &s.ErrMsg
	}
	v_src := reflect.Indirect(reflect.ValueOf(src))
	v_src.FieldByName("StatusCode").Set(reflect.ValueOf(v_statusCode))
	v_src.FieldByName("StatusMsg").Set(reflect.ValueOf(v_statusMsg))
	return src
}
