package p0067

import (
	"strings"
	"testing"
)

func TestByteConversion(t *testing.T) {
	container := &strings.Builder{}
	a := rune(1)
	b := byte(1)
	container.WriteRune(a)
	container.WriteByte(b)
}

func TestSolution(t *testing.T) {
	addBinary("1", "111")
}
