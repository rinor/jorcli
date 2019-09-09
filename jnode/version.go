package jnode

import (
	"bytes"
	"os/exec"
)

// jnodeStd ...
func jnodeStd(arg ...string) ([]byte, error) {
	var (
		cmd    *exec.Cmd
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	cmd = exec.Command(jnodeName, arg...)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return stderr.Bytes(), err
	}
	return stdout.Bytes(), nil
}

// Version - get jormungandr version.
//
//  jormungandr --version | STDOUT
func Version() ([]byte, error) {
	return jnodeStd("--version")
}

// VersionSource - get jormungandr source version.
//
//  jormungandr --source-version | STDOUT
func VersionSource() ([]byte, error) {
	return jnodeStd("--source-version")
}

// VersionFull - get jormungandr full version.
//
//  jormungandr --full-version | STDOUT
func VersionFull() ([]byte, error) {
	return jnodeStd("--full-version")
}
