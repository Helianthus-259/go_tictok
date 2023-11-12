package redis

import (
	"context"
	"fmt"
	"time"
	"user/dal/model"
	"user/pkg/constants"
)

func HSetUserInfo(c context.Context, key string, value map[string]interface{}) error {
	err := HSet(c, key, value)
	if err != nil {
		return err
	}
	_, err = Expire(c, key, time.Hour*168)
	return err
}

func HSetUserCountInfo(c context.Context, key string, value map[string]interface{}) error {
	return HSet(c, key, value) // 需要频繁更改的信息，不设置过期时间
}

func HGetUserInfo(c context.Context, key string) (*model.User, error) {
	userMap, err := HGetAll(c, key)
	if err != nil {
		return nil, err
	}
	if len(userMap) == 0 {
		return nil, nil
	}
	return model.CreateUserInfo(userMap)
}

func HGetUserCountInfo(c context.Context, key string) (*model.UserCount, error) {
	userMap, err := HGetAll(c, key)
	fmt.Println("userMap:", userMap)
	if err != nil {
		return nil, err
	}
	if len(userMap) == 0 {
		return nil, nil
	}
	return model.CreateUserCountInfo(userMap)
}

func AddFollowCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.FollowCount, 1)
	return err
}

func AddFollowerCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.FollowerCount, 1)
	return err
}

func SubFollowCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.FollowCount, -1)
	return err
}

func SubFollowerCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.FollowerCount, -1)
	return err
}

func AddWorkCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.WorkCount, 1)
	return err
}

func AddFavoriteCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.FavoriteCount, 1)
	return err
}

func AddTotalFavoriteCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.TotalFavorited, 1)
	return err
}

func SubFavoriteCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.FavoriteCount, -1)
	return err
}

func SubTotalFavoriteCount(c context.Context, key string) error {
	_, err := IncrHMCount(c, key, constants.TotalFavorited, -1)
	return err
}
