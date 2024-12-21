package cmd

import (
	"fmt"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/spf13/cobra"
)

func handleLoginCommand() {
	var code string

	authRequest := internal.NewAuthRequest()
	fmt.Printf(
		"Open the below link in browser to login with anilist: \n\n\n%s",
		authRequest.GetAuthURL(),
	)
	fmt.Println("\n\n\nCopy the code from the browser and enter it below:")
	fmt.Print("Enter code: ")
	fmt.Scanln(&code)

	if code == "" {
		ErrorMessage("Please provide a valid token")
	}

	err := authRequest.Login(code)
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
