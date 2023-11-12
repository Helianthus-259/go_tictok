package service

import (
	"context"
	"errors"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
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
			return nil, err
		}
		if userInfo == nil {
			userInfo, err = db.GetUserInfoById(s.ctx, targetId)
			if err != nil {
				err = errors.New("Mysql GetUserInfoById:" + err.Error())
				return nil, err
			}
			// Set UserInfo into Redis
			if err = redis.HSetUserInfo(s.ctx, userKey, model.CreateMapUserInfo(userInfo)); err != nil {
				err = errors.New("Redis HSetUserInfo err:" + err.Error())
				return nil, err
			}
		}
		// Get UserCount Info From Redis
		userCountInfo, err := redis.HGetUserCountInfo(s.ctx, userKey)
		if err != nil {
			err = errors.New("Redis HGetUserCountInfo err:" + err.Error())
			return nil, err
		}
		if userCountInfo == nil {
			// Get UserCount Info from mysql
			userCountInfo, err = db.GetUserCountInfoByUserId(s.ctx, targetId)
			if logger.CheckError(err, "Redis HGetUserCountInfo err") {
				return nil, err
			}
			// Set UserCountInfo into Redis
			err = redis.HSetUserCountInfo(s.ctx, userKey, model.CreateMapUserCount(userCountInfo))
			if logger.CheckError(err, "Redis HSetUserCountInfo err") {
				return nil, err
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
	return
}
