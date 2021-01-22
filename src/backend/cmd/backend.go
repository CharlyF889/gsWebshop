package cmd

import (
	"fmt"

	"gosvelteHex/src/backend/pkg/handlers"
	"gosvelteHex/src/backend/pkg/providers"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "backend",
	Short: "Run backend service",
	Long:  `run backend service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backend service is running")

		e := echo.New()

		handlers.RegisterRoute(e, providers.InitProvider())

		e.Start(":8181")
	},
}
