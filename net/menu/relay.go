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
	"net/url"

	"github.com/futcity/controller/server/api"
	"github.com/futcity/tgbot/auth"
	"github.com/futcity/tgbot/utils"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// RelayMenu Main menu struct
type RelayMenu struct {
	log *utils.Log
	ath *auth.Authorization
}

// NewRelayMenu make new struct
func NewRelayMenu(log *utils.Log, ath *auth.Authorization) *RelayMenu {
	return &RelayMenu{
		log: log,
		ath: ath,
	}
}

// UpdateMessage update telegram message
func (r *RelayMenu) UpdateMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI, menu *MenuLevel) bool {
	var intro = "<b>Умные Розетки</b>"
	return r.SendMessage(bot, message, intro, menu)
}

// SendMessage sends message with buttons to user
func (r *RelayMenu) SendMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string, mn *MenuLevel) bool {
	var keys [][]tgbotapi.KeyboardButton

	//
	// Check privileges
	//
	CheckPriv := func() bool {
		for _, bld := range *r.ath.User(message.From.UserName).Buildings() {
			if bld == mn.building {
				return true
			}
		}
		return false
	}
	if !CheckPriv() {
		return false
	}

	//
	// Get building adress
	//
	var building, err = r.ath.Building(mn.Building())
	if err != nil {
		return r.SendError(bot, message, err.Error())
	}

	//
	// Switch relay
	//
	if message.Text != "Обновить" && message.Text != "Назад" {
		var devResp api.DeviceResponse
		var relResp api.RelayResponse
		err = utils.HTTPGetData(building.IP(), building.Port(), "/user/"+r.ath.User(message.From.UserName).Key()+"/device/desc/"+
			url.QueryEscape(message.Text), 5, &devResp)
		if err != nil {
			r.SendError(bot, message, err.Error())
		}
		if !devResp.Result {
			r.SendError(bot, message, devResp.Error)
		}

		err = utils.HTTPGetData(building.IP(), building.Port(), "/user/"+r.ath.User(message.From.UserName).Key()+"/relay/"+devResp.Name+"/switch", 5, &relResp)
		if err != nil {
			r.SendError(bot, message, err.Error())
		}
		if !relResp.Result {
			r.SendError(bot, message, relResp.Error)
		}
	}

	//
	// Get data from controller
	//
	var resp api.RelayDevResponse
	err = utils.HTTPGetData(building.IP(), building.Port(), "/user/"+r.ath.User(message.From.UserName).Key()+"/relay", 5, &resp)
	if err != nil {
		return r.SendError(bot, message, err.Error())
	}

	//
	// Make buttons from configs
	//
	text += "\n\n"
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
	}
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
func (r *RelayMenu) SendError(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string) bool {
	var msg = tgbotapi.NewMessage(message.Chat.ID, "<b>Ошибка: </b>"+text)

	msg.ParseMode = "HTML"
	msg.ChatID = message.Chat.ID
	bot.Send(msg)

	return false
}
