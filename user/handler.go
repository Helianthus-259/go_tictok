package main

import (
	"context"
	"fmt"
	demouser "rpc/kitex_gen/user"
	"user/pack"
	"user/pkg/errno"
	"user/pkg/logger"
	"user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

func (u UserServiceImpl) AddFollowCount(ctx context.Context, request *demouser.AddFollowCountRequest) (r *demouser.AddFollowCountResponse, err error) {
	r = new(demouser.AddFollowCountResponse)
	// call service function
	err = service.NewAddFollowCountService(ctx).AddFollowCount(request)
	if logger.CheckError(err, "Service AddFollowCount err") {
		r.BaseResp = pack.BuildBaseResp(err)
		return r, nil
	}
	r.BaseResp = pack.BuildBaseResp(errno.Success)
	return r, nil
}

func (u UserServiceImpl) SubFollowCount(ctx context.Context, request *demouser.SubFollowCountRequest) (r *demouser.SubFollowCountResponse, err error) {
	r = new(demouser.SubFollowCountResponse)
	// call service function
	err = service.NewSubFollowCountService(ctx).SubFollowCount(request)
	if logger.CheckError(err, "Service SubFollowCount err") {
		r.BaseResp = pack.BuildBaseResp(err)
		return r, nil
	}
	r.BaseResp = pack.BuildBaseResp(errno.Success)
	return r, nil
}

func (u UserServiceImpl) GetUserList(ctx context.Context, request *demouser.GetUserListRequest) (r *demouser.GetUserListResponse, err error) {
	r = new(demouser.GetUserListResponse)
	// call service function
	userList, err := service.NewGetUserListService(ctx).GetUserList(request)
	if logger.CheckError(err, "Service GetUserList err") {
		r.BaseResp = pack.BuildBaseResp(err)
		return r, nil
	}
	// pack response
	r.BaseResp = pack.BuildBaseResp(errno.Success)
	r.UserList = userList

	return r, nil
}

func (u UserServiceImpl) AddUserWorkCount(ctx context.Context, request *demouser.AddUserWorkCountRequest) (r *demouser.AddUserWorkCountResponse, err error) {
	r = new(demouser.AddUserWorkCountResponse)
	// call service function
	err = service.NewAddUserWorkCountService(ctx).AddUserWorkCount(request)
	if logger.CheckError(err, "Service AddUserWorkCount err") {
		r.BaseResp = pack.BuildBaseResp(err)
		return r, nil
	}
	// pack response
	r.BaseResp = pack.BuildBaseResp(errno.Success)
	return r, nil
}

func (u UserServiceImpl) UpdateUserFavoriteCount(ctx context.Context, request *demouser.UpdateUserFavoriteCountRequest) (r *demouser.UpdateUserFavoriteCountResponse, err error) {
	r = new(demouser.UpdateUserFavoriteCountResponse)
	// call service function
	err = service.NewUpdateUserFavoriteCountService(ctx).UpdateUserFavoriteCount(request)
	if logger.CheckError(err, "Service UpdateUserFavoriteCount err") {
		r.BaseResp = pack.BuildBaseResp(err)
		return r, nil
	}
	// pack response
	r.BaseResp = pack.BuildBaseResp(errno.Success)
	return r, nil
}

func (u UserServiceImpl) Register(ctx context.Context, request *demouser.RegisterRequest) (r *demouser.RegisterResponse, err error) {
	r = new(demouser.RegisterResponse)
	fmt.Println(request)
	uid, err := service.NewRegisterService(ctx).Register(request)
	if err != nil {
		r.BaseResp = pack.BuildBaseResp(err)
		return r, nil
	}
	r.UserId = uid
	r.BaseResp = pack.BuildBaseResp(errno.Success)
	return r, nil
}

func (u UserServiceImpl) Login(ctx context.Context, request *demouser.LoginRequest) (r *demouser.LoginResponse, err error) {
	r = new(demouser.LoginResponse)

	uid, err := service.NewLoginService(ctx).Login(request)
	if err != nil {
		fmt.Println("err:", err)
		r.BaseResp = pack.BuildBaseResp(err)
		return r, nil
	}

	r.UserId = uid
	r.BaseResp = pack.BuildBaseResp(errno.Success)
	fmt.Println(r)
	return r, nil
}

func (u UserServiceImpl) UserIndex(ctx context.Context, request *demouser.UserIndexRequest) (r *demouser.UserIndexResponse, err error) {
	r = new(demouser.UserIndexResponse)
	user, err := service.NewUserIndexService(ctx).UserIndex(request)
	if err != nil {
		r.BaseResp = pack.BuildBaseResp(err)
		return r, nil
	}
	r.User = user
	r.BaseResp = pack.BuildBaseResp(errno.Success)
	return r, nil
}
