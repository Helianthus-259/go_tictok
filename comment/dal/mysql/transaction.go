package mysql

import (
	"comment/pkg/logger"
	"gorm.io/gorm"
)

// Transaction 事务的操作

func MyTransaction(conn *gorm.DB, f func(tx *gorm.DB) error) error {
	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := f(tx)
	if logger.CheckError(err, "Transaction Err, DataBase Rollback") {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func MyTransactionWithResult(conn *gorm.DB, f func(tx *gorm.DB) (interface{}, error)) (interface{}, error) {
	tx := conn.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	result, err := f(tx)
	if logger.CheckError(err, "Transaction Err, DataBase Rollback") {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return result, err
}
