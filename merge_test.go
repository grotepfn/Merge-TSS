package main

import (
	"testing"
)

func TestMerge(t *testing.T) {
	input := [][2]int{
		{-3, -2},
		{0, 1},
		{4, 5},
		{2, 3},
		{4, 20}}
	expectedOutput := [][2]int{
		{-3, -2},
		{0, 1},
		{2, 3},
		{4, 20}}
	input = Merge(input)

	if (input == nil) != (expectedOutput == nil) {
		t.Fatalf(`Wrong expectedOutput nil`)
	}

	if len(input) != len(expectedOutput) {
		t.Fatalf(`Wrong expectedOutput length`)
	}

	for i := range input {
		if input[i][0] != expectedOutput[i][0] || input[i][1] != expectedOutput[i][1] {
			t.Fatalf(`Wrong order`)
		}
	}

}

func TestMergeDuplicate(t *testing.T) {
	input := [][2]int{
		{-3, -2},
		{0, 1},
		{0, 1},
		{4, 5},
		{2, 3},
		{4, 20}}

	expectedOutput := [][2]int{
		{-3, -2},
		{0, 1},
		{2, 3},
		{4, 20}}
	input = Merge(input)

	if (input == nil) != (expectedOutput == nil) {
		t.Fatalf(`Wrong expectedOutput nil`)
	}

	if len(input) != len(expectedOutput) {
		t.Fatalf(`Wrong expectedOutput length`)
	}

	for i := range input {
		if input[i][0] != expectedOutput[i][0] || input[i][1] != expectedOutput[i][1] {
			t.Fatalf(`Wrong order`)
		}
	}

}
