package main

import (
	"os"
	"pingo/packet"

	"github.com/spf13/cobra"
)

func main() {
	var echoRequestsCount int
	var bufferSize uint16

	rootCmd := &cobra.Command{
		Use:  "pingo [TARGET_NAME]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			targetName := args[0]
			return packet.StartPinging(targetName, echoRequestsCount, bufferSize)
		},
		SilenceUsage: true,
	}

	rootCmd.Flags().IntVarP(&echoRequestsCount, "count", "n", 4, "Number of echo requests to send.")
	rootCmd.Flags().Uint16VarP(&bufferSize, "size", "l", 32, "Send buffer size.")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
