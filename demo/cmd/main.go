package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"sofa-demo/internal/calc"
	"sofa-demo/internal/db"
	"sofa-demo/internal/parsecsv"
	"time"
)

const demoTime int = 5 // Time in minutes for one run of the demo

func main() {
	fmt.Println("SOFA Demo App - Starting...")

	// Connect to DB
	dbpool, err := db.CreateDBPool()
	if err != nil {
		fmt.Printf("Error connecting to DB, %v.", err)
		return
	}
	defer db.CloseDBPool(dbpool)

	// Load the trailer data
	trailers := db.GetTrailerInfo(dbpool)
	fmt.Println(trailers)

	// Open CSV file
	file, err := parsecsv.FileOpen("./data/testdata.csv")
	if err != nil {
		fmt.Printf("Error opening data file, %v.\n", err)
		return
	}
	defer parsecsv.FileClose(file)

	// Create the csv reader
	reader := csv.NewReader(file)

	// Count the number of data records in the file
	recordCount, err := parsecsv.CountFileRecords(reader)
	fmt.Printf("Number of records in the file is %d.\n", recordCount)

	// Rewind the csv reader to the start of the file
	file.Seek(0, io.SeekStart)

	// Burn the header
	_, err = reader.Read()
	if err != nil {
		fmt.Printf("Data read error: %v.", err)
	}

	// Calculate time interval to insert each record over the demo time
	demoTimeMilli := demoTime * 60 * 1000
	tickTime := demoTimeMilli / recordCount

	// Set up ticker to insert data into the DB
	ticker := time.NewTicker(time.Duration(tickTime) * time.Millisecond)
	dataFinished := make(chan bool)

	// Create the trailer load variables used to store calculations
	var trailerLoad, trailerLoadLast calc.TrailerLoad
	var trailerInput calc.TrailerInput

	// Demo loop
	go func() {
		count := 0
		for {
			select {
			case <-ticker.C:
				// Read the next record
				record, err := reader.Read()
				if err != nil {
					fmt.Printf("Data read error: %v.\n", err)
				} else {
					// Parse the current record into struct
					err := db.InsertOneTrailerInput(dbpool, record)
					if err != nil {
						fmt.Printf("Error inserting input into db, %v.\n", err)
					} else {
						// DEBUG
						count++
						fmt.Printf("Insert DB trailer input record %v.\n", count)
					}

					// Convert csv record to trailerInput struct
					trailerInput.UnmarshalRecord(record)

					// Run the calculation for the inserted record.
					go func() {
						t := trailers[trailerInput.VIN]
						trailerInput.CalcLoadsFromInput(&trailerLoad, &trailerLoadLast, &t)
						trailerLoadLast = trailerLoad

						err := db.InsertOneTrailerLoad(dbpool, &trailerLoad)
						if err != nil {
							fmt.Printf("Error inserting load into db, %v.\n", err)
						} else {
							fmt.Printf("Insert DB trailer load record %v.\n", count)
						}
					}()

				}

			// All records have been read
			case <-dataFinished:
				return
			}
		}
	}()

	// Sleep is calculated to be the end of the record inputs
	time.Sleep(time.Duration(demoTimeMilli) * time.Millisecond)
	ticker.Stop()
	dataFinished <- true

	fmt.Println("Demo complete")

}
