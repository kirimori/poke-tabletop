package pokeapi

// Pokemon is the return object for retrieving a pokemon by name
type Pokemon struct {
	Abilities []struct {
		Ability  Link `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int    `json:"base_experience"`
	Forms          []Link `json:"forms"`
	GameIndices    []struct {
		GameIndex int  `json:"game_index"`
		Version   Link `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item           Link `json:"item"`
		VersionDetails []struct {
			Rarity  int  `json:"rarity"`
			Version Link `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move                Link `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int  `json:"level_learned_at"`
			MoveLearnMethod Link `json:"move_learn_method"`
			VersionGroup    Link `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string      `json:"name"`
	Order   int         `json:"order"`
	Species SpeciesLink `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int  `json:"base_stat"`
		Effort   int  `json:"effort"`
		Stat     Link `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int  `json:"slot"`
		Type Link `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}
