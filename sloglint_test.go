package sloglint_test

import (
	"testing"

	"go.tmz.dev/sloglint"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	t.Run("mixed arguments", func(t *testing.T) {
		analyzer := sloglint.New(nil)
		analysistest.Run(t, testdata, analyzer, "mixed_args")
	})

	t.Run("key-value pairs only", func(t *testing.T) {
		analyzer := sloglint.New(&sloglint.Options{KVOnly: true})
		analysistest.Run(t, testdata, analyzer, "kv_only")
	})

	t.Run("attributes only", func(t *testing.T) {
		analyzer := sloglint.New(&sloglint.Options{AttrOnly: true})
		analysistest.Run(t, testdata, analyzer, "attr_only")
	})
}
