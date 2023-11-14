package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
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

	users, err := db.CheckSameAccount(s.ctx, db.DB, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, errno.ServiceErr
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	UserInfo := &model.User{
		ID:       0,
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
	//uid, err = db.TransactionAdd(s.ctx, UserInfo, userCount)
	result, err := db.MyTransactionWithResult(db.DB, func(tx *gorm.DB) (interface{}, error) {
		uid, err = db.AddUser(s.ctx, tx, UserInfo)
		if logger.CheckError(err, "Mysql Add User err") {
			return nil, errno.MysqlErr
		}
		err = db.AddUserCount(s.ctx, tx, userCount)
		if logger.CheckError(err, "Mysql Add UserCount err") {
			return nil, errno.MysqlErr
		}
		return uid, nil
	})
	if logger.CheckError(err, "Mysql Add User err") {
		return 0, errno.MysqlErr
	}
	uid = result.(int64)
	UserInfo.ID = uint(uid)
	// 将 userInfo 存储 redis
	userKey := model.CreateUserKey(UserInfo.ID)
	if err = redis.HSetUserInfo(s.ctx, userKey, model.CreateMapUserInfo(UserInfo)); logger.CheckError(err, "Redis HSetUserInfo err") {
		return 0, errno.RedisErr
	}
	// 将 userCount 存储 redis
	if err = redis.HSetUserCountInfo(s.ctx, userKey, model.CreateMapUserCount(userCount)); logger.CheckError(err, "Redis HSetUserCountInfo err") {
		return 0, errno.RedisErr
	}
	return uid, nil
}
