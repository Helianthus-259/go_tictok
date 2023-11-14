package pack

import (
	"errors"
	"message/pkg/errno"
	demomessage "rpc/kitex_gen/message"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *demomessage.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *demomessage.BaseResp {
	return &demomessage.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
