// Code generated by Kitex v0.7.3. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "rpc/kitex_gen/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FollowAction(ctx context.Context, request *relation.FollowActionRequest, callOptions ...callopt.Option) (r *relation.FollowActionResponse, err error)
	FollowList(ctx context.Context, request *relation.FollowListRequest, callOptions ...callopt.Option) (r *relation.FollowListResponse, err error)
	FansList(ctx context.Context, request *relation.FansListRequest, callOptions ...callopt.Option) (r *relation.FansListResponse, err error)
	FriendList(ctx context.Context, request *relation.FriendListRequest, callOptions ...callopt.Option) (r *relation.FriendListResponse, err error)
	IsFollowTarget(ctx context.Context, request *relation.IsFollowTargetRequest, callOptions ...callopt.Option) (r *relation.IsFollowTargetResponse, err error)
	IsFollowManyTargets(ctx context.Context, request *relation.IsFollowManyTargetsRequest, callOptions ...callopt.Option) (r *relation.IsFollowManyTargetsResponse, err error)
	IsFriend(ctx context.Context, request *relation.IsFriendRequest, callOptions ...callopt.Option) (r *relation.IsFriendResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) FollowAction(ctx context.Context, request *relation.FollowActionRequest, callOptions ...callopt.Option) (r *relation.FollowActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowAction(ctx, request)
}

func (p *kRelationServiceClient) FollowList(ctx context.Context, request *relation.FollowListRequest, callOptions ...callopt.Option) (r *relation.FollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowList(ctx, request)
}

func (p *kRelationServiceClient) FansList(ctx context.Context, request *relation.FansListRequest, callOptions ...callopt.Option) (r *relation.FansListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FansList(ctx, request)
}

func (p *kRelationServiceClient) FriendList(ctx context.Context, request *relation.FriendListRequest, callOptions ...callopt.Option) (r *relation.FriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendList(ctx, request)
}

func (p *kRelationServiceClient) IsFollowTarget(ctx context.Context, request *relation.IsFollowTargetRequest, callOptions ...callopt.Option) (r *relation.IsFollowTargetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFollowTarget(ctx, request)
}

func (p *kRelationServiceClient) IsFollowManyTargets(ctx context.Context, request *relation.IsFollowManyTargetsRequest, callOptions ...callopt.Option) (r *relation.IsFollowManyTargetsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFollowManyTargets(ctx, request)
}

func (p *kRelationServiceClient) IsFriend(ctx context.Context, request *relation.IsFriendRequest, callOptions ...callopt.Option) (r *relation.IsFriendResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFriend(ctx, request)
}