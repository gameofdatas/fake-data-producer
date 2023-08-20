package cmd

import (
	"log"

	console "fake-data-producer/cmd/console"
	csv "fake-data-producer/cmd/csv"
	"fake-data-producer/cmd/kafka"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fake-produce",
	Short: "generate and produce fake data",
	Long:  `produce fake data to different queues like kafka/rabbitmq`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(kafka.RunCmd)
	rootCmd.AddCommand(csv.RunCmd)
	rootCmd.AddCommand(console.RunCmd)
}
