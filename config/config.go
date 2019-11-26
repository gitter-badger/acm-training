package config

import (
	flag "github.com/spf13/pflag"
)

var version *bool

// InitCommandLineFlags initializes the common command flags used by the
// application, including db configuration, version, etc.
func InitCommandLineFlags() {
	version = flag.BoolP("version", "v", false, "Show the version information of the application.")
	flag.Parse()
}

// VersionPresent returns whether the version flag (--version) is present or not.
func VersionPresent() bool {
	return *version
}
