package main

import (
	"net/http"

	"github.com/spf13/cobra"

	"school_project/internal/app/router"
)

var rootCmd = &cobra.Command{
	Use:   "schoolAPI",
	Short: "Start app",
	Run: func(cmd *cobra.Command, args []string) {
		defineRouter()
	},
}

func defineRouter() {
	r := router.NewRouter()
	if err := http.ListenAndServe(":3001", r); err != nil {
		print(err)
	}
}