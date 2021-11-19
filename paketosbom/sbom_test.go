package paketosbom_test

import (
	"testing"

	"github.com/paketo-buildpacks/packit/paketosbom"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testPaketoSBOM(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	context("GetSBOMChecksumAlgorithm", func() {
		context("given an algorithm string with the exact name of a CycloneDX algorithm", func() {
			it("returns the same algorithm name", func() {
				algorithm512, err := paketosbom.GetSBOMChecksumAlgorithm("SHA-512")
				Expect(err).ToNot(HaveOccurred())
				Expect(algorithm512).To(Equal(paketosbom.SHA512))
			})
		})
		context("given an algorithm string with a lowercase version of a CycloneDX algorithm", func() {
			it("returns the Cyclonedx algorithm name", func() {
				algorithm512, err := paketosbom.GetSBOMChecksumAlgorithm("sha-512")
				Expect(err).ToNot(HaveOccurred())
				Expect(algorithm512).To(Equal(paketosbom.SHA512))
			})
			context("it also does not contain a dash", func() {
				it("returns the Cyclonedx algorithm name", func() {
					algorithm512, err := paketosbom.GetSBOMChecksumAlgorithm("sha512")
					Expect(err).ToNot(HaveOccurred())
					Expect(algorithm512).To(Equal(paketosbom.SHA512))
				})
			})
		})
		context("failure cases", func() {
			context("when the attempted SBOM checksum algorithm is not supported", func() {
				it("persists a build.toml", func() {
					_, err := paketosbom.GetSBOMChecksumAlgorithm("RANDOM-ALG")
					Expect(err).To(MatchError("failed to get supported SBOM checksum algorithm: RANDOM-ALG is not valid"))
				})
			})
		})
	})
}
