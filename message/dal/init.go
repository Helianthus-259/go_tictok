package dal

import (
	"message/dal/mysql"
	"message/dal/redis"
)

// Init init dal
func Init() {
	mysql.Init() // mysql init
	redis.InitRedis()
}
