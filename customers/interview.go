// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

type EmailFilter struct {
	domain string
}

func ImportCustomers(r io.Reader) (map[EmailFilter]int, error) {
	customers := make(map[EmailFilter]int)
	reader := csv.NewReader(r)
	// reader.ReuseRecord = true
	for {
		l, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		// fmt.Printf("l: %v\n", l)
		ef, err := parseLine(l)
		if err != nil {
			l, p := reader.FieldPos(2)
			log.Printf("Line %v, col %v - %v", l, p, err)
			continue
		}
		customers[*ef]++
	}

	return customers, nil
}

func parseLine(l []string) (*EmailFilter, error) {
	email := l[2]
	split := strings.Split(email, "@")

	if len(split) <= 1 {
		return nil, fmt.Errorf("Line doesn't contain email: %v", l)
	}
	domain := split[1]
	e := EmailFilter{domain}

	return &e, nil
}
