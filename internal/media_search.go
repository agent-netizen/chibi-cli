package internal

type MediaSearch struct {
    Data struct {
        Page struct {
            Media []struct {
                Id    int `json:"id"`
                Title struct {
                    UserPreferred string `json:"userPreferred"`
                } `json:"title"`
                AverageScore float64 `json:"averageScore"`
            } `json:"media"`
        } `json:"page"`
    } `json:"data"`
}

func (m *MediaSearch) Get(searchQuery string, mediaType string, perPage int) error {
    switch mediaType {
    case "anime", "a", "":
        mediaType = "ANIME"
    case "manga", "m":
        mediaType = "MANGA"
    }

    query :=
        `query($searchQuery: String, $perPage: Int, $mediaType: MediaType) {
        Page(perPage: $perPage) {
            media(search: $searchQuery, type: $mediaType) {
                id
                title {
                    userPreferred
                }
                type
                averageScore
            }
        }
    }`

    anilistClient := NewAnilistClient()
    err := anilistClient.ExecuteGraqhQL(
        query,
        map[string]interface{}{
            "searchQuery": searchQuery,
            "perPage":     perPage,
            "mediaType":   mediaType,
        },
        &m,
    )
    return err
}

func NewMediaSearch() *MediaSearch {
    return &MediaSearch{}
}
