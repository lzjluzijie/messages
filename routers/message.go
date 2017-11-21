package routers

import (
	"encoding/json"

	"github.com/lzjluzijie/messages/models"
	macaron "gopkg.in/macaron.v1"
)

//PostMessage creates new messages.
func PostMessage(c *macaron.Context) {
	text := c.Query("message")
	if text == "" {
		c.Write([]byte("Message is empty!"))
		return
	}

	name := c.Query("name")
	ip := c.RemoteAddr()

	models.InsertMessage(text, name, ip)
	c.Write([]byte("Submitted successfully. Thank you!"))
	return
}

//GetMessage shows all messages in the database.
func GetMessage(c *macaron.Context) {
	key := c.Query("key")
	if key != "YourKey" {
		c.Write([]byte("404 not found."))
		return
	}

	msgs, err := models.GetMessage()
	if err != nil {
		c.Write([]byte(err.Error()))
		return
	}

	j, _ := json.Marshal(msgs)
	c.Write(j)
}
