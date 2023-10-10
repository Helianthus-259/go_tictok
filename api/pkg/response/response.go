package response

import (
	"common-components/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	StatusCode int64       `json:"code"`
	StatusMsg  string      `json:"message"`
	Data       interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		Data:       data,
	})
}
