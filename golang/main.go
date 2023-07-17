package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var filePath1 = "../database/file1.ndjson"
var filePath2 = "../database/file2.ndjson"
var filePath3 = "../database/file3.ndjson"

func main() {
	// init schedule time
	start := time.Now()

	// create wg to wait for all the goroutines to finish
	var wg sync.WaitGroup

	// create a channel to receive the data
	dataCh := make(chan map[string]interface{})

	// especify the file path
	filePaths := []string{filePath1, filePath2, filePath3}

	// init a goroutine for each file
	for _, filePath := range filePaths {
		wg.Add(1)
		go func(filePath string) {
			// read the file
			readNDJSON(filePath, dataCh)
			wg.Done()
		}(filePath)
	}

	// init a goroutine to close the channel when all the files are read
	go func() {
		wg.Wait()
		close(dataCh)
	}()

	// process the received data
	for data := range dataCh {
		// process the data
		for key, value := range data {
			fmt.Println(key, value)
		}
	}

	// print the time taken to read the file
	fmt.Println("time taken:", time.Since(start))
}

// readNDJSON reads the NDJSON file and returns the it for the channel
func readNDJSON(filePah string, dataCh chan<- map[string]interface{}) {
	// open the NDJSON file
	file, err := os.Open(filePah)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// read the NDJSON file line by line
	for scanner.Scan() {
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
