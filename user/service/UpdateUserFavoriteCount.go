package service

import (
	"context"
	"errors"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
	"user/pkg/constants"
	"user/pkg/logger"
)

type UpdateUserFavoriteCountService struct {
	ctx context.Context
}

// NewUpdateUserFavoriteCountService  new UpdateUserFavoriteCountService
func NewUpdateUserFavoriteCountService(ctx context.Context) *UpdateUserFavoriteCountService {
	return &UpdateUserFavoriteCountService{
		ctx: ctx,
	}
}
func (s *UpdateUserFavoriteCountService) UpdateUserFavoriteCount(request *demouser.UpdateUserFavoriteCountRequest) (err error) {
	userCntKey := model.CreateUserCountKey(uint(request.GetUserId()))
	authorCntKey := model.CreateUserCountKey(uint(request.GetAuthorId()))
	// Update UserCount From Redis
	switch request.GetActionType() {
	// Add FavoriteCount
	case constants.FAVORITE:
		// Add FavoriteCount in Redis
		err = redis.AddFavoriteCount(s.ctx, userCntKey)
		if logger.CheckError(err, "Redis AddFavoriteCount err") {
			return
		}
		err = redis.AddTotalFavoriteCount(s.ctx, authorCntKey)
		if logger.CheckError(err, "Redis AddTotalFavoriteCount err") {
			return
		}
		// Add FavoriteTotalCount in Mysql
		err = db.AddFavoriteCountByUserId(s.ctx, request.GetUserId())
		if logger.CheckError(err, "Mysql AddFavoriteCount err") {
			return
		}
		err = db.AddTotalFavoriteCountByUserId(s.ctx, request.GetAuthorId())
		if logger.CheckError(err, "Mysql AddTotalFavoriteCount err") {
			return
		}
	case constants.CANCELFAVORITE:
		// Sub FavoriteCount in Redis
		err = redis.SubFavoriteCount(s.ctx, userCntKey)
		if logger.CheckError(err, "Redis SubFavoriteCount err") {
			return
		}
		err = redis.SubTotalFavoriteCount(s.ctx, authorCntKey)
		if logger.CheckError(err, "Redis SubTotalFavoriteCount err") {
			return
		}
		// Sub FavoriteTotalCount in Mysql
		err = db.SubFavoriteCountByUserId(s.ctx, request.GetUserId())
		if logger.CheckError(err, "Mysql AddFavoriteCount err") {
			return
		}
		err = db.SubTotalFavoriteCountByUserId(s.ctx, request.GetAuthorId())
		if logger.CheckError(err, "Mysql AddTotalFavoriteCount err") {
			return
		}
	default:
		err = errors.New("no Find Favorite ActionType")
		return
	}
	return
}
