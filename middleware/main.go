package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type subscription struct {
	Object  string  `json:"object"`
	Entries []entry `json:"entry"`
}

type entry struct {
	Messaging []messageEvent `json:"messaging"`
}

type messageEvent struct {
	Message    string `json:"message"`
	SenderData sender `json:"sender"`
}

type sender struct {
	ID string `json:"id"`
}

type response struct {
	Message   responseMessage `json:"message"`
	Recipient recipient       `json:"recipient"`
}

type recipient struct {
	ID string
}

type responseMessage struct {
	Text string
}

// Home is a basic return message for home
func Home(c echo.Context) error {
	message := "Hello World!"
	return c.String(http.StatusOK, message)
}

// Validate verifies that the verification token is authenticate for a webhook GET request
func Validate(c echo.Context) error {
	token := os.Getenv("TOKEN")

	mode := c.QueryParam("hub.mode")
	verifyToken := c.QueryParam("hub.verify_token")
	challenge := c.QueryParam("hub.challenge")

	fmt.Println("Query", c.QueryString())

	if mode != "" && verifyToken != "" {
		if mode == "subscribe" && verifyToken == token {
			fmt.Println("TOKEN VERIFIED")
			return c.String(http.StatusOK, challenge)
		}

		return c.String(http.StatusForbidden, "Verify token does not match")
	}

	return c.String(http.StatusBadRequest, "Bad request")
}

// Translate translates and parses a webhook POST request
func Translate(c echo.Context) error {
	s := new(subscription)

	if err := c.Bind(s); err != nil {
		fmt.Println(http.StatusBadRequest, "Bind failure")
		return c.String(http.StatusBadRequest, "Bad request")
	}

	fmt.Println("SUBSCRIPTION", s)
	fmt.Println("OBJECT", s.Object)

	// check if event is from a page subscription
	if s.Object == "page" {
		for _, e := range s.Entries {
			event := e.Messaging[0]

			ProcessMessage(event)
		}

		return c.String(http.StatusOK, "EVENT_RECEIEVED")
	}

	return c.String(http.StatusBadRequest, "Bad request")
}

// ProcessMessage processes message event data
func ProcessMessage(msgEvent messageEvent) {
	API := os.Getenv("FACEBOOK_API")
	TOKEN := os.Getenv("PAGE_ACCESS_TOKEN")
	msg := msgEvent.Message

	if msg == "" {
		log.Fatal("Received message is blank")
	}

	rmsg := responseMessage{Text: "Hello World!"}
	rcp := recipient{ID: msgEvent.SenderData.ID}

	resp := response{
		Message:   rmsg,
		Recipient: rcp,
	}

	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(&resp)

	if err != nil {
		log.Fatal(err)
	}

	url := API + "?access_token=" + TOKEN

	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR", err)
	}

	fmt.Println("RESPONSE", res)
}
