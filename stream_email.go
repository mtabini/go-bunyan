package bunyan

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/beefsack/go-rate"
	"net/mail"
	"net/smtp"
	"strings"
	"text/template"
	"time"
)

type EmailStream struct {
	*Stream
	recipient   string
	mailServer  string
	template    *template.Template
	rateLimiter *rate.RateLimiter
}

func NewEmailStream(minLogLevel LogLevel, filter StreamFilter, templateSource string, recipient, mailServer string, minimumInterval time.Duration) (result *EmailStream) {
	t, err := template.New("email").Parse(templateSource)

	if err != nil {
		panic(fmt.Sprintf("Unable to compile email template: %s", err.Error()))
	}

	result = &EmailStream{
		Stream: &Stream{
			MinLogLevel: minLogLevel,
			Filter:      filter,
		},
		recipient:   recipient,
		mailServer:  mailServer,
		template:    t,
		rateLimiter: rate.New(1, minimumInterval),
	}

	return
}

func (s *EmailStream) Publish(l *LogEntry) {
	if ok, _ := s.rateLimiter.Try(); !ok {
		// No more than 1 e-mail for each period!

		return
	}

	encodeRFC2047 := func(String string) string {
		addr := mail.Address{String, ""}
		return strings.Trim(addr.String(), " <>")
	}

	if s.shouldPublish(l) {
		var output bytes.Buffer
		err := s.template.ExecuteTemplate(&output, "email", l)

		if err != nil {
			println(fmt.Sprintf("Error compiling exception template: %s", err))
		}

		header := make(map[string]string)
		header["From"] = "Telemetry API <noreply@telemetryapp.com>"
		header["To"] = s.recipient
		header["Subject"] = encodeRFC2047("Telemetry API Exception Report")
		header["MIME-Version"] = "1.0"
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
		header["Content-Transfer-Encoding"] = "base64"

		message := ""
		for k, v := range header {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		message += "\r\n" + base64.StdEncoding.EncodeToString(output.Bytes())

		c, err := smtp.Dial(s.mailServer)

		defer c.Close()

		if err != nil {
			println(fmt.Sprintf("Error connecting to SMTP server: %s", err))
			return
		}

		c.Mail("Telemetry Composerd <noreply@telemetryapp.com>")
		c.Rcpt(s.recipient)

		wc, err := c.Data()

		if err != nil {
			println(fmt.Sprintf("Error streaming to SMTP server: %s", err))
			return
		}

		defer wc.Close()

		_, err = bytes.NewBuffer([]byte(message)).WriteTo(wc)

		if err != nil {
			println(fmt.Sprintf("Error writing to SMTP server: %s", err))
			return
		}
	}
}
