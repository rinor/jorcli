// Package jnode provides jormungandr binary and block0 config helpers.
package jnode

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

var (
	jnodeName = "jormungandr"
)

// Jnode contains the node commandline config parameters.
// Contains also some extra configs for flexibily.
type Jnode struct {
	GenesisBlock     string   // --genesis-block (bin file)
	GenesisBlockHash string   // --genesis-block-hash
	Config           string   // --config <file>
	Storage          string   // --storage <storage>
	Secrets          []string // --secret <secret>
	EnableExplorer   bool     // --enable-explorer
	// Extra
	WorkingDir string
	Stdout     io.Writer
	Stderr     io.Writer

	cmd *exec.Cmd
}

// NewJnode returns a Jnode with some defaults.
func NewJnode() *Jnode {
	return &Jnode{
		WorkingDir: os.TempDir(),
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
	}
}

// Run the node and wait.
func (jnode *Jnode) Run() error {
	var arg []string

	if jnode.Config == "" {
		return fmt.Errorf("parameter missing : %s", "jnode.Config")
	}
	arg = append(arg, "--config", jnode.Config)

	if jnode.GenesisBlock != "" {
		arg = append(arg, "--genesis-block", jnode.GenesisBlock)
	}

	if jnode.GenesisBlockHash != "" {
		arg = append(arg, "--genesis-block-hash", jnode.GenesisBlockHash)
	}

	if len(jnode.Secrets) > 0 {
		for i := range jnode.Secrets {
			arg = append(arg, "--secret", jnode.Secrets[i])
		}
	}

	if jnode.Storage != "" {
		arg = append(arg, "--storage", jnode.Storage)
	}

	if jnode.EnableExplorer {
		arg = append(arg, "--enable-explorer")
	}

	jnode.cmd = exec.Command(jnodeName, arg...)

	jnode.cmd.Dir = jnode.WorkingDir
	jnode.cmd.Stdout = jnode.Stdout
	jnode.cmd.Stderr = jnode.Stderr

	if err := jnode.cmd.Start(); err != nil {
		return err
	}

	return jnode.cmd.Wait()
}

// Stop the node if running.
func (jnode *Jnode) Stop() error {
	if jnode.cmd.Process == nil {
		return nil // no need for error. Keep it simple
	}
	return jnode.cmd.Process.Kill()
}

// Pid provided for the running node process.
func (jnode *Jnode) Pid() int {
	if jnode.cmd.Process == nil {
		return 0
	}
	return jnode.cmd.Process.Pid
}

// BinName set the executable name/full path if not the default one.
func BinName(name string) {
	jnodeName = name
}
