package kmap_test

import (
	"github.com/cucumber/godog"
	"testing"
)

func stepdefs(ctx *godog.ScenarioContext) {

}

func TestFeatures(t *testing.T) {
	if r := (godog.TestSuite{
		ScenarioInitializer: stepdefs,
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     []string{"features"},
			Randomize: -1,
			TestingT:  t,
		},
	}).Run(); r != 0 {
		t.Fatalf("godog exited with non-zero exit code '%d'", r)
	}
}
