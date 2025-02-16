package out_field

import "DeliciousTown/app/model"

type CuisineField struct {
	Id                int64       `json:"id"`
	Name              string      `json:"name"`
	CookingTechniques string      `json:"cooking_techniques"`
	Feature           string      `json:"feature"`
	Producer          string      `json:"producer"`
	Taste             string      `json:"taste"`
	Remark            string      `json:"remark"`
	CreatedAt         model.XTime `json:"created_at"`
}

type CuisineListField struct {
	Id        int64       `json:"id"`
	Name      string      `json:"name"`
	Producer  string      `json:"producer"`
	Taste     string      `json:"taste"`
	CreatedAt model.XTime `json:"created_at"`
}
