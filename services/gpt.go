package services

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
	"log"
	"os"
)

var ChatGPTClient *openai.Client

func InitChatGpt() {
	key := os.Getenv("OPENAI_KEY")
	if key == "" {
		log.Fatalf("No OPENAI_KEY key found! ")
	}
	c := openai.NewClient(key)

	ChatGPTClient = c
	log.Println("GPT client initialized!")
}

func SendPrompt(prompt string) (string, error) {
	ctx := context.Background()
	var messages []openai.ChatCompletionMessage
	oPrompt := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: prompt,
	}
	messages = append(messages, oPrompt)
	continues := []string{
		"please continue",
		"please continue",
		"please continue",
	}
	for _, p := range continues {
		message := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: p,
		}
		messages = append(messages, message)
	}
	req := &openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
	}
	res, err := ChatGPTClient.CreateChatCompletion(ctx, *req)
	if err != nil {
		return "", err
	}
	c := GetContentFromResponse(res)
	return c, nil
}

func GetContentFromResponse(response openai.ChatCompletionResponse) string {
	content := response.Choices[0].Message.Content
	return content
}
