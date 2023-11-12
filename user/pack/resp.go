package pack

import (
	"errors"
	demouser "rpc/kitex_gen/user"
	"user/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *demouser.BaseResp {
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

func baseResp(err errno.ErrNo) *demouser.BaseResp {
	return &demouser.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
