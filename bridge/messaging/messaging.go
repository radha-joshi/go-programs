package messaging

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Message struct {
	formattedText string
}

type MessageSender interface {
	Send(message *Message)
}

type smsSender struct {
}

func (s *smsSender) Send(message *Message) {
	log.Info().Msg(fmt.Sprintf("Sending by SMS: %s", message.formattedText))
}

type emailSender struct {
}

func (s *emailSender) Send(message *Message) {
	log.Info().Msg(fmt.Sprintf("Sending by Email: %s", message.formattedText))
}

func CreateMessageSender(senderType string) MessageSender {
	if senderType == "sms" {
		return &smsSender{}
	} else {
		return &emailSender{}
	}
}

type MessageFormatter interface {
	FormatMessage(message *Message)
}

type textFormatter struct {
}

func (f *textFormatter) FormatMessage(message *Message) {
	log.Info().Msg("Applying Text Format")
}

type htmlFormatter struct {
}

func (f *htmlFormatter) FormatMessage(message *Message) {
	log.Info().Msg("Applying HTML Format")
}

func CreateMessageFormatter(formatType string) MessageFormatter {
	if formatType == "text" {
		return &textFormatter{}
	} else {
		return &htmlFormatter{}
	}
}

type MessageService struct {
	sender    MessageSender
	formatter MessageFormatter
}

func (ms *MessageService) Send(text string) {
	log.Info().Msg("Sending message")
	message := &Message{
		formattedText: text,
	}
	ms.formatter.FormatMessage(message)
	ms.sender.Send(message)
}

func CreateMessageService(sender MessageSender, formatter MessageFormatter) *MessageService {
	return &MessageService{
		sender:    sender,
		formatter: formatter,
	}
}
