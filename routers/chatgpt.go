package routers

import (
	"fmt"

	"github.com/Wind-318/wind-chimes/openai"
	"github.com/gin-gonic/gin"
)

func Chatgpt(c *gin.Context) {
	chat := &openai.Chat{}
	chat.SetAuthorizationKey("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")
	chat.AddMessageAsSystem("You are azur lane akashi.")
	chat.AddMessageAsUser("Hello akashi, introduce yourself.")

	resp, err := chat.NewChat()
	if err != nil {
		// Handle error
	}

	for _, choice := range resp.Choices {
		fmt.Println(choice.Msg.Content)
	}
}
