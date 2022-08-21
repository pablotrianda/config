package config_test

import (
	"fmt"

	"github.com/pablotrianda/config"
)

// Test if the filename is load correctly
func Example_LoadFileData() {
	cfg, _ := config.Load("testdata/app/config.yaml")

	fmt.Println(cfg.ConfigFilePath)

	// Output:
	// testdata/app/config.yaml
}

func Example_GetErrorToLoadFile() {
	_, error := config.Load("to/non-exist-fie.yaml")

	if error != nil {
		fmt.Printf("%v", error)
	}

	// Output:
	// Cannot read the file
}

// Test get all data
func Example_GetAllData() {
	cfg, _ := config.Load("testdata/app/config.yaml")

	fmt.Println(cfg.Info)

	// Output:
	// map[command:map[path:/here/we/go] here:goes some:thing]
}

func Example_GetEspecificData() {
	cfg, _ := config.Load("testdata/app/config.yaml")
	value, _ := cfg.Get("some")
	fmt.Print(value)

	// Output:
	// thing
}

func Example_GetErrorToGetNonExistData() {
	cfg, _ := config.Load("testdata/app/config.yaml")
	_, error := cfg.Get("non-exist")
	fmt.Printf("%v", error)

	// Output:
	// Key non exist
}
