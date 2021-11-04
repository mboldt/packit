package paketobom_test

import (
	"errors"
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitPaketoBom(t *testing.T) {
	suite := spec.New("paketobom", spec.Report(report.Terminal{}))
	suite("bom", testPaketoBom)
	suite.Run(t)
}

type errorReader struct{}

func (r errorReader) Read(p []byte) (int, error) {
	return 0, errors.New("failed to read")
}
