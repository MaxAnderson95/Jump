package cmd

import (
	"strings"
	"fmt"
	"os"
	"os/exec"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connects to a host",
	Long:  `This command connects to a given host`,
	Run: func(cmd *cobra.Command, args []string) {
		connect(args[0])
	},
	Args: cobra.ExactArgs(1),
}

func connect(h string) {
	jumpUser := getJumpUser()
	connStr := jumpUser + "@" + hosts[strings.ToLower(h)]

	fmt.Println("Attempting to connect to", h, "...")
	cmd := exec.Command("ssh", connStr)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func init() {
	RootCmd.AddCommand(connectCmd)
}