package schema

import (
	"time"
)

type PrizeDistPoolRow struct {
	DistPoolId             int64     `gorm:"column:dist_pool_id" json:"distPoolId" db:"dist_pool_id"`
	PromotionId            int64     `json:"promotionId" db:"promotion_id"`
	PrizeDenomId           int64     `json:"prizeDenomId" db:"prize_denom_id"`
	TotalSupply            uint64    `json:"totalSupply" db:"total_supply"`
	RemainingQty           uint64    `json:"remainingQty" db:"remaining_qty"`
	IsActive               bool      `json:"isActive" db:"is_active"`
	CreatedAt              time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt              time.Time `json:"updatedAt" db:"updated_at"`
	//temp
	// Name string `json:"prizeName" db:"prize_name" gorm:"column:name"`
	// Type string `json:"prizeType" db:"prizeType" gorm:"column:type"`
}

type PrizeDistPoolInsertRow struct {
	ID                     int64
	DistPoolId             int64     `gorm:"column:dist_pool_id" json:"distPoolId" db:"dist_pool_id"`
	PromotionId            int64     `json:"promotionId" db:"promotion_id"`
	PrizeDenomId           int64     `json:"prizeDenomId" db:"prize_denom_id"`
	TotalSupply            uint64    `json:"totalSupply" db:"total_supply"`
	RemainingQty           uint64    `json:"remainingQty" db:"remaining_qty"`
	IsActive               bool      `json:"isActive" db:"is_active"`
	CreatedAt              time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt              time.Time `json:"updatedAt" db:"updated_at"`
	//temp
	// Name string `json:"prizeName" db:"prize_name" gorm:"column:name"`
	// Type string `json:"prizeType" db:"prizeType" gorm:"column:type"`
}

func (r *PrizeDistPoolRow) TableName() string {
	return "distribution_pool"
}
