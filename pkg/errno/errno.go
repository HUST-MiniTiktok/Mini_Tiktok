package errno

import (
	"errors"
	"fmt"
)

const (
	// Common
	SuccessCode                = 0
	ServiceErrCode             = 10000 + iota
	ParamErrCode
	AuthorizationFailedErrCode

	// User
	UserIsNotExistErrCode
	UserAlreadyExistErrCode
	UserPasswordErrCode

	// Video
	VideoIsNotExistErrCode
	
	// Comment
	CommentIsNotExistErrCode
)

const (
	// Common
	SuccessMsg                = "Success"
	ServiceErrMsg             = "Service is unable to start successfully"
	ParamErrMsg               = "Wrong Parameter has been given"
	AuthorizationFailedErrMsg = "Authorization failed"

	// User
	UserIsNotExistErrMsg    = "User is not exist"
	UserAlreadyExistErrMsg = "User already exists"
	UserPasswordErrMsg     = "Password is wrong"

	// Video
	VideoIsNotExistErrMsg = "Video is not exist"

	// Comment
	CommentIsNotExistErrMsg = "Comment is not exist"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	// Common
	Success = ErrNo{SuccessCode, SuccessMsg}
	ServiceErr = ErrNo{ServiceErrCode, ServiceErrMsg}
	ParamErr = ErrNo{ParamErrCode, ParamErrMsg}
	AuthorizationFailedErr = ErrNo{AuthorizationFailedErrCode, AuthorizationFailedErrMsg}

	// User
	UserIsNotExistErr = ErrNo{UserIsNotExistErrCode, UserIsNotExistErrMsg}
	UserAlreadyExistErr = ErrNo{UserAlreadyExistErrCode, UserAlreadyExistErrMsg}
	UserPasswordErr = ErrNo{UserPasswordErrCode, UserPasswordErrMsg}

	// Video
	VideoIsNotExistErr = ErrNo{VideoIsNotExistErrCode, VideoIsNotExistErrMsg}

	// Comment
	CommentIsNotExistErr = ErrNo{CommentIsNotExistErrCode, CommentIsNotExistErrMsg}
)

// convert error to Errno
func ToErrno(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}