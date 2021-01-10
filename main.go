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
	"fmt"

	"github.com/futcity/tgbot/auth"
	"github.com/futcity/tgbot/net"
	"github.com/futcity/tgbot/net/menu"
	"github.com/futcity/tgbot/utils"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	container.Provide(utils.NewLog)
	container.Provide(utils.NewConfigs)

	container.Provide(auth.NewAuthorization)
	container.Provide(menu.NewMainMenu)
	container.Provide(menu.NewModMenu)
	container.Provide(menu.NewRelayMenu)
	container.Provide(net.NewTgBot)

	container.Provide(NewApp)

	err := container.Invoke(func(app *App) {
		app.Start()
	})

	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
	}
}
