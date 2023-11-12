package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
	"user/pkg/constants"
)

type User struct {
	ID              uint           `gorm:"primarykey" json:"id,string"`
	CreatedAt       time.Time      `json:"created_at,string"`
	UpdatedAt       time.Time      `json:"updated_at,string"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,string"`
	Username        string         `json:"username" gorm:"uniqueIndex:username;size:40;not null"` // 设置唯一索引，判断用户名是否重复
	Password        string         `json:"password" gorm:"not null"`
	Avatar          string         `json:"avatar" gorm:"type:varchar(255);not null"`           // 用户头像
	BackgroundImage string         `json:"background_image" gorm:"type:varchar(255);not null"` //背景图片
	//IsFollow        bool           `json:"is_follow,string" gorm:"not null"`                   // true-已关注，false-未关注 在数据库中这个字段没用
	Signature string `json:"signature" gorm:"type:varchar(255);"` // 个人简介
}

func (u *User) TableName() string {
	return constants.UserTableName
}

func CreateUserKey(userId uint) string {
	userStr := strconv.FormatInt(int64(userId), 10)
	return "user_info::" + userStr
}

func CreateUserCountKey(userId uint) string {
	userStr := strconv.FormatInt(int64(userId), 10)
	return "user_count::" + userStr
}

func CreateMapUserInfo(userInfo *User) map[string]interface{} {
	userStr, _ := json.Marshal(userInfo)
	userMap := make(map[string]interface{})
	_ = json.Unmarshal(userStr, &userMap)
	delete(userMap, "created_at")
	delete(userMap, "updated_at")
	delete(userMap, "deleted_at")
	fmt.Println("userMap:", userMap)
	return userMap
}

func CreateUserInfo(userMap map[string]string) (*User, error) {
	userStr, _ := json.Marshal(userMap)
	userInfo := new(User)
	err := json.Unmarshal(userStr, userInfo)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("userInfo:", userInfo)
	return userInfo, err
}

func CreateMapUserCount(userCount *UserCount) map[string]interface{} {
	userStr, _ := json.Marshal(userCount)
	userCountMap := make(map[string]interface{})
	_ = json.Unmarshal(userStr, &userCountMap)
	delete(userCountMap, "created_at")
	delete(userCountMap, "updated_at")
	delete(userCountMap, "deleted_at")
	delete(userCountMap, "User")
	fmt.Println("userCountMap:", userCountMap)
	return userCountMap
}

func CreateUserCountInfo(userCountMap map[string]string) (*UserCount, error) {
	userStr, _ := json.Marshal(userCountMap)
	userCountInfo := new(UserCount)
	err := json.Unmarshal(userStr, userCountInfo)
	return userCountInfo, err
}
