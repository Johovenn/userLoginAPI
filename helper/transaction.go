package helper

import "gorm.io/gorm"

func RollbackOrCommit(tx *gorm.DB, err error) {
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
