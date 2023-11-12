// Code generated by Kitex v0.7.3. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	relation "rpc/kitex_gen/relation"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FollowAction":        kitex.NewMethodInfo(followActionHandler, newRelationServiceFollowActionArgs, newRelationServiceFollowActionResult, false),
		"FollowList":          kitex.NewMethodInfo(followListHandler, newRelationServiceFollowListArgs, newRelationServiceFollowListResult, false),
		"FansList":            kitex.NewMethodInfo(fansListHandler, newRelationServiceFansListArgs, newRelationServiceFansListResult, false),
		"FriendList":          kitex.NewMethodInfo(friendListHandler, newRelationServiceFriendListArgs, newRelationServiceFriendListResult, false),
		"IsFollowTarget":      kitex.NewMethodInfo(isFollowTargetHandler, newRelationServiceIsFollowTargetArgs, newRelationServiceIsFollowTargetResult, false),
		"IsFollowManyTargets": kitex.NewMethodInfo(isFollowManyTargetsHandler, newRelationServiceIsFollowManyTargetsArgs, newRelationServiceIsFollowManyTargetsResult, false),
		"IsFriend":            kitex.NewMethodInfo(isFriendHandler, newRelationServiceIsFriendArgs, newRelationServiceIsFriendResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "relation",
		"ServiceFilePath": `idl\relation.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func followActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFollowActionArgs)
	realResult := result.(*relation.RelationServiceFollowActionResult)
	success, err := handler.(relation.RelationService).FollowAction(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFollowActionArgs() interface{} {
	return relation.NewRelationServiceFollowActionArgs()
}

func newRelationServiceFollowActionResult() interface{} {
	return relation.NewRelationServiceFollowActionResult()
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFollowListArgs)
	realResult := result.(*relation.RelationServiceFollowListResult)
	success, err := handler.(relation.RelationService).FollowList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFollowListArgs() interface{} {
	return relation.NewRelationServiceFollowListArgs()
}

func newRelationServiceFollowListResult() interface{} {
	return relation.NewRelationServiceFollowListResult()
}

func fansListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFansListArgs)
	realResult := result.(*relation.RelationServiceFansListResult)
	success, err := handler.(relation.RelationService).FansList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFansListArgs() interface{} {
	return relation.NewRelationServiceFansListArgs()
}

func newRelationServiceFansListResult() interface{} {
	return relation.NewRelationServiceFansListResult()
}

func friendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFriendListArgs)
	realResult := result.(*relation.RelationServiceFriendListResult)
	success, err := handler.(relation.RelationService).FriendList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFriendListArgs() interface{} {
	return relation.NewRelationServiceFriendListArgs()
}

func newRelationServiceFriendListResult() interface{} {
	return relation.NewRelationServiceFriendListResult()
}

func isFollowTargetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceIsFollowTargetArgs)
	realResult := result.(*relation.RelationServiceIsFollowTargetResult)
	success, err := handler.(relation.RelationService).IsFollowTarget(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceIsFollowTargetArgs() interface{} {
	return relation.NewRelationServiceIsFollowTargetArgs()
}

func newRelationServiceIsFollowTargetResult() interface{} {
	return relation.NewRelationServiceIsFollowTargetResult()
}

func isFollowManyTargetsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceIsFollowManyTargetsArgs)
	realResult := result.(*relation.RelationServiceIsFollowManyTargetsResult)
	success, err := handler.(relation.RelationService).IsFollowManyTargets(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceIsFollowManyTargetsArgs() interface{} {
	return relation.NewRelationServiceIsFollowManyTargetsArgs()
}

func newRelationServiceIsFollowManyTargetsResult() interface{} {
	return relation.NewRelationServiceIsFollowManyTargetsResult()
}

func isFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceIsFriendArgs)
	realResult := result.(*relation.RelationServiceIsFriendResult)
	success, err := handler.(relation.RelationService).IsFriend(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceIsFriendArgs() interface{} {
	return relation.NewRelationServiceIsFriendArgs()
}

func newRelationServiceIsFriendResult() interface{} {
	return relation.NewRelationServiceIsFriendResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FollowAction(ctx context.Context, request *relation.FollowActionRequest) (r *relation.FollowActionResponse, err error) {
	var _args relation.RelationServiceFollowActionArgs
	_args.Request = request
	var _result relation.RelationServiceFollowActionResult
	if err = p.c.Call(ctx, "FollowAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, request *relation.FollowListRequest) (r *relation.FollowListResponse, err error) {
	var _args relation.RelationServiceFollowListArgs
	_args.Request = request
	var _result relation.RelationServiceFollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FansList(ctx context.Context, request *relation.FansListRequest) (r *relation.FansListResponse, err error) {
	var _args relation.RelationServiceFansListArgs
	_args.Request = request
	var _result relation.RelationServiceFansListResult
	if err = p.c.Call(ctx, "FansList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendList(ctx context.Context, request *relation.FriendListRequest) (r *relation.FriendListResponse, err error) {
	var _args relation.RelationServiceFriendListArgs
	_args.Request = request
	var _result relation.RelationServiceFriendListResult
	if err = p.c.Call(ctx, "FriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFollowTarget(ctx context.Context, request *relation.IsFollowTargetRequest) (r *relation.IsFollowTargetResponse, err error) {
	var _args relation.RelationServiceIsFollowTargetArgs
	_args.Request = request
	var _result relation.RelationServiceIsFollowTargetResult
	if err = p.c.Call(ctx, "IsFollowTarget", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFollowManyTargets(ctx context.Context, request *relation.IsFollowManyTargetsRequest) (r *relation.IsFollowManyTargetsResponse, err error) {
	var _args relation.RelationServiceIsFollowManyTargetsArgs
	_args.Request = request
	var _result relation.RelationServiceIsFollowManyTargetsResult
	if err = p.c.Call(ctx, "IsFollowManyTargets", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFriend(ctx context.Context, request *relation.IsFriendRequest) (r *relation.IsFriendResponse, err error) {
	var _args relation.RelationServiceIsFriendArgs
	_args.Request = request
	var _result relation.RelationServiceIsFriendResult
	if err = p.c.Call(ctx, "IsFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}