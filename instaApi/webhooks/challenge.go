package webhooks

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RunChallengeWebhook(e *echo.Echo) {

    e.GET("/webhook/messaging-webhook", ChallengeHandler)
}

const APP_TOKEN = "b497b303-c2da-4963-95b8-1b34ce6fa993"

func ChallengeHandler(c echo.Context) error {
    // log.Println("Challenge handler")
    // log.Println(c.Request().Method)

    // allNames := c.QueryParams()


    // log.Println(allNames)

    log.Println("Challenge recebido")

    mode := c.QueryParam("hub.mode")
    token := c.QueryParam("hub.verify_token")
    challenge := c.QueryParam("hub.challenge")

    log.Println(mode, token, challenge)

    if mode != "" && token != "" {
        if mode == "subscribe" && token == APP_TOKEN {
            log.Println("Código devolvido: ", challenge)
            return c.String(http.StatusOK, challenge)
        }
    }

    return c.NoContent(http.StatusNotFound)
}