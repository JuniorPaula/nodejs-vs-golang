package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var filePath1 = "../database/file1.ndjson"
var filePath2 = "../database/file2.ndjson"
var filePath3 = "../database/file3.ndjson"
var outputFilePath = "../database/output-gmail.ndjson"

func main() {
	// init schedule time
	start := time.Now()

	// create wg to wait for all the goroutines to finish
	var wg sync.WaitGroup

	// create a channel to receive the data
	dataCh := make(chan map[string]interface{})

	// especify the file path
	filePaths := []string{filePath1, filePath2, filePath3}

	// create a semaphore to limit the number of goroutines
	sem := make(chan struct{}, 3)

	// init a goroutine for each file
	for _, filePath := range filePaths {
		wg.Add(1)
		go func(filePath string) {

			// read the file
			sem <- struct{}{}
			readNDJSON(filePath, dataCh)
			<-sem
			wg.Done()
		}(filePath)
	}

	// init a goroutine to close the channel when all the files are read
	go func() {
		wg.Wait()
		close(dataCh)
	}()

	// process the received data
	var filteredData []map[string]interface{}
	for data := range dataCh {
		// check if the email is from gmail
		if email, ok := data["email"].(string); ok {
			if !strings.HasSuffix(email, "@gmail.com") {
				continue
			}
		}

		// append the data to the filtered data
		filteredData = append(filteredData, data)
	}

	// save the filtered data to a file
	err := saveFilteredData(filteredData)
	if err != nil {
		log.Fatal("error to save data", err)
	}

	// print the time taken to read the file
	fmt.Println("time taken:", time.Since(start))
}

// readNDJSON reads the NDJSON file and returns the it for the channel
func readNDJSON(filePath string, dataCh chan<- map[string]interface{}) {
	// open the NDJSON file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create file NDJSON of the filtered emails
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	var counter int

	// read the NDJSON file line by line
	for scanner.Scan() {
		counter++
		fmt.Println(filePath, " found ", counter, " so far")
		// get the line
		line := scanner.Bytes()

		// parse the line into a JSON object
		var obj map[string]interface{}
		err := json.Unmarshal(line, &obj)
		if err != nil {
			log.Println("error to parse json", err)
			continue
		}

		// send the object to the channel
		dataCh <- obj
	}

	// check if there is any error while reading the file
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func saveFilteredData(filteredData []map[string]interface{}) error {
	// create a file to save the filtered data
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// create the writer
	writer := bufio.NewWriter(outputFile)

	// write the data to the file
	for _, data := range filteredData {
		// parse the data to JSON
		line, err := json.Marshal(data)
		if err != nil {
			return err
		}

		_, err = writer.Write(line)
		if err != nil {
			return err
		}
		writer.WriteString("\n")
	}
	// flush the writer
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
