package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)


var RootCmd = &cobra.Command {
	Use: "aiops",
	RunE: func(c *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}