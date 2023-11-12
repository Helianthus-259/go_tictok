// Code generated by hertz generator.

package video

import (
	"context"

	video "api/biz/model/video"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// VideoFeed .
// @router /douyin/feed/ [GET]
func VideoFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.VideoFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.VideoFeedResponse)

	c.JSON(consts.StatusOK, resp)
}

// VideoPublish .
// @router /douyin/publish/action/ [POST]
func VideoPublish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.VideoPublishRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.VideoPublishResponse)

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.PublishListResponse)

	c.JSON(consts.StatusOK, resp)
}
