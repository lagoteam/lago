package models

import (
	"github.com/goravel/framework/support/carbon"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type BaseModel struct {
	TableID
	TableTimestamps
}

type TableID struct {
	ID uint `gorm:"primaryKey;column:id;type:bigint(20) unsigned;autoIncrement;comment:ID;" json:"id,omitempty"`
}

type TableTenantId struct {
	TenantId uint `gorm:"column:tenant_id;type:bigint(20) unsigned;default:0;comment:租户ID;index;" json:"tenant_id,omitempty"`
}

type TableSoftDeletes struct {
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:datetime;comment:删除时间;index;autoDeleteTime;" json:"delete_time,omitempty"`
}

type TableTimestamps struct {
	CreateTime carbon.DateTime `gorm:"column:create_time;type:datetime;comment:创建时间;index;autoCreateTime;" json:"create_time,omitempty"`
	UpdateTime carbon.DateTime `gorm:"column:update_time;type:datetime;comment:更新时间;index;autoUpdateTime;" json:"update_time,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (m BaseModel) GetStringID() string {
	return cast.ToString(m.ID)
}

// GetStringID 获取 ID 的字符串格式
func (t TableID) GetStringID() string {
	return cast.ToString(t.ID)
}
