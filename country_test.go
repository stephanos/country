package country

import (
	. "github.com/101loops/bdd"
)

var _ = Describe("Countries", func() {

	It("is complete", func() {
		Check(Countries, HasLen, 252)
	})

	It("contains correctly mapped information", func() {
		Check(DE.CallingCode, Equals, "49")
		Check(DE.Capital, Equals, "Berlin")
		Check(DE.ContinentCode, Equals, "EU")
		Check(DE.CurrencyCode, Equals, "EUR")
		Check(DE.CurrencyName, Equals, "Euro")
		Check(DE.ISO, Equals, "DE")
		Check(DE.ISO3, Equals, "DEU")
		Check(DE.Name, Equals, "Germany")
		Check(DE.TLD, Equals, ".de")
	})
})
