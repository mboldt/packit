package paketobom_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitPaketoBom(t *testing.T) {
	suite := spec.New("paketobom", spec.Report(report.Terminal{}))
	suite("bom", testPaketoBom)
	suite.Run(t)
}
