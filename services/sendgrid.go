package services

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

var SendGridClient *sendgrid.Client
var SendGridEnabled bool

func InitSendGrid() {
	APIKey := os.Getenv("SENDGRID_API_KEY")
	SendGridEnabled = os.Getenv("SG_ENABLED") == "true"
	if !SendGridEnabled {
		log.Println("SendGrid client disabled.")
		return
	}
	if APIKey == "" {
		log.Fatalf("No SENDGRID_API_KEY key found! ")
	}
	SendGridClient = sendgrid.NewSendClient(APIKey)
	log.Println("SendGrid client initialized!")
}

func SendEmail(content string, toEmail string) bool {
	if !SendGridEnabled {
		log.Println("SendGrid is not enabled. No email will be sent.")
		return true
	}
	fromEmail := os.Getenv("SENDGRID_FROM_EMAIL")
	from := mail.NewEmail("Brady from Personal Finance GPT", fromEmail)
	subject := fmt.Sprintf("You've received a personalized finance plan!")
	to := mail.NewEmail("Friend", toEmail)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := fmt.Sprintf(content)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	res, err := SendGridClient.Send(message)
	if err != nil {
		log.Println(err)
	}

	responseCode := res.StatusCode

	if responseCode >= 400 {
		log.Printf("api response: HTTP %d: %s", res.StatusCode, res.Body)
		return false
	}
	return true
}
