package main

import "testing"

func TestPart1(t *testing.T) {
	scanner := makeScanner("sample.txt")
	result := calculatePart1(scanner)
	if result != SAMPLE_TARGET_P1 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P1)
	}
}

func TestPart2(t *testing.T) {
	scanner := makeScanner("sample.txt")
	result := calculatePart2(scanner)
	if result != SAMPLE_TARGET_P2 {
		t.Errorf("Part2 = %v; want %v", result, SAMPLE_TARGET_P2)
	}
}