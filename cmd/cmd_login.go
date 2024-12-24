package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func handleLoginCommand() {
	var code string

	authRequest := internal.NewAuthRequest()
	fmt.Printf("Open the below link in browser to login with anilist: \n\n")

	fmt.Print(
		OTHER_MESSAGE_TEMPLATE.Render(authRequest.GetAuthURL()),
	)

	fmt.Print("\n\n")

	huh.NewText().
		Title("Paste your token here:").
		CharLimit(2000).
		Value(&code).
		Validate(func(s string) error {
			if s == "" {
				return errors.New("please provide a valid token")
			}
			return nil
		}).
		Run()

	if code == "" {
		fmt.Println(
			ERROR_MESSAGE_TEMPLATE.Render("please provide a valid token"),
		)
		os.Exit(0)
	}

	err := authRequest.Login(strings.TrimSpace(code))
	if err != nil {
		ErrorMessage(err.Error())
	}
	fmt.Println("Login Successful")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with anilist",
	Run: func(cmd *cobra.Command, args []string) {
		handleLoginCommand()
	},
}
