package cmd

import (
    "fmt"
    "os"
    "strconv"

    "github.com/CosmicPredator/chibi/internal"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/lipgloss/table"
    "github.com/spf13/cobra"
)

var listMediaType string

func handleLs() {
    var mediaType string

    switch listMediaType {
    case "anime", "a":
        mediaType = "ANIME"
    case "manga", "m":
        mediaType = "MANGA"
    }

    mediaList := internal.NewMediaList()
    err := mediaList.Get(mediaType)
    if err != nil {
        fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(
            "Some error happened, please try again...",
        ))
        os.Exit(0)
    }

    if len(mediaList.Data.MediaListCollection.Lists) == 0 {
        fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(
            "No entries found, go add some :)",
        ))
        os.Exit(0)
    }

    rows := [][]string{}
    for _, i := range mediaList.Data.MediaListCollection.Lists[0].Entries {
        var progress string
        if mediaType == "ANIME" {
            progress = fmt.Sprintf("%d/%d", i.Progress, i.Media.Episodes)
        } else {
            progress = fmt.Sprintf("%d/%d", i.Progress, i.Media.Chapters)
        }

        rows = append(rows, []string{
            strconv.Itoa(i.Media.Id),
            i.Media.Title.UserPreferred,
            progress,
        })
    }

    t := table.New().
        Border(lipgloss.NormalBorder()).
        BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
        StyleFunc(func(row, col int) lipgloss.Style {
            switch {
            case row == -1:
                return lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true).Align(lipgloss.Center)
            default:
                return lipgloss.NewStyle().Align(lipgloss.Center).PaddingLeft(2).PaddingRight(2)
            }
        }).
        Headers("ID", "TITLE", "PROGRESS").
        Rows(rows...)

    fmt.Println(t)
}


var mediaListCmd = &cobra.Command{
    Use: "ls",
    Short: "List your current anime/manga list",
    Run: func(cmd *cobra.Command, args []string) {
        handleLs()
    },
}

func init() {
    mediaListCmd.Flags().StringVarP(
        &listMediaType, "type", "t", "anime", "Type of media. for anime, pass 'anime' or 'a', for manga, use 'manga' or 'm'",
    )
}
