package config_test

import (
	"fmt"

	"github.com/pablotrianda/config"
)

// Test if the filename is load correctly
func Example_LoadFileData() {
	cfg := config.Load("testdata/app/config.yaml")

	fmt.Println(cfg.ConfigFilePath)

	// Output:
	// testdata/app/config.yaml
}

// Test get all data
func Example_GetAllData() {
	cfg := config.Load("testdata/app/config.yaml")

	fmt.Println(cfg.Info)

	// Output:
	// map[command:map[path:/here/we/go] here:goes some:thing]
}

func Example_GetEspecificData() {
	cfg := config.Load("testdata/app/config.yaml")

	fmt.Print(cfg.Get("some"))

	// Output:
	// thing
}
