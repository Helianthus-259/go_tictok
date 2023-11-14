package response

import (
	"api/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	StatusCode int64       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Data       interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err errno.ErrNo, data interface{}) {
	c.JSON(consts.StatusOK, Response{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
		Data:       data,
	})
}
