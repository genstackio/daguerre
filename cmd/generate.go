package cmd

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/services"
	"github.com/spf13/cobra"
	"log"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate <platform.json> [<schema>]",
	Short: "Generate diagrams/schemas for your cloud architecture",
	Args:  cobra.MinimumNArgs(2),
	Long: `Daguerre generate architecture diagrams for your cloud platform
based on a description file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := services.Generate(&commons.Order{
			Input:  args[0],
			Schema: args[1],
		}); nil != err {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
