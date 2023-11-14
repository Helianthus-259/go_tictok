package pack

import (
	"errors"
	"relation/pkg/errno"
	demorelation "rpc/kitex_gen/relation"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *demorelation.BaseResp {
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

func baseResp(err errno.ErrNo) *demorelation.BaseResp {
	return &demorelation.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
