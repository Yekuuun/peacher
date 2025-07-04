<p align="center">
  <img src="https://github.com/Yekuuun/peacher/blob/main/assets/peacher.webp" alt="Description de l'image" width="300">
</p>
<div align="center">
  <h1>Peacher</h1>
  <p><strong>A base POC for MYM & OF phishing. under construction...</strong></p>
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

<img src="https://github.com/Yekuuun/peacher/blob/main/assets/mym.png" alt="DebugInfo" />


<br>

### Built with ⚡

<img src="https://skillicons.dev/icons?i=css,html,go,docker,ngrok" />
