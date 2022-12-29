package main

import "time"

type MsgDetail struct {
	MsgText  string
	UserName string
	Time     time.Time
}

type GroupData []MsgDetail

type GroupStorage map[string]GroupData
