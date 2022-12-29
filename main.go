// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	gpt3 "github.com/sashabaranov/go-gpt3"
)

var bot *linebot.Client
var client *gpt3.Client
var summaryQueue GroupStorage

func main() {
	var err error
	summaryQueue = make(GroupStorage)
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)

	port := os.Getenv("PORT")
	apiKey := os.Getenv("ChatGptToken")

	if apiKey != "" {
		client = gpt3.NewClient(apiKey)
	}

	http.HandleFunc("/callback", callbackHandler)
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

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
				reply := "msg ID:" + message.ID + ":" + "Get:" + message.Text + " , \n OK!"

				// If chatbot in a group, start to save string
				if event.Source.GroupID != "" {
					userName := event.Source.UserID
					userProfile, err := bot.GetProfile(event.Source.UserID).Do()
					if err != nil {
						userName = userProfile.DisplayName
					}

					q := summaryQueue[event.Source.GroupID]
					m := MsgDetail{
						MsgText:  message.Text,
						UserName: userName,
						Time:     time.Now(),
					}
					log.Println("Save msg:", m)
					summaryQueue[event.Source.GroupID] = append(q, m)
					log.Println("All msg:", q)
				}

				// Directly to ChatGPT
				if strings.Contains(message.Text, "gpt:") {
					reply = CompleteContext(message.Text)

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				} else if strings.EqualFold(message.Text, ":list_all") {
					q := summaryQueue[event.Source.GroupID]
					for _, m := range q {
						reply = reply + fmt.Sprintf("[%s]: %s . %s\n", m.UserName, m.MsgText, m.Time.Local().UTC().Format("2006-01-02 15:04:05"))
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				} else if strings.EqualFold(message.Text, ":sum_all") {
					oriContext := ""
					q := summaryQueue[event.Source.GroupID]
					for _, m := range q {
						oriContext = oriContext + fmt.Sprintf("[%s]: %s . %s\n", m.UserName, m.MsgText, m.Time.Local().UTC().Format("2006-01-02 15:04:05"))
					}
					oriContext = fmt.Sprintf("幫我總結 `%s`", oriContext)
					reply = CompleteContext(oriContext)

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ori"+reply), linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				}

			// Handle only on Sticker message
			case *linebot.StickerMessage:
				var kw string
				for _, k := range message.Keywords {
					kw = kw + "," + k
				}

				outStickerResult := fmt.Sprintf("收到貼圖訊息: %s, pkg: %s kw: %s  text: %s", message.StickerID, message.PackageID, kw, message.Text)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outStickerResult)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
