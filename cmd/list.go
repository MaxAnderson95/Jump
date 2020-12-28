package cmd

import (
	"strings"
	"fmt"
	"sort"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists hosts",
	Long:  `This command lists the available jump hosts`,
	Run: func(cmd *cobra.Command, args []string) {
		listHosts()
	},
}

func getHosts() []string {
	keys := make([]string, 0)
	for k := range hosts {
		keys = append(keys, strings.ToUpper(k))
	}
	sort.Strings(keys)
	return keys
}

func listHosts() {
	for _, k := range getHosts() {
		fmt.Println(k)
	}
}

func init() {
	RootCmd.AddCommand(listCmd)
}