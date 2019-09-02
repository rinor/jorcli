package jcli

// JormungandrVersion - get jormungandr version.
// Provided just for convenience.
//
//  jormungandr --version | STDOUT
func JormungandrVersion() ([]byte, error) {
	return execStd(nil, "jormungandr", "--version")
}

// Version - get jcli version.
//
//  jcli --version | STDOUT
func Version() ([]byte, error) {
	return execStd(nil, "jcli", "--version")
}

// VersionSource - get jcli source version.
//
//  jcli --source-version | STDOUT
func VersionSource() ([]byte, error) {
	return execStd(nil, "jcli", "--source-version")
}

// VersionFull - get jcli full version.
//
//  jcli --full-version | STDOUT
func VersionFull() ([]byte, error) {
	return execStd(nil, "jcli", "--full-version")
}
