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
	Question string `form:"question"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home Here"))
}

func GetFinancialHandler(c *gin.Context) {
	var chatPrompt ChatPrompt

	if c.ShouldBind(&chatPrompt) == nil {
		log.Println(chatPrompt.Question)
	}


// }

// func ask_chat(prompt string) string {
	openai_key := os.Getenv("OPENAI_KEY")
	client := openai.NewClient(openai_key)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: "system",
				Content: "You are a helpful financial advisor for native americans who answers questions about college and scholarships.",
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
			c.String(200, "Success")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		ai_response = append(ai_response, response.Choices[0].Delta.Content)

		c.JSON(http.StatusOK, gin.H{
			"message": strings.Join(ai_response, ""),
		})
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
