package main

import "testing"

func TestPart1Tx01(t *testing.T) {
	scanner := makeScanner("sample_01.txt")
	result := calculate(scanner, 4)
	if result != SAMPLE_TARGET_P1_01 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P1_01)
	}
}
func TestPart1Tx02(t *testing.T) {
	scanner := makeScanner("sample_02.txt")
	result := calculate(scanner, 4)
	if result != SAMPLE_TARGET_P1_02 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P1_02)
	}
}
func TestPart1Tx03(t *testing.T) {
	scanner := makeScanner("sample_03.txt")
	result := calculate(scanner, 4)
	if result != SAMPLE_TARGET_P1_03 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P1_03)
	}
}
func TestPart1Tx04(t *testing.T) {
	scanner := makeScanner("sample_04.txt")
	result := calculate(scanner, 4)
	if result != SAMPLE_TARGET_P1_04 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P1_04)
	}
}

func TestPart2Tx01(t *testing.T) {
	scanner := makeScanner("sample_01.txt")
	result := calculate(scanner, 14)
	if result != SAMPLE_TARGET_P2_01 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P2_01)
	}
}
func TestPart2Tx02(t *testing.T) {
	scanner := makeScanner("sample_02.txt")
	result := calculate(scanner, 14)
	if result != SAMPLE_TARGET_P2_02 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P2_02)
	}
}
func TestPart2Tx03(t *testing.T) {
	scanner := makeScanner("sample_03.txt")
	result := calculate(scanner, 14)
	if result != SAMPLE_TARGET_P2_03 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P2_03)
	}
}
func TestPart2Tx04(t *testing.T) {
	scanner := makeScanner("sample_04.txt")
	result := calculate(scanner, 14)
	if result != SAMPLE_TARGET_P2_04 {
		t.Errorf("Part1 = %v; want %v", result, SAMPLE_TARGET_P2_04)
	}
}