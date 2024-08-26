package main

import (
	"reflect"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		input    string
		expected Player
		hasError bool
	}{
		{"1 2 3 4 5", Player{[5]int{1, 2, 3, 4, 5}}, false},
		{"10 20 30 40 50", Player{[5]int{10, 20, 30, 40, 50}}, false},
		{"1 2 3 4", Player{}, true},
		{"1 2 3 4 5 6", Player{}, true},
		{"a b c d e", Player{}, true},
	}

	for _, test := range tests {
		result, err := NewPlayer(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("NewPlayer(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("NewPlayer(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestNewPick(t *testing.T) {
	tests := []struct {
		input    string
		expected Pick
		hasError bool
	}{
		{"1 2 3 4 5", Pick{[5]int{1, 2, 3, 4, 5}}, false},
		{"10 20 30 40 50", Pick{[5]int{10, 20, 30, 40, 50}}, false},
		{"1 2 3 4", Pick{}, true},
		{"1 2 3 4 5 6", Pick{}, true},
		{"a b c d e", Pick{}, true},
	}

	for _, test := range tests {
		result, err := NewPick(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("NewPick(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("NewPick(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestCountWinners(t *testing.T) {
	players := []Player{
		{[5]int{1, 2, 3, 4, 5}},
		{[5]int{6, 7, 8, 9, 10}},
		{[5]int{1, 3, 5, 7, 9}},
		{[5]int{2, 4, 6, 8, 10}},
	}

	tests := []struct {
		pick     Pick
		expected [4]int
	}{
		{Pick{[5]int{1, 2, 3, 4, 5}}, [4]int{1, 1, 0, 1}},
		{Pick{[5]int{6, 7, 8, 9, 10}}, [4]int{1, 1, 0, 1}},
		{Pick{[5]int{1, 3, 5, 7, 9}}, [4]int{1, 1, 0, 1}},
		{Pick{[5]int{2, 4, 6, 8, 10}}, [4]int{1, 1, 0, 1}},
		{Pick{[5]int{1, 2, 6, 7, 8}}, [4]int{2, 2, 0, 0}},
		{Pick{[5]int{11, 12, 13, 14, 15}}, [4]int{0, 0, 0, 0}},
	}

	for _, test := range tests {
		result := test.pick.CountWinners(players)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CountWinners(%v) = %v, want %v", test.pick, result, test.expected)
		}
	}
}
