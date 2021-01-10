///////////////////////////////////////////////////////////////////
//
// Future City Project
//
// Copyright (C) 2020-2021 Sergey Denisov. GPLv3
//
// Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
//
///////////////////////////////////////////////////////////////////

package configs

type UserCfg struct {
	Login     string
	Key       string
	Buildings []string
}

type UsersCfg struct {
	Users []UserCfg
}
