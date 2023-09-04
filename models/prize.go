package models

import (
	"roulette-api-server/config"
	"roulette-api-server/models/schema"

	_ "github.com/go-sql-driver/mysql"
)

func QueryPrizes(prizes *[]schema.PrizeRow) (err error) {
	err = config.DB.Table("prize").Find(prizes).Error
	return
}

func CreatePrize(prize *schema.PrizeRow) (err error) {
	err = config.DB.Table("prize").Create(prize).Error
	return
}

func QueryPrize(prize *schema.PrizeRow) (err error) {
	err = config.DB.Table("prize").Where("prize_id = ?", prize.PrizeId).First(prize).Error
	return
}

func UpdatePrize(prize *schema.PrizeRow) (err error) {
	err = config.DB.Table("prize").Where("prize_id = ?", prize.PrizeId).Update(prize).Error
	return
}

func DeletePrize(prize *schema.PrizeRow) (err error) {
	err = config.DB.Table("prize").Where("prize_id = ?", prize.PrizeId).Delete(prize).Error
	return
}