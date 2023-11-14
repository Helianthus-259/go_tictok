package main

import (
	"context"
	rpcComment "rpc/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

func (c CommentServiceImpl) CommentAction(ctx context.Context, request *rpcComment.CommentActionRequest) (r *rpcComment.CommentActionResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentServiceImpl) CommentList(ctx context.Context, request *rpcComment.CommentListRequest) (r *rpcComment.CommentListResponse, err error) {
	//TODO implement me
	panic("implement me")
}
