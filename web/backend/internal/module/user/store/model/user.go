package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;uniqueIndex;comment:GitHub 用户名"`
	Token    string `gorm:"type:varchar(100);not null;uniqueIndex;comment:GitHub 用户令牌"`
}

//type VerifiedPublisher struct {
//	gorm.Model
//}
