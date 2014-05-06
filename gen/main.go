package main

import (
	"encoding/csv"
	"fmt"
	"go/format"
	"os"
	"strconv"
	"strings"
)

type continent struct {
	Name    string
	Slug    string
	Code    string
	Demonym string
}

type country struct {
	CallingCode   string
	Capital       string
	CurrencyCode  string
	CurrencyName  string
	ContinentCode string
	ISO           string
	ISO3          string
	ISONum        int
	Name          string
	Population    int
	TLD           string
}

var (
	continents = map[string]continent{
		"AS": {"Asia", "asian", "AS", "asian"},
		"AF": {"Africa", "africa", "AF", "african"},
		"AN": {"Antarctica", "antarctica", "AN", ""},
		"EU": {"Europe", "europe", "EU", "european"},
		"OC": {"Oceania", "oceania", "OC", "oceanian"},
		"NA": {"North America", "north-america", "NA", "north american"},
		"SA": {"South America", "south-america", "SA", "south american"},
	}
)

func main() {
	saveCode("country", generateCountries(loadCountries()))
	saveCode("continent", generateContinents())
}

func generateContinents() (out string) {
	out += `package country

	var (
	`

	for _, cont := range continents {
		name := strings.Replace(cont.Name, " ", "", -1)
		data := `
		Continent{
			Code: "` + cont.Code + `",
			Slug: "` + cont.Slug + `",
			Name: "` + cont.Name + `",
			Demonym: "` + cont.Demonym + `",
		}`
		out += fmt.Sprintf(`
		// %s - the continent
		%s = %s`, name, name, data) + "\n"
	}

	out += `
	// Continents maps a continent code to its continent.
	Continents = map[string]Continent{
	`
	for _, cont := range continents {
		name := strings.Replace(cont.Name, " ", "", -1)
		out += fmt.Sprintf("%q: %s,\n", cont.Code, name)
	}
	out += `
		}
	)
	`

	return
}

func loadCountries() (res []country) {
	file, err := os.Open("countries.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.Comment = '#'
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		continent, ok := continents[record[8]]
		if !ok {
			panic("continent not found for " + record[10])
		}
		isoNum, _ := strconv.Atoi(record[2])
		res = append(res, country{
			ISO:           record[0],
			ISO3:          record[1],
			ISONum:        isoNum,
			Name:          record[4],
			Capital:       record[5],
			TLD:           record[9],
			ContinentCode: continent.Code,
			CurrencyCode:  record[10],
			CurrencyName:  record[11],
			CallingCode:   record[12],
		})
	}

	return
}

func generateCountries(countries []country) (out string) {
	out += `package country

	var (
	`

	for _, cnt := range countries {
		code := cnt.ISO
		data := `
		Country{
			Capital: "` + cnt.Capital + `",
			CallingCode: "` + cnt.CallingCode + `",
			CurrencyCode: "` + cnt.CurrencyCode + `",
			CurrencyName: "` + cnt.CurrencyName + `",
			ContinentCode: "` + cnt.ContinentCode + `",
			ISO: "` + cnt.ISO + `",
			ISO3: "` + cnt.ISO3 + `",
			ISONum: ` + strconv.Itoa(cnt.ISONum) + `,
			Name: "` + cnt.Name + `",
			TLD: "` + cnt.TLD + `",
		}`
		out += fmt.Sprintf(`
		// %s is %q.
		%s = %s`, code, cnt.Name, code, data) + "\n"
	}

	out += `
	// Countries maps an ISO code to a country.
	Countries = map[string]Country {
	`
	for _, cnt := range countries {
		out += fmt.Sprintf("%q: %s,\n", cnt.ISO, cnt.ISO)
	}
	out += `
		}
	)
	`

	return
}

func saveCode(name, code string) {
	file, err := os.OpenFile("../"+name+"_gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := format.Source([]byte(code))
	if err != nil {
		file.Write([]byte(code))
		panic(err)
	}

	_, err = file.Write(bytes)
	if err != nil {
		panic(err)
	}
}
