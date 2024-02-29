package tests

import (
	"github.com/inocencio/grinvon/conv"
	"testing"
)

func TestSplitFloat(t *testing.T) {
	intPart, fracPart := conv.SplitFloat(1234.56, 2)

	if intPart != 1234 {
		t.Errorf("Unexpected intPart: got %v want %v", intPart, 1234)
	}
	if fracPart != 56 {
		t.Errorf("Unexpected fracPart: got %v want %v", fracPart, 56)
	}

	intPart, fracPart = conv.SplitFloat(123456, 10)
	if intPart != 123456 {
		t.Errorf("Unexpected intPart: got %v want %v", intPart, 123456)
	}
	if fracPart != 0 {
		t.Errorf("Unexpected fracPart: got %v want %v", fracPart, 0)
	}
}
