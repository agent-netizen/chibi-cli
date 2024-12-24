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

func parseMediaStatus(status string) string {
	switch status {
	case "watching", "reading", "w", "r":
		return "CURRENT"
	case "planning", "p":
		return "PLANNING"
	case "completed", "c":
		return "COMPLETED"
	case "dropped", "d":
		return "DROPPED"
	case "paused", "ps":
		return "PAUSED"
	case "repeating", "rp":
		return "REPEATING"
	default:
		return "CURRENT"
	}
}

func (ml *MediaList) Get(mediaType string, status string) error {
	anilistClient := NewAnilistClient()
	tokenConfig := types.NewTokenConfig()
	err := tokenConfig.ReadFromJsonFile()

	if err != nil {
		return err
	}

	query :=
		`query($userId: Int, $type: MediaType, $status: MediaListStatus) {
        MediaListCollection(userId: $userId, type: $type, status: $status) {
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
			"status": parseMediaStatus(status),
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
