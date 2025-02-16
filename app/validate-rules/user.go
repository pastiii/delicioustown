package validate_rules

type User struct {
	Name string `from:"name" json:"name" validate:"required,max=20,min=3,gsv"`
	Age  int    `from:"age" json:"age"`
	Sex  string `json:"sex"`
	Info  []int `json:"info"`
}