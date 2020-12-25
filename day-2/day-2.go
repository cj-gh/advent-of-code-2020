package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SplitString(r rune) bool {
	return r == '-' || r == ' ' || r == ':'
}

func main() {

	// open file
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	// defer closing file
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	countValid1 := 0
	countValid2 := 0

	for scanner.Scan() {
		lineString := scanner.Text()
		splitString := strings.FieldsFunc(lineString, SplitString)
		lowerBound, _ := strconv.Atoi(splitString[0])
		upperBound, _ := strconv.Atoi(splitString[1])
		char := splitString[2]
		passWord := splitString[3]
		count := strings.Count(passWord, char)

		if count >= lowerBound && count <= upperBound {
			countValid1 += 1
		}

		pos1 := string(passWord[lowerBound-1]) == char
		pos2 := string(passWord[upperBound-1]) == char

		if (pos1 || pos2) && !(pos1 && pos2) {
			countValid2 += 1
		}

	}

	fmt.Println(countValid1, countValid2)
}
