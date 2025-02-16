package repositories

import (
	"DeliciousTown/app/model"
	"DeliciousTown/global"
)

func CheckMenuName(name string) (int64, error) {
	var menu model.Menu
	var count int64
	err := global.GvaMysqlClient.Model(menu).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func CreateMenu(menu model.Menu) (bool, error) {
	err := global.GvaMysqlClient.Create(&menu).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
