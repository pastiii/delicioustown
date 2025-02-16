package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BaseFiled struct {
	ID        int64          `json:"id,omitempty" gorm:"size:11;primaryKey;autoIncrement:true"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type XTime struct {
	time.Time
}

//自定义序列化规则
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

//应用于value方法
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

//应用于Scan方法
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
