package p0093

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	restoreIpAddresses("25525511135")
	val, err := strconv.Atoi("525")
	if err == nil {
		fmt.Println(val)
	}
}
