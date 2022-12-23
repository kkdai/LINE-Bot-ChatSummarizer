package main

import "time"

type MsgDetail struct {
	MsgText string
	UserID  string
	Time    time.Time
}

type GroupData []MsgDetail

type GroupStorage map[string]GroupData
