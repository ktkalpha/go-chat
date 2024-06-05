package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

type message struct {
	User string `json:"user"`
	Data string `json:"data"`
}

var messages = []message{
	{User: "koko", Data: "i like chicken"}, {User: "momo", Data: "i like pizza"},
}

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages)
}
func postMessages(c *gin.Context) {
	var newMsg message
	if err := c.BindJSON(&newMsg); err != nil {
		return
	}
	messages = append(messages, newMsg)
	if len(messages) > 10 {
		messages = messages[1:]
	}
	c.IndentedJSON(http.StatusCreated, newMsg)
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST"},
		MaxAge:       12 * time.Hour,
	}))
	router.GET("/messages", getMessages)
	router.POST("/messages", postMessages)

	router.Run("localhost:8080")
}
