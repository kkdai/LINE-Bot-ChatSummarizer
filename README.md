LINE Bot Chat Summarizer: Use ChatGPT to summarize your group chat as a chat summarizer with the LINE Bot Chat Summarizer.

==============

[![Join the chat at https://gitter.im/kkdai/LINE-Bot-ChatSummarizer](https://badges.gitter.im/kkdai/LINE-Bot-ChatSummarizer.svg)](https://gitter.im/kkdai/LINE-Bot-ChatSummarizer?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

 [![GoDoc](https://godoc.org/github.com/kkdai/LINE-Bot-ChatSummarizer.svg?status.svg)](https://godoc.org/github.com/kkdai/LINE-Bot-ChatSummarizer)  

 ![Go](https://github.com/kkdai/LINE-Bot-ChatSummarizer/workflows/Go/badge.svg)

[![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/LINE-Bot-ChatSummarizer)](https://goreportcard.com/report/github.com/kkdai/LineBotTemplate)

How to use this
=============

### 1. Got A LINE Bot API devloper account

- [Make sure you already registered on LINE developer console](https://developers.line.biz/console/), if you need use LINE Bot.

- Create new Messaging Channel
- Get `Channel Secret` on "Basic Setting" tab.
- Issue `Channel Access Token` on "Messaging API" tab.
- Open LINE OA manager from "Basic Setting" tab.
- Go to Reply setting on OA manager, enable "webhook"

### 2. Just Deploy this on Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

- Input `Channel Secret` and `Channel Access Token`.
- Remember your heroku, ID.

### 3. Go to LINE Bot Dashboard, setup basic API

- Setup your basic account information. Here is some info you will need to know.
- `Callback URL`: <https://{YOUR_HEROKU_SERVER_ID}.herokuapp.com/callback>

It all done.

### Video Tutorial

- [How to deploy LINE BotTemplate](https://www.youtube.com/watch?v=0BIknEz1f8k)
- [Hoe to modify your LINE BotTemplate code](https://www.youtube.com/watch?v=ckij73sIRik)

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
