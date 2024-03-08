package main

import (
	"github.com/rs/zerolog/log"

	"company.com/bridge/messaging"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	var sender messaging.MessageSender = messaging.CreateMessageSender("sms")
	var formatter messaging.MessageFormatter = messaging.CreateMessageFormatter("text")
	var messageService *messaging.MessageService = messaging.CreateMessageService(sender, formatter)

	var message = "Test Message"

	messageService.Send(message)

	sender = messaging.CreateMessageSender("email")
	formatter = messaging.CreateMessageFormatter("html")
	messageService = messaging.CreateMessageService(sender, formatter)

	message = "Test Message"

	messageService.Send(message)

	sender = messaging.CreateMessageSender("sms")
	formatter = messaging.CreateMessageFormatter("html")
	messageService = messaging.CreateMessageService(sender, formatter)

	message = "Test Message"

	messageService.Send(message)
}
