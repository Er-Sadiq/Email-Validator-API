package authbag

import "regexp"

// CheckEmailSyntax validates the format of an email address using regex
func CheckMailSyntax(email string) bool { // Correct function name and capitalization
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
