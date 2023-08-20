package cmd

import (
	"fmt"
	"path/filepath"
	"strconv"

	"fake-data-producer/config"
	"fake-data-producer/pkg/fakedata"
	"fake-data-producer/pkg/writer"
	"fake-data-producer/version"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "console",
	Short: "generate and produce fake data to console",
	Long:  `generate and produce fake data to console`,
	RunE:  cmd,
}

var (
	configDir string
	file      string
	nrMessage string
)

func init() {
	RunCmd.Flags().StringVarP(&configDir, "config-dir", "c", ".", "Directory containing the configuration file")
	RunCmd.Flags().StringVarP(&file, "file", "f", "config.json", "Name of the configuration file")
	RunCmd.Flags().StringVarP(&nrMessage, "nr-messages", "m", "-1", "Number of messages to generate (-1 for infinite)")
	RunCmd.MarkFlagRequired("nr-messages")
	RunCmd.MarkFlagRequired("config-dir")
	RunCmd.MarkFlagRequired("file")
}

func cmd(_ *cobra.Command, _ []string) error {
	fmt.Printf("Starting the fake data generator to Console with:\nVersion: %s %s %s\nOS/Arch: %s\nGo Version: %s\n",
		version.Version,
		version.BuildDate,
		version.GitCommit,
		version.OsArch,
		version.GoVersion)

	configData, err := config.ReadConfigFile(filepath.Join(configDir, file))
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	consoleWriter, err := writer.NewWriter("console", nil)
	if err != nil {
		return fmt.Errorf("cannot create console writer: %w", err)
	}

	messageCount, err := strconv.Atoi(nrMessage)
	if err != nil {
		return err
	}

	dataStruct, err := fakedata.CreateStructFromJSON(configData)
	if err != nil {
		return err
	}

	count := 0
	goInfinite := messageCount <= 0

	for {
		_, value, err := fakedata.CreateData(dataStruct, configData)
		if err != nil {
			return fmt.Errorf("error creating fake data: %w", err)
		}

		err = consoleWriter.Produce(nil, value)
		if err != nil {
			return err
		}

		count++
		if !goInfinite && count == messageCount {
			break
		}
	}
	return nil
}
