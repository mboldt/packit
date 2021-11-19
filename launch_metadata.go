package packit

// LaunchMetadata represents the launch metadata details persisted in the
// launch.toml file according to the buildpack lifecycle specification:
// https://github.com/buildpacks/spec/blob/main/buildpack.md#launchtoml-toml.
type LaunchMetadata struct {
	// Processes is a list of processes that will be returned to the lifecycle to
	// be executed during the launch phase.
	Processes []Process

	// Slices is a list of slices that will be returned to the lifecycle to be
	// exported as separate layers during the export phase.
	Slices []Slice

	// Labels is a map of key-value pairs that will be returned to the lifecycle to be
	// added as config label on the image metadata. Keys must be unique.
	Labels map[string]string

	// BOM is the Bill-of-Material entries containing information about the
	// dependencies provided to the launch environment.
	SBOM []SBOMEntry
}

func (l LaunchMetadata) isEmpty() bool {
	return (len(l.Processes) == 0 &&
		len(l.Slices) == 0 &&
		len(l.Labels) == 0 &&
		len(l.SBOM) == 0)
}

func (b BuildMetadata) isEmpty() bool {
	return (len(b.SBOM) == 0 &&
		len(b.Unmet) == 0)
}
