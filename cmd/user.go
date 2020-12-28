package cmd

import (
	"fmt"
	"os/user"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Sets gets or clears the user for Jumping",
	Long:  `This command sets gets or clears the user for Jumping`,
}

var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the current user configured for jumps",
	Long:  `Gets the current user configured for jumps`,
	Run:   func(cmd *cobra.Command, args []string) {
		fmt.Println("The current user for connections is:", getJumpUser())
		fmt.Println("Use 'Jump User Set [username]' to set the username")
	},
}

var userSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the username for jumps",
	Long:  "Sets the username for jumps",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {
		setJumpUser(args[0])
	},
}

var userClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the username for jumps",
	Long:  "Clears the username for jumps, and sets back to the default of the currently logged in user.",
	Run:   func(cmd *cobra.Command, args []string) {
		clearJumpUser()
	},
}

func getJumpUser() string {
	configUser := viper.GetString("User")
	if configUser == "" {
		return getOSUser()
	}
	return configUser
}

func setJumpUser(u string) {
	viper.Set("User", u)
	viper.WriteConfig()
}

func clearJumpUser() {
	setJumpUser("")
}

func getOSUser() string {
	osUser, err := user.Current()
    if err != nil {
        panic(err)
    }
	return osUser.Username
}

func init() {
	RootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userGetCmd)
	userCmd.AddCommand(userSetCmd)
	userCmd.AddCommand(userClearCmd)
}