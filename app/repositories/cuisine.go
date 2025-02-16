package repositories

import (
	"DeliciousTown/app/model"
	OutField "DeliciousTown/app/setting/out-field"
	ValidateRules "DeliciousTown/app/validate-rules"
	"DeliciousTown/app/vo/admin"
	"DeliciousTown/global"
	"errors"
)

func CheckCuisineName(name string) (count int64, err error) {
	var cuisine model.Cuisine
	err = global.GvaMysqlClient.Model(&cuisine).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func AddCuisine(cuisine model.Cuisine) (bool bool, err error) {
	err = global.GvaMysqlClient.Create(&cuisine).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetCuisineInfoById(id int64) (OutField.CuisineField, error) {
	var cuisine model.Cuisine
	var cuisineInfo OutField.CuisineField
	err := global.GvaMysqlClient.Model(&cuisine).Where("id = ?", id).Scan(&cuisineInfo).Error
	if err != nil {
		return cuisineInfo, err
	}

	return cuisineInfo, nil
}

func GetCuisineList(params ValidateRules.CuisineList) (list map[string]interface{}, err error) {
	var cuisine model.Cuisine
	var outField []OutField.CuisineListField
	var total int64
	offset, size := GetOffsetAndPageSize(params.Page, params.PageSize)
	query := global.GvaMysqlClient.Model(cuisine)
	if len(params.Name) > 0 {
		query = query.Where("name like ?", "%"+params.Name+"%")
	}

	query.Count(&total)
	err = query.Limit(size).Offset(offset).Scan(&outField).Error
	if err != nil {
		return list, err
	}

	list = admin.GetList(total, outField)
	return list, nil
}

func EditCuisine(id int64, data interface{}) (bool, error) {
	var cuisine model.Cuisine
	err := global.GvaMysqlClient.Model(cuisine).Where("id = ?", id).Updates(data).Error
	if err != nil {
		err = errors.New("菜系信息更新失败")
		return false, err
	}

	return true, nil
}

func DelCuisine(id int64) (bool, error) {
	var cuisine model.Cuisine
	err := global.GvaMysqlClient.Where("id = ?", id).Delete(&cuisine).Error
	if err != nil {
		err = errors.New("删除失败")
		return false, err
	}

	return true, nil
}
