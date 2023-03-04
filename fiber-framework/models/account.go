package models

import "github.com/uptrace/bun"

type Account struct {
	bun.BaseModel `bun:"table:push.app_account"`

	ID       int64   `json:"id" bun:"id,pk,autoincrement"`
	Address  string  `json:"address"`
	DeviceID int64   `json:"device_id"`
	Device   *Device `bun:"rel:belongs-to,join:device_id=id"`
}

// type Device struct {
// 	bun.BaseModel `bun:"table:push.device"`

// 	ID                    int64     `json:"id" bun:"id,pk,autoincrement"`
// 	Token                 string    `json:"token" bun:"token"`
// 	SubscribeTx           bool      `json:"subscribe_tx" bun:"subscribe_tx"`
// 	SubscribeAnnouncement bool      `json:"subscribe_announcement" bun:"subscribe_announcement"`
// 	TimeStamp             time.Time `json:"timestamp" bun:"timestamp,default:current_timestamp"`
// }
