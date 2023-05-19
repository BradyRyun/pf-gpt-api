package services

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
	"log"
	"os"
)

var ChatGPTClient *openai.Client
var GPTEnabled bool

func InitChatGpt() {
	GPTEnabled = os.Getenv("GPT_ENABLED") == "true"
	if !GPTEnabled {
		log.Println("Bypassing chatgpt client initialization")
		return
	}
	key := os.Getenv("OPENAI_KEY")
	if key == "" {
		log.Fatalf("No OPENAI_KEY key found! ")
	}
	c := openai.NewClient(key)

	ChatGPTClient = c
	log.Println("GPT client initialized!")
}

func SendPrompt(prompt string) (string, error) {
	if !GPTEnabled {
		var TestResponse = `Based on your financial situation, here are some actionable steps you can take to improve your overall wealth management:\n\n- Pay off your credit card debt as soon as possible to avoid paying high interest charges. Consider creating a budget and cutting back on non-essential expenses to put more money towards the debt.\n- Continue saving and investing aggressively, with a focus on tax-advantaged accounts like an IRA or 401(k). Consider increasing your investment contributions if possible.\n- Build up your emergency fund to at least 3-6 months' worth of living expenses, to protect yourself against unexpected expenses or job loss.\n- Consider exploring the possibility of homeownership, as it can be a smart long-term investment. Evaluate your finances to see if you can afford a down payment and monthly mortgage payments.\n- Review and adjust your investment strategy regularly, based on changes in your goals, risk tolerance, and financial situation.\n- Consider speaking with a certified advisor to get personalized financial advice based on your unique situation.`
		log.Println("Bypassing chatgpt client initialization")
		return TestResponse, nil
	}
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
