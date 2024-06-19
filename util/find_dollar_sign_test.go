package util

import (
	"testing"
)

func TestFindDollarSignIndexInUni16Text(t *testing.T) {
	text := "Hello, $ hello こんにちは $, สวัสดีครับ $"
	indexes := FindDollarSignIndexInUTF16Text(text)
	if len(indexes) != 3 {
		t.Errorf("Expected 3, but got %d", len(indexes))
	}
}
