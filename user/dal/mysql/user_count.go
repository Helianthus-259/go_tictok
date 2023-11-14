package mysql

import (
	"context"
	"gorm.io/gorm"
	"user/dal/model"
)

// AddUserCount  Dao Function: Add a new user to the database
func AddUserCount(ctx context.Context, Db *gorm.DB, userCount *model.UserCount) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Create(userCount)
	err = result.Error
	return
}

// GetUserCountInfoByUserId  Dao Function: Get UserCountInfo By user_id From Database
func GetUserCountInfoByUserId(ctx context.Context, Db *gorm.DB, userId int64) (userCount *model.UserCount, err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).First(&userCount)
	err = result.Error
	return
}

// AddFollowCountByUserId Dao Function:Add FollowCountByUserId By user_id From Database
func AddFollowCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("follow_count", gorm.Expr("follow_count+ ?", 1))
	err = result.Error
	return
}

// AddFollowerCountByUserId Dao Function:Add FollowCounterByUserId By user_id From Database
func AddFollowerCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("follower_count", gorm.Expr("follower_count+ ?", 1))
	err = result.Error
	return
}

// SubFollowCountByUserId Dao Function:Sub FollowCountByUserId By user_id From Database
func SubFollowCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("follow_count", gorm.Expr("follow_count- ?", 1))
	err = result.Error
	return
}

// SubFollowerCountByUserId Dao Function:Sub FollowCounterByUserId By user_id From Database
func SubFollowerCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("follower_count", gorm.Expr("follower_count- ?", 1))
	err = result.Error
	return
}

// AddWorkCountByUserId Dao Function:Add WorkCounterByUserId By user_id From Database
func AddWorkCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("work_count", gorm.Expr("work_count+ ?", 1))
	err = result.Error
	return
}

// AddFavoriteCountByUserId Dao Function:Add FavoriteCountByUserId By user_id From Database
func AddFavoriteCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1))
	err = result.Error
	return
}

// AddTotalFavoriteCountByUserId Dao Function:Add TotalFavoriteCountByUserId By user_id From Database
func AddTotalFavoriteCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("total_favorited", gorm.Expr("total_favorited+ ?", 1))
	err = result.Error
	return
}

// SubFavoriteCountByUserId Dao Function:Sub FavoriteCountByUserId By user_id From Database
func SubFavoriteCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("favorite_count", gorm.Expr("favorite_count- ?", 1))
	err = result.Error
	return
}

// SubTotalFavoriteCountByUserId Dao Function:Sub TotalFavoriteCountByUserId By user_id From Database
func SubTotalFavoriteCountByUserId(ctx context.Context, Db *gorm.DB, userId int64) (err error) {
	result := Db.WithContext(ctx).Table("user_count").Where("user_id=?", userId).Update("total_favorited", gorm.Expr("total_favorited- ?", 1))
	err = result.Error
	return
}
