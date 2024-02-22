package main

import (
	// "encoding/json"

	"fmt"

	"github.com/labstack/echo/v4"

	"apiMessages/instaApi"
	"apiMessages/instaApi/on_message"
)


func main() {
    // Create a new Echo instance
    e := echo.New()

    client := instaApi.NewClient(e)

    client.MessageEvents.OnMessage(onMessageEvent)
    client.MessageEvents.OnReaction(onReactionEvent)
    client.MessageEvents.OnRead(onReadEvent)

    // Start the server on port 8080
    e.Logger.Fatal(e.Start(":8080"))
}

func onMessageEvent(payload on_message.OnMessagePayload) {
    fmt.Printf("%#v\n", payload)
}

func onReactionEvent(payload on_message.OnReactionPayload) {
    fmt.Printf("%#v\n", payload)
}

func onReadEvent(payload on_message.OnReadPayload) {
    fmt.Printf("%#v\n", payload)
}
