package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "shelper",
}

func Execute() {
	rootCmd.Execute()
}
