// Code generated by hertz generator.

package favorite

import (
	"context"

	favorite "api/biz/model/favorite"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteList .
// @router /douyin/favorite/action/ [POST]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(favorite.FavoriteListResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .
// @router /douyin/favorite/list/ [GET]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(favorite.FavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}