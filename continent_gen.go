package country

var (

	// Asia - the continent
	Asia = Continent{
		Code:    "AS",
		Slug:    "asian",
		Name:    "Asia",
		Demonym: "asian",
	}

	// Africa - the continent
	Africa = Continent{
		Code:    "AF",
		Slug:    "africa",
		Name:    "Africa",
		Demonym: "african",
	}

	// Antarctica - the continent
	Antarctica = Continent{
		Code:    "AN",
		Slug:    "antarctica",
		Name:    "Antarctica",
		Demonym: "",
	}

	// Europe - the continent
	Europe = Continent{
		Code:    "EU",
		Slug:    "europe",
		Name:    "Europe",
		Demonym: "european",
	}

	// Oceania - the continent
	Oceania = Continent{
		Code:    "OC",
		Slug:    "oceania",
		Name:    "Oceania",
		Demonym: "oceanian",
	}

	// NorthAmerica - the continent
	NorthAmerica = Continent{
		Code:    "NA",
		Slug:    "north-america",
		Name:    "North America",
		Demonym: "north american",
	}

	// SouthAmerica - the continent
	SouthAmerica = Continent{
		Code:    "SA",
		Slug:    "south-america",
		Name:    "South America",
		Demonym: "south american",
	}

	// Continents maps a continent code to its continent.
	Continents = map[string]Continent{
		"AS": Asia,
		"AF": Africa,
		"AN": Antarctica,
		"EU": Europe,
		"OC": Oceania,
		"NA": NorthAmerica,
		"SA": SouthAmerica,
	}
)
