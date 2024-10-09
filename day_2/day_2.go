package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func extract_games(line string) (index int, sets []string) {
	parts := strings.Split(line, ":")
	if len(parts) == 1 {
		fmt.Printf("Found single length line: '%s'\n", line)
		var v []string
		return 0, v
	}
	part_one := parts[0]

	game_label, err := strconv.Atoi(strings.Split(part_one, " ")[1])
	if err != nil {
		log.Fatalf("Failed to extract game number from string: '%s'", part_one)
	}

	return game_label, strings.Split(parts[1], ";")
}

func validate_set(set string) bool {
	entries := strings.Split(set, ",")
	var colour_limits = map[string]int {
		"red":12,
		"green":13,
		"blue":14,
	}

	fmt.Println(set)

	for i := 0; i < len(entries); i++ {
		entry := entries[i]
		bits := strings.Split(entry, " ")
		count, _ := strconv.Atoi(bits[1])
		if count > colour_limits[bits[2]] {
			fmt.Printf("Found colour '%s' with value '%d', exceeds limit '%d'\n", bits[2], count, colour_limits[bits[2]])
			return false
		}
	}
	return true
}

func parse_line(line string) (value *int, is_valid bool) {
	game, sets := extract_games(line)
	
	for i := 0; i < len(sets); i++ {
		set := sets[i]
		is_valid := validate_set(set)
		if !is_valid {
			fmt.Printf("Line is not valid: %s\n", line)
			return nil, false
		}
	}
	fmt.Printf("Valid line: %s\n", line)
	return &game, true
}

func part_one(lines []string) int {
	var total int
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		value, is_valid := parse_line(line)
		if is_valid {
			total += *value
			fmt.Printf("Total value is now '%d'\n", total)
		}
	}
	return total
}

type minima struct {
	red, green, blue int
}

func extract_values(set string) map[string]int {
	entries := strings.Split(set, ",")
	values := make(map[string]int)
	for i := 0; i < len(entries); i++ {
		entry := entries[i]
		bits := strings.Split(entry, " ")
		count, _ := strconv.Atoi(bits[1])
		colour := bits[2]
		values[colour] = count
	}
	return values
}

func get_minima_from_line(line string) (v minima) {
	_, sets := extract_games(line)

	var minimal minima = minima{0, 0, 0}
	for i := 0; i < len(sets); i++ {
		set := sets[i]
		values := extract_values(set)
		var red, green, blue int
		var ok bool
		red, ok = values["red"]
		if ok && (red > minimal.red) {
			minimal.red = red
		}
		green, ok = values["green"]
		if ok && (green > minimal.green) {
			minimal.green = green
		}
		blue, ok = values["blue"]
		if ok && (blue > minimal.blue) {
			minimal.blue = blue
		}
	}
	return minimal
}

func calculate_power(v minima) int {
	return v.red * v.green * v.blue 

}

func part_two(lines []string) int {
	var total int = 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		minimal_values := get_minima_from_line(line)
		fmt.Println(minimal_values)

		power := calculate_power(minimal_values)
		total += power
		fmt.Printf("New power to add ('%d'), total is now '%d'\n", power, total)
	}
	return total
}

func main()  {
	filename := "input"

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	var answer int
	answer = part_one(lines)
	fmt.Printf("Answer to part One is: %d\n", answer)

	answer = part_two(lines)
	fmt.Printf("Answer to part Two is: %d\n", answer)
}
