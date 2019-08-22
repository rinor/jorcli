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
