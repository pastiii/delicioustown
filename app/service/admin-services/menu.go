package admin_services

import (
	"DeliciousTown/app/model"
	"DeliciousTown/app/repositories"
	ValidateRules "DeliciousTown/app/validate-rules"
	"errors"
)

func CreateMenu(params ValidateRules.CreateMenu) (bool, error) {
	count, err := repositories.CheckMenuName(params.Name)
	if err != nil {
		return false, err
	}

	if count > 0 {
		err = errors.New("菜名已存在")
		return false, err
	}

	// todo 待添加事务，食材尚不作为必选参数，关联新版本数据表
	menu := model.Menu{
		Name:      params.Name,
		CuisineId: params.CuisineId,
		Level:     params.Level,
		Tag:       params.Tag,
	}
	res, _ := repositories.CreateMenu(menu)
	if !res {
		err = errors.New("添加菜单失败")
		return false, err
	}

	return true, nil
}

func EditMenu(params ValidateRules.EditMenu) (bool, error) {
	// 此处修改为校验 name+Id 同ID问题校验
	count, err := repositories.CheckMenuName(params.Name)
	if err != nil {
		return false, err
	}

	if count > 0 {
		err = errors.New("名称已存在")
		return false, err
	}

	menu := model.Menu{
		Name:      params.Name,
		CuisineId: params.CuisineId,
		Level:     params.Level,
		Tag:       params.Tag,
	}
	res, _ := repositories.CreateMenu(menu)
	if !res {
		err = errors.New("添加菜单失败")
		return false, err
	}

	return true, nil
}

func HandleIngredientData(ingredientData []*ValidateRules.Ingredients) {

}
