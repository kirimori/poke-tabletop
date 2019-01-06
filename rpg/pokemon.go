package rpg

// Pokemon is a representation of an RPG character
type Pokemon struct {
	Name    string   `json:"name"`
	Height  int      `json:"height"`
	Weight  int      `json:"weight"`
	Stats   []Stat   `json:"stats"`
	Level   int      `json:"level"`
	Types   []string `json:"types"`
	Attacks []Attack `json:"attacks"`
}
