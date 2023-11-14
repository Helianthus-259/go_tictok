package dal

import (
	"favorite/dal/mysql"
	"favorite/dal/redis"
)

// Init init dal
func Init() {
	mysql.Init() // mysql init
	redis.InitRedis()
}
