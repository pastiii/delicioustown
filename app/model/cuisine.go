package model

type Cuisine struct {
	BaseFiled
	Name string `json:"name,omitempty" form:"name" gorm:"varchar(20),not null"`
	CookingTechniques string `json:"cooking_techniques,omitempty" form:"cooking_techniques" gorm:"varchar(100),not null"`
	Feature string `json:"feature,omitempty" form:"feature" gorm:"varchar(100),not null"`
	Producer string `json:"producer,omitempty" form:"producer" gorm:"varchar(20),not null"`
	Taste string `json:"taste,omitempty" form:"taste" gorm:"varchar(20),not null"`
	Remark string `json:"remark,omitempty" form:"taste" gorm:"varchar(200),not null"`
}
