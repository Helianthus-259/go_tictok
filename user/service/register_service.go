package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	demouser "rpc/kitex_gen/user"
	"user/dal/model"
	db "user/dal/mysql"
	"user/dal/redis"
	"user/pkg/errno"
	"user/pkg/logger"
)

type RegisterService struct {
	ctx context.Context
}

// NewRegisterService New RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{
		ctx: ctx,
	}
}

// Register Register User
func (s *RegisterService) Register(req *demouser.RegisterRequest) (uid int64, err error) {

	users, err := db.CheckSameAccount(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	UserInfo := &model.User{
		Username: req.Username,
		Password: password,
	}

	userCount := &model.UserCount{
		FollowCount:    0,
		FollowerCount:  0,
		TotalFavorited: 0,
		WorkCount:      0,
		FavoriteCount:  0,
	}
	uid, err = db.TransactionAdd(s.ctx, UserInfo, userCount)
	if logger.CheckError(err, "Mysql Add User err") {
		return 0, errno.ServiceErr
	}
	UserInfo.ID = uint(uid)
	// 将 userInfo 存储 redis
	userKey := model.CreateUserKey(UserInfo.ID)
	if err = redis.HSetUserInfo(s.ctx, userKey, model.CreateMapUserInfo(UserInfo)); err != nil {
		err = errors.New("Redis HSetUserInfo err:" + err.Error())
	}
	// 将 userCount 存储 redis
	if err = redis.HSetUserCountInfo(s.ctx, userKey, model.CreateMapUserCount(userCount)); err != nil {
		err = errors.New("Redis HSetUserInfo err:" + err.Error())
	}
	return
}
