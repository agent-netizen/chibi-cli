package internal

type Profile struct {
	Data struct {
		Viewer struct {
			Name       string `json:"name"`
			SiteUrl    string `json:"siteUrl"`
			Id         int    `json:"id"`
			Statistics struct {
				Anime struct {
					Count          int `json:"count"`
					MinutesWatched int `json:"minutesWatched"`
				} `json:"anime"`
				Manga struct {
					Count        int `json:"count"`
					ChaptersRead int `json:"chaptersRead"`
				} `json:"manga"`
			} `json:"statistics"`
		} `json:"Viewer"`
	} `json:"data"`
}

func (p *Profile) Get() error {
	anilistClient := NewAnilistClient()
	query :=
		`query {
        Viewer {
            id
            name
            statistics {
                anime {
                    count
                    minutesWatched
                }
                manga {
                    count
                    chaptersRead
                }
            }
            siteUrl
        }
    }`

	err := anilistClient.ExecuteGraqhQL(
		query, make(map[string]interface{}), &p,
	)

	if err != nil {
		return err
	}

	return nil
}

func NewProfile() *Profile {
	return &Profile{}
}
