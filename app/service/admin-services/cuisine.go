package admin_services

import (
	"DeliciousTown/app/model"
	"DeliciousTown/app/repositories"
	OutField "DeliciousTown/app/setting/out-field"
	ValidateRules "DeliciousTown/app/validate-rules"
	"errors"
)

func AddCuisine(params ValidateRules.AddCuisine) (bool, error) {
	res, err := repositories.CheckCuisineName(params.Name)
	if err != nil {
		return false, err
	}

	if res > 0 {
		err = errors.New("菜系已存在请勿重复添加")
		return false, err
	}

	cuisine := model.Cuisine{
		Name:              params.Name,
		CookingTechniques: params.CookingTechniques,
		Feature:           params.Feature,
		Producer:          params.Producer,
		Remark:            params.Remark,
		Taste:             params.Taste,
	}
	ret, _ := repositories.AddCuisine(cuisine)
	if !ret {
		err = errors.New("菜系添加失败")
		return false, err
	}

	return true, nil
}

func GetCuisineInfo(id int64) (OutField.CuisineField, error) {
	info, err := repositories.GetCuisineInfoById(id)
	if err != nil {
		return info, err
	}

	return info, nil
}

func GetCuisineList(params ValidateRules.CuisineList) (map[string]interface{}, error) {
	ret, err := repositories.GetCuisineList(params)
	return ret, err
}

func EditCuisine(params ValidateRules.EditCuisine) (bool, error) {
	//todo 添加名称唯一校验
	cuisine := map[string]interface{}{
		"name":               params.Name,
		"cooking_techniques": params.CookingTechniques,
		"feature":            params.Feature,
		"producer":           params.Producer,
		"taste":              params.Taste,
		"remark":             params.Remark,
	}
	res, err := repositories.EditCuisine(params.Id, cuisine)
	return res, err
}

func DelCuisine(id int64) (bool, error) {
	res, err := repositories.DelCuisine(id)
	return res, err
}
