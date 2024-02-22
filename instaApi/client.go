package instaApi

import (
	"apiMessages/instaApi/on_message"
	"apiMessages/instaApi/webhooks"

	"github.com/labstack/echo/v4"
)

type Client struct {
	MessageEvents *on_message.OnMessageEvent
}

func NewClient(e *echo.Echo) Client {
	client := Client{
		MessageEvents: on_message.NewOnMessageEvent(e),
	}

	webhooks.RunChallengeWebhook(e)

	return client
}