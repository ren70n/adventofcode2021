package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	dataFile := "data1.txt"

	data := readFile(dataFile)

	prev,_ := strconv.Atoi(data[0])

	incCnt :=0

	for _, x := range data[1:] {
		value, _ := strconv.Atoi(x)
		
		// fmt.Println(value)
		if value>prev{
			incCnt++
		}
		prev = value

	}

	fmt.Println(incCnt)
}

func readFile(name string)[]string{

	file, err := os.Open(name)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var data []string

	
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
//https://adventofcode.com/2020/day/1/input
