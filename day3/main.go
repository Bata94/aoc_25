package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkJoultageFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	sumJoultage := 0

	for {
		textLine, err := scanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		textLine = strings.TrimSuffix(textLine, "\n")
		if textLine == "" {
			continue
		}

		firstNumber := ""
		firstNumberPos := 0
		joultageFound := false

		for x := 9; x >= 1; x-- {
			if joultageFound {
				break
			}
			for i := 0; i < len(textLine); i++ {
				if textLine[i:i+1] == fmt.Sprint(x) {
					if firstNumber == "" {
						firstNumber = fmt.Sprint(x)
						firstNumberPos = i
						continue
					} else if firstNumberPos < i {
						jStr := firstNumber + textLine[i:i+1]
						j, err := strconv.Atoi(jStr)
						if err != nil {
							panic(err)
						}
						sumJoultage += j
						joultageFound = true
						fmt.Println(j)
						break
					}
				}
			}
		}
	}

	fmt.Printf("FileName: %s, SumJoultage: %v\n", fileName, sumJoultage)
}

func main() {
	checkJoultageFile("input_example.txt")
	// checkJoultageFile("input_1.txt")
}
