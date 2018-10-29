package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Washi",
	Short: "Start our Washi server for Authentication",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

	}
}
