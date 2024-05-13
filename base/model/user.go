package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Phone     string         `gorm:"type:varchar(50);"`
	Email     string         `gorm:"type:varchar(50);"`
	Password  string         `gorm:"type:varchar(30);"`
	Salt      string         `gorm:"type:varchar(20);"`
	CreatedAt time.Time      // 自动设置创建时间
	UpdatedAt time.Time      // 自动设置更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` // 支持软删除
}

type UserInfo struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      User   `gorm:"foreignKey:UserID"`
	Nick      string `gorm:"type:varchar(100)"`
	Avatar    string `gorm:"type:varchar(255)"`
	Sign      string `gorm:"type:varchar(255)"`
	Birth     time.Time
	CreatedAt time.Time      // 自动设置创建时间
	UpdatedAt time.Time      // 自动设置更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` // 支持软删除
}
