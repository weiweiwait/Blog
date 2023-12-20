package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 创建公共model
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"primary_key"`
	CreatedBy  string `json:"created_by" gorm:"created_by"`
	ModifiedBy string `json:"modified_by" gorm:"modified_by"`
	CreatedOn  uint32 `json:"created_on" gorm:"created_on"`
	ModifiedOn uint32 `json:"modified_on" gorm:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on" gorm:"deleted_on"`
	IsDel      uint8  `json:"is_del" gorm:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
