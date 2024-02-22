package on_message

import (
	"time"

	"github.com/labstack/echo/v4"
)

var MessageReceived OnMessageEvent

// UserCreatedPayload is the data for when a user is created
type OnMessagePayload struct {
    SenderID string
    RecipientID string
    Timestamp  time.Time
    Message Message
}

type OnReactionPayload struct {
    SenderID string
    RecipientID string
    Timestamp  time.Time
    Reaction Reaction
}

type OnReadPayload struct {
    SenderID string
    RecipientID string
    Timestamp  time.Time
    Read Read
}

type OnMessageEvent struct {
    messagesHandlers []func(OnMessagePayload)
    reactionHandlers []func(OnReactionPayload)
    readHandlers []func(OnReadPayload)

    echo *echo.Echo
}

// Register adds an event handler for this event
func (u *OnMessageEvent) OnMessage(handler func(OnMessagePayload)) {
    u.messagesHandlers = append(u.messagesHandlers, handler)
}

func (u *OnMessageEvent) OnReaction(handler func(OnReactionPayload)) {
    u.reactionHandlers = append(u.reactionHandlers, handler)
}

func (u *OnMessageEvent) OnRead(handler func(OnReadPayload)) {
    u.readHandlers = append(u.readHandlers, handler)
}

// Trigger sends out an event with the payload
func (u OnMessageEvent) TriggerMessage(payload OnMessagePayload) {
    for _, handler := range u.messagesHandlers {
        go handler(payload)
    }
}

func (u OnMessageEvent) TriggerReaction(payload OnReactionPayload) {
    for _, handler := range u.reactionHandlers {
        go handler(payload)
    }
}

func (u OnMessageEvent) TriggerRead(payload OnReadPayload) {
    for _, handler := range u.readHandlers {
        go handler(payload)
    }
}

func NewOnMessageEvent(e *echo.Echo) *OnMessageEvent {
    handler := &OnMessageEvent{
        messagesHandlers: make([]func(OnMessagePayload), 0),
        reactionHandlers: make([]func(OnReactionPayload), 0),
        readHandlers: make([]func(OnReadPayload), 0),

        echo: e,
    }

    runMessageHook(e, handler)

    return handler
}

