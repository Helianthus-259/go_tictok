package service

import (
	"context"
	"gorm.io/gorm"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
	"user/pkg/constants"
	"user/pkg/errno"
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
			return errno.UpdateUserCountFailedErr
		}
		err = redis.AddTotalFavoriteCount(s.ctx, authorCntKey)
		if logger.CheckError(err, "Redis AddTotalFavoriteCount err") {
			return errno.UpdateUserCountFailedErr
		}
		// Add FavoriteTotalCount in Mysql
		err = db.MyTransaction(db.DB, func(tx *gorm.DB) (err error) {
			err = db.AddFavoriteCountByUserId(s.ctx, tx, request.GetUserId())
			if logger.CheckError(err, "Mysql AddFavoriteCount err") {
				return errno.UpdateUserCountFailedErr
			}
			err = db.AddTotalFavoriteCountByUserId(s.ctx, tx, request.GetAuthorId())
			if logger.CheckError(err, "Mysql AddTotalFavoriteCount err") {
				return errno.UpdateUserCountFailedErr
			}
			return nil
		})
		//err = db.AddFavoriteCountByUserId(s.ctx, request.GetUserId())
		//if logger.CheckError(err, "Mysql AddFavoriteCount err") {
		//	return errno.UpdateUserCountFailedErr
		//}
		//err = db.AddTotalFavoriteCountByUserId(s.ctx, request.GetAuthorId())
		//if logger.CheckError(err, "Mysql AddTotalFavoriteCount err") {
		//	return errno.UpdateUserCountFailedErr
		//}
	case constants.CANCELFAVORITE:
		// Sub FavoriteCount in Redis
		err = redis.SubFavoriteCount(s.ctx, userCntKey)
		if logger.CheckError(err, "Redis SubFavoriteCount err") {
			return errno.UpdateUserCountFailedErr
		}
		err = redis.SubTotalFavoriteCount(s.ctx, authorCntKey)
		if logger.CheckError(err, "Redis SubTotalFavoriteCount err") {
			return errno.UpdateUserCountFailedErr
		}
		// Sub FavoriteTotalCount in Mysql
		err = db.MyTransaction(db.DB, func(tx *gorm.DB) (err error) {
			err = db.SubFavoriteCountByUserId(s.ctx, tx, request.GetUserId())
			if logger.CheckError(err, "Mysql AddFavoriteCount err") {
				return errno.UpdateUserCountFailedErr
			}
			err = db.SubTotalFavoriteCountByUserId(s.ctx, tx, request.GetAuthorId())
			if logger.CheckError(err, "Mysql AddTotalFavoriteCount err") {
				return errno.UpdateUserCountFailedErr
			}
			return nil
		})
		//err = db.SubFavoriteCountByUserId(s.ctx, request.GetUserId())
		//if logger.CheckError(err, "Mysql AddFavoriteCount err") {
		//	return errno.UpdateUserCountFailedErr
		//}
		//err = db.SubTotalFavoriteCountByUserId(s.ctx, request.GetAuthorId())
		//if logger.CheckError(err, "Mysql AddTotalFavoriteCount err") {
		//	return errno.UpdateUserCountFailedErr
		//}
	default:
		return errno.ServiceErr
	}
	return nil
}
