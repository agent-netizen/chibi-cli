package cmd

import (
    "fmt"
    "strconv"

    "github.com/CosmicPredator/chibi/internal"
    "github.com/charmbracelet/lipgloss"
    "github.com/spf13/cobra"
)

func getUserProfile() {
    profile := internal.NewProfile()
    err := profile.Get()
    if err != nil {
        panic(err)
    }

    keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF79C6"))
    valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD"))

    fmt.Printf("%-20s : %s\n", keyStyle.Render("ID"), valueStyle.Render(strconv.Itoa(profile.Data.Viewer.Id)))
    fmt.Printf("%-20s : %s\n", keyStyle.Render("Name"), valueStyle.Render(profile.Data.Viewer.Name))
    fmt.Printf("%-20s : %s\n", keyStyle.Render("Total Anime"), valueStyle.Render(strconv.Itoa(profile.Data.Viewer.Statistics.Anime.Count)))
    fmt.Printf("%-20s : %s\n", keyStyle.Render("Total Manga"), valueStyle.Render(strconv.Itoa(profile.Data.Viewer.Statistics.Manga.Count)))
    fmt.Printf("%-20s : %s\n", keyStyle.Render("Total Days Watched"), valueStyle.Render(fmt.Sprintf("%.2f", float32(profile.Data.Viewer.Statistics.Anime.MinutesWatched)/1440)))
    fmt.Printf("%-20s : %s\n", keyStyle.Render("Total Chapters Read"), valueStyle.Render(strconv.Itoa(profile.Data.Viewer.Statistics.Manga.ChaptersRead)))
    fmt.Printf("%-20s : %s\n", keyStyle.Render("URL"), valueStyle.Render(profile.Data.Viewer.SiteUrl))
}

var profileCmd = &cobra.Command{
    Use:   "profile",
    Short: "Get's your AniList profile (requires login)",
    Run: func(cmd *cobra.Command, args []string) {
        getUserProfile()
    },
}
