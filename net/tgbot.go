/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2019 Sergey Denisov.
/*
/* Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
/* Github: https://github.com/LittleBuster
/*	       https://github.com/futcamp
/*
/* This library is free software; you can redistribute it and/or
/* modify it under the terms of the GNU General Public Licence 3
/* as published by the Free Software Foundation; either version 3
/* of the Licence, or (at your option) any later version.
/*
/*******************************************************************/

package net

import (
	"fmt"

	"github.com/futcity/tgbot/auth"
	"github.com/futcity/tgbot/net/menu"
	"github.com/futcity/tgbot/utils"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// TgBot Telegram Bot struct
type TgBot struct {
	level map[string]*menu.MenuLevel
	main  *menu.MainMenu
	mod   *menu.ModMenu
	relay *menu.RelayMenu
	log   *utils.Log
	ath   *auth.Authorization
}

// NewTgBot make new struct
func NewTgBot(main *menu.MainMenu, mod *menu.ModMenu, log *utils.Log, r *menu.RelayMenu,
	ath *auth.Authorization) *TgBot {
	return &TgBot{
		main:  main,
		log:   log,
		mod:   mod,
		relay: r,
		ath:   ath,
	}
}

// StartBot starting bot loop
func (tb *TgBot) StartBot(key string) error {
	var bot, errb = tgbotapi.NewBotAPI(key)
	if errb != nil {
		return errb
	}
	bot.Debug = false

	var u = tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var updates, erru = bot.GetUpdatesChan(u)
	if erru != nil {
		return erru
	}

	tb.level = make(map[string]*menu.MenuLevel)
	for _, user := range tb.ath.Users() {
		tb.level[user.Login()] = &menu.MenuLevel{}
		tb.level[user.Login()].SetLevel(menu.MainLevel)
	}

	// Main loop
	for update := range updates {
		if update.Message != nil {
			if tb.ath.User(update.Message.From.UserName) == nil {
				tb.SendError(bot, update.Message, "Доступ запрещён. Зарегистрируйтесь.")
				continue
			}
			tb.UpdateMessage(update.Message, bot)
		}
	}

	return nil
}

// UpdateMessage update messages handler
func (tb *TgBot) UpdateMessage(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	tb.log.Info("TGBOT", "New message \""+msg.Text+"\" from user \""+msg.From.UserName+
		"\" chat \""+fmt.Sprintf("%d", msg.Chat.ID)+"\"")

	tb.ChangeLevel(msg)

	for {
		switch tb.level[msg.From.UserName].Level() {
		case menu.MainLevel:
			if !tb.main.UpdateMessage(msg, bot, tb.level[msg.From.UserName]) {
				return
			} else {
				return
			}

		case menu.ModLevel:
			if !tb.mod.UpdateMessage(msg, bot, tb.level[msg.From.UserName]) {
				tb.level[msg.From.UserName].SetLevel(menu.MainLevel)
			} else {
				return
			}
			break

		case menu.RelayLevel:
			if !tb.relay.UpdateMessage(msg, bot, tb.level[msg.From.UserName]) {
				tb.level[msg.From.UserName].SetLevel(menu.RelayLevel)
			} else {
				return
			}
			break
		}
	}
}

// ChangeLevel change menu level handler
func (tb *TgBot) ChangeLevel(msg *tgbotapi.Message) {
	var m = tb.level[msg.From.UserName]

	switch m.Level() {
	case menu.MainLevel:
		m.SetBuilding(msg.Text)
		m.SetLevel(menu.ModLevel)
		break

	//
	// Modules level
	//

	case menu.ModLevel:
		switch msg.Text {
		case "Назад":
			m.SetLevel(menu.MainLevel)
			break

		case "Розетки":
			m.SetLevel(menu.RelayLevel)
			break
		}

	//
	// Relay levels
	//

	case menu.RelayLevel:
		switch msg.Text {
		case "Обновить":
			break
		case "Назад":
			m.SetLevel(menu.ModLevel)
			break
		}
		break
	}

}

// SendError Send error tg message
func (tb *TgBot) SendError(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string) {
	var imgMsg = tgbotapi.NewPhotoShare(message.Chat.ID, "https://www.chip.pl/uploads/2009/02/5217f8b6060d0ac29f2bc1f143a02485.png")
	imgMsg.Caption = text
	imgMsg.ChatID = message.Chat.ID
	bot.Send(imgMsg)
}
