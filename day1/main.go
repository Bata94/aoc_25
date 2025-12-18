package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkPwd(fileName string, method2 bool) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewReader(file)

	lineCounter := 0
	dialEnds := []int{0,99}
	dialPos := 50
	pwdCounter := 0

	for {
		textLine, err := scanner.ReadString('\n')
		if err == io.EOF {
				if len(textLine) != 0 {
						fmt.Print(textLine) // Print last line if not empty
				}
				break
		}
		if err != nil {
				panic(err)
		}

		textLine = strings.TrimSuffix(textLine, "\n")
		clicksAbs, err := strconv.Atoi(textLine[1:])
		if err != nil {
			panic(err)
		}

		// fmt.Printf("DialPos: %v, Clicks: %s %v, PwdCounter: %v\n", dialPos, textLine[:1], clicksAbs, pwdCounter)

		var clicks int = clicksAbs % 100
		var rotation int = clicksAbs / 100
		dialPosOld := dialPos

		if method2 && rotation != 0 {
			pwdCounter += rotation
		}

		switch textLine[0] {
			case 'L':
				dialPos -= clicks
			case 'R':
				dialPos += clicks
			default:
				panic(fmt.Sprint("Invalid input: ", textLine[0]))
		}

		if dialPos < dialEnds[0] {
			dialPos = dialEnds[1] + dialPos + 1
			if method2 && dialPos != 0 && dialPosOld != 0 {
				pwdCounter += 1
			}
		} else if dialPos > dialEnds[1] {
			dialPos = dialEnds[0] + dialPos - dialEnds[1] - 1
			if method2 && dialPos != 0 && dialPosOld != 0 {
				pwdCounter += 1
			}
		}

		if dialPos < dialEnds[0] || dialPos > dialEnds[1] {
			panic(fmt.Sprint("Invalid dialPos: ", dialPos))
		}

		if dialPos == 0 {
			pwdCounter += 1
		}

		lineCounter += 1
	}

	fmt.Printf("InputFile: %s, Pwd: %v, with endposition: %v, read Lines %v\n", fileName, pwdCounter, dialPos, lineCounter)
}

func main() {
	fmt.Println("Method 1")

	checkPwd("input_example.txt", false)
	checkPwd("input_1.txt", false)

	fmt.Println("Method 2")

	checkPwd("input_example.txt", true)
	checkPwd("input_1.txt", true)
}
