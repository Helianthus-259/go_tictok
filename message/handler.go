package main

import (
	"context"
	rpcMessage "rpc/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

func (m MessageServiceImpl) MessageList(ctx context.Context, request *rpcMessage.MessageListRequest) (r *rpcMessage.MessageListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageServiceImpl) ChatAction(ctx context.Context, request *rpcMessage.ChatActionRequest) (r *rpcMessage.ChatActionResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageServiceImpl) GetFriendLatestMessage(ctx context.Context, request *rpcMessage.GetFriendLatestMessageRequest) (r *rpcMessage.GetFriendLatestMessageResponse, err error) {
	//TODO implement me
	panic("implement me")
}
