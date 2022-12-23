package main

import "time"

type MsgDetail struct {
	MsgText string
	UserID  string
	Time    time.Time
}

type GroupData []MsgDetail

type GroupStorage map[string]GroupData

func (m MsgDetail) SaveMsg(text string, uid string, time time.Time) {
	m.MsgText = text
	m.UserID = uid
	m.Time = time
}
