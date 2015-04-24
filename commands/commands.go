package commands

import (
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "gognutil",
	Short: "Gognutil are some of the GNU Core Utils implemented in Go.",
	Long:  `Gognutil are some of the GNU Core Utils implemented in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	mainCmd.AddCommand(seqCmd())
}

// Execute is the main entry point.
func Execute() {
	mainCmd.Execute()
}
