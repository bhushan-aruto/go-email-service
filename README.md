# 📄 **Go Email Service** 📧

## 📌 **Overview** 🚀
A **lightweight and efficient** Go-based email service that allows sending **OTP and welcome emails** using **SMTP and RabbitMQ** for message queuing. The service listens for email requests and processes them asynchronously.

---

## ✨ **Features** 🎉
- ✅ **Supports OTP and Welcome Emails** ✉️
- ✅ **Uses RabbitMQ for message queuing** 🐇
- ✅ **Template-based email rendering** 📑
- ✅ **Secure SMTP authentication** 🔐
- ✅ **Handles JSON email requests** 🛠️

---

## 🛠️ **Tech Stack & Dependencies** 💻

- 🔹 **Go** – High-performance backend programming language 🚀
- 🔹 **RabbitMQ** – Message broker for handling email requests asynchronously 🐇
- 🔹 **SMTP** – For sending emails securely 📩

### Required Go Packages 📦

```sh
# Install dependencies
go mod tidy
```

---

## 📂 **Project Structure** 📁
```
├── cmd
│   └── main.go          # Entry point of the application
├── consumer
│   └── consumer.go      # Email queue consumer
├── email
│   ├── sender.go        # SMTP email sender
│   ├── templates
│   │   ├── otp.html     # OTP email template
│   │   └── welcome.html # Welcome email template
│   └── templates.go     # Template rendering logic
├── internal
│   └── models
│       └── email.go     # Email struct definition
├── queue
│   └── rabbitmq.go      # RabbitMQ connection handler
├── go.mod               # Go module file
├── go.sum               # Dependency lock file
└── README.md            # Project documentation
```

---

## ⚙️ **Setup & Configuration** 🛠️

### 1️⃣ **Set Up Environment Variables** 🌍
Create a `.env` file and configure the required variables:
```sh
ROOT_EMAIL=your_email@example.com
ROOT_EMAIL_PASSWORD=your_password
SMTP_HOST=smtp.example.com
SMTP_PORT=587
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
QUEUE_NAME=email_queue
```

### 2️⃣ **Run RabbitMQ** 🐇
Make sure RabbitMQ is running:
```sh
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
```

### 3️⃣ **Start the Email Consumer** 📩
Run the email consumer service:
```sh
go run cmd/main.go
```

---

## 📜 **How It Works** 🔄

1. **Producer sends an email request** 📨
   - A message is sent to RabbitMQ with JSON payload:
     ```json
     {
       "to": "user@example.com",
       "subject": "Your OTP Code",
       "email_type": "otp",
       "data": { "otp": "123456" }
     }
     ```
2. **Consumer receives the request** 🎧
   - The `consumer.go` listens to the queue, processes messages, and invokes the appropriate email template.
3. **SMTP sends the email** 📧
   - `sender.go` formats and sends the email using the SMTP credentials.

---

## ⚠️ **Notes & Troubleshooting** 📝

- Ensure RabbitMQ is running before starting the consumer service.
- Double-check environment variables if emails are not being sent.
- Check logs for debugging (`log.Println` is used for error handling).

---

🚀 **Happy Coding!** 🚀

