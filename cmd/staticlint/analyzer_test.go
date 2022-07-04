package staticlint

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestOsExitAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), osExitAnalyzer, "./...")
}

func TestDuplicateAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), duplicateAnalyzer, "./...")
}
