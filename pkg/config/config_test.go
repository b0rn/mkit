package config

import (
	"os"
	"testing"
)

func TestLoadEnvVars(t *testing.T) {
	if err := LoadEnvVars("env", "./testdata/test.env"); err != nil {
		t.Errorf("expected no error when running LoadEnvVars but got %v", err)
	}
	if e := os.Getenv("A"); e != "5" {
		t.Errorf("expected env variable A to be 5 but got %s", e)
	}
	if e := os.Getenv("B"); e != "foo" {
		t.Errorf("expected env variable A to be 5 but got %s", e)
	}
	if err := LoadEnvVars("env", "./falsefile"); err == nil {
		t.Errorf("must return an error when env file is incorrect")
	}
}

type cfg struct {
	A struct {
		B int    `yaml:"b"`
		C string `yaml:"c"`
		D string `yaml:"d"`
	} `yaml:"a"`
}

func TestBuildConfig(t *testing.T) {
	var c cfg
	if err := BuildConfig("yaml", "./testdata/config.yaml", &c); err != nil {
		t.Fatalf("expected no error but got %v", err)
	}
	if (c == cfg{}) {
		t.Fatal("config is default value")
	}
	if c.A.B != 5 {
		t.Errorf("expected A.B to be 5 but got %v", c.A.B)
	}
	if c.A.C != "foo" {
		t.Errorf("expected A.B to be foo but got %v", c.A.C)
	}
	if c.A.D != "bar" {
		t.Errorf("expected A.B to be bar but got %v", c.A.D)
	}
	if err := BuildConfig("yaml", "./incorrectFile", c); err == nil {
		t.Fatal("must return an error when file is incorrect")
	}
}
