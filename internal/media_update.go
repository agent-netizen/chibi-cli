package internal

import (
    "fmt"
    "os"

    "github.com/charmbracelet/lipgloss"
)

func getTotalEps(mediaId int) (int, error) {
    query := `query ($id: Int) {
        Media(id: $id) {
            episodes
            chapters
            type
        }
    }`

    var reponseData struct {
        Data struct {
            Media struct {
                Episodes int `json:"episodes"`
                Chapters int `json:"chapters"`
                Type string `json:"type"`           
            } `json:"media"`
        } `json:"data"`
    }

    anilistClient := NewAnilistClient()
    err := anilistClient.ExecuteGraqhQL(
        query,
        map[string]interface{} {
            "id": mediaId,
        },
        &reponseData,
    )

    if err != nil {
        return 0, err
    }

    if reponseData.Data.Media.Type == "ANIME" {
        return reponseData.Data.Media.Episodes, nil
    } else {
        return reponseData.Data.Media.Chapters, nil
    }
}

type MediaUpdate struct {
    Data struct {
        SaveMediaListEntry struct {
            MediaId int `json:"mediaId"`
        } `json:"SaveMediaListEntry"`
    } `json:"data"`
}

func (mu *MediaUpdate) Get(mediaId int, progress int) error {
    total, err := getTotalEps(mediaId)
    if err != nil {
        return err
    }

    if progress > total {
        fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(
            fmt.Sprintf("Entered value is greater than total episodes / chapters, which is %d", total),
        ))
        os.Exit(0)
    }

    mutation := `mutation($id: Int, $progress: Int) {
        SaveMediaListEntry(mediaId: $id, progress: $progress) {
            mediaId
        }
    }`

    err = NewAnilistClient().ExecuteGraqhQL(
        mutation,
        map[string]interface{} {
            "id": mediaId,
            "progress": progress,
        },
        &mu,
    )
    return err
}

func NewMediaUpdate() *MediaUpdate {
    return &MediaUpdate{}
}