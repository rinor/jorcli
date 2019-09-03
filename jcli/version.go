package jcli

// Version - get jcli version.
//
//  jcli --version | STDOUT
func Version() ([]byte, error) {
	return jcli(nil, "--version")
}

// VersionSource - get jcli source version.
//
//  jcli --source-version | STDOUT
func VersionSource() ([]byte, error) {
	return jcli(nil, "--source-version")
}

// VersionFull - get jcli full version.
//
//  jcli --full-version | STDOUT
func VersionFull() ([]byte, error) {
	return jcli(nil, "--full-version")
}
