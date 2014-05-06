package country

import "fmt"

// Country represents meta data about a country.
type Country struct {
	CallingCode   string
	Capital       string
	ContinentCode string
	ContinentName string
	CurrencyCode  string
	CurrencyName  string
	ISO           string
	ISO3          string
	Name          string
	TLD           string
}

func (c *Country) String() {
	fmt.Sprintf("%s[%s]", c.Name, c.ISO)
}

// Continent returns the country's Continent.
func (c *Country) Continent() (ret Continent) {
	if cont, ok := Continents[c.ContinentCode]; ok {
		ret = cont
	}
	return
}
