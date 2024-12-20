package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
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

type updateMediaFields struct {
    CompletedAtDate int
    CompletedAtMonth int
    CompletedAtYear int
    Notes string
    Score float32    
}

func updateMediaEntry() (*updateMediaFields, error) {
    mediaFields := &updateMediaFields{}
    currDate := fmt.Sprintf("%d/%d/%d\n", time.Now().Day(), time.Now().Month(), time.Now().Year())
    var scoreString string

    huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("Date of completion").
                Value(&currDate).
                Description("Date should be in format DD/MM/YYYY").
                Validate(func(s string) error {
                    layout := "02/01/2006"
                    _, err := time.Parse(layout, strings.TrimSpace(s))
                    return err
                }),
        ),
        huh.NewGroup(
            huh.NewText().
                Title("Notes").
                Description("Note: you can add multiple lines").
                Value(&mediaFields.Notes),
        ),
        huh.NewGroup(
            huh.NewInput().
                Title("Score").
                Description("If your score is in emoji, type 1 for 😞, 2 for 😐 and 3 for 😊").
                Prompt("> ").
                Validate(func(s string) error {
                    _, err := strconv.ParseFloat(s, 64)
                    return err
                }).
                Value(&scoreString),
        ),
    ).Run()
    completedDate, err := time.Parse("02/01/2006", strings.TrimSpace(currDate))
    if err != nil {
        return nil, err
    }
    scoreFloat, err := strconv.ParseFloat(scoreString, 32)
    if err != nil {
        return nil, err
    }

    mediaFields.CompletedAtDate = completedDate.Day()
    mediaFields.CompletedAtMonth = int(completedDate.Month())
    mediaFields.CompletedAtYear = completedDate.Year()
    mediaFields.Score = float32(scoreFloat)

    return mediaFields, nil
}

type MediaUpdate struct {
    Data struct {
        SaveMediaListEntry struct {
            MediaId int `json:"mediaId"`
        } `json:"SaveMediaListEntry"`
    } `json:"data"`
}

func (mu *MediaUpdate) Get(isMediaAdd bool, mediaId int, progress int, status string, startDate string) error {
    if status == "" {
        status = "COMPLETED"
    }

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

    mutation := 
    `mutation(
        $id: Int, 
        $progress: Int,
        $score: Float,
        $notes: String,
        $cDate: Int,
        $cMonth: Int,
        $cYear: Int,
        $sDate: Int,
        $sMonth: Int,
        $sYear: Int,
        $status: MediaListStatus
    ) {
        SaveMediaListEntry(
            mediaId: $id, 
            progress: $progress,
            status: $status,
            score: $score,
            notes: $notes,
            completedAt: {
                day: $cDate,
                month: $cMonth,
                year: $cYear
            },
            startedAt: {
                day: $sDate,
                month: $sMonth,
                year: $sYear
            }
        ) {
            mediaId
        }
    }`

    if isMediaAdd {
        variables := map[string]interface{} {
            "id": mediaId,
            "status": status,
        }

        if startDate != "" {
            startDateRaw, err := time.Parse("02/01/2006", startDate)
            if err != nil {
                return err
            }

            if status == "CURRENT" {
                variables["sDate"] = startDateRaw.Day()
                variables["sMonth"] = int(startDateRaw.Month())
                variables["sYear"] = startDateRaw.Year()
            }
        }

        err = NewAnilistClient().ExecuteGraqhQL(
            mutation,
            variables,
            &mu,
        )
        return err
    }

    var canEditList bool = false

    if progress == total {
        huh.NewConfirm().
            Title("Seems like you completed the anime/manga. Do you want to mark this as completed?").
            Affirmative("Yes!").
            Negative("No").
            Value(&canEditList).
            Run()
    }

    if canEditList {
        mediaFields, err := updateMediaEntry()
        if err != nil {
            return err
        }

        err = NewAnilistClient().ExecuteGraqhQL(
            mutation,
            map[string]interface{} {
                "id": mediaId,
                "progress": progress,
                "score": mediaFields.Score,
                "notes": mediaFields.Notes,
                "cDate": mediaFields.CompletedAtDate,
                "cMonth": mediaFields.CompletedAtMonth,
                "cYear": mediaFields.CompletedAtYear,
            },
            &mu,
        )
        return err
    }

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