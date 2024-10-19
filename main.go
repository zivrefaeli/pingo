package main

import (
	"os"
	"pingo/packet"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "pingo",
		RunE: func(cmd *cobra.Command, args []string) error {
			targetName := args[0]
			return packet.StartPinging(targetName, 4, 32)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
