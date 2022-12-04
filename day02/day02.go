package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RpsCounter(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	my_moves := "XYZ"
	my_total := 0

	options := map[string]map[string]string{
		"X": {
			"strong": "C",
			"weak":   "B",
		},
		"Y": {
			"strong": "A",
			"weak":   "C",
		},
		"Z": {
			"strong": "B",
			"weak":   "A",
		},
	}

	for scanner.Scan() {
		strData := scanner.Text()
		if options[strData[2:3]]["strong"] == strData[0:1] {
			my_total += 6
		}
		if options[strData[2:3]]["strong"] != strData[0:1] && options[strData[2:3]]["weak"] != strData[0:1] {
			my_total += 3
		}
		my_total += strings.Index(my_moves, strData[2:3]) + 1
	}
	fmt.Println("My Total: ", my_total)
}

func RpsCounter2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	my_total := 0

	options := map[string]map[string]int{
		"A": {
			"win":  2,
			"lose": 3,
			"draw": 1,
		},
		"B": {
			"win":  3,
			"lose": 1,
			"draw": 2,
		},
		"C": {
			"win":  1,
			"lose": 2,
			"draw": 3,
		},
	}

	for scanner.Scan() {
		strData := scanner.Text()
		if strData[2:3] == "X" {
			my_total += options[strData[0:1]]["lose"]
		}
		if strData[2:3] == "Y" {
			my_total += 3
			my_total += options[strData[0:1]]["draw"]

		}
		if strData[2:3] == "Z" {
			my_total += 6
			my_total += options[strData[0:1]]["win"]

		}
		fmt.Println("Current Line: ", strData)
		fmt.Println("My Total: ", my_total)

	}

}

func main() {
	RpsCounter2("day02.txt")
}
