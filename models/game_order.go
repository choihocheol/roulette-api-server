package models

import (
	"roulette-api-server/config"
	"roulette-api-server/models/schema"
	"roulette-api-server/types"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func QueryOrderById(order *schema.OrderRow) (err error) {
	err = config.DB.Table("game_order").Where("order_id = ?", order.OrderId).Find(order).Error
	return
}

func QueryOrderDetailById(order *types.ResGetLatestOrderByAddr) (err error) {
	sql := `
		SELECT * FROM game_order
		WHERE order_id = ?
	`
	if err = config.DB.Raw(sql, order.OrderId).Scan(order).Error; err != nil {return}
	
	sql = `
		SELECT * FROM prize
		WHERE prize_id=?
	`
	if err = config.DB.Raw(sql, order.PrizeId).Scan(&order.Prize).Error; err != nil {return}
	
	sql = `
		SELECT * FROM prize_denom
		WHERE prize_denom_id=?
	`
	if err = config.DB.Raw(sql, order.Prize.PrizeDenomId).Scan(&order.Prize.PrizeDenom).Error; err != nil {return}

	return
}

func QueryLatestOrderByAddr(order *types.ResGetLatestOrderByAddr) (err error) {
	sql := `
		SELECT * FROM game_order
		WHERE addr=? AND game_id=?
		ORDER BY order_id DESC LIMIT 1
	`
	if err = config.DB.Raw(sql, order.Addr, order.GameId).Scan(order).Error; err != nil {return}
	
	sql = `
		SELECT * FROM prize
		WHERE prize_id=?
	`
	if err = config.DB.Raw(sql, order.PrizeId).Scan(&order.Prize).Error; err != nil {return}
	
	sql = `
		SELECT * FROM prize_denom
		WHERE prize_denom_id=?
	`
	if err = config.DB.Raw(sql, order.Prize.PrizeDenomId).Scan(&order.Prize.PrizeDenom).Error; err != nil {return}

	if order.Status == 1 {
		order.IsWin = false
		order.PrizeId = 0
		order.Prize = types.ResOrderPrize{}
	}
	return
}

func UpdateOrder(order *schema.OrderRow) (err error) {
	err = config.DB.Table("game_order").Where("order_id = ?", order.OrderId).Update(order).Error
	return
}

func CreateOrderWithTx(tx *gorm.DB, order *schema.OrderRow) error {
	if tx == nil {
		tx = config.DB
	}

	err := tx.Table("game_order").Create(order).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}


func QueryOrders(orders *[]schema.OrderRow) (err error) {
	err = config.DB.Table("game_order").Find(orders).Error
	return
}

func CreateOrder(order *schema.OrderRow) (err error) {
	err = config.DB.Table("game_order").Create(order).Error
	return
}

func QueryOrder(order *schema.OrderRow) (err error) {
	err = config.DB.Table("game_order").Where("order_id = ?", order.OrderId).First(order).Error
	return
}

func QueryInProgressGameCnt(order *schema.OrderRow) (*types.Count, error) {
	sql := `
		SELECT COUNT(*) as cnt
		FROM game_order
		WHERE addr = ? AND
			status = 1 AND
			promotion_id = ? AND
			game_id = ?;
	`

	var res types.Count
	err := config.DB.Raw(sql, order.Addr, order.PromotionId, order.GameId).Scan(&res).Error
	return &res, err
}

func DeleteOrder(order *schema.OrderRow) (err error) {
	err = config.DB.Table("game_order").Where("order_id = ?", order.OrderId).Delete(order).Error
	return
}

func QueryGameWinningResults(results *[](*types.ResGetGameWinningResults)) (err error) {
	q := "SELECT " +
		"    G.addr, P.title, G.used_ticket_qty,  " +
		"	 D.name as 'prize_name', R.amount as 'prize_amount', G.status, " +
		"	 G.claimed_at, G.claim_finished_at,  " +
		"	 D.type as 'prize_type', R.prize_denom_id, G.prize_id, " +
		"	 G.order_id, G.account_id " +
		"  FROM GAME_ORDER G " +
		"   LEFT JOIN (SELECT title, promotion_id FROM promotion) P ON G.promotion_id = P.promotion_id " +
		"   LEFT JOIN prize R ON G.prize_id = R.prize_id " +
		"   LEFT JOIN (SELECT name, type, prize_denom_id FROM prize_denom) D ON R.prize_denom_id = D.prize_denom_id " +
		"  WHERE G.is_win = 1 " +
		"  ORDER BY G.claimed_at DESC, P.title ASC, G.addr ASC "

	if err = config.DB.Raw(q).Scan(results).
		Error; err != nil {
		return
	}
	return
}

func UpdateOrderStatusReset(order *schema.OrderRow) (err error) {
	err = config.DB.Table("game_order").Where("order_id = ?", order.OrderId).
		UpdateColumns(map[string]interface{}{
			"status":            order.Status,
			"claimed_at":        nil,
			"claim_finished_at": nil,
		}).Error
	return
}

func UpdateOrderStatusClaimed(order *schema.OrderRow) (err error) {
	err = config.DB.Table("game_order").Where("order_id = ?", order.OrderId).
		UpdateColumns(map[string]interface{}{
			"status":            order.Status,
			"claimed_at":        order.ClaimedAt,
			"claim_finished_at": nil,
		}).Error
	return
}
