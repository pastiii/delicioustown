package repositories

import (
	"DeliciousTown/app/model"
	OutField "DeliciousTown/app/setting/out-field"
	ValidateRules "DeliciousTown/app/validate-rules"
	"DeliciousTown/global"
)

type UserDao struct {
	UserId   int
	Account  string
	Password string
	Status   string
}

func CreateUser(user model.AdminUser) bool {
	err := global.GvaMysqlClient.Create(&user)
	if err.Error != nil {
		return false
	}

	return true
}

func (u *UserDao) GetCountByAccount() int64 {
	var count int64
	var user model.AdminUser
	err := global.GvaMysqlClient.Where("account = ?", u.Account).Find(&user).Count(&count)
	if err.Error != nil {
		return 0
	}

	return count
}

func (u *UserDao) GetUserByID() (*model.AdminUser, error) {
	var user model.AdminUser
	err := global.GvaMysqlClient.Where("id=?", u.UserId).First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (u *UserDao) EditRecord(updateData map[string]interface{}) bool {
	var user model.AdminUser
	err := global.GvaMysqlClient.Model(&user).Where("id=?", u.UserId).Updates(updateData).Error
	if err != nil {
		return false
	}

	return true
}

func UserList(params ValidateRules.AdminUserList) (count int64, list []OutField.AdminUserListField) {
	var user model.AdminUser
	var users []OutField.AdminUserListField
	var total int64
	offset, size := GetOffsetAndPageSize(params.Page, params.PageSize)
	query := global.GvaMysqlClient.Model(&user)
	if len(params.Status) != 0 {
		query = query.Where("status = ?", params.Status)
	}

	if len(params.Account) != 0 {
		query = query.Where("account like ?", "%"+params.Account+"%")
	}

	query.Count(&total)
	query.Offset(offset).Limit(size).Scan(&users)
	return total, users
}

func (u UserDao) DelUser() bool {
	var user model.AdminUser
	err := global.GvaMysqlClient.Where("id=?", u.UserId).Delete(&user).Error
	if err != nil {
		return false
	}

	return true
}

func (u UserDao) GetUserByParams() (model.AdminUser, error) {
	var user model.AdminUser
	query := global.GvaMysqlClient
	if len(u.Status) != 0 {
		query = query.Where("status = ?", u.Status)
	}

	if len(u.Account) != 0 {
		query = query.Where("account = ?", u.Account)
	}

	if len(u.Password) != 0 {
		query = query.Where("password = ?", u.Password)
	}

	err := query.First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
