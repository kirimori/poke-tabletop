package pokeapi

// EvolutionChainLink is a link to an evolution chain
type EvolutionChainLink struct {
	Link
}

// Get retrieves the pokemon species from the link
func (ecl *EvolutionChainLink) Get() (ec EvolutionChain, err error) {
	err = ecl.Link.Get(&ec)
	return ec, err
}

type EvolutionDetails struct {
	Gender                interface{} `json:"gender"`
	HeldItem              interface{} `json:"held_item"`
	Item                  Link        `json:"item"`
	KnownMove             interface{} `json:"known_move"`
	KnownMoveType         interface{} `json:"known_move_type"`
	Location              interface{} `json:"location"`
	MinAffection          interface{} `json:"min_affection"`
	MinBeauty             interface{} `json:"min_beauty"`
	MinHappiness          int         `json:"min_happiness"`
	MinLevel              interface{} `json:"min_level"`
	NeedsOverworldRain    bool        `json:"needs_overworld_rain"`
	PartySpecies          interface{} `json:"party_species"`
	PartyType             interface{} `json:"party_type"`
	RelativePhysicalStats interface{} `json:"relative_physical_stats"`
	TimeOfDay             string      `json:"time_of_day"`
	TradeSpecies          interface{} `json:"trade_species"`
	Trigger               Link        `json:"trigger"`
	TurnUpsideDown        bool        `json:"turn_upside_down"`
}

type EvolvesTo struct {
	EvolutionDetails []EvolutionDetails `json:"evolution_details"`
	EvolvesTo        []EvolvesTo        `json:"evolves_to"`
	IsBaby           bool               `json:"is_baby"`
	Species          SpeciesLink        `json:"species"`
}

type EvolutionChain struct {
	BabyTriggerItem interface{} `json:"baby_trigger_item"`
	Chain           EvolvesTo   `json:"chain"`
	ID              int         `json:"id"`
}
