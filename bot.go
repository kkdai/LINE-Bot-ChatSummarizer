package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// Handle only on text message
			case *linebot.TextMessage:
				// 預設訊息
				reply := "msg ID:" + message.ID + ":" + "Get:" + message.Text + " , \n OK!"

				// 如果聊天機器人在群組中，開始儲存訊息。
				if isGroupEvent(event) {
					// 先取得使用者 Display Name (也就是顯示的名稱)
					userName := event.Source.UserID
					userProfile, err := bot.GetProfile(event.Source.UserID).Do()
					if err == nil {
						userName = userProfile.DisplayName
					}

					// event.Source.GroupID 就是聊天群組的 ID，並且透過聊天群組的 ID 來放入 Map 之中。
					m := MsgDetail{
						MsgText:  message.Text,
						UserName: userName,
						Time:     time.Now(),
					}
					summaryQueue.AppendGroupInfo(getGroupID(event), m)
				}

				// Directly to ChatGPT
				if strings.Contains(message.Text, ":gpt") {
					reply = gptCompleteContext(message.Text)

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				} else if strings.EqualFold(message.Text, ":list_all") {
					q := summaryQueue.ReadGroupInfo(getGroupID(event))
					for _, m := range q {
						reply = reply + fmt.Sprintf("[%s]: %s . %s\n", m.UserName, m.MsgText, m.Time.Local().UTC().Format("2006-01-02 15:04:05"))
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				} else if strings.EqualFold(message.Text, ":sum_all") {
					// 把聊天群組裡面的訊息都捲出來（依照先後順序）
					oriContext := ""
					q := summaryQueue.ReadGroupInfo(getGroupID(event))
					for _, m := range q {
						// [xxx]: 他講了什麼... 時間
						oriContext = oriContext + fmt.Sprintf("[%s]: %s . %s\n", m.UserName, m.MsgText, m.Time.Local().UTC().Format("2006-01-02 15:04:05"))
					}

					// 取得使用者暱稱
					userName := event.Source.UserID
					userProfile, err := bot.GetProfile(event.Source.UserID).Do()
					if err == nil {
						userName = userProfile.DisplayName
					}

					// 訊息內先回，再來總結。
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("好的，總結文字已經發給您了"+userName)).Do(); err != nil {
						log.Print(err)
					}

					// 就是請 ChatGPT 幫你總結
					oriContext = fmt.Sprintf("幫我總結 `%s`", oriContext)
					reply = gptCompleteContext(oriContext)

					// 因為 ChatGPT 可能會很慢，所以這邊後來用 SendMsg 來發送私訊給使用者。
					if _, err = bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				}

			// Handle only on Sticker message
			case *linebot.StickerMessage:
				var kw string
				for _, k := range message.Keywords {
					kw = kw + "," + k
				}

				if isGroupEvent(event) {
					// 在群組中，一樣紀錄起來不回覆。
					outStickerResult := fmt.Sprintf("貼圖訊息: %s ", kw)
					userName := event.Source.UserID
					userProfile, err := bot.GetProfile(event.Source.UserID).Do()
					if err == nil {
						// 使用者有加 OA 為好友，反之將拿不到 Profile
						userName = userProfile.DisplayName
					}
					m := MsgDetail{
						MsgText:  outStickerResult,
						UserName: userName,
						Time:     time.Now(),
					}
					summaryQueue.AppendGroupInfo(getGroupID(event), m)
				} else {
					outStickerResult := fmt.Sprintf("貼圖訊息: %s, pkg: %s kw: %s  text: %s", message.StickerID, message.PackageID, kw, message.Text)

					// 1 on 1 就回覆
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outStickerResult)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}

func textHander(event *linebot.Event, message string) {

}

func isGroupEvent(event *linebot.Event) bool {
	return event.Source.GroupID != "" || event.Source.RoomID != ""
}

func getGroupID(event *linebot.Event) string {
	if event.Source.GroupID != "" {
		return event.Source.GroupID
	} else if event.Source.RoomID != "" {
		return event.Source.RoomID
	}

	return ""
}
