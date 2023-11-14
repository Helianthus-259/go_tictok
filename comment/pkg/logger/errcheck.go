package logger

import (
	"fmt"
	"go.uber.org/zap"
)

// CheckError 专门用来处理错误信息
func CheckError(err error, msg string, args ...interface{}) bool {
	if err != nil {
		if len(args) > 0 {
			zap.L().Error(fmt.Sprintf(msg, args...), zap.Error(err))
		} else {
			zap.L().Error(msg, zap.Error(err))
		}
		return true
	}
	return false
}
