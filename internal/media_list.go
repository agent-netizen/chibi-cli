package internal

import "github.com/CosmicPredator/chibi/types"

type MediaList struct {
	Data struct {
		MediaListCollection struct {
			Lists []struct {
				Entries []struct {
					Progress        int `json:"progress"`
					ProgressVolumes int `json:"progressVolumes"`
					Media           struct {
						Id    int `json:"id"`
						Title struct {
							UserPreferred string `json:"userPreferred"`
						} `json:"title"`
						Chapters int `json:"chapters"`
						Volumes  int `json:"volumes"`
						Episodes int `json:"episodes"`
					} `json:"media"`
				} `json:"entries"`
			} `json:"lists"`
		} `json:"MediaListCollection"`
	} `json:"data"`
}

func (ml *MediaList) Get(mediaType string) error {
	anilistClient := NewAnilistClient()
	tokenConfig := types.NewTokenConfig()
	err := tokenConfig.ReadFromJsonFile()

	if err != nil {
		return err
	}

	query :=
		`query($userId: Int, $type: MediaType) {
        MediaListCollection(userId: $userId, type: $type, status: CURRENT) {
            lists {
                entries {
                    progress
                    progressVolumes
                    media {
                        id
                        title {
                            userPreferred
                        }
                        chapters
                        volumes
                        episodes
                    }
                }
            }
        }
    }`

	err = anilistClient.ExecuteGraqhQL(
		query,
		map[string]interface{}{
			"type":   mediaType,
			"userId": tokenConfig.UserId,
		},
		&ml,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewMediaList() *MediaList {
	return &MediaList{}
}
