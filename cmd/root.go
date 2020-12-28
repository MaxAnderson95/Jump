package cmd

import (
	"fmt"
	"os"
	"os/user"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//RootCmd is the root command
var RootCmd = &cobra.Command{
	Use:   "Jump",
	Short: "An example of cobra",
	Long: `This application shows how to create modern CLI 
applications in go using Cobra CLI library`,
}

var hosts map[string]string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	usr, err := user.Current()
    if err != nil {
        panic(err)
    }

	viper.SetConfigName("JumpConfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(usr.HomeDir)

	err2 := viper.ReadInConfig()
	if err2 != nil {
		fmt.Println("No configuration imported:\n", err)
	}

	hosts = viper.GetStringMapString("Hosts")
}