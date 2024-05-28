package authbag

import "strings"

var espDomains = []string{
	"gmail.com",
	"yahoo.com",
	"outlook.com",
	"hotmail.com",
	"aol.com",
	"icloud.com",
	"mail.com",
	"zoho.com",
	"protonmail.com",
	"gmx.com",
	"yandex.com",
	"comcast.net",
	"verizon.net",
	"msn.com",
	"att.net",
	"rediffmail.com",
	"mail.ru",
	"lycos.com",
	"tutanota.com",
	"fastmail.com",
}

// IsEmailFromESP checks if the email is from a known email service provider
func IsEmailFromESP(email string) bool {
	domain := strings.Split(email, "@")[1]
	for _, espDomain := range espDomains {
		if domain == espDomain {
			return true
		}
	}
	return false
}
