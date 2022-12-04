package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func partA(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	alphaB := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	total := 0
	var commons []string
	common := ""

	for scanner.Scan() {
		strData := scanner.Text()
		midRucksack := (len(strData) / 2)
		compartment1 := strData[:midRucksack]
		compartment2 := strData[midRucksack:]

		for i := 0; i < len(compartment1); i++ {
			common = ""
			comLen := len(common)
			for j := 0; j < len(compartment2); j++ {
				if compartment1[i] == compartment2[j] {
					common += compartment1[i : i+1]
				}
			}
			if comLen < len(common) {
				commons = append(commons, common[0:1])
				break
			}
		}
	}
	for term := 0; term < len(commons); term++ {
		total += strings.Index(alphaB, commons[term]) + 1
	}
	fmt.Println("Total: ", total)

}

func partB(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	alphaB := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	total := 0
	common := ""
	var commons []string
	line1 := ""
	line2 := ""
	line3 := ""
	lineCount := 0

	for scanner.Scan() {
		strData := scanner.Text()
		lineCount += 1

		if lineCount == 1 {
			line1 = strData
		}

		if lineCount == 2 {
			line2 = strData
		}

		if lineCount == 3 {
			line3 = strData
		}

		if lineCount == 3 {
			lineCount = 0
			for i := 0; i < len(line1); i++ {
				common = ""
				for j := 0; j < len(line2); j++ {
					for k := 0; k < len(line3); k++ {
						if line1[i] == line2[j] && line1[i] == line3[k] {
							common += string(line1[i])
						}
					}
				}
				if 0 < len(common) {
					commons = append(commons, common[0:1])
					break
				}
			}
		}
	}
	for term := 0; term < len(commons); term++ {
		total += strings.Index(alphaB, commons[term]) + 1
	}
	fmt.Println("Total: ", total)
}

func main() {
	partA("day03.txt")
	partB("day03.txt")
}
