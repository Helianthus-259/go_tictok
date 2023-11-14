package pack

import (
	"errors"
	"favorite/pkg/errno"
	demofavorite "rpc/kitex_gen/favorite"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *demofavorite.BaseResp {
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

func baseResp(err errno.ErrNo) *demofavorite.BaseResp {
	return &demofavorite.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
