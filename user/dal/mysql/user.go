package mysql

import (
	"context"
	"gorm.io/gorm"
	"user/dal/model"
)

// AddUser Dao Function: Add a new user to the database
func AddUser(ctx context.Context, user *model.User) (id int64, err error) {
	result := DB.WithContext(ctx).Table("user").Create(user)
	err = result.Error
	id = int64(user.ID)
	return
}

// LoginProofread Dao Function: Verify username and password
func LoginProofread(ctx context.Context, name string, password string) (user *model.User, err error) {
	result := DB.WithContext(ctx).Table("user").Where("username=?", name).Where("password=?", password).First(&user)
	err = result.Error
	return
}

// GetUserInfoById Dao Function: Get UserInfo By user_id From Database
func GetUserInfoById(ctx context.Context, userId int64) (user *model.User, err error) {
	result := DB.WithContext(ctx).Table("user").Where("id=?", userId).First(&user)
	err = result.Error
	return
}

// CheckSameAccount Dao Function: Check if there are users with duplicate usernames From Database
func CheckSameAccount(ctx context.Context, name string) (user []model.User, err error) {
	result := DB.WithContext(ctx).Table("user").Where("username=?", name).Find(&user)
	err = result.Error
	return
}

// TransactionAdd Dao Function: Transaction Add a new user to the database
func TransactionAdd(ctx context.Context, user *model.User, userCount *model.UserCount) (id int64, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.WithContext(ctx).Table("user").Create(user).Error; err != nil {
			return err
		}
		userCount.UserId = user.ID
		if err = tx.WithContext(ctx).Table("user_count").Create(userCount).Error; err != nil {
			return err
		}
		return nil
	})
	id = int64(user.ID)
	return id, err
}
