package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

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

	// read all integers
	var intSlice []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		intSlice = append(intSlice, i)
	}

	// check for errors when reading
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// sort the slice
	sort.Ints(intSlice)

	// solve puzzle
	for _, n := range intSlice {
		for _, m := range intSlice {
			for _, k := range intSlice {
				if n+m+k == 2020 {
					fmt.Println(n, m, k, n*m*k)
				}
			}
		}
	}
}
