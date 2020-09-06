package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var baseUrl = "https://api.telegram.org/bot"

func GetBody(methodName string) ([]byte, error) {
	resp, err := http.Get(baseUrl + botToken + methodName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Method getUpdates receives incoming updates
func GetUpdates() ([]UpdateT, error) {
	methodName := "/getUpdates"
	body, err := GetBody(methodName)
	if err != nil {
		return nil, err
	}

	var updatesResult UpdatesResultT
	json.Unmarshal(body, &updatesResult)

	return updatesResult.Result, nil

}

// Send text message to chat with chat_id
func SendMessage(chatId int, text string) error {
	methodName := "/sendMessage"
	query := methodName + "?chat_id=" + strconv.Itoa(chatId) + "&text=" + text

	body, err := GetBody(query)
	if err != nil {
		return err
	}
	var sendResult SendResultT
	err = json.Unmarshal(body, &sendResult)

	if !sendResult.Ok {
		return errors.New("Message is not sent correctly")
	}

	return nil
}
