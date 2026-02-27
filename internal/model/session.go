package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Session struct {
	Id            int64          `user_service:"column:id;primaryKey;comment:自增id"`
	Uuid          string         `user_service:"column:uuid;uniqueIndex;type:char(20);comment:会话uuid"`
	SendId        string         `user_service:"column:send_id;Index;type:char(20);not null;comment:创建会话人id"`
	ReceiveId     string         `user_service:"column:receive_id;Index;type:char(20);not null;comment:接受会话人id"`
	ReceiveName   string         `user_service:"column:receive_name;type:varchar(20);not null;comment:名称"`
	Avatar        string         `user_service:"column:avatar;type:char(255);default:default_avatar.png;not null;comment:头像"`
	LastMessage   string         `user_service:"column:last_message;type:TEXT;comment:最新的消息"`
	LastMessageAt sql.NullTime   `user_service:"column:last_message_at;type:datetime;comment:最近接收时间"`
	CreatedAt     time.Time      `user_service:"column:created_at;Index;type:datetime;comment:创建时间"`
	DeletedAt     gorm.DeletedAt `user_service:"column:deleted_at;Index;type:datetime;comment:删除时间"`
}

func (Session) TableName() string {
	return "session"
}
