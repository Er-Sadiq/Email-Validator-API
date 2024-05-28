package authbag

import (
	"net"
	"net/smtp"
	"strings"
)

// CheckDomain verifies if the domain has valid DNS records
func CheckDomain(domain string) bool {
	_, err := net.LookupMX(domain)
	return err == nil
}

// VerifyEmailSMTP connects to the email server to verify the email address
func VerifyMailSMTP(email string) bool {
	domain := email[strings.LastIndex(email, "@")+1:]
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return false
	}

	client, err := smtp.Dial(mxRecords[0].Host + ":25")
	if err != nil {
		return false
	}
	defer client.Close()

	client.Hello("example.com")
	client.Mail("verify@example.com")
	rcptErr := client.Rcpt(email)
	if rcptErr != nil {
		return false
	}

	return true
}
