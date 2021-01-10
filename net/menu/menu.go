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

const (
	MainLevel = iota
	ModLevel
	RelayLevel
)

type MenuLevel struct {
	level    int
	device   string
	building string
}

func (m *MenuLevel) Device() string {
	return m.device
}

func (m *MenuLevel) SetDevice(device string) {
	m.device = device
}

func (m *MenuLevel) Level() int {
	return m.level
}

func (m *MenuLevel) SetLevel(level int) {
	m.level = level
}

func (m *MenuLevel) SetBuilding(building string) {
	m.building = building
}

func (m *MenuLevel) Building() string {
	return m.building
}
