// backend-bakashi-go/types/episode.go
package types

type Episode struct {
    ID          string `json:"id"`
    AnimeID     string `json:"animeId"`
    Season      int    `json:"season"`
    Title       string `json:"title"`
    Slug        string `json:"slug"`
    Image       string `json:"image"`
    VideoUrl    string `json:"videoUrl"`
    ReleaseDate string `json:"releaseDate"`
    IsLancamento bool  `json:"isLancamento"`
}
