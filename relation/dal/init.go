package dal

import (
	"relation/dal/mysql"
	"relation/dal/redis"
)

// Init init dal
func Init() {
	mysql.Init() // mysql init
	redis.InitRedis()
}
