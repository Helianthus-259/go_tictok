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

type SubFollowCountService struct {
	ctx context.Context
}

// NewSubFollowCountService  new AddFollowCountService
func NewSubFollowCountService(ctx context.Context) *SubFollowCountService {
	return &SubFollowCountService{
		ctx: ctx,
	}
}

func (s *SubFollowCountService) SubFollowCount(request *demouser.SubFollowCountRequest) (err error) {
	userCountKey := model.CreateUserCountKey(uint(request.GetUserId()))
	targetCountKey := model.CreateUserCountKey(uint(request.GetTargetId()))
	// Update UserCount from Redis
	err = redis.SubFollowCount(s.ctx, userCountKey)
	if logger.CheckError(err, "Redis SubFollowCount err") {
		return errno.UpdateUserCountFailedErr
	}
	err = redis.SubFollowerCount(s.ctx, targetCountKey)
	if logger.CheckError(err, "Redis SubFollowerCount err") {
		return errno.UpdateUserCountFailedErr
	}
	// Update UserCount From Mysql
	err = db.MyTransaction(db.DB, func(tx *gorm.DB) (err error) {
		err = db.SubFollowCountByUserId(s.ctx, tx, request.GetUserId())
		if logger.CheckError(err, "Mysql SubFollowCount err") {
			return errno.UpdateUserCountFailedErr
		}
		err = db.SubFollowerCountByUserId(s.ctx, tx, request.GetTargetId())
		if logger.CheckError(err, "Mysql SubFollowerCount err") {
			return errno.UpdateUserCountFailedErr
		}
		return nil
	})
	//err = db.SubFollowCountByUserId(s.ctx, request.GetUserId())
	//if logger.CheckError(err, "Mysql SubFollowCount err") {
	//	return errno.UpdateUserCountFailedErr
	//}
	//err = db.SubFollowerCountByUserId(s.ctx, request.GetTargetId())
	//if logger.CheckError(err, "Mysql SubFollowerCount err") {
	//	return errno.UpdateUserCountFailedErr
	//}
	return nil
}
