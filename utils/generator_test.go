package utils

import (
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	generated := GenerateShortURL()
	if len(generated) != 6 {
		t.Errorf("expected URL length 6, got %d", len(generated))
	}

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range generated {
		if !containsChar(charset, char) {
			t.Errorf("unexpected character '%c' in generated URL", char)
		}
	}
}

func containsChar(charset string, char rune) bool {
	for _, c := range charset {
		if c == char {
			return true
		}
	}
	return false
}
