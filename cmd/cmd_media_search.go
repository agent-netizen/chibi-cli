package cmd

import (
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/CosmicPredator/chibi/internal"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/lipgloss/table"
    "github.com/spf13/cobra"
)

var pageSize int
var searchMediaType string

func getMediaSearch(searchQuery string) {
    CheckIfTokenExists()
    
    if pageSize < 0 || pageSize > 50 {
        fmt.Println("page count must be lesser than 50 and greater than 0")
        os.Exit(0)
    }

    mediaSearch := internal.NewMediaSearch()
    err := mediaSearch.Get(searchQuery, searchMediaType, pageSize)
    if err != nil {
        ErrorMessage(err.Error())
    }
    rows := [][]string{}

    for _, i := range mediaSearch.Data.Page.Media {
        rows = append(rows, []string{
            strconv.Itoa(i.Id),
            i.Title.UserPreferred,
            fmt.Sprintf("%.2f", i.AverageScore),
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
        Headers("ID", "TITLE", "SCORE").
        Rows(rows...)

    fmt.Println(t)
}

var mediaSearchCmd = &cobra.Command{
    Use:   "search [query...]",
    Short: "Search for anime and manga",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            fmt.Println("No seach queries provided")
            os.Exit(0)
        }

        getMediaSearch(strings.Join(args, " "))
    },
}

func init() {
    mediaSearchCmd.Flags().StringVarP(
        &searchMediaType, "type", "t", "anime", "Type of media. for anime, pass 'anime' or 'a', for manga, use 'manga' or 'm'")
    mediaSearchCmd.Flags().IntVarP(&pageSize, "page", "p", 10, "The number of results to be returned")
}
