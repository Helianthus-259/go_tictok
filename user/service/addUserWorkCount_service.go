package service

import (
	"context"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
	"user/pkg/logger"
)

type AddUserWorkCountService struct {
	ctx context.Context
}

// NewAddUserWorkCountService  new AddUserWorkCountService
func NewAddUserWorkCountService(ctx context.Context) *AddUserWorkCountService {
	return &AddUserWorkCountService{
		ctx: ctx,
	}
}

func (s *AddUserWorkCountService) AddUserWorkCount(request *demouser.AddUserWorkCountRequest) (err error) {
	userCountKey := model.CreateUserCountKey(uint(request.GetUserId()))
	// Update UserCount From Redis
	err = redis.AddWorkCount(s.ctx, userCountKey)
	if logger.CheckError(err, "Redis Add WorkCount err") {
		return
	}
	// Update UserCount From Mysql
	err = db.AddWorkCountByUserId(s.ctx, request.GetUserId())
	if logger.CheckError(err, "Mysql Add WorkCount err") {
		return
	}
	return
}
