package kmap_test

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/noah-friedman/kmap"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

func iInitializeTheKmap(ctx context.Context) context.Context {
	k, e := kmap.NewKmap(ctx.Value("size").(int), ctx.Value("args").([]int)...)

	return context.WithValue(context.WithValue(ctx, "kmap", k), "err", e)
}

func iRandomlyGenerateTheArgumentsToTheKmap(ctx context.Context) context.Context {
	size := int(math.Pow(2, float64(ctx.Value("size").(int))))

	var args []int
	for i := 0; i < size; i++ {
		if rand.Int31n(2) == 1 {
			args = append(args, i)
		}
	}

	return context.WithValue(ctx, "args", args)
}

func theArgumentsToTheKmapAre(ctx context.Context, args *godog.Table) (context.Context, error) {
	var r []int
	for _, v := range args.Rows[0].Cells {
		if i, e := strconv.Atoi(v.Value); e != nil {
			return ctx, e
		} else {
			r = append(r, i)
		}
	}

	return context.WithValue(ctx, "args", r), nil
}

func theKmapSizeIs(ctx context.Context, size int) context.Context {
	return context.WithValue(ctx, "size", size)
}

func theKmapValuesShouldMatch(ctx context.Context, expected *godog.Table) error {
	k := ctx.Value("kmap").(*kmap.Kmap)

	if l := len(expected.Rows); l != k.Rows {
		return fmt.Errorf("expected %d rows, found %d", l, k.Rows)
	}
	if l := len(expected.Rows[0].Cells); l != k.Cols {
		return fmt.Errorf("expected %d cols, found %d", l, k.Cols)
	}

	for row, actual := range k.Values {
		exp := expected.Rows[row]
		for col, actual := range actual {
			exp := exp.Cells[col].Value == "1"
			if actual != exp {
				var expArr [][]bool
				for _, v := range expected.Rows {
					var r []bool
					for _, v := range v.Cells {
						r = append(r, v.Value == "1")
					}
					expArr = append(expArr, r)
				}

				return fmt.Errorf("row %d, col %d: expected %t, got %t\nexpected:\n%v\nactual:\n%v\n", row, col, exp, actual, expArr, k.Values)
			}
		}
	}

	return nil
}

func theKmapValuesShouldMatchTheArguments(ctx context.Context) error {
	var (
		k        = ctx.Value("kmap").(*kmap.Kmap)
		actual   = k.Minterms()
		expected = make([]bool, len(actual))
	)
	for _, v := range ctx.Value("args").([]int) {
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

func theMintermsMethodShouldOutput(ctx context.Context, expected *godog.Table) error {
	actual := ctx.Value("kmap").(*kmap.Kmap).Minterms()
	for i, exp := range expected.Rows[0].Cells {
		if exp := exp.Value == "1"; exp != actual[i] {
			return fmt.Errorf("expected minterm %d to be %t but found %t\nexpected:\n%v\nactual:\n%v\n", i, exp, actual[i], expected, actual)
		}
	}
	return nil
}

func thePropertyOfTheKmapShouldBe(ctx context.Context, prop string, expected int64) error {
	if actual := reflect.ValueOf(*(ctx.Value("kmap").(*kmap.Kmap))).FieldByName(prop).Int(); expected != actual {
		return fmt.Errorf("expected %d, got %d", expected, actual)
	}

	return nil
}

func anErrorShouldHaveOccurred(ctx context.Context) (context.Context, error) {
	if ctx.Value("err") == nil {
		return ctx, fmt.Errorf("expected an error occurred but found no error")
	}

	return context.WithValue(ctx, "err", nil), nil
}

func iParseTheString(s string) error {
	return godog.ErrPending
}

func theParsingResultShouldBe(expected *godog.Table) error {
	return godog.ErrPending
}

func Stepdefs(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		for k, v := range map[string]interface{}{
			"kmap": (*kmap.Kmap)(nil),
			"size": 0,
			"args": []int(nil),
			"err":  error(nil),
		} {
			ctx = context.WithValue(ctx, k, v)
		}

		return ctx, nil
	})

	ctx.Step(`^I initialize the k-map$`, iInitializeTheKmap)
	ctx.Step(`^I randomly generate the arguments to the k-map$`, iRandomlyGenerateTheArgumentsToTheKmap)
	ctx.Step(`^the arguments to the k-map are$`, theArgumentsToTheKmapAre)
	ctx.Step(`^the k-map size is (\d+)$`, theKmapSizeIs)
	ctx.Step(`^the k-map values should match$`, theKmapValuesShouldMatch)
	ctx.Step(`^the k-map values should match the arguments$`, theKmapValuesShouldMatchTheArguments)
	ctx.Step(`^the Minterms method should output$`, theMintermsMethodShouldOutput)
	ctx.Step(`^the "([^"]*)" property of the k-map should be (\d+)$`, thePropertyOfTheKmapShouldBe)
	ctx.Step(`^an error should have occurred$`, anErrorShouldHaveOccurred)
	ctx.Step(`^I parse the string "([^"]*)"$`, iParseTheString)
	ctx.Step(`^the parsing result should be$`, theParsingResultShouldBe)

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		switch e := ctx.Value("err"); e.(type) {
		case error:
			return ctx, fmt.Errorf("the following error occured at some point: %s", e.(error).Error())
		}

		return ctx, nil
	})
}

func TestFeatures(t *testing.T) {
	if r := (godog.TestSuite{
		ScenarioInitializer: Stepdefs,
		Options: &godog.Options{
			Concurrency: runtime.NumCPU(),
			Format:      "pretty",
			Paths:       []string{"features"},
			Randomize:   -1,
			TestingT:    t,
			Tags:        "~@wip",
		},
	}).Run(); r != 0 {
		t.Fatalf("godog exited with non-zero exit code '%d'", r)
	}
}
