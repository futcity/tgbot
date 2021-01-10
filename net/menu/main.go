//*******************************************************************
//
// Future City Inc.
//
// Copyright (C) 2020 Sergey Denisov.
//
// Written by Sergey Denisov aka LittleBuster(DenisovS21@gmail.com)
// Github:  https://github.com/LittleBuster
//          https://github.com/futcity
//
// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public Licence 3
// as published by the Free Software Foundation; either version 3
// of the Licence, or(at your option) any later version.
//
//*******************************************************************

package menu

import (
	"github.com/futcity/tgbot/auth"
	"github.com/futcity/tgbot/utils"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// MainMenu Main menu struct
type MainMenu struct {
	log *utils.Log
	ath *auth.Authorization
}

// NewMainMenu make new struct
func NewMainMenu(log *utils.Log, ath *auth.Authorization) *MainMenu {
	return &MainMenu{
		log: log,
		ath: ath,
	}
}

// UpdateMessage update telegram message
func (m *MainMenu) UpdateMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI, menu *MenuLevel) bool {

	var intro = "Проект \"Город будущего\"\n\n" +
		"Любительский проект по автоматизации и удаленному управлению гродскими квартирами, дачными участками, частными домами и офисами."
	return m.SendMessage(bot, intro, message)
}

// SendMessage sends message with buttons to user
func (m *MainMenu) SendMessage(bot *tgbotapi.BotAPI, text string, message *tgbotapi.Message) bool {
	var keys [][]tgbotapi.KeyboardButton

	for _, row := range *m.ath.User(message.From.UserName).Buildings() {
		var tgRow []tgbotapi.KeyboardButton
		tgRow = append(tgRow, tgbotapi.NewKeyboardButton(row))
		keys = append(keys, tgRow)
	}

	var keybd = tgbotapi.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       keys,
	}

	var imgMsg = tgbotapi.NewPhotoShare(message.Chat.ID, "http://pvo74.ru/upload/iblock/a80/foto-k-novosti.jpg")
	imgMsg.Caption = text
	imgMsg.ChatID = message.Chat.ID
	imgMsg.ReplyMarkup = keybd
	bot.Send(imgMsg)

	return true
}

// SendError Send error tg message
func (m *MainMenu) SendError(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string) {
	var msg = tgbotapi.NewMessage(message.Chat.ID, "<b>Ошибка: </b>"+text)
	msg.ParseMode = "HTML"
	msg.ChatID = message.Chat.ID
	bot.Send(msg)
}
