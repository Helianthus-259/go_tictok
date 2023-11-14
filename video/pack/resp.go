package pack

import (
	"errors"
	demovideo "rpc/kitex_gen/video"
	"video/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *demovideo.BaseResp {
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

func baseResp(err errno.ErrNo) *demovideo.BaseResp {
	return &demovideo.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
