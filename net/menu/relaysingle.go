///////////////////////////////////////////////////////////////////
//
// Future City Project
//
// Copyright (C) 2020-2021 Sergey Denisov. GPLv3
//
// Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
//
///////////////////////////////////////////////////////////////////

package menu

import (
	"fmt"

	"github.com/futcity/controller/server/api"
	"github.com/futcity/tgbot/auth"
	"github.com/futcity/tgbot/utils"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// RelaySingleMenu Main menu struct
type RelaySingleMenu struct {
	log *utils.Log
	ath *auth.Authorization
}

// NewRelaySingleMenu make new struct
func NewRelaySingleMenu(log *utils.Log, ath *auth.Authorization) *RelaySingleMenu {
	return &RelaySingleMenu{
		log: log,
		ath: ath,
	}
}

// UpdateMessage update telegram message
func (r *RelaySingleMenu) UpdateMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI, menu *MenuLevel) bool {
	var intro = "<b>Управление розеткой</b>"
	return r.SendMessage(bot, message, intro, menu)
}

// SendMessage sends message with buttons to user
func (r *RelaySingleMenu) SendMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string, mn *MenuLevel) bool {
	var keys [][]tgbotapi.KeyboardButton

	// Get building
	var building, err = r.ath.Building(mn.Building())
	if err != nil {
		r.SendError(bot, message, err.Error())
		return false
	}

	fmt.Println("HERE!")

	// Get data from controller
	var devResp api.DeviceResponse
	err = utils.HTTPGetData(building.IP(), building.Port(), "/user/"+r.ath.User(message.From.UserName).Key()+"/device/desc/"+mn.Device(), 5, &devResp)
	if err != nil {
		r.SendError(bot, message, err.Error())
		return false
	}
	fmt.Println("HERE2!")

	var resp api.RelayDevResponse
	err = utils.HTTPGetData(building.IP(), building.Port(), "/user/"+r.ath.User(message.From.UserName).Key()+"/relay/"+devResp.Name, 5, &resp)
	if err != nil {
		r.SendError(bot, message, err.Error())
		return false
	}
	fmt.Println("HERE3!")
	// Make buttons from configs
	/*text += "\n\n"
	for _, relay := range resp.Relays {
		var tgRow []tgbotapi.KeyboardButton
		tgRow = append(tgRow, tgbotapi.NewKeyboardButton(relay.Description))
		keys = append(keys, tgRow)

		text += relay.Description + ": "
		if relay.Status {
			text += "<b>ВКЛ</b>"
		} else {
			text += "<b>ОТК</b>"
		}
		text += "\n"
	}*/
	keys = append(keys, []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Назад"),
		tgbotapi.NewKeyboardButton("Обновить"),
	})

	var keybd = tgbotapi.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       keys,
	}

	var msg = tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ParseMode = "HTML"
	msg.ChatID = message.Chat.ID
	msg.ReplyMarkup = keybd
	bot.Send(msg)

	return true
}

// SendError Send error tg message
func (r *RelaySingleMenu) SendError(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string) {
	var msg = tgbotapi.NewMessage(message.Chat.ID, "<b>Ошибка: </b>"+text)
	msg.ParseMode = "HTML"
	msg.ChatID = message.Chat.ID
	bot.Send(msg)
}
