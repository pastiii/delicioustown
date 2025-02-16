package validate_rules

//todo 添加菜品材料相关配置
type CreateMenu struct {
	Name        string        `json:"name" form:"name" validate:"required,max=30,min=2"`
	CuisineId   int64         `json:"cuisine_id" form:"cuisine_id" validate:"required"`
	Level       int64         `json:"level" form:"level" validate:"omitempty"`
	Tag         string        `json:"tag" form:"tag" validate:"required,max=10"`
	Ingredients []Ingredients `json:"ingredients" form:"ingredients" validate:"omitempty"`
}

//todo 编辑菜品材料相关配置
type EditMenu struct {
	Name        string        `json:"name" form:"name" validate:"required,max=30,min=2"`
	CuisineId   int64         `json:"cuisine_id" form:"cuisine_id" validate:"required"`
	Level       int64         `json:"level" form:"level" validate:"omitempty"`
	Tag         string        `json:"tag" form:"tag" validate:"required,max=10"`
	Ingredients []Ingredients `json:"ingredients" form:"ingredients" validate:"omitempty"`
}

type Ingredients struct {
	MenuId       int64 `json:"menu_id" form:"menu_id" validate:"required"`
	IngredientId int64 `json:"ingredient_id" form:"ingredient_id" validate:"required"`
	Num          int64 `json:"num" form:"num" validate:"required"`
	Level        int64 `json:"level" form:"level" validate:"omitempty"`
}
