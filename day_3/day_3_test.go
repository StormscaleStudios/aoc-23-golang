package main

import (
    "os"
    "log"
    "strings"
    "testing"
)

// note: this test gave false positives, because I was initially using 0 as empty value, and -1 for tokens
//  these masking values led to the wrong identification of numbers like 903 and 760 (becoming 9, 3 and 76)
//  test passed because there are no 0's in the test data
func TestDay3PartOne(t *testing.T) {
	filename := "test_input"

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
    result := part_one(lines)
    answer := 4361
    
    if (result != answer) {
        t.Errorf("Expected %d, got %d", answer, result)
    }
}

func TestDay3PartTwo(t *testing.T) {
	filename := "test_input"

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
    result := part_two(lines)
    answer := 467835
    
    if (result != answer) {
        t.Errorf("Expected %d, got %d", answer, result)
    }
}
