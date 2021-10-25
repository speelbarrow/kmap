package kmap_test

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/noah-friedman/kmap"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

var (
	k        *kmap.Kmap
	size     int
	kmapArgs []int
)

func iInitializeTheKmap() error {
	var e error
	if k, e = kmap.NewKmap(size, kmapArgs...); e != nil {
		return e
	}
	return nil
}

func iRandomlyGenerateTheArgumentsToTheKmap() {
	size := int(math.Pow(2, float64(size)))
	for i := 0; i < size; i++ {
		if rand.Int31n(2) == 1 {
			kmapArgs = append(kmapArgs, i)
		}
	}
}

func theArgumentsToTheKmapAre(args *godog.Table) error {
	for _, v := range args.Rows[0].Cells {
		if i, e := strconv.Atoi(v.Value); e != nil {
			return e
		} else {
			kmapArgs = append(kmapArgs, i)
		}
	}

	return nil
}

func theKmapSizeIs(s int) {
	size = s
}

func theKmapValuesShouldMatch(expected *godog.Table) error {
	if l := len(expected.Rows); l != k.Rows {
		return fmt.Errorf("expected %d rows, found %d", l, k.Rows)
	}
	if l := len(expected.Rows[0].Cells); l != k.Cols {
		return fmt.Errorf("expected %d cols, found %d", l, k.Cols)
	}

	for row, actual := range k.Values {
		expected := expected.Rows[row]
		for col, actual := range actual {
			expected := expected.Cells[col].Value == "1"
			if actual != expected {
				return fmt.Errorf("row %d, col %d: expected %t, got %t", row, col, expected, actual)
			}
		}
	}

	return nil
}

func theKmapValuesShouldMatchTheArguments() error {
	var (
		actual   = k.Minterms()
		expected = make([]bool, len(actual))
	)
	for _, v := range kmapArgs {
		expected[v] = true
	}

	if !reflect.DeepEqual(expected, actual) {
		var exp [][]bool
		for i, s := k.Cols, k.Rows*k.Cols; i <= s; i += k.Cols {
			exp = append(exp, expected[i-k.Cols:i])
		}
		return fmt.Errorf("expected a k-map with args %v, found k-map with args %v", exp, k.Values)
	}

	return nil
}

func theMintermsMethodShouldOutput(expected *godog.Table) error {
	actual := k.Minterms()
	for i, exp := range expected.Rows[0].Cells {
		if exp := exp.Value == "1"; exp != actual[i] {
			return fmt.Errorf("expected minterm %d to be %t but found %t\nexpected:\n%v\nactual:\n%v\n", i, exp, actual[i], expected, actual)
		}
	}
	return nil
}

func thePropertyOfTheKmapShouldBe(prop string, expected int64) error {
	if actual := reflect.ValueOf(*k).FieldByName(prop).Int(); expected != actual {
		return fmt.Errorf("expected %d, got %d", expected, actual)
	}

	return nil
}

func stepdefs(ctx *godog.ScenarioContext) {
	ctx.Step(`^I initialize the k-map$`, iInitializeTheKmap)
	ctx.Step(`^I randomly generate the arguments to the k-map$`, iRandomlyGenerateTheArgumentsToTheKmap)
	ctx.Step(`^the arguments to the k-map are$`, theArgumentsToTheKmapAre)
	ctx.Step(`^the k-map size is (\d+)$`, theKmapSizeIs)
	ctx.Step(`^the k-map values should match$`, theKmapValuesShouldMatch)
	ctx.Step(`^the k-map values should match the arguments$`, theKmapValuesShouldMatchTheArguments)
	ctx.Step(`^the Minterms method should output$`, theMintermsMethodShouldOutput)
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
