package authbag

import (
	"net"
	"strings"
)

func CheckAuth(email string) (bool, bool, bool, string, string) {
	var MX, SPF, DMARC bool
	var spfRec, dmarcRec string

	mxRecord, err := net.LookupMX(email)
	if err != nil {
		// log.Printf("Error looking up MX record: %v\n", err)
	}
	if len(mxRecord) > 0 {
		MX = true
	}

	txtRecord, err := net.LookupTXT(email)
	if err != nil {
		//log.Printf("Error looking up TXT record: %v \n", err)
	} else {
		for _, rec := range txtRecord {
			if strings.HasPrefix(rec, "v=spf1") {
				SPF = true
				spfRec = rec
				break
			}
		}
	}

	dmarcRecord, err := net.LookupTXT("_dmarc." + strings.Split(email, "@")[1])
	if err != nil {
		// log.Printf("Error looking up DMARC record: %v\n", err)
	} else {
		for _, rec := range dmarcRecord {
			if strings.HasPrefix(rec, "v=DMARC1") {
				DMARC = true
				dmarcRec = rec
				break
			}
		}
	}

	return MX, SPF, DMARC, spfRec, dmarcRec
}
