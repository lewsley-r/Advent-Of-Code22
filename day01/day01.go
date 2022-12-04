package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadingCalories(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf_score := 0
	elf_number := 0
	counter := 0
	current_score := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if elf_score < current_score {
				elf_score = current_score
				elf_number = counter
			}
			current_score = 0
			counter += 1
		}
		current_line := scanner.Text()
		num, _ := strconv.Atoi(current_line)
		current_score += num
	}
	fmt.Println("elf_score: ", elf_score)
	fmt.Println("elf_index: ", elf_number)

}

func topThreeCalories(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf_score_1 := 0
	elf_score_2 := 0
	elf_score_3 := 0

	current_score := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if elf_score_1 < current_score {
				elf_score_3 = elf_score_2
				elf_score_2 = elf_score_1
				elf_score_1 = current_score

			} else if elf_score_2 < current_score {
				elf_score_2 = current_score

			} else if elf_score_3 < current_score {
				elf_score_3 = current_score
			}
			fmt.Println(current_score)
			current_score = 0
		}
		current_line := scanner.Text()
		num, _ := strconv.Atoi(current_line)
		current_score += num
	}
	top_elves_total := elf_score_1 + elf_score_2 + elf_score_3

	fmt.Println("elf_score_1: ", elf_score_1)
	fmt.Println("elf_score_2: ", elf_score_2)
	fmt.Println("elf_score_3: ", elf_score_3)
	fmt.Println("TOP 3 ELVES TOTAL: ", top_elves_total)

}

func main() {
	// ReadingCalories("day1_a_input.txt")
	topThreeCalories("day1_a_input.txt")

}
