package pokeapi

// SpeciesLink is a link to a pokemon species
type SpeciesLink struct {
	Link
}

// Get retrieves the pokemon species from the link
func (sl *SpeciesLink) Get() (s Species, err error) {
	err = sl.Link.Get(&s)
	return s, err
}

// Species is details about a pokemon species
type Species struct {
	BaseHappiness      int                `json:"base_happiness"`
	CaptureRate        int                `json:"capture_rate"`
	Color              Link               `json:"color"`
	EggGroups          []Link             `json:"egg_groups"`
	EvolutionChain     EvolutionChainLink `json:"evolution_chain"`
	EvolvesFromSpecies SpeciesLink        `json:"evolves_from_species"`
	FlavorTextEntries  []struct {
		FlavorText string `json:"flavor_text"`
		Language   Link   `json:"language"`
		Version    Link   `json:"version"`
	} `json:"flavor_text_entries"`
	FormDescriptions []interface{} `json:"form_descriptions"`
	FormsSwitchable  bool          `json:"forms_switchable"`
	GenderRate       int           `json:"gender_rate"`
	Genera           []struct {
		Genus    string `json:"genus"`
		Language Link   `json:"language"`
	} `json:"genera"`
	Generation           Link   `json:"generation"`
	GrowthRate           Link   `json:"growth_rate"`
	Habitat              Link   `json:"habitat"`
	HasGenderDifferences bool   `json:"has_gender_differences"`
	HatchCounter         int    `json:"hatch_counter"`
	ID                   int    `json:"id"`
	IsBaby               bool   `json:"is_baby"`
	Name                 string `json:"name"`
	Names                []struct {
		Language Link   `json:"language"`
		Name     string `json:"name"`
	} `json:"names"`
	Order             int `json:"order"`
	PalParkEncounters []struct {
		Area      Link `json:"area"`
		BaseScore int  `json:"base_score"`
		Rate      int  `json:"rate"`
	} `json:"pal_park_encounters"`
	PokedexNumbers []struct {
		EntryNumber int  `json:"entry_number"`
		Pokedex     Link `json:"pokedex"`
	} `json:"pokedex_numbers"`
	Shape     Link `json:"shape"`
	Varieties []struct {
		IsDefault bool        `json:"is_default"`
		Pokemon   PokemonLink `json:"pokemon"`
	} `json:"varieties"`
}
