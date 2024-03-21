// Copyright 2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Teonet telegram bot module

package main

import (
	"errors"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	waitCommand = iota // Wait a command state
	waitToken          // Wait login token state
)

// Bot receiver
type Bot struct {
	*tgbotapi.BotAPI
	BotState
	teo *Teonet
}

// NewBot create new telegram bot
func NewBot(token string, teo *Teonet) (bot *Bot, err error) {
	bot = new(Bot)

	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		err = errors.New("can't create telegram bot, err: " + err.Error())
		return
	}
	bot.BotState.m = make(map[string]BotStateData)
	bot.BotAPI = b
	bot.teo = teo
	// b.Debug = true

	log.Printf("Authorized on account %s", b.Self.UserName)

	return
}

// Run receive and process users messages
func (b *Bot) Run() {
	conf := tgbotapi.NewUpdate(0)
	conf.Timeout = 30

	updates := b.GetUpdatesChan(conf)
	// com := &Command{b.teo}

	for update := range updates {
		// Ignore any non-Message Updates
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Commands arguments
		// args := update.Message.CommandArguments()

		// Send answer to message
		sendAnswer := func(update tgbotapi.Update, state int, text string, markdown ...bool) {
			// Set new state
			b.SetState(update.Message.From.UserName, state)

			// Send answer to user
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			if len(markdown) > 0 && markdown[0] {
				msg.ParseMode = tgbotapi.ModeMarkdownV2
			}
			// msg.ReplyToMessageID = update.Message.MessageID
			b.Send(msg)
		}

		// Check command
		var state int
		var answer string
		switch update.Message.Command() {
		// fortune command
		case "fortune":
			var err error
			answer, err = b.teo.Fortune()
			state = waitCommand
			if err != nil {
				answer = err.Error()
			}

		// Check State argument, Text or Unknown command
		default:
			text := update.Message.Text
			// Check command state
			state = b.State(update.Message.From.UserName)
			switch state {
			// Set login token
			// case waitToken:
			// 	continue

			// Text or Unknown command
			default:
				answer = "I don't know anything about '" + text + "'.\n\n" +
					"I can tell you some fortune message:\n\n"
				msg, err := b.teo.Fortune()
				if err != nil {
					msg = err.Error()
				}
				answer += msg
			}
		}

		// Send answer to user and sent new user state
		sendAnswer(update, state, answer)
	}
}

// BotState is telegram bot state receiver
type BotState struct {
	m map[string]BotStateData
	sync.RWMutex
}

// BotStateData is BotState data structure
type BotStateData struct {
	state int
}

// SetState save users state
func (s *BotState) SetState(user string, state int) {
	s.Lock()
	defer s.Unlock()
	d := s.m[user]
	d.state = state
	s.m[user] = d
}

// State return users state
func (s *BotState) State(user string) (state int) {
	s.RLock()
	defer s.RUnlock()
	d, ok := s.m[user]
	if !ok {
		state = waitCommand
	}
	state = d.state
	return
}
