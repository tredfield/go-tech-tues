package cmd

import (
	"github.com/spf13/cobra"

	"go-tech-tues/pkg/api"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		api.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
