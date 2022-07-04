package staticlint

import (
	"strings"

	"github.com/gostaticanalysis/nakedreturn"
	"github.com/gostaticanalysis/nilerr"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

func main() {

	var staticlints []*analysis.Analyzer

	staticlints = append(staticlints, osExitAnalyzer)
	staticlints = append(staticlints, duplicateAnalyzer)
	staticlints = append(staticlints, printf.Analyzer)
	staticlints = append(staticlints, shadow.Analyzer)
	staticlints = append(staticlints, structtag.Analyzer)
	staticlints = append(staticlints, structtag.Analyzer)

	staticlints = append(staticlints, nakedreturn.Analyzer)
	staticlints = append(staticlints, nilerr.Analyzer)

	for _, v := range staticcheck.Analyzers {
		if strings.Contains(v.Analyzer.Name, "SA") || strings.Contains(v.Analyzer.Name, "ST") {
			staticlints = append(staticlints, v.Analyzer)
		}
	}

	multichecker.Main(
		staticlints...,
	)
}
