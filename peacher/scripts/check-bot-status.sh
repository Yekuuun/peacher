#!/bin/bash

#loading environment variables.
set -a
source ../.env
set +a

#calling telegram api.
echo "TELEGRAM BOT STATUS :"
curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/getMe" | jq -r '.ok'