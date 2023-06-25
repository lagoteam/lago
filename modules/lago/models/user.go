package models

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
	"github.com/shopspring/decimal"
)

type User struct {
	TableID
	Username          string          `json:"username"`
	Password          string          `json:"password"`
	Email             string          `json:"email"`
	Mobile            string          `json:"mobile"`
	Nickname          string          `json:"nickname"`
	Avatar            string          `json:"avatar"`
	Signature         string          `json:"signature"`
	RealName          string          `json:"real_name"`
	IdentityCard      string          `json:"identity_card"`
	Gender            int             `json:"gender"`
	Age               int             `json:"age"`
	Birthday          carbon.DateTime `json:"birthday"`
	Money             decimal.Decimal `json:"money"`
	Monetary          decimal.Decimal `json:"monetary"`
	Level             int             `json:"level"`
	Integral          int             `json:"integral"`
	HistoryIntegral   int             `json:"history_integral"`
	VipLevel          int             `json:"vip_level"`
	VipExp            int             `json:"vip_exp"`
	VipTime           carbon.DateTime `json:"vip_time"`
	EmailVerifiedTime carbon.DateTime `json:"email_verified_time"`
	InviteCode        string          `json:"invite_code"`
	Status            int             `json:"status"`
	TeamId            uint            `json:"team_id"`
	PromotionId       uint            `json:"promotion_id"`
	PromotionTime     carbon.DateTime `json:"promotion_time"`
	TableTenantId
	TableTimestamps
	TableSoftDeletes
}

func (r *User) TableName() string {
	prefix := facades.Config().Env("DB_PREFIX", "").(string)
	return prefix + "user"
}
