package models

type TelegramUpdate struct {
	UpdateId int             `json:"update_id"`
	Message  TelegramMessage `json:"message"`
}

type TelegramMessage struct {
	Text string       `json:"text"`
	Chat TelegramChat `json:"chat"`
}

type TelegramChat struct {
	Id int `json:"id"`
}
