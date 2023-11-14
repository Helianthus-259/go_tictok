package service

import (
	"context"
	"errors"
	rpcRelation "rpc/kitex_gen/relation"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
	"user/pkg/errno"
	"user/pkg/logger"
	rpcClient "user/rpc"
)

type UserIndexService struct {
	ctx context.Context
}

// NewUserIndexService  new UserIndexService
func NewUserIndexService(ctx context.Context) *UserIndexService {
	return &UserIndexService{
		ctx: ctx,
	}
}

// UserIndex  Get UserIndex by ID
func (s *UserIndexService) UserIndex(req *demouser.UserIndexRequest) (*demouser.User, error) {

	// Call RPC Service Check If Follow
	var isFollow bool
	if req.UserId == req.MyUserId {
		isFollow = false
	} else {
		resp, err := rpcClient.IsFollowTarget(s.ctx, &rpcRelation.IsFollowTargetRequest{UserId: req.MyUserId, TargetId: req.UserId})
		if logger.CheckError(err, "rpcRelation.IsFollowTargetRequest Failed") {
			err = errno.ConvertErr(err)
			return nil, err
		}
		isFollow = resp
	}

	// Get UserInfo from Redis
	userKey := model.CreateUserKey(uint(req.UserId))
	userInfo, err := redis.HGetUserInfo(s.ctx, userKey)
	if logger.CheckError(err, "Redis HGetUserInfo err") {
		return nil, errno.RedisErr
	}
	// Get UserInfo from mysql
	if userInfo == nil {
		userInfo, err = db.GetUserInfoById(s.ctx, db.DB, req.GetUserId())
		if logger.CheckError(err, "Mysql GetUserInfoById err") {
			return nil, errno.UserIsNotExistErr
		}
		if userInfo == nil {
			return nil, errno.UserIsNotExistErr
		}
	}
	//Get UserCountInfo From Redis
	userCountInfo, err := redis.HGetUserCountInfo(s.ctx, userKey)
	if err != nil {
		err = errors.New("Redis HGetUserCountInfo err:" + err.Error())
		return nil, errno.RedisErr
	}
	// Get UserCountInfo from Mysql
	if userCountInfo == nil {
		userCountInfo, err = db.GetUserCountInfoByUserId(s.ctx, db.DB, req.GetUserId())
		if logger.CheckError(err, "Mysql GetUserCountInfoByUserId err") {
			return nil, errno.UserIsNotExistErr
		}
		if userCountInfo == nil {
			return nil, errno.UserIsNotExistErr
		}
	}
	// set UserInfo into Redis
	err = redis.HSetUserInfo(s.ctx, userKey, model.CreateMapUserInfo(userInfo))
	if logger.CheckError(err, "Redis HSetUserInfo err") {
		return nil, errno.RedisErr
	}
	// Set UserCountInfo into Redis
	err = redis.HSetUserCountInfo(s.ctx, userKey, model.CreateMapUserCount(userCountInfo))
	if logger.CheckError(err, "Redis HSetUserCountInfo err") {
		return nil, errno.RedisErr
	}
	// pack
	User := &demouser.User{
		Id:              int64(userInfo.ID),
		Name:            userInfo.Username,
		FollowCount:     &userCountInfo.FollowCount,
		FollowerCount:   &userCountInfo.FollowerCount,
		IsFollow:        isFollow,
		Avatar:          &userInfo.Avatar,
		BackgroundImage: &userInfo.BackgroundImage,
		Signature:       &userInfo.Signature,
		TotalFavorited:  &userCountInfo.TotalFavorited,
		WorkCount:       &userCountInfo.WorkCount,
		FavoriteCount:   &userCountInfo.FavoriteCount,
	}
	return User, nil
}
