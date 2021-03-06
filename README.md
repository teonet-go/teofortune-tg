# Teonet fortune telegram-bot microservice

This is simple [Teonet](https://github.com/teonet-go/teonet) telegram-bot micriservice application which get fortune message from [Teonet Fortune](https://github.com/teonet-go/teofortune) microservice and show it in Telegram.

[![GoDoc](https://godoc.org/github.com/teonet-go/teofortune-tg?status.svg)](https://godoc.org/github.com/teonet-go/teofortune-tg/)
[![Go Report Card](https://goreportcard.com/badge/github.com/teonet-go/teofortune-tg)](https://goreportcard.com/report/github.com/teonet-go/teofortune-tg)

This microservice use telegram bot api. To create your telergam bot api token see [How to Create and Connect a Telegram Chatbot](https://sendpulse.com/knowledge-base/chatbot/create-telegram-chatbot). We have created the `@teofortune` Telegram bot and run it in Teonet Cloud.

## Run the Teonet fortune telegram-bot microservice

There are various ways to start this Teonet microservice application:

_In code blow we use preinstalled in Teonet Cloud teofortune microservice address: '-fortune=8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF'. 
Change this address to your application address. Or you can use this address, but than you will connect to `teofortune` microservice application running in Teonet Cloud. The address prints after you start Teonet application in string:_  
`Teonet address: 8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF`

### 1. From sources

```bash
git clone https://github.com/teonet-go/teofortune-tg
cd teofortune-tg
go run . -token=your-telegram-token -fortune=8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF -loglevel=debug
```

### 2. Install binary from github

```bash
go install github.com/teonet-go/teofortune-tg .
teofortune -token=your-telegram-token -fortune=8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF -loglevel=debug
```

### 3. Using docker

```bash
docker run -d -it --network=host --restart=always --name teofortune-tg -v \
$HOME/.config/teonet/teofortune-tg:/root/.config/teonet/teofortune-tg ghcr.io/teonet-go/\
teofortune-tg:latest teofortune-tg -token=your-telegram-token \
-fortune=8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF -loglevel=debug
```

## How to use

There is preinstalled teofortune-tg telegram bot with name @teofortune.
Open @teofortune in Telegram and write him something.

## Licence

[BSD](LICENSE)
