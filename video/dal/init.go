package dal

import (
	"video/dal/mysql"
	"video/dal/redis"
)

// Init init dal
func Init() {
	mysql.Init() // mysql init
	redis.InitRedis()
}
