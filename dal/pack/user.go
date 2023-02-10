package pack

import (
	"douyin/pkg/errno"
	"errors"
)

func BuildCodeAndMsg(err error) (statusCode int64, statusMsg *string) {
	if err == nil {
		return errno.Success.ErrCode, &errno.Success.ErrMsg
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return e.ErrCode, &e.ErrMsg
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return s.ErrCode, &s.ErrMsg
}

func User() {

}
