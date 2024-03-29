// Copyright 2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Teonet fortune telegram-bot microservice. This is simple Teonet telegram-bot
// micriservice application which get fortune message from Teonet Fortune
// microservice and show it in Telegram.
package main

import (
	"flag"
	"log"
	"time"

	"github.com/teonet-go/teonet"
)

const (
	appShort   = "teofortune-tg"
	appName    = "Teonet fortune telegram-bot microservice application"
	appLong    = ""
	appVersion = "0.6.1"
)

var appStartTime = time.Now()
var token, fortune, monitor string

// Params is teonet command line parameters
var Params struct {
	appShort    string
	port        int
	stat        bool
	hotkey      bool
	showPrivate bool
	loglevel    string
	logfilter   string
}

func main() {

	// Application logo
	teonet.Logo(appName, appVersion)

	// Parse application command line parameters
	flag.StringVar(&Params.appShort, "name", appShort, "application short name")
	flag.IntVar(&Params.port, "p", 0, "local port")
	flag.BoolVar(&Params.stat, "stat", false, "show statistic")
	flag.BoolVar(&Params.hotkey, "hotkey", false, "start hotkey menu")
	flag.BoolVar(&Params.showPrivate, "show-private", false, "show private key")
	flag.StringVar(&Params.loglevel, "loglevel", "NONE", "log level")
	flag.StringVar(&Params.logfilter, "logfilter", "", "log filter")
	//
	flag.StringVar(&token, "token", "", "telegram token")
	flag.StringVar(&fortune, "fortune", "", "fortune microservice address")
	flag.StringVar(&monitor, "monitor", "", "monitor address")
	//
	flag.Parse()

	// Check requered parameters
	teonet.CheckRequeredParams("token", "fortune")

	// Initialize and run Teonet
	teo, err := newTeonet()
	if err != nil {
		log.Panic(err)
		return
	}

	// Initialize telegram bot
	bot, err := NewBot(token, teo)
	if err != nil {
		log.Panic(err)
		return
	}

	// Run bot
	bot.Run()
}
