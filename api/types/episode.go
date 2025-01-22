package types

// Estrutura para representar os episódios
type Episode struct {
	ID          string `json:"id"`
	AnimeID     string `json:"animeId"`
	Season      int    `json:"season"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	VideoURL    string `json:"videoUrl"`
	ReleaseDate string `json:"releaseDate"`
	IsLancamento bool  `json:"isLancamento"`
}

// Estrutura para representar a lista de episódios
type Episodes struct {
	Episodes []Episode `json:"episodes"`
}
