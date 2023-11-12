package service

import (
	"context"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
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
		return
	}
	err = redis.AddFollowerCount(s.ctx, targetCountKey)
	if logger.CheckError(err, "Redis AddFollowerCount err") {
		return
	}
	// Update UserCount From Mysql
	err = db.AddFollowCountByUserId(s.ctx, request.GetUserId())
	if logger.CheckError(err, "Mysql AddFollowCount err") {
		return
	}
	err = db.AddFollowerCountByUserId(s.ctx, request.GetTargetId())
	if logger.CheckError(err, "Mysql AddFollowerCount err") {
		return
	}
	return
}
