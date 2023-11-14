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

type GetUserListService struct {
	ctx context.Context
}

// NewGetUserListService  new GetUserListService
func NewGetUserListService(ctx context.Context) *GetUserListService {
	return &GetUserListService{
		ctx: ctx,
	}
}

func (s *GetUserListService) GetUserList(request *demouser.GetUserListRequest) (userList []*demouser.User, err error) {
	TargetId := request.GetTargetId()
	for i, targetId := range TargetId {
		userKey := model.CreateUserKey(uint(targetId))
		// Get UserInfo From Redis
		userInfo, err := redis.HGetUserInfo(s.ctx, userKey)
		if err != nil {
			err = errors.New("Redis HSetUserInfo err:" + err.Error())
			return nil, errno.RedisErr
		}
		if userInfo == nil {
			userInfo, err = db.GetUserInfoById(s.ctx, db.DB, targetId)
			if err != nil {
				err = errors.New("Mysql GetUserInfoById:" + err.Error())
				return nil, errno.MysqlErr
			}
			// Set UserInfo into Redis
			if err = redis.HSetUserInfo(s.ctx, userKey, model.CreateMapUserInfo(userInfo)); err != nil {
				err = errors.New("Redis HSetUserInfo err:" + err.Error())
				return nil, errno.RedisErr
			}
		}
		// Get UserCount Info From Redis
		userCountInfo, err := redis.HGetUserCountInfo(s.ctx, userKey)
		if err != nil {
			err = errors.New("Redis HGetUserCountInfo err:" + err.Error())
			return nil, errno.RedisErr
		}
		if userCountInfo == nil {
			// Get UserCount Info from mysql
			userCountInfo, err = db.GetUserCountInfoByUserId(s.ctx, db.DB, targetId)
			if logger.CheckError(err, "Mysql HGetUserCountInfo err") {
				return nil, errno.MysqlErr
			}
			// Set UserCountInfo into Redis
			err = redis.HSetUserCountInfo(s.ctx, userKey, model.CreateMapUserCount(userCountInfo))
			if logger.CheckError(err, "Redis HSetUserCountInfo err") {
				return nil, errno.RedisErr
			}
		}
		userList[i] = &demouser.User{
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
	}
	return userList, nil
}
