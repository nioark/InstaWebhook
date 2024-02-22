package on_message

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
)

type InstagramMessage struct {
    Object string   `json:"object"`
    Entry  []Entry  `json:"entry"`
}

type Entry struct {
    Time      int64       `json:"time"`
    ID        string      `json:"id"`
    Messaging []Messaging `json:"messaging"`
}

type Sender struct {
    ID string `json:"id"`
}

type Recipient struct {
    ID string `json:"id"`
}

type Messaging struct {
    Sender    Sender    `json:"sender"`
    Recipient Recipient `json:"recipient"`
    Timestamp string    `json:"timestamp"`
    Message   Message   `json:"message"`
    Reaction   Reaction   `json:"reaction"`
    Read Read `json:"read"`

}

type Payload struct {
    URL string `json:"url"`   
}

type Attachment struct {
    Type string `json:"type"`   
    Payload Payload `json:"payload"`
}

type Story struct {
    URL string `json:"url"`   
    ID string `json:"id"`
}

type ReplyTo struct {
    Mid string `json:"mid"`   
    Story Story
}

type Message struct {
    Mid  string `json:"mid"`
    Text string `json:"text"`
    Attachments []Attachment `json:"attachments"`
    IsDeleted bool `json:"is_deleted"`
    IsEcho bool `json:"is_echo"`
    IsUnsupported bool `json:"is_unsupported"`
    ReplyTo ReplyTo `json:"reply_to"`
}

type Reaction struct {
    Mid  string `json:"mid"`   
    Action string `json:"action"`
    Reaction string `json:"reaction"`
    Emoji string `json:"emoji"`
}

type Read struct {
    Mid string `json:"mid"`   
}

type messageHandler struct {
    echo *echo.Echo   
    events *OnMessageEvent
}

func runMessageHook(e *echo.Echo, event *OnMessageEvent ) {
    messageHandler := messageHandler{
        echo: e,
        events: event,
    }

    e.POST("/webhook/messaging-webhook", messageHandler.MessageHandler)   
}

func (h *messageHandler) MessageHandler(c echo.Context) error {
    log.Println("Message handler event")
    log.Println(c.Request().Method)

    buf := new(bytes.Buffer)
    buf.ReadFrom(c.Request().Body)
    bodyBytes := buf.String()

    var result InstagramMessage
    json.Unmarshal([]byte(bodyBytes), &result)

    eventMessage := result.Entry[0].Messaging[0]

    if (!reflect.DeepEqual(eventMessage.Message, Message{})){
        log.Println("Evento -> mensagem")

        payload := OnMessagePayload{
            SenderID: eventMessage.Sender.ID,
            RecipientID: eventMessage.Recipient.ID,
            Timestamp: time.Unix(result.Entry[0].Time, 0),
            Message: eventMessage.Message,
        }

        h.events.TriggerMessage(payload)
    } else if (eventMessage.Reaction != Reaction{}){
        log.Println("Evento -> reaction")
        log.Println(eventMessage.Reaction)

        payload := OnReactionPayload{
            SenderID: eventMessage.Sender.ID,
            RecipientID: eventMessage.Recipient.ID,
            Timestamp: time.Unix(result.Entry[0].Time, 0),
            Reaction: eventMessage.Reaction,
        }

        h.events.TriggerReaction(payload)
    } else if (eventMessage.Read != Read{}){
        log.Println("Evento -> lido")
        log.Println(eventMessage.Read)

        payload := OnReadPayload{
            SenderID: eventMessage.Sender.ID,
            RecipientID: eventMessage.Recipient.ID,
            Timestamp: time.Unix(result.Entry[0].Time, 0),
            Read: eventMessage.Read,
        }

        h.events.TriggerRead(payload)
    }

    return c.String(http.StatusOK, "EVENT_RECEIVED")


    // Check if the webhook object is "page"
    // if body.Object == "page" {
    //     // Send a 200 OK response
    //     return c.String(http.StatusOK, "EVENT_RECEIVED")
    //     return c.String(http.StatusOK, "EVENT_RECEIVED")


    //     // Determine which webhooks were triggered and get sender PSIDs and locale, message content, and more.
    //     // Add your logic here
    // } else {
    //     // Return a '404 Not Found' if the event is not from a page subscription
    //     return c.NoContent(http.StatusNotFound)
    // }
}