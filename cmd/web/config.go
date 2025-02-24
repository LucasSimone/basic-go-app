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
func setEnvConfig() {
	// Open the .env file
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	// Defer file to be closed on function end
	defer file.Close()

	// Scanner creation
	fileScanner := bufio.NewScanner(file)

	// Read in the .env file line by line
	for fileScanner.Scan() {
		//Find the first = to split the line
		splitIndex := strings.Index(fileScanner.Text(),"=")
		if splitIndex > -1 {
			//Split line into Key and value removing any surrounding whitespace
			key := strings.TrimSpace(fileScanner.Text()[:splitIndex])
			value := strings.TrimSpace(fileScanner.Text()[splitIndex+1:])
			//Set the variable as a os env variable
			setEnvVariable(key, value)
		}else{
			log.Printf("Invalid line in .env: no = found")
		}		
		
	}

	if err := fileScanner.Err(); err != nil {
		log.Println(err)
	}
}

// Takes a key and a value and sets an according
// os enviroment variable
func setEnvVariable(key, value string) {
	err := os.Setenv(key, trimQuotes(value))

	if err != nil {
		log.Fatal(err)
	}
}

// Takes a string and trims any prefix and suffix quotation marks
func trimQuotes(str string) string {
	if str[0] == '\'' || str[0] == '"' {
		str = str[1 : len(str)-1]
	}
	return str
}
