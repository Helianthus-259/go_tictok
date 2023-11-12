package service

import (
	"context"
	"errors"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
	"user/pkg/errno"
	"user/pkg/logger"
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

	// Get UserInfo from Redis
	userKey := model.CreateUserKey(uint(req.UserId))
	userInfo, err := redis.HGetUserInfo(s.ctx, userKey)
	if err != nil {
		err = errors.New("Redis HGetUserInfo err:" + err.Error())
		return nil, err
	}
	// Get UserInfo from mysql
	if userInfo == nil {
		userInfo, err = db.GetUserInfoById(s.ctx, req.GetUserId())
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
		return nil, err
	}
	// Get UserCountInfo from Mysql
	if userCountInfo == nil {
		userCountInfo, err = db.GetUserCountInfoByUserId(s.ctx, req.GetUserId())
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
		return nil, err
	}
	// Set UserCountInfo into Redis
	err = redis.HSetUserCountInfo(s.ctx, userKey, model.CreateMapUserCount(userCountInfo))
	if logger.CheckError(err, "Redis HSetUserCountInfo err") {
		return nil, err
	}
	// pack
	User := &demouser.User{
		Id:              int64(userInfo.ID),
		Name:            userInfo.Username,
		FollowCount:     &userCountInfo.FollowCount,
		FollowerCount:   &userCountInfo.FollowerCount,
		IsFollow:        false,
		Avatar:          &userInfo.Avatar,
		BackgroundImage: &userInfo.BackgroundImage,
		Signature:       &userInfo.Signature,
		TotalFavorited:  &userCountInfo.TotalFavorited,
		WorkCount:       &userCountInfo.WorkCount,
		FavoriteCount:   &userCountInfo.FavoriteCount,
	}
	return User, nil
}
