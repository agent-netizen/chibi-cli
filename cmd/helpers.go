package cmd

import (
	"fmt"
	"os"

	"github.com/CosmicPredator/chibi/types"
)

func CheckIfTokenExists() {
	message := ERROR_MESSAGE_TEMPLATE.Render("Not logged in. Please login with anilist to continue...")
	accessToken := types.NewTokenConfig()
	if err := accessToken.ReadFromJsonFile(); err != nil {
		fmt.Println(message)
		os.Exit(0)
	}

	if accessToken.AccessToken == "" {
		fmt.Println(message)
		os.Exit(0)
	}
}

func ErrorMessage(errorString string) {
	fmt.Println(
		ERROR_MESSAGE_TEMPLATE.Render(
			fmt.Sprintf("An internal error occured: %s", errorString),
		),
	)
	os.Exit(0)
}
