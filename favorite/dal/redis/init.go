package redis

import (
	"context"
	"favorite/pkg/constants"
	"favorite/pkg/logger"
	"github.com/go-redis/redis/v8"
	"time"
)

var Rdb *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     constants.RedisDefaultAddr,
		Password: "",
		DB:       1,
	})
	err := rdb.Ping(context.Background()).Err()
	_ = logger.CheckError(err, "redis init error")
	Rdb = rdb
}

func GetKeys(c context.Context, pattern string) ([]string, error) {
	return Rdb.Keys(c, pattern).Result()
}

func KeyExist(c context.Context, key string) (int64, error) {
	return Rdb.Exists(c, key).Result()
}

func Set(c context.Context, key, value string, expire time.Duration) error {
	err := Rdb.Set(c, key, value, expire).Err()
	return err
}

func Get(c context.Context, key string) (string, error) {
	result, err := Rdb.Get(c, key).Result()
	return result, err
}

func Del(c context.Context, key string) error {
	return Rdb.Del(c, key).Err()
}

func Expire(c context.Context, key string, expireTime time.Duration) (bool, error) {
	// 有目标key返回true，没有目标key，返回false
	return Rdb.Expire(c, key, expireTime).Result()
}

func HSet(c context.Context, key string, value interface{}) error {
	return Rdb.HSet(c, key, value).Err()
}

func HGet(c context.Context, key string, filed string) error {
	return Rdb.HGet(c, key, filed).Err()
}

func HGetAll(c context.Context, key string) (map[string]string, error) {
	return Rdb.HGetAll(c, key).Result()
}

func IncrHMCount(c context.Context, key, field string, incr int64) (int64, error) {
	return Rdb.HIncrBy(c, key, field, incr).Result()
}

func SAdd(c context.Context, key string, value interface{}) (int64, error) {
	// 返回存入数据的数量
	return Rdb.SAdd(c, key, value).Result()
}

func SGetAll(c context.Context, key string) ([]string, error) {
	return Rdb.SMembers(c, key).Result()
}

func SDel(c context.Context, key string, value interface{}) (int64, error) {
	return Rdb.SRem(c, key, value).Result()
}

func SIsExist(c context.Context, key string, value interface{}) (bool, error) {
	return Rdb.SIsMember(c, key, value).Result()
}

func ZSet(c context.Context, key string, score []float64, member []interface{}) (int64, error) {
	// 返回存入数据的数量
	z := make([]*redis.Z, 0)
	for i := 0; i < len(score); i++ {
		z = append(z, &redis.Z{Score: score[i], Member: member[i]})
	}
	return Rdb.ZAdd(c, key, z...).Result()
}

func ZGetRangeWithScores(c context.Context, key string, start, stop int64) ([]redis.Z, error) {
	return Rdb.ZRangeWithScores(c, key, start, stop).Result()
}

func ZGetRangeByScoreWithScores(c context.Context, key string, min, max string, offset, count int64) ([]redis.Z, error) {
	return Rdb.ZRangeByScoreWithScores(c, key, &redis.ZRangeBy{Min: min, Max: max, Offset: offset, Count: count}).Result()
}

func ZGetRevRangeByScoreWithScores(c context.Context, key string, min, max string, offset, count int64) ([]redis.Z, error) {
	return Rdb.ZRevRangeByScoreWithScores(c, key, &redis.ZRangeBy{Min: min, Max: max, Offset: offset, Count: count}).Result()
}

func ZDel(c context.Context, key string, value interface{}) (int64, error) {
	return Rdb.ZRem(c, key, value).Result()
}

func LPush(c context.Context, key string, values ...interface{}) (int64, error) {
	return Rdb.LPush(c, key, values).Result()
}

func LPopCount(c context.Context, key string, count int) ([]string, error) {
	return Rdb.LPopCount(c, key, count).Result()
}
