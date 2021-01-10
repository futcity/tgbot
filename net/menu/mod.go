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
	"github.com/futcity/controller/server/api"
	"github.com/futcity/tgbot/auth"
	"github.com/futcity/tgbot/utils"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// ModMenu Main menu struct
type ModMenu struct {
	log *utils.Log
	ath *auth.Authorization
}

// NewModMenu make new struct
func NewModMenu(log *utils.Log, ath *auth.Authorization) *ModMenu {
	return &ModMenu{
		log: log,
		ath: ath,
	}
}

// UpdateMessage update telegram message
func (m *ModMenu) UpdateMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI, menu *MenuLevel) bool {
	var intro = "<b>Модули умного дома</b>"
	return m.SendMessage(bot, message, intro, menu)
}

// SendMessage sends message with buttons to user
func (m *ModMenu) SendMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string, mn *MenuLevel) bool {
	var keys [][]tgbotapi.KeyboardButton

	// Check privileges
	CheckPriv := func() bool {
		for _, bld := range *m.ath.User(message.From.UserName).Buildings() {
			if bld == mn.building {
				return true
			}
		}
		return false
	}
	if !CheckPriv() {
		return false
	}

	// Get building
	var building, err = m.ath.Building(mn.Building())
	if err != nil {
		m.SendError(bot, message, err.Error())
		return false
	}

	// Get data from controller
	var resp api.GroupResponse
	err = utils.HTTPGetData(building.IP(), building.Port(), "/user/"+m.ath.User(message.From.UserName).Key()+"/groups", 5, &resp)
	if err != nil {
		m.SendError(bot, message, err.Error())
		return false
	}

	// Make buttons from configs
	for _, row := range resp.Groups {
		var tgRow []tgbotapi.KeyboardButton
		tgRow = append(tgRow, tgbotapi.NewKeyboardButton(row))
		keys = append(keys, tgRow)
	}
	keys = append(keys, []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton("Назад")})

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
func (m *ModMenu) SendError(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string) {
	var msg = tgbotapi.NewMessage(message.Chat.ID, "<b>Ошибка: </b>"+text)
	msg.ParseMode = "HTML"
	msg.ChatID = message.Chat.ID
	bot.Send(msg)
}
