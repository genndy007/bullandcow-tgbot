package main

// Telegram Updates Result
type UpdatesResultT struct {
	Ok     bool      `json:"ok"`
	Result []UpdateT `json:"result"`
}

type SendResultT struct {
	Ok     bool     `json:"ok"`
	Result MessageT `json:"result"`
}

// Telegram Update
type UpdateT struct {
	UpdateId int      `json:"update_id"`
	Message  MessageT `json:"message"`
}

// Telegram Message
type MessageT struct {
	MessageId int    `json:"message_id"`
	Chat      ChatT  `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

// Telegram Chat
type ChatT struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}
