package dal

import (
	"comment/dal/mysql"
	"comment/dal/redis"
)

// Init init dal
func Init() {
	mysql.Init() // mysql init
	redis.InitRedis()
}
