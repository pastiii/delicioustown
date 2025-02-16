package model

// Ingredient 食材信息
type Ingredient struct {
	BaseFiled
	Name  string `json:"name,omitempty" form:"name" gorm:"varchar(30),not null"`
	Class int64  `json:"class,omitempty" form:"class" gorm:"not null"`
	Level int64  `json:"level,omitempty" form:"level" gorm:"not null"`
	Price int64  `json:"price,omitempty" form:"price" gorm:"not null"`
}

// FoodIngredient 菜品食材 修改逻辑，暂时不是创建与等级进行关联的食谱升级体系，修改为统一的升级体系
type FoodIngredient struct {
	Name  string `json:"name,omitempty" form:"name" gorm:"varchar(30),not null"`
	Class int64  `json:"class,omitempty" form:"class" gorm:"not null"`
	Level int64  `json:"level,omitempty" form:"level" gorm:"not null"`
	Price int64  `json:"price,omitempty" form:"price" gorm:"not null"`
}
