package model

import (
	"time"

	"gorm.io/gorm"
)

/*
1.密文2.密文不可反解.
1.对称加密
2.非对称加密
3. md5信息摘要算法
密码如果不可以反解,用户找回密码
*/

type BaseModel struct {
	ID        int32     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:null;type:varchar(6) comment 'female女,male男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1普通用户,2管理员'"`
}
