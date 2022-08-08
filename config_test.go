package config_test

import (
	"fmt"

	"github.com/pablotrianda/config"
)

// Test if the filename is load correctly
func Example_LoadFileData() {
	cfg := config.Load("testdata/app/config.yaml")

	fmt.Println(cfg)

	// Output:
	// testdata/app/config.yaml
}
