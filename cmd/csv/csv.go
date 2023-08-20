package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"fake-data-producer/config"
	"fake-data-producer/pkg/fakedata"
	"fake-data-producer/pkg/writer"
	"fake-data-producer/version"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "csv",
	Short: "generate and produce fake data to csv",
	Long:  `generate and produce fake data to csv`,
	RunE:  cmd,
}

var (
	configDir  string
	file       string
	outputDir  string
	outputFile string
	nrMessage  string
)

func init() {
	RunCmd.Flags().StringVarP(&configDir, "config-dir", "c", ".", "Directory containing the configuration file")
	RunCmd.Flags().StringVarP(&file, "file", "f", "config.json", "Name of the configuration file")
	RunCmd.Flags().StringVarP(&outputDir, "output-dir", "o", ".", "Directory for the output CSV file")
	RunCmd.Flags().StringVarP(&outputFile, "output-filename", "n", "output.csv", "Name of the output CSV file")
	RunCmd.Flags().StringVarP(&nrMessage, "nr-messages", "m", "-1", "Number of messages to generate (-1 for infinite)")
	RunCmd.MarkFlagRequired("output-dir")
	RunCmd.MarkFlagRequired("output-filename")
	RunCmd.MarkFlagRequired("nr-messages")
	RunCmd.MarkFlagRequired("config-dir")
	RunCmd.MarkFlagRequired("file")
}

func cmd(_ *cobra.Command, _ []string) error {
	fmt.Printf("Starting the fake data generator to CSV with:\nVersion: %s %s %s\nOS/Arch: %s\nGo Version: %s\n",
		version.Version,
		version.BuildDate,
		version.GitCommit,
		version.OsArch,
		version.GoVersion)

	configData, err := config.ReadConfigFile(filepath.Join(configDir, file))
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	csvConfig := &writer.CSVConfig{
		Path:     outputDir,
		FileName: outputFile,
	}
	csvWriter, err := writer.NewWriter("csv", csvConfig)
	if err != nil {
		return fmt.Errorf("cannot create CSV writer: %w", err)
	}
	defer csvWriter.Close()

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
	var headerGenerated bool
	var sortedKeys []string

	for {
		_, value, err := fakedata.CreateData(dataStruct, configData)
		if err != nil {
			return fmt.Errorf("error creating fake data: %w", err)
		}

		var data map[string]interface{}
		err = json.Unmarshal(value, &data)
		if err != nil {
			return err
		}

		if !headerGenerated {
			// Generate and write the header row only once
			for key := range data {
				sortedKeys = append(sortedKeys, key)
			}
			sort.Strings(sortedKeys)
			csvHeader := strings.Join(sortedKeys, "|")
			err = csvWriter.Produce([]byte(csvHeader), nil)
			if err != nil {
				return fmt.Errorf("producing error: %w", err)
			}
			headerGenerated = true
		}

		var csvData []string
		for _, key := range sortedKeys {
			csvValue := fmt.Sprintf("%v", data[key])
			csvData = append(csvData, csvValue)
		}
		csvDataRow := strings.Join(csvData, "|")
		err = csvWriter.Produce(nil, []byte(csvDataRow))
		if err != nil {
			return fmt.Errorf("producing error: %w", err)
		}

		count++
		if !goInfinite && count == messageCount {
			break
		}
	}
	return nil
}
