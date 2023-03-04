package entities

type Account struct {
	ID       int64  `json:"id"`
	Address  string `json:"address"`
	DeviceID int64  `json:"device_id"`
}

type Device struct {
	ID                    int64  `json:"id"`
	Token                 string `json:"token"`
	SubscribeTx           bool   `json:"subscribe_tx"`
	SubscribeAnnouncement bool   `json:"subscribe_announcement"`
}
