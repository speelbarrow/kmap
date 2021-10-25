package kmap_test

import (
	"github.com/cucumber/godog"
	"testing"
)

func iInitializeTheKmap() error {
	return godog.ErrPending
}

func iRandomlyGenerateTheArgumentsToTheKmap() error {
	return godog.ErrPending
}

func theArgumentsToTheKmapAre(arg1 *godog.Table) error {
	return godog.ErrPending
}

func theKmapSizeIs(arg1 int) error {
	return godog.ErrPending
}

func theKmapValuesShouldMatch(arg1 *godog.Table) error {
	return godog.ErrPending
}

func theKmapValuesShouldMatchTheArguments() error {
	return godog.ErrPending
}

func theMintermsMethodShouldMatch(arg1 *godog.Table) error {
	return godog.ErrPending
}

func thePropertyOfTheKmapShouldBe(arg1 string, arg2 int) error {
	return godog.ErrPending
}

func stepdefs(ctx *godog.ScenarioContext) {
	ctx.Step(`^I initialize the k-map$`, iInitializeTheKmap)
	ctx.Step(`^I randomly generate the arguments to the k-map$`, iRandomlyGenerateTheArgumentsToTheKmap)
	ctx.Step(`^the arguments to the k-map are$`, theArgumentsToTheKmapAre)
	ctx.Step(`^the k-map size is (\d+)$`, theKmapSizeIs)
	ctx.Step(`^the k-map values should match$`, theKmapValuesShouldMatch)
	ctx.Step(`^the k-map values should match the arguments$`, theKmapValuesShouldMatchTheArguments)
	ctx.Step(`^the Minterms method should match$`, theMintermsMethodShouldMatch)
	ctx.Step(`^the "([^"]*)" property of the k-map should be (\d+)$`, thePropertyOfTheKmapShouldBe)

}

func TestFeatures(t *testing.T) {
	if r := (godog.TestSuite{
		ScenarioInitializer: stepdefs,
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     []string{"features"},
			Randomize: -1,
			TestingT:  t,
			Tags:      "~@wip",
		},
	}).Run(); r != 0 {
		t.Fatalf("godog exited with non-zero exit code '%d'", r)
	}
}
