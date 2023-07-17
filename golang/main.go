package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var filePath1 = "../database/file1.ndjson"

func main() {
	// init schedule time
	start := time.Now()

	// open the NDJSON file
	file, err := os.Open(filePath1)
	if err != nil {
		panic(err)
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

		// print the out
		for k, v := range obj {
			fmt.Printf("%s: %v\n", k, v)
		}
	}

	// check if there is any error while reading the file
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// print the time taken to read the file
	fmt.Println("time taken:", time.Since(start))
}
