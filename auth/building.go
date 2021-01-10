///////////////////////////////////////////////////////////////////
//
// Future City Project
//
// Copyright (C) 2020-2021 Sergey Denisov. GPLv3
//
// Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
//
///////////////////////////////////////////////////////////////////

package auth

type Building struct {
	name string
	ip   string
	port int
}

func NewBuilding(name string, ip string, port int) *Building {
	return &Building{
		name: name,
		ip:   ip,
		port: port,
	}
}

func (b *Building) Name() string {
	return b.name
}

func (b *Building) IP() string {
	return b.ip
}

func (b *Building) Port() int {
	return b.port
}
