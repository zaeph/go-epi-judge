package replace_and_remove_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/replace_and_remove"
)

func TestReplaceAndRemove(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "replace_and_remove.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Size           int
		S              []string
		ExpectedResult []string
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Size,
			&tc.S,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			result, err := replaceAndRemoveWrapper(tc.Size, tc.S)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func replaceAndRemoveWrapper(size int, s []string) ([]string, error) {
	allChars := make([]rune, len(s))
	for i, c := range []rune(strings.Join(s, "")) {
		allChars[i] = c
	}
	resSize := ReplaceAndRemove(size, allChars)

	if resSize >= size {
		return nil, errors.New("result can't be greater than the original size")
	}
	result := make([]string, resSize)
	for i := 0; i < resSize; i++ {
		result[i] = string(allChars[i])
	}

	return result, nil
}
