package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	fmt.Fprintf(os.Stderr, "starting tests!\n")
	Result := m.Run()
	fmt.Fprintf(os.Stderr, "finished tests!\n")
	fmt.Fprintf(os.Stderr, "Result: %v ", Result)
	if Result == 0 {
		os.Exit(0)
	}
	os.Exit(1)
}
