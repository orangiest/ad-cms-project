package mapper

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/model"
	"strconv"
)

// 添加资金记录
func InsertFund(fund model.Funds) error {

	err := model.DB.Create(&fund).Error

	return err

}

// 查找全部资金记录
func SelectFunds() (*[]model.Funds,error) {
	var funds []model.Funds

	if err := model.DB.Find(&funds).Error; err != nil {
		clog.Error("SelectFunds", err)
		return nil, err
	}

	return &funds, nil

}

// 查找单条资金记录
func SelectFundById(id string) (*model.Funds,error) {
	var fund model.Funds
	if err := model.DB.Where("system_id = ?", id).First(&fund).Error; err != nil {
		clog.Error("SelectMaterial", err)
		return nil, err
	}

	return &fund, nil
}

// 修改资金记录
func UpdateFund(u model.Funds) (*model.Funds,error) {
	var fund *model.Funds
	fund, err := SelectFundById(strconv.Itoa(u.SystemID))
	if err != nil {
		clog.Error("UpdateFund", err)
		return nil, err
	}

	if fund != nil {
		model.DB.Model(&fund).Update(u)
	}

	return fund, nil
}


// 删除资金记录
func DeleteFund(id string) error {
	err := model.DB.Where("system_id = ?", id).Delete(model.Funds{}).Error

	return err
}