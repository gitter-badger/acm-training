package config

import (
	flag "github.com/spf13/pflag"
)

var version *bool
var dbFile *string

// InitCommandLineFlags initializes the common command flags used by the
// application, including db configuration, version, etc.
func InitCommandLineFlags() {
	version = flag.BoolP("version", "v", false, "Show the version information of the application.")
	dbFile = flag.String("dbFile", "acm-training.db", "The sqlite3 database file location to save data.")
	flag.Parse()
}

// VersionPresent returns whether the version flag (--version) is present or not.
func VersionPresent() bool {
	return *version
}

// GetDbFile returns the location of db file to be saved.
func GetDbFile() string {
	return *dbFile
}
