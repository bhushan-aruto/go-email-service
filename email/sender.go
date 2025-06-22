package email

import (
	"fmt"
	"os"

	"github.com/bhushan-aruto/go-email-service/internal/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func sendEmail(to string, subject string, htmlBody string) error {
	fromEmail := os.Getenv("ROOT_EMAIL")
	apiKey := os.Getenv("SENDGRID_API_KEY")

	if fromEmail == "" || apiKey == "" {
		return fmt.Errorf("missing SENDGRID_API_KEY or ROOT_EMAIL")
	}

	from := mail.NewEmail("Aspiration Matters", fromEmail)
	toEmail := mail.NewEmail("", to)
	message := mail.NewSingleEmail(from, subject, toEmail, "", htmlBody)

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)

	if err != nil {
		return fmt.Errorf("sendgrid error: %v", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("sendgrid API response error: %v - %v", response.StatusCode, response.Body)
	}

	return nil
}

func SendOtpEmail(message *models.Email) error {
	template, err := parseTemplate("otp")

	if err != nil {
		return err
	}

	htmlTemplate, err := renderTemplate(template, message.Data)

	if err != nil {
		return err
	}

	if err := sendEmail(message.To, message.Subject, htmlTemplate); err != nil {
		return err
	}
	return nil
}

func WelcomeEmail(message *models.Email) error {
	template, err := parseTemplate("welcome")

	if err != nil {
		return err
	}

	htmlTemplate, err := renderTemplate(template, message.Data)

	if err != nil {
		return err
	}

	if err := sendEmail(message.To, message.Subject, htmlTemplate); err != nil {
		return err
	}
	return nil
}
