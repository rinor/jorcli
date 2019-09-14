// Package jcli provides jormungandr jcli binary helpers.
package jcli

import (
	"bytes"
	"os/exec"
)

var (
	jcliName = "jcli"
)

// jcli executes "stdin | 'jcliName' args | stdout"
func jcli(stdin []byte, arg ...string) ([]byte, error) {
	var (
		cmd    *exec.Cmd
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	cmd = exec.Command(jcliName, arg...)
	// TODO: check also exec.CommandContext(ctx context.Context, name string, arg ...string)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if stdin != nil /* && len(stdin) > 0 */ {
		cmd.Stdin = bytes.NewBuffer(stdin)
	}

	if err := cmd.Run(); err != nil {
		return stderr.Bytes(), err
	}
	return stdout.Bytes(), nil
}

// BinName set the executable name/path if not the default one.
func BinName(name string) {
	jcliName = name
}
