package admin_services

import (
	"DeliciousTown/app/model"
	"DeliciousTown/app/repositories"
	OutField "DeliciousTown/app/setting/out-field"
	"DeliciousTown/app/util/common"
	ValidateRules "DeliciousTown/app/validate-rules"
	"errors"
)

func CreateAdminUser(params ValidateRules.AdminUser) bool {
	token := common.GetPassword(params.Password)
	AdminUser := model.AdminUser{
		Account:  params.Account,
		Password: token,
	}

	return repositories.CreateUser(AdminUser)
}

func CheckAccount(params ValidateRules.AdminUser) bool {
	dao := repositories.UserDao{
		Account: params.Account,
	}

	count := dao.GetCountByAccount()
	if count > 0 {
		return false
	}

	return true
}

func GetUserInfo(userId int) (*model.AdminUser, bool) {
	dao := repositories.UserDao{
		UserId: userId,
	}
	adminUser, err := dao.GetUserByID()
	if err != nil {
		return adminUser, false
	}

	return adminUser, true
}

func EditPass(params ValidateRules.EditPass) (bool bool, err error) {
	dao := repositories.UserDao{
		UserId: params.UserId,
	}

	oldPass := common.GetPassword(params.OldPassword)
	adminUser, ok := dao.GetUserByID()
	if ok != nil || adminUser.Password != oldPass {
		err = errors.New("用户原始密码错误")
		return false, err
	}

	if params.OldPassword == params.NewPassword {
		return true, nil
	}

	updateData := make(map[string]interface{})
	updateData["password"] = common.GetPassword(params.NewPassword)
	res := dao.EditRecord(updateData)
	if !res {
		err = errors.New("密码更新失败,请联系管理员")
		return false, err
	}

	return true, nil
}

func EditUserStatus(params ValidateRules.EditUserStatus) (bool bool, err error) {
	dao := repositories.UserDao{
		UserId: params.UserId,
	}

	adminUser, ok := dao.GetUserByID()
	if ok != nil {
		err = errors.New("用户原始密码错误")
		return false, err
	}

	if adminUser.Status == params.Status {
		return true, nil
	}

	updateData := make(map[string]interface{})
	updateData["status"] = params.Status
	res := dao.EditRecord(updateData)
	if !res {
		err = errors.New("用户状态更新失败")
		return false, err
	}

	return true, nil
}

func UserList(params ValidateRules.AdminUserList) (count int64, data []OutField.AdminUserListField) {
	total, list := repositories.UserList(params)
	return total, list
}

func DelUser(userId int) (bool bool, err error) {
	dao := repositories.UserDao{
		UserId: userId,
	}
	ok := dao.DelUser()
	if !ok {
		err := errors.New("删除失败,请联系管理员")
		return false, err
	}

	return true, nil
}
