package mapper

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"strconv"
	"time"
)

func InsertUser(user vo.UserInput) error {

	i := model.Users{
		Username:   user.Username,
		Password:   user.Password,
		Type:       user.Type,
		CreateTime: int(time.Now().Unix()),
		UpdateTime: int(time.Now().Unix()),
	}

	err := model.DB.Create(&i).Error

	return err

}

func SelectUserByCondition(userInput vo.UserInput) (*model.Users, error) {
	var user model.Users

	if err := model.DB.Where("username = ? and password = ?", userInput.Username, userInput.Password).First(&user).Error; err != nil {
		clog.Error("DB record not found ", err)
		return nil, err
	}

	return &user, nil
}

func SelectUsers() (*[]model.Users, error) {

	var users []model.Users
	if err := model.DB.Find(&users).Error; err != nil {
		clog.Error("DB record not found ", err)
		return nil, err
	}

	return &users, nil

}

func SelectUser(id string) (*model.Users, error) {
	var user model.Users
	if err := model.DB.Where("system_id = ?", id).First(&user).Error; err != nil {
		clog.Error("DB record not found ", err)
		return nil, err
	}

	return &user, nil
}

func UpdateUser(i vo.UserInput) (*model.Users, error) {

	var user *model.Users
	user, err := SelectUser(strconv.Itoa(int(i.SystemId)))
	if err != nil {
		clog.Error("UpdateUser", err)
		return nil, err
	}

	update := make(map[string]interface{})
	update["password"] = i.Password
	update["username"] = i.Username
	update["type"] = i.Type
	update["update_time"] = int(time.Now().Unix())

	if user != nil {
		model.DB.Model(user).Update(update)
	}

	//if i.Type != nil {
	//
	//}



	return user, nil

}

func DeleteUser(id string) error {

	err := model.DB.Where("system_id = ?", id).Delete(model.Users{}).Error

	return err
}