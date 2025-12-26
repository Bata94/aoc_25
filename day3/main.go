package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
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

		joultage := 0
		joultageLS := []int{}
		joultageLSSorted := []int{}

		for _, char := range textLine {
			joultageLS = append(joultageLS, int(char-'0'))
		}

		joultageLSSorted = append(joultageLSSorted, joultageLS...)
		sort.Ints(joultageLSSorted)
		slices.Reverse(joultageLSSorted)

		if joultageLSSorted[0] == joultageLSSorted[1] {
			joultage = joultageLSSorted[0] * 10 + joultageLSSorted[1]
		} else {
			indexFirstNumber := slices.IndexFunc(joultageLS, func(i int) bool {
				return i == joultageLSSorted[0]
			})
			indexSecondNumber := slices.IndexFunc(joultageLS, func(i int) bool {
				return i == joultageLSSorted[1]
			})

			if indexFirstNumber < indexSecondNumber {
				joultage = joultageLSSorted[0] * 10 + joultageLSSorted[1]
			} else if indexFirstNumber == len(joultageLSSorted)-1 {
				joultage = joultageLSSorted[1] * 10 + joultageLSSorted[0]
			} else {
				subJoultageLS := joultageLS[indexFirstNumber+1:]
				subJoultageLSSorted := []int{}
				subJoultageLSSorted = append(subJoultageLSSorted, subJoultageLS...)
				sort.Ints(subJoultageLSSorted)
				slices.Reverse(subJoultageLSSorted)

				joultage = joultageLSSorted[0] * 10 + subJoultageLSSorted[0]
			}
		}

		sumJoultage += joultage
		// fmt.Println(textLine, joultageLS, joultageLSSorted, joultage)
	}

	fmt.Printf("FileName: %s, SumJoultage: %v\n", fileName, sumJoultage)
}

func main() {
	checkJoultageFile("input_example.txt")
	checkJoultageFile("input_1.txt")
}
