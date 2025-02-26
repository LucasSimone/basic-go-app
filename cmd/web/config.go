package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Checks if a go enviroment variable exists. If it does it
// returns the enviroment variable. Else it returns the default value
func getEnv(key, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}

// Reads in the .env variables and sets them as os enviroment
// varibales
func setEnvConfig(path string) {

	// Open the .env file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open ", path, "file.")
	}
	// Defer file to be closed on function end
	defer file.Close()

	// Scanner creation
	fileScanner := bufio.NewScanner(file)

	// Read in the .env file line by line
	for fileScanner.Scan() {
		// Get the next line
		line := fileScanner.Text()

		// Check if the line is empty or a comment
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into the key and value parts
		key, value, found := strings.Cut(line, "=")

		if !found {
			log.Println("Invalid line in .env: ", line)
		}

		// Trim any surroundng white space
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		// Time any quotations from the value
		value = strings.Trim(value, `"'`)

		// Set the enviroment variable and log any errors
		err := os.Setenv(key, value)
		if err != nil {
			log.Fatal(err)
		}

	}

	if err := fileScanner.Err(); err != nil {
		log.Println(err)
	}
}
