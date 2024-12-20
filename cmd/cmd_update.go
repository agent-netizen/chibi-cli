package cmd

import (
    "fmt"
    "strconv"

    "github.com/CosmicPredator/chibi/internal"
    "github.com/charmbracelet/lipgloss"
    "github.com/spf13/cobra"
)


//TODO: Update progress relatively. For example "+2", "-10" etc.,
var progress int

func handleUpdate(mediaId int) {
    if progress == 0 {
        fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(
            "The flag 'progress' should be greater than 0.",
        ))
    }

    mediaUpdate := internal.NewMediaUpdate()
    err := mediaUpdate.Get(false, mediaId, progress, "", "")

    if err != nil {
        fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(
            "Some error happened. please try again...",
        ))
    }
    fmt.Println(
        SUCCESS_MESSAGE_TEMPLATE.Render(
            "Done ✅",
        ),
    )
}

var mediaUpdateCmd = &cobra.Command{
    Use: "update [id]",
    Short: "Update a list entry",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(
                "Invalid media id. please provide a valid one...",
            ))
        }
        handleUpdate(id)
    },
}

func init() {
    mediaUpdateCmd.Flags().IntVarP(
        &progress,
        "progress",
        "p",
        0,
        "The number of episodes/chapter to update",
    )
}