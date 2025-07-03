#!/bin/bash

set -a 
source ../.env
set +a

#get informations
curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/getUpdates" | jq