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

	"github.com/line/line-bot-sdk-go/v7/linebot"
	gpt3 "github.com/sashabaranov/go-openai"
)

var bot *linebot.Client
var client *gpt3.Client
var summaryQueue GroupDB
var stickerRedeemable bool
var enableRedeem string

const RedeemStickerPID = "789"
const RedeemStickerSID = "10856"

func main() {
	stickerRedeemable = false
	var err error

	// Enable new feature (YES, default no)
	enableRedeem = os.Getenv("REDEEM_ENABLE")

	//  如果有預設 DABTASE_URL 就建立 PostGresSQL; 反之則建立 Mem DB
	pSQL := os.Getenv("DATABASE_URL")
	if pSQL != "" {
		summaryQueue = NewPGSql(pSQL)
	} else {
		summaryQueue = NewMemDB()
	}

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

func IsRedemptionEnabled() bool {
	return enableRedeem == "YES"
}
