package pokeapi

// Link is a link for further information in the api
type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func removeEmptyStrings(ss []string) []string {
	cleanSS := []string{}
	for _, s := range ss {
		if len(s) > 0 {
			cleanSS = append(cleanSS, s)
		}
	}
	return cleanSS
}
