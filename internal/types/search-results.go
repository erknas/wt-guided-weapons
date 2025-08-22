package types

type SearchResult struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type SearchResults struct {
	Results []SearchResult `json:"results"`
}
