package writer

import (
	"fmt"
)

type ConsoleWriter struct {
}

func NewConsoleWriter() (WriterInterface, error) {
	return &ConsoleWriter{}, nil
}

func (cw *ConsoleWriter) Produce(headers []byte, data []byte) error {
	fmt.Println(string(data))
	return nil
}

func (cw *ConsoleWriter) Close() {}

// Flush is not applicable for CSV, return 0
func (cw *ConsoleWriter) Flush(i int) int {
	return 0
}
