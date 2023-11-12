package dal

import (
	"user/dal/mysql"
	"user/dal/redis"
)

// Init init dal
func Init() {
	mysql.Init() // mysql init
	redis.InitRedis()
}
