
# Email-Validation-API

EmailAuthValidator_API is a Go-based API service designed to verify the authenticity of email addresses using seven distinct factors. This API aims to enhance the reliability and security of email validation processes by leveraging various validation techniques including syntax checks, domain reputation, ESP identification, MX record verification, SPF record validation, SMTP checks, and DMARC record verification.
## Features

 - Syntax Validation: Checks if the email address   follows the correct format.
 - Common Domains Check: Identifies if the email domain is a commonly used domain.
 - ESP Detection: Determines if the email provider is a well-known ESP (Email Service Provider).
 - MX Record Verification: Verifies the existence and validity of MX (Mail Exchange) records for the domain.
 - SPF Record Validation: Checks the domain's SPF (Sender Policy Framework) records for email spoofing prevention.
 - SMTP Verification: Attempts to connect to the SMTP server to verify the existence of the email address.
 - DMARC Record Verification: Ensures the domain has a DMARC (Domain-based Message Authentication, Reporting, and Conformance) policy set up.


## Installation

Clone git repo

```bash
  git clone https://github.com/Er-Sadiq/EmailAuthValidator_API.git
  cd into the dir 
  go mod tidy
```
```bash
 Go run main.go
````

## Screenshots

![GoEmailAuth](https://github.com/Er-Sadiq/EmailAuthValidator_API/assets/125464939/6bbf06eb-3d1f-4d17-ac06-7a425805d8e3)

## False Email 
![Annotation 2024-05-28 231859](https://github.com/Er-Sadiq/EmailAuthValidator_API/assets/125464939/ea5269c1-61bc-4d5b-bfbe-e8653d477c0e)

## Genuine Email

![Annotation 2024-05-28 232838](https://github.com/Er-Sadiq/EmailAuthValidator_API/assets/125464939/b001bee3-d17d-4e2a-b871-56acad6e9200)






## License

This project is licensed under the MIT License. See the LICENSE file for details.

