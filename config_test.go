package config_test

import (
	"testing"

	"github.com/pablotrianda/config"
	"github.com/stretchr/testify/assert"
)

// Test if the filename is load correctly
func TestLoadFileData(t *testing.T) {
	cfg, _ := config.Load("testdata/app/config.yaml")

	assert.Equal(t, "testdata/app/config.yaml", cfg.ConfigFilePath)
}

func TestGetErrorToLoadFile(t *testing.T) {
	_, error := config.Load("to/non-exist-fie.yaml")

	assert.NotNil(t, error)
}

// Test get all data
func TestGetAllData(t *testing.T) {
	cfg, error := config.Load("testdata/app/config.yaml")

	assert.Nil(t, error)
	assert.NotEmpty(t, cfg)
}

func TestGetEspecificData(t *testing.T) {
	cfg, _ := config.Load("testdata/app/config.yaml")
	value, err := cfg.Get("some")

	assert.Nil(t, err)
	assert.Equal(t, "thing", value)
}

func TestErrorToGetNonExistData(t *testing.T) {
	cfg, _ := config.Load("testdata/app/config.yaml")
	_, error := cfg.Get("non-exist")

	assert.ErrorContains(t, error, "Key non exist")
}
