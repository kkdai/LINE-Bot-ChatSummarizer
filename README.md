LINE Bot 聊天摘要生成器：使用 ChatGPT 將你的群聊作為聊天摘要生成器，與 LINE Bot 聊天摘要生成器一起使用
==============

[![Join the chat at https://gitter.im/kkdai/LINE-Bot-ChatSummarizer](https://badges.gitter.im/kkdai/LINE-Bot-ChatSummarizer.svg)](https://gitter.im/kkdai/LINE-Bot-ChatSummarizer?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![GoDoc](https://godoc.org/github.com/kkdai/LINE-Bot-ChatSummarizer.svg?status.svg)](https://godoc.org/github.com/kkdai/LINE-Bot-ChatSummarizer)  ![Go](https://github.com/kkdai/LINE-Bot-ChatSummarizer/workflows/Go/badge.svg) [![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/LINE-Bot-ChatSummarizer)](https://goreportcard.com/report/github.com/kkdai/LineBotTemplate)

English version, please check [ENG Version](README.en.md)

如何使用
=============

### 獲取 LINE Bot API 開發者帳戶

如果你想使用 LINE Bot，請確保在 <https://developers.line.biz/console/> 註冊 LINE 開發者控制台。

在「基本設定」選項卡上創建新的消息通道並獲取「通道密鑰」。

在「消息 API」選項卡上獲取「通道訪問權杖」。

從「基本設定」選項卡中打開 LINE OA 管理器，然後轉到 OA 管理器的回復設定。在那裡啟用「webhook」。

### 獲取 OpenAI API 權杖

在 <https://openai.com/api/> 註冊帳戶。

一旦你有了帳戶，就可以在帳戶設定頁面找到你的 API 權杖。

如果你想在開發中使用 OpenAI API，你可以在 API 文檔頁面中找到更多信息和說明。

請注意，OpenAI API 只面向滿足某些條件的用戶開放。你可以在 API 文檔頁面中找到有關 API 的使用條件和限制的更多信息。

### 部署在 Heroku 上

在 <https://www.herokucdn.com/deploy/button.svg> 上點擊「部署」按鈕。

輸入「Channel Secret」、「Channel Access Token」和「ChatGPT Access Token」。

記住你的 Heroku 伺服器 ID。

在 LINE Bot 儀表板中設置基本 API：

設置你的基本帳戶信息，包括「回調 URL」在 <https://{YOUR_HEROKU_SERVER_ID}.herokuapp.com/callback>。

這就是它！你完成了。

License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

<http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
