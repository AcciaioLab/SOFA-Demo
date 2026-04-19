package parsecsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func FileOpen(fname string) (*os.File, error) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening csv data file, %v.", err)
		return nil, err
	}
	return file, err
}

func FileClose(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Printf("Error closing file: %v.\n", err)
	}
	fmt.Println("Data record file closed.")
}

func CountFileRecords(r *csv.Reader) (int, error) {
	count := 0
	var err error
	for {
		_, err = r.Read()
		if err == io.EOF {
			break
		}
		count++
	}
	// Take one off for the header
	count--

	return count, err
}
