package model

import (
	"gorm.io/gorm"
	"time"
)

type ContactApply struct {
	Id          int64          `user_service:"column:id;primaryKey;comment:自增id"`
	Uuid        string         `user_service:"column:uuid;uniqueIndex;type:char(20);comment:申请id"`
	UserId      string         `user_service:"column:user_id;index;type:char(20);not null;comment:申请人id"`
	ContactId   string         `user_service:"column:contact_id;index;type:char(20);not null;comment:被申请id"`
	ContactType int8           `user_service:"column:contact_type;not null;comment:被申请类型，0.用户，1.群聊"`
	Status      int8           `user_service:"column:status;not null;comment:申请状态，0.申请中，1.通过，2.拒绝，3.拉黑"`
	Message     string         `user_service:"column:message;type:varchar(100);comment:申请信息"`
	LastApplyAt time.Time      `user_service:"column:last_apply_at;type:datetime;not null;comment:最后申请时间"`
	DeletedAt   gorm.DeletedAt `user_service:"column:deleted_at;index;type:datetime;comment:删除时间"`
}

func (ContactApply) TableName() string {
	return "contact_apply"
}
