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

type BuildCfg struct {
	Name string
	IP   string
	Port int
}

type BldsCfg struct {
	Buildings []BuildCfg
}
