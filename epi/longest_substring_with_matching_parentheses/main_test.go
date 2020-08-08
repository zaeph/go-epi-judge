package longest_substring_with_matching_parentheses_test

import (
	"os"
	"testing"

	"github.com/stefantds/go-epi-judge/config"
)

var cfg *config.Config

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.Parse()
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}
