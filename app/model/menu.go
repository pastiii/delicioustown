package model

type Menu struct {
	BaseFiled
	Name      string `json:"name,omitempty" form:"name" gorm:"varchar(30),not null"`
	CuisineId int64  `json:"cuisine_id,omitempty" form:"cuisine_id" gorm:"not null"`
	Level     int64  `json:"level,omitempty" form:"level" gorm:"not null"`
	Tag       string `json:"tag,omitempty" form:"tag" gorm:"varchar(10),not null"`
}
