package main

import "fmt"

func main() {
	updates, err := GetUpdates()
	PrintError(err)
	numUpdates := len(updates)
	for {
		updates, err = GetUpdates()
		PrintError(err)
		if len(updates) > numUpdates {
			for i := numUpdates; i < len(updates); i++ {
				// react to new msgs
				msgTxt := updates[i].Message.Text
				chatId := updates[i].Message.Chat.Id
				if msgTxt == "/start" {
					err = SendMessage(chatId, "Hello, it's bot for playing")
					PrintError(err)
				}
				if msgTxt == "/rules" {
					err = SendMessage(chatId, "Here are the rules")
					PrintError(err)
				}
			}
		}
		numUpdates = len(updates)
	}
}

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
