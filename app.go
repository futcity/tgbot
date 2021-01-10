///////////////////////////////////////////////////////////////////
//
// Future City Project
//
// Copyright (C) 2020-2021 Sergey Denisov. GPLv3
//
// Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
//
///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/futcity/tgbot/auth"
	"github.com/futcity/tgbot/configs"
	"github.com/futcity/tgbot/net"
	"github.com/futcity/tgbot/utils"
)

// App Application main module
type App struct {
	log *utils.Log
	cfg *utils.Configs
	ath *auth.Authorization
	bot *net.TgBot
}

// NewApp Make new struct
func NewApp(l *utils.Log, c *utils.Configs, ath *auth.Authorization, bot *net.TgBot) *App {
	return &App{
		log: l,
		cfg: c,
		ath: ath,
		bot: bot,
	}
}

// Start applications
func (a *App) Start() {
	//
	// Init logger
	//
	a.log.SetPath("./")

	//
	// Loading configs
	//
	var ac = &configs.AppCfg{}
	var uc = &configs.UsersCfg{}
	var bc = &configs.BldsCfg{}

	var err = a.cfg.LoadFromFile(ac, "tgbot.conf")
	if err != nil {
		a.log.Error("APP", "Fail to load tgbot configs", err.Error())
		return
	}
	err = a.cfg.LoadFromFile(uc, "users.conf")
	if err != nil {
		a.log.Error("APP", "Fail to load users configs", err.Error())
		return
	}
	err = a.cfg.LoadFromFile(bc, "buildings.conf")
	if err != nil {
		a.log.Error("APP", "Fail to load buildings configs", err.Error())
		return
	}

	//
	// Applying configs
	//
	for _, bld := range bc.Buildings {
		a.ath.AddBuilding(auth.NewBuilding(bld.Name, bld.IP, bld.Port))
		a.log.Info("APP", "Add new building \""+bld.Name+"\"")
	}

	for _, usr := range uc.Users {
		a.log.Info("APP", "Add new user login \""+usr.Login+"\"")

		var u = auth.NewUser(usr.Login, usr.Key)
		for _, bld := range usr.Buildings {
			u.AddBuilding(bld)
			a.log.Info("APP", "Add new user building \""+bld+"\"")
		}
		a.ath.AddUser(u)
	}

	//
	// Free configs
	//
	uc = nil
	bc = nil

	//
	// Starting server
	//
	a.log.Info("APP", "Starting bot...")
	err = a.bot.StartBot(ac.Key)
	if err != nil {
		a.log.Error("APP", "Fail to start bot", err.Error())
	}
}
