package cmd

import (
    "fmt"
    "os"
    "strconv"

    "github.com/CosmicPredator/chibi/internal"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/lipgloss/table"
    "github.com/charmbracelet/x/term"
    "github.com/spf13/cobra"
)

var listMediaType string

func handleLs() {
    CheckIfTokenExists()
    
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
        ErrorMessage(err.Error())
    }

    if len(mediaList.Data.MediaListCollection.Lists) == 0 {
        ErrorMessage(err.Error())
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

    // get size of terminal
    tw, _, err := term.GetSize((os.Stdin.Fd()))
    if err != nil {
        ErrorMessage(err.Error())
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
        Rows(rows...).Width(tw)

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
