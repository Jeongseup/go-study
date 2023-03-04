package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Device model info
// @Description ID account information
type Device struct {
	bun.BaseModel `swaggerignore:"true" bun:"table:push.device"`

	ID                    int64     `json:"id" bun:"id,pk,autoincrement"`
	Token                 string    `json:"token" bun:"token"`
	SubscribeTx           bool      `json:"subscribe_tx" bun:"subscribe_tx"`
	SubscribeAnnouncement bool      `json:"subscribe_announcement" bun:"subscribe_announcement"`
	TimeStamp             time.Time `json:"timestamp" bun:"timestamp,default:current_timestamp"`
}
