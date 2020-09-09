package main

import (
	"fmt"
	"strconv"
	"strings"
)

var rules string = "One player, the Chooser (a bot in our case), thinks of a four-letter number and the other player, the Guesser, tries to guess it" +

	"At each turn the Guesser tries a four-letter number, and the Chooser says how close it is to the answer by giving:" +

	"The number of Bulls - digits correct in the right position." +
	"The number of Cows - digits correct but in the wrong position." +
	"The Guesser tries to guess the answer in the fewest number of turns."

func main() {
	ChatIdAndNumber := make(map[int]int)
	updates, err := GetUpdates()
	PrintError(err)
	numUpdates := len(updates)
	for {
		updates, err = GetUpdates()
		PrintError(err)
		if len(updates) > numUpdates {
			for i := numUpdates; i < len(updates); i++ {
				// react to new msgs
				msgTxt := strings.TrimSpace(updates[i].Message.Text)
				chatId := updates[i].Message.Chat.Id
				if msgTxt == "/start" {
					err = SendMessage(chatId, "Hello, it's bot for playing Bulls and Cows")
					PrintError(err)
				}
				if msgTxt == "/newgame" {
					err = SendMessage(chatId, "Ok, your game starts now, bot came up with a number for you!")
					PrintError(err)
					number := GenerateNumber()
					ChatIdAndNumber[chatId] = number
				}
				if msgTxt == "/rules" {
					err = SendMessage(chatId, rules)
					PrintError(err)
				}
				if number, ok := ChatIdAndNumber[chatId]; ok {
					if len(msgTxt) == 4 {
						guess := msgTxt
						numberComp := strconv.Itoa(number)
						bullsAndCows := CheckBullsAndCows(guess, numberComp)
						bulls := strconv.Itoa(bullsAndCows[0])
						cows := strconv.Itoa(bullsAndCows[1])
						err = SendMessage(chatId,
							"You have now: "+bulls+" bulls and "+
								cows+" cows")
						PrintError(err)

						if bullsAndCows[0] == 4 {
							err = SendMessage(chatId, "You won a game, I really used number "+numberComp)
							PrintError(err)
							delete(ChatIdAndNumber, chatId)
						}
					}
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
