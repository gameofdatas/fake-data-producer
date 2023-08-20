package writer

import (
	"encoding/csv"
	"os"
	"strings"
)

type CSVWriter struct {
	file   *os.File
	writer *csv.Writer
}

func NewCSVWriter(csvConfig *CSVConfig) (WriterInterface, error) {
	filePath := csvConfig.Path + "/" + csvConfig.FileName
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)

	return &CSVWriter{
		file:   file,
		writer: writer,
	}, nil
}

func (cw *CSVWriter) Produce(headers []byte, data []byte) error {
	if len(headers) > 0 {
		cw.writer.Write(strings.Split(string(headers), "|"))
	}

	if len(data) > 0 {
		cw.writer.Write(strings.Split(string(data), "|"))
	}

	return nil
}

func (cw *CSVWriter) Close() {
	cw.writer.Flush()
	cw.file.Close()
}

// Flush is not applicable for CSV, return 0
func (cw *CSVWriter) Flush(i int) int {
	return 0
}
