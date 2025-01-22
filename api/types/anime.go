package types

// Estrutura para representar os animes
type Anime struct {
	IsRelease        bool     `json:"isRelease"`
	IsPopularSeason  bool     `json:"isPopularSeason"`
	NewReleases      bool     `json:"newReleases"`
	IsPopular        bool     `json:"isPopular"`
	IsNextSeason     bool     `json:"isNextSeason"`
	IsThumbnail      bool     `json:"isThumbnail"`
	IsMovie          bool     `json:"isMovie"`
	ID               string   `json:"id"`
	Slug             string   `json:"slug"`
	Name             string   `json:"name"`
	ReleaseYear      string   `json:"releaseYear"`
	Image            string   `json:"image"`
	Synopsis         string   `json:"synopsis"`
	Rating           int      `json:"rating"`
	Score            float64  `json:"score"`
	Genres           []string `json:"genres"`
	AiringDay        string   `json:"airingDay"`
	Episodes         int      `json:"episodes"`
	Season           int      `json:"season"`
	AudioType        string   `json:"audioType"`
	ThumbnailImage   string   `json:"thumbnailImage"`
}

// Estrutura para representar a lista de animes
type Animes struct {
	Animes []Anime `json:"animes"`
}
