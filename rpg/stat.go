package rpg

// Stat is an ability score for defining things like charisma, dexterity, etc
type Stat struct {
	Ability string `json:"ability"`
	Score   int    `json:"score"`
}

// Modifier is the calculated Ability Modifier for the particular stat requested
func (s Stat) Modifier() int {
	return int(s.Score/2 - 5)
}
