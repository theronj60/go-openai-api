package Controllers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type ChatPrompt struct {
	Question string `form:"question" binding:"required"`
}

func HomeHandler(c *gin.Context) {
	c.String(200, "Success")

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome home buddy",
	})
}

func GetFinancialHandler(c *gin.Context) {
	var chatPrompt ChatPrompt
	// chatPrompt.Question = c.PostForm("question")
	if err := c.Bind(&chatPrompt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if c.ShouldBindJSON(&chatPrompt) == nil {
	// 	log.Println(chatPrompt.Question)
	// }

	log.Println("question: " + chatPrompt.Question)
	openai_key := os.Getenv("OPENAI_KEY")
	client := openai.NewClient(openai_key)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are a helpful financial advisor for native americans who answers questions about college and scholarships.",
			},
			{
				Role:    "system",
				Content: "Your priority is to help native americans.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: chatPrompt.Question,
			},
		},
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		// return
	}

	defer stream.Close()

	fmt.Printf("Stream response: ")
	var ai_response []string
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			// c.String(200, "Success")
			c.JSON(http.StatusOK, gin.H{
				"message": strings.Join(ai_response, ""),
			})
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		ai_response = append(ai_response, response.Choices[0].Delta.Content)

		fmt.Printf(response.Choices[0].Delta.Content)
	}

	// response, err := client.CreateChatCompletion(
	// 	context.Background(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role: "system",
	// 				Content: "You are a helpful financial advisor who writes blog posts for small businesses. generate a blog title and post from the users {question}",
	// 			},
	// 			{
	// 				Role:    openai.ChatMessageRoleUser,
	// 				Content: prompt,
	// 			},
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	log.Printf("ChatCompletion error: %v\n", err)
	// 	ResponseErr()
	// 	return "Could not complete prompt"
	// }
	// fmt.Println(response.Choices[0].Message.Content)
	// return response.Choices[0].Message.Content
}
