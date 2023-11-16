package customerimporter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ImportCustomers(t *testing.T) {
	dummyData := `first_name,last_name,email,gender,ip_address
Mildred,Hernandez,mhernandez0@github.io,Female,38.194.51.128
Bonnie,Ortiz,bortiz1@cyberchimps.com,Female,197.54.209.129
Dennis,Henry,dhenry2@hubpages.com,Male,155.75.186.217
Justin,Hansen,jhansen3@360.cn,Male,251.166.224.119
Carlos,Garcia,cgarcia4@statcounter.com,Male,57.171.52.110
Ernest,Reid,ereid5@rediff.com,Male,243.219.170.46
Gary,Henderson,ghenderson6@acquirethisname.com,Male,30.97.220.14
`

	reader := strings.NewReader(dummyData)
	c, err := ImportCustomers(reader)
	if err != nil {
		panic(err)
	}

	assert.Len(t, c, 7)

}

func Test_ImportCustomersFromFile(t *testing.T) {
	file, err := os.Open("customers.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	c, err := ImportCustomers(reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("domain count: %v\n", len(c))
}

func Test_CanParseFilterShouldParse(t *testing.T) {
	cases := [][]string{
		{"Gary", "Nguyen", "gnguyenrr@angelfire.com", "Male", "199.130.125.248"},
		{"Carl", "Hill", "chill5@dagondesign.com", "Male", "52.37.120.209"},
		{"Laura", "Hawkins", "lhawkins6@google.com.au", "Female", "201.138.248.3"},
		{"Randy", "Carter", "rcarter7@mlb.com", "Male", "227.100.213.163"},
		{"Virginia", "Burke", "vburke8@twitter.com", "Female", "27.118.216.19"},
		{"Peter", "Richards", "prichards9@wordpress.com", "Male", "111.160.156.207"},
		{"Angela", "Freeman", "afreemana@sfgate.com", "Female", "115.54.7.174"},
	}
	for _, r := range cases {
		ef, err := parseLine(r)
		assert.NotNil(t, ef)
		assert.Nil(t, err)
	}
}
