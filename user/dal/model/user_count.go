package model

import (
	"gorm.io/gorm"
	"time"
	"user/pkg/constants"
)

type UserCount struct {
	ID             uint           `gorm:"primarykey" json:"id,string"`
	CreatedAt      time.Time      `json:"created_at,string"`
	UpdatedAt      time.Time      `json:"updated_at,string"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,string"`
	UserId         uint           `json:"user_id,string" gorm:"uniqueIndex,not null"`
	User           User           `gorm:"ForeignKey:UserId"`
	FollowCount    int64          `json:"follow_count,string" `    // 关注总数
	FollowerCount  int64          `json:"follower_count,string" `  // 粉丝总数
	TotalFavorited int64          `json:"total_favorited,string" ` // 获赞数量
	WorkCount      int64          `json:"work_count,string" `      // 作品数
	FavoriteCount  int64          `json:"favorite_count,string" `  // 点赞总数
}

func (u *UserCount) TableName() string {
	return constants.UserCountTableName
}
