package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkProductIDFile(inputFileName string, method2 bool) {
	file, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)

	productIDRanges := []string{}
	sumInvIDs := 0

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

		// fmt.Println(textLine)

		productIDRanges = append(productIDRanges, strings.Split(textLine, ",")...)
	}

	fmt.Println(len(productIDRanges), productIDRanges)

	for _, r := range productIDRanges {
		if r == "" {
			continue
		}

		rangeValuesStr := strings.Split(r, "-")
		lowerRangeEnd, err := strconv.Atoi(rangeValuesStr[0])
		if err != nil {
			panic(err)
		}
		upperRangeEnd, err := strconv.Atoi(rangeValuesStr[1])
		if err != nil {
			panic(err)
		}
		invaldIDs := []int{}

		for i := lowerRangeEnd; i <= upperRangeEnd; i++ {
			id := fmt.Sprint(i)
			idLen := len(id)

			if !method2 {
				if id[:len(id)/2] == id[len(id)/2:] {
					invaldIDs = append(invaldIDs, i)
				}
			} else {
				for x := 1; x <= idLen/2; x++ {
					if idLen % x != 0 {
						continue
					}

					c := strings.Count(id, string(id[:x]))

					if c == idLen/x {
						invaldIDs = append(invaldIDs, i)
					}
				}
			}
		}
		RemoveDuplicates(&invaldIDs)
		// fmt.Printf("Range: %v-%v, InvalidIDs: %v\n", lowerRangeEnd, upperRangeEnd, invaldIDs)

		for _, i := range invaldIDs {
			sumInvIDs += i
		}
	}

	fmt.Println(sumInvIDs)
}

func RemoveDuplicates(xs *[]int) {
	found := make(map[int]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func main() {
	fmt.Println("Method 1")
	checkProductIDFile("input_example.txt", false)
	checkProductIDFile("input_1.txt", false)

	fmt.Println("Method 2")
	checkProductIDFile("input_example.txt", true)
	checkProductIDFile("input_1.txt", true)
}
