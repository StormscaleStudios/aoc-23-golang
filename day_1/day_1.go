package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse_line(line string) (value int, err error) {
	numbers := "1234567890"
	
	first_index := strings.IndexAny(line, numbers)
	last_index := strings.LastIndexAny(line, numbers)
	
	if (first_index == -1) || (last_index == -1) {
		return -1, errors.New("Unable to extract value from string")
	}

	code := fmt.Sprintf("%s%s", string(line[first_index]), string(line[last_index]))

	i, err := strconv.Atoi(code)
	if err != nil {
		log.Fatal("Failed to convert code to integer value")
	}

	return i, nil
}

func part_one(lines []string) int {

	var codes []int
	for i := 0; i < len(lines); i++ {
		data, err := parse_line(lines[i])
		if err != nil {
			continue
		}
		fmt.Println(data)
		codes = append(codes, data)
	} 

	var code_sum int
	for i := 0; i < len(codes); i++ {
		code_sum += codes[i]
		fmt.Println(code_sum)
	}

	return code_sum
}

func match_values(line string, index int) (match *string, err error) {
	valid_numbers := [18]string{"1", "one", "2", "two", "3", "three", "4", "four", "5", "five", "6", "six", "7", "seven", "8", "eight", "9", "nine"}
	for i := 0; i < 18; i++ {
		number := valid_numbers[i]
		if (index + len(number)) > len(line) {
			continue
		}
		line_slice := line[index:index + len(number)]

		if line_slice == number {
			return &line_slice, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No match found at index %d", index))
}

func convert_to_value(match string) string {
	numbers := "1234567890"
	check := strings.IndexAny(match, numbers)
	if check != -1 {
		return match
	}

	switch m := match; m {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}

func parse_line_complex(line string) (value *int, err error) {
	
	fmt.Println(line)

	var matches []string
	for line_index := 0; line_index < len(line); line_index++ {
		match, err := match_values(line, line_index)
		if err != nil {
			continue
		}
		matches = append(matches, *match)

	}

	fmt.Println(matches)
	
	if len(matches) == 0 {
		return nil, errors.New("No matches found!")
	}

	first_match := matches[0]
	last_match := matches[len(matches) - 1]

	fmt.Printf("%s %s\n", first_match, last_match)

	first_code := convert_to_value(first_match)
	last_code := convert_to_value(last_match)
	
	code := fmt.Sprintf("%s%s", first_code, last_code)
	i, err := strconv.Atoi(code)
	if err != nil {
		log.Fatal("Failed to convert code to integer value")
	}
	return &i, nil
}

func part_two(lines []string) int {

	var codes []int
	for i := 0; i < len(lines); i++ {
		data, err := parse_line_complex(lines[i])
		if err != nil {
			continue
		}
		fmt.Println(data)
		codes = append(codes, *data)
	} 

	var code_sum int
	for i := 0; i < len(codes); i++ {
		code_sum += codes[i]
		fmt.Println(code_sum)
	}

	return code_sum
}

func main() {
	filename := "input_a"

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	//part_one_output := part_one(lines)
	//fmt.Printf("Answer to part One is '%d'\n", part_one_output)

	part_two_output := part_two(lines)
	fmt.Printf("Answer to part Two is '%d'\n", part_two_output)

}


