package country

import (
	. "github.com/101loops/bdd"
)

var _ = Describe("Continents", func() {

	It("is complete", func() {
		Check(Continents, HasLen, 7)
	})

	It("contains correctly mapped information", func() {
		Check(NorthAmerica.Name, Equals, "North America")
		Check(NorthAmerica.Code, Equals, "NA")
		Check(NorthAmerica.Demonym, Equals, "north american")
		Check(NorthAmerica.Slug, Equals, "north-america")
	})
})
