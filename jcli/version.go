package jcli

// JormungandrVersion - get jormungandr version.
// Provided just for convenience.
func JormungandrVersion() ([]byte, error) {
	return execStd(nil, "jormungandr", "--version")
}

// Version - get jcli version.
func Version() ([]byte, error) {
	return execStd(nil, "jcli", "--version")
}

// VersionSource - get jcli source version.
func VersionSource() ([]byte, error) {
	return execStd(nil, "jcli", "--source-version")
}

// VersionFull - get jcli full version.
func VersionFull() ([]byte, error) {
	return execStd(nil, "jcli", "--full-version")
}
