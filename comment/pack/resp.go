package pack

import (
	"comment/pkg/errno"
	"errors"
	democomment "rpc/kitex_gen/comment"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *democomment.BaseResp {
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

func baseResp(err errno.ErrNo) *democomment.BaseResp {
	return &democomment.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
