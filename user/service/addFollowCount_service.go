package service

import (
	"context"
	"gorm.io/gorm"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
	"user/pkg/errno"
	"user/pkg/logger"
)

type AddFollowCountService struct {
	ctx context.Context
}

// NewAddFollowCountService  new AddFollowCountService
func NewAddFollowCountService(ctx context.Context) *AddFollowCountService {
	return &AddFollowCountService{
		ctx: ctx,
	}
}

func (s *AddFollowCountService) AddFollowCount(request *demouser.AddFollowCountRequest) (err error) {
	userCountKey := model.CreateUserCountKey(uint(request.GetUserId()))
	targetCountKey := model.CreateUserCountKey(uint(request.GetTargetId()))
	// Update UserCount from Redis
	err = redis.AddFollowCount(s.ctx, userCountKey)
	if logger.CheckError(err, "Redis AddFollowCount err") {
		return errno.UpdateUserCountFailedErr
	}
	err = redis.AddFollowerCount(s.ctx, targetCountKey)
	if logger.CheckError(err, "Redis AddFollowerCount err") {
		return errno.UpdateUserCountFailedErr
	}
	// Update UserCount From Mysql
	err = db.MyTransaction(db.DB, func(tx *gorm.DB) (err error) {
		err = db.AddFollowCountByUserId(s.ctx, tx, request.GetUserId())
		if logger.CheckError(err, "Mysql AddFollowCount err") {
			return errno.UpdateUserCountFailedErr
		}
		err = db.AddFollowerCountByUserId(s.ctx, tx, request.GetTargetId())
		if logger.CheckError(err, "Mysql AddFollowerCount err") {
			return errno.UpdateUserCountFailedErr
		}
		return nil
	})
	//err = db.AddFollowCountByUserId(s.ctx, db.DB, request.GetUserId())
	//if logger.CheckError(err, "Mysql AddFollowCount err") {
	//	return errno.UpdateUserCountFailedErr
	//}
	//err = db.AddFollowerCountByUserId(s.ctx, request.GetTargetId())
	//if logger.CheckError(err, "Mysql AddFollowerCount err") {
	//	return errno.UpdateUserCountFailedErr
	//}
	return nil
}
