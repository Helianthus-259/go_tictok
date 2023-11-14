package main

import (
	"context"
	rpcRelation "rpc/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

func (r2 RelationServiceImpl) FollowAction(ctx context.Context, request *rpcRelation.FollowActionRequest) (r *rpcRelation.FollowActionResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (r2 RelationServiceImpl) FollowList(ctx context.Context, request *rpcRelation.FollowListRequest) (r *rpcRelation.FollowListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (r2 RelationServiceImpl) FansList(ctx context.Context, request *rpcRelation.FansListRequest) (r *rpcRelation.FansListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (r2 RelationServiceImpl) FriendList(ctx context.Context, request *rpcRelation.FriendListRequest) (r *rpcRelation.FriendListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (r2 RelationServiceImpl) IsFollowTarget(ctx context.Context, request *rpcRelation.IsFollowTargetRequest) (r *rpcRelation.IsFollowTargetResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (r2 RelationServiceImpl) IsFollowManyTargets(ctx context.Context, request *rpcRelation.IsFollowManyTargetsRequest) (r *rpcRelation.IsFollowManyTargetsResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (r2 RelationServiceImpl) IsFriend(ctx context.Context, request *rpcRelation.IsFriendRequest) (r *rpcRelation.IsFriendResponse, err error) {
	//TODO implement me
	panic("implement me")
}
