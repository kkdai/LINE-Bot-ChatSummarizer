package main

import "time"

type GroupDB interface {
	ReadGroupInfo(string) GroupData
	AppendGroupInfo(string, MsgDetail)
}
type MsgDetail struct {
	MsgText  string
	UserName string
	Time     time.Time
}

type GroupData []MsgDetail
