package main

// Telegram Update
type UpdateT struct {
	UpdateId int      `json:"update_id"`
	Message  MessageT `json:"message"`
}

// Telegram Message
type MessageT struct {
	MessageId int   `json:"message_id"`
	Chat      ChatT `json:"chat"`
	Date      int   `json:"date"`
}

// Telegram Chat
type ChatT struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}
