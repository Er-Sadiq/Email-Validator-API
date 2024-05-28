package main

import (
	"fmt"
	"strings"
	authbag "validate/modules"
)

func main() {

	var authCount int = 0

	fmt.Print("Input An Email String & Let Us Validate It For You :")
	var email string
	fmt.Scan(&email)

	fmt.Println()
	fmt.Println("Validating Your Input....")
	fmt.Println()

	// 1 Checking for syntax
	if !authbag.CheckMailSyntax(email) {
		fmt.Println("Invalid Email Syntax ✘")
	} else {
		fmt.Println("Valid Email Syntax ✔")
		authCount = authCount + 10
	}

	// 2 Checking for most Common Email Servive Providers
	if authbag.IsEmailFromESP(email) {
		fmt.Println("Vaild Email Service Provider ✔")
		authCount = authCount + 20
	} else {
		fmt.Println("InVaild Email Service Provider ✘")
	}

	// 3 Checking for domain
	domain := email[strings.LastIndex(email, "@")+1:]
	if !authbag.CheckDomain(domain) {
		fmt.Println("Invalid domain ✘")
		return
	} else {
		fmt.Println("Valid domain ✔")
		authCount = authCount + 20
	}

	// 4 Checking for Mail SMTP
	if authbag.VerifyMailSMTP(email) {
		fmt.Println("Email SMTP is Valid ✔")
		authCount = authCount + 20
	} else {
		fmt.Println("Email SMTP is Invalid ✘")
	}

	// Calling Auth for MX,SFP,DMARC
	MX, SPF, DMARC, spfRec, dmarcRec := authbag.CheckAuth(email)

	// 5 Checking for mail exchange records
	if MX {
		fmt.Println("Vaild Mail Exchange Records were Found ✔")
		authCount = authCount + 10

	} else {
		fmt.Println("Mail Exchange Records Found were Invalid ✘")
	}

	// 6  Checking SPF (Sender Policy Framework)
	if SPF {
		fmt.Println("Vaild SPF Records  were Found ✔")
		authCount = authCount + 10
	} else {
		fmt.Println("SPF Records Found were Invalid ✘")
	}

	// 7  Checking for Domain-based Message Authentication, Reporting, and Conformance
	if DMARC {
		fmt.Println("Vaild DMARC Records were Found ✔")
		authCount = authCount + 20
	} else {
		fmt.Println("DMARC Records Found were Invalid ✘")
	}

	fmt.Print("\n Record Log :")
	fmt.Println(spfRec)
	fmt.Println(dmarcRec)

	fmt.Println()
	fmt.Println("The Authenticity of Your Eamil is ", authCount, "%")
}
