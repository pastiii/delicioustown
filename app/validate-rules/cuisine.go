package validate_rules

type AddCuisine struct {
	Name              string `json:"name" form:"name" validate:"required,max=30"`
	CookingTechniques string `json:"cooking_techniques" form:"cooking_techniques" validate:"required,max=100"`
	Feature           string `json:"feature" form:"feature" validate:"required,max=100"`
	Producer          string `json:"producer" form:"producer" validate:"required,max=20"`
	Taste             string `json:"taste" form:"taste" validate:"required,max=20"`
	Remark            string `json:"remark" form:"remark" validate:"required,max=200"`
}

type OnlyId struct {
	CuisineId int64 `form:"cuisine_id" json:"cuisine_id" validate:"required"`
}

type CuisineList struct {
	Name     string `json:"name" form:"name" validate:"omitempty"`
	Page     int    `form:"page" json:"page" validate:"omitempty"`
	PageSize int    `form:"page_size" json:"page_size" validate:"omitempty"`
}

type EditCuisine struct {
	Id                int64  `form:"cuisine_id" json:"cuisine_id" validate:"required"`
	Name              string `json:"name" form:"name" validate:"required,max=30"`
	CookingTechniques string `json:"cooking_techniques" form:"cooking_techniques" validate:"required,max=100"`
	Feature           string `json:"feature" form:"feature" validate:"required,max=100"`
	Producer          string `json:"producer" form:"producer" validate:"required,max=20"`
	Taste             string `json:"taste" form:"taste" validate:"required,max=20"`
	Remark            string `json:"remark" form:"remark" validate:"required,max=200"`
}
