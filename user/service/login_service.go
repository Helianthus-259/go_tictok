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
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService  new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{
		ctx: ctx,
	}
}

// Login check user info
func (s *LoginService) Login(req *demouser.LoginRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username

	users, err := db.CheckSameAccount(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}

	//fmt.Println(u)
	// 将 userInfo 存储 redis
	if err = redis.HSetUserInfo(s.ctx, model.CreateUserKey(u.ID), model.CreateMapUserInfo(&u)); err != nil {
		err = errors.New("Redis HSetUserInfo err:" + err.Error())
		return 0, err
	}
	// 将 userCount 存储 redis
	if err = redis.HSetUserCountInfo(s.ctx, model.CreateUserKey(u.ID), model.CreateMapUserCount(&model.UserCount{
		UserId:         u.ID,
		FollowCount:    0,
		FollowerCount:  0,
		TotalFavorited: 0,
		WorkCount:      0,
		FavoriteCount:  0,
	})); err != nil {
		err = errors.New("Redis HSetUserInfo err:" + err.Error())
		return 0, err
	}

	return int64(u.ID), nil
}
