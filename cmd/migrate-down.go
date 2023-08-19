package cmd

import (
	"github.com/spf13/cobra"
	"go-starter-kit/pkg/bootstrap/migrate"
)

func init() {
	rootCmd.AddCommand(migrateDownCmd)
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate-down",
	Short: "Database migration down",
	Long:  "Database migration down",
	Run: func(cmd *cobra.Command, args []string) {
		migrateDown()
	},
}

func migrateDown() {
	migrate.Down()
}
