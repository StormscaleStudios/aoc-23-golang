package main

import (
	"log"
	"os"
	"strings"
)

type Neighbour struct {
	X int
	Y int
}

func part_one(lines []string) int {
	lines = lines[:len(lines)-1]
	columnCount := len(lines[0])
	rowCount := len(lines)
	log.Printf("number of columns: %d", columnCount)
	log.Printf("number of rows: %d", rowCount)

	// initialise grid
	var grid [][]int = make([][]int, rowCount)
	for i := 0; i < rowCount; i++ {
		grid[i] = make([]int, columnCount)
	}

	// populate grid with markers and parse input values
	var in_island bool
	var island_count int
	var island_values = make(map[int]int)
	for y := 0; y < rowCount; y++ {
		log.Printf("%v", lines[y])
		for x := 0; x < columnCount; x++ {
			if x == 0 {
				in_island = false
				island_count++
			}
			value := int(lines[y][x])
			parsed_value := parse_position(value)
			if parsed_value >= 0 {
				in_island = true
				grid[y][x] = island_count
			} else {
				grid[y][x] = parsed_value
			}
			if parsed_value < 0 && in_island {
				in_island = false
				island_count++
			}
			if in_island {
				if island_values[island_count] == 0 {
					island_values[island_count] = parsed_value
				} else {
					previous_value := island_values[island_count]
					island_values[island_count] = 10*previous_value + parsed_value
				}
			}
		}
	}

	// find neighbouring islands
	var set = make(map[int]bool)
	for y := 0; y < rowCount; y++ {
		for x := 0; x < columnCount; x++ {
			if grid[y][x] == -2 {
				for dy := -1; dy < 2; dy++ {
					if y+dy < 0 || y+dy >= rowCount {
						continue
					}
					for dx := -1; dx < 2; dx++ {
						if x+dx < 0 || x+dx >= columnCount {
							continue
						}
						value := grid[y+dy][x+dx]
						if value > 0 {
							set[value] = true
						}
					}
				}
			}
		}
	}

	result := 0
	for k, _ := range set {
		result += island_values[k]
	}

	log.Printf("%v", grid)
	log.Printf("%v", island_values)
	log.Printf("%v", set)
	log.Printf("%v", result)

	return result
}

func part_two(lines []string) int {
	lines = lines[:len(lines)-1]
	columnCount := len(lines[0])
	rowCount := len(lines)
	log.Printf("number of columns: %d", columnCount)
	log.Printf("number of rows: %d", rowCount)

	// initialise grid
	var grid [][]int = make([][]int, rowCount)
	for i := 0; i < rowCount; i++ {
		grid[i] = make([]int, columnCount)
	}

	// populate grid with markers and parse input values
	var in_island bool
	var island_count int
	var island_values = make(map[int]int)
	for y := 0; y < rowCount; y++ {
		log.Printf("%v", lines[y])
		for x := 0; x < columnCount; x++ {
			if x == 0 {
				in_island = false
				island_count++
			}
			value := int(lines[y][x])
			parsed_value := parse_position(value)
			if parsed_value >= 0 {
				in_island = true
				grid[y][x] = island_count
			} else {
				grid[y][x] = parsed_value
			}
			if parsed_value < 0 && in_island {
				in_island = false
				island_count++
			}
			if in_island {
				if island_values[island_count] == 0 {
					island_values[island_count] = parsed_value
				} else {
					previous_value := island_values[island_count]
					island_values[island_count] = 10*previous_value + parsed_value
				}
			}
		}
	}

	// find neighbouring islands
	result := 0
	var set = make(map[int]bool)
	for y := 0; y < rowCount; y++ {
		for x := 0; x < columnCount; x++ {
			if grid[y][x] == -2 {
				for k := range set {
					delete(set, k)
				}

				for dy := -1; dy < 2; dy++ {
					if y+dy < 0 || y+dy >= rowCount {
						continue
					}
					for dx := -1; dx < 2; dx++ {
						if x+dx < 0 || x+dx >= columnCount {
							continue
						}
						value := grid[y+dy][x+dx]
						if value > 0 {
							set[value] = true
						}
					}
				}

				value := 1
				if len(set) != 2 {
					continue
				}
				for k := range set {
					value *= island_values[k]
				}
				result += value
			}
		}
	}

	log.Printf("%v", grid)
	log.Printf("%v", island_values)
	log.Printf("%v", set)
	log.Printf("%v", result)

	return result
}

// label empty spaces as -1, characters as -2, and islands as their index, starting at 0
func parse_position(value int) int {
	switch {
	case value == 46:
		return -1
	case value >= 48 && value < 58:
		return value - 48
	default:
		return -2
	}
}

func parse_position_specific(value int) int {
	switch {
	case value == 42:
		return -2
	case value >= 48 && value < 58:
		return value - 48
	default:
		return -1
	}
}

func main() {
	filename := "input"

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	answer_one := part_one(lines)
	answer_two := part_two(lines)

	log.Printf("Answer from part one: %d", answer_one)
	log.Printf("Answer from part two: %d", answer_two)
}
