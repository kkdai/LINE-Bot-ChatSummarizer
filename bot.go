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
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// Handle only on text message
			case *linebot.TextMessage:
				// Directly to ChatGPT
				if strings.Contains(message.Text, ":gpt") {
					// New feature.
					if IsRedemptionEnabled() {
						if stickerRedeemable {
							handleGPT(event, message.Text)
							stickerRedeemable = false
						} else {
							handleRedeemRequestMsg(event)
						}
					} else {
						// Original one
						handleGPT(event, message.Text)
					}
				} else if strings.EqualFold(message.Text, ":list_all") && isGroupEvent(event) {
					handleListAll(event)
				} else if strings.EqualFold(message.Text, ":sum_all") && isGroupEvent(event) {
					handleSumAll(event)
				} else if isGroupEvent(event) {
					// 如果聊天機器人在群組中，開始儲存訊息。
					handleStoreMsg(event, message.Text)
				}

			// Handle only on Sticker message
			case *linebot.StickerMessage:
				var kw string
				for _, k := range message.Keywords {
					kw = kw + "," + k
				}

				log.Println("Sticker: PID=", message.PackageID, " SID=", message.StickerID)
				if IsRedemptionEnabled() {
					if message.PackageID == RedeemStickerPID && message.StickerID == RedeemStickerSID {
						stickerRedeemable = true
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你的賦能功能啟動了！")).Do(); err != nil {
							log.Print(err)
						}
					}
				}

				if isGroupEvent(event) {
					// 在群組中，一樣紀錄起來不回覆。
					outStickerResult := fmt.Sprintf("貼圖訊息: %s ", kw)
					handleStoreMsg(event, outStickerResult)
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

func handleSumAll(event *linebot.Event) {
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
	reply := gptCompleteContext(oriContext)

	// 因為 ChatGPT 可能會很慢，所以這邊後來用 SendMsg 來發送私訊給使用者。
	if _, err = bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func handleListAll(event *linebot.Event) {
	reply := ""
	q := summaryQueue.ReadGroupInfo(getGroupID(event))
	for _, m := range q {
		reply = reply + fmt.Sprintf("[%s]: %s . %s\n", m.UserName, m.MsgText, m.Time.Local().UTC().Format("2006-01-02 15:04:05"))
	}

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func handleGPT(event *linebot.Event, message string) {
	reply := gptCompleteContext(message)

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func handleRedeemRequestMsg(event *linebot.Event) {
	// 先取得使用者 Display Name (也就是顯示的名稱)
	userName := event.Source.UserID
	userProfile, err := bot.GetProfile(event.Source.UserID).Do()
	if err == nil {
		userName = userProfile.DisplayName
	}

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(userName+":你需要買貼圖，開啟這個功能"), linebot.NewStickerMessage(RedeemStickerPID, RedeemStickerSID)).Do(); err != nil {
		log.Print(err)
	}
}

func handleStoreMsg(event *linebot.Event, message string) {
	// 先取得使用者 Display Name (也就是顯示的名稱)
	userName := event.Source.UserID
	userProfile, err := bot.GetProfile(event.Source.UserID).Do()
	if err == nil {
		userName = userProfile.DisplayName
	}

	// event.Source.GroupID 就是聊天群組的 ID，並且透過聊天群組的 ID 來放入 Map 之中。
	m := MsgDetail{
		MsgText:  message,
		UserName: userName,
		Time:     time.Now(),
	}
	summaryQueue.AppendGroupInfo(getGroupID(event), m)
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
