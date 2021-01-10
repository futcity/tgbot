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

import "errors"

type Authorization struct {
	users     map[string]*User
	buildings map[string]*Building
}

func NewAuthorization() *Authorization {
	return &Authorization{
		users:     make(map[string]*User),
		buildings: make(map[string]*Building),
	}
}

func (a *Authorization) AddUser(usr *User) {
	a.users[usr.Login()] = usr
}

func (a *Authorization) User(login string) *User {
	return a.users[login]
}

func (a *Authorization) Users() []*User {
	var users []*User

	for _, usr := range a.users {
		users = append(users, usr)
	}

	return users
}

func (a *Authorization) AddBuilding(bld *Building) {
	a.buildings[bld.Name()] = bld
}

func (a *Authorization) Building(name string) (*Building, error) {
	var bld = a.buildings[name]
	if bld == nil {
		return nil, errors.New("Building not found")
	}
	return bld, nil
}
