// Package jnode provides jormungandr binary and block0 config helpers.
package jnode

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
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

	cmd    *exec.Cmd
	chStop chan struct{}
}

// NewJnode returns a Jnode with some defaults.
func NewJnode() *Jnode {
	return &Jnode{
		WorkingDir: os.TempDir(),
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		chStop:     make(chan struct{}),
	}
}

// Start the node and wait.
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

	// return jnode.cmd.Wait()

	go func() {
		err := jnode.cmd.Wait()
		if err != nil {
			log.Printf("jnode.Run : %v", err)
		}
		select {
		case <-jnode.chStop:
		default:
			close(jnode.chStop)
		}
	}()

	return nil
}

// Wait for the node to stop.
func (jnode *Jnode) Wait() {
	<-jnode.chStop
}

// Stop the node if running.
func (jnode *Jnode) Stop() error {
	if jnode.cmd.Process == nil {
		return fmt.Errorf("%s : exec: not started", "jnode.Stop")
	}
	return jnode.cmd.Process.Kill()
}

// StopAfter seconds.
func (jnode *Jnode) StopAfter(seconds int) {
	if jnode.cmd.Process == nil {
		return // fmt.Errorf("%s : exec: not started", "jnode.Stop")
	}

	go func() {
		select {
		case <-jnode.chStop:
		case <-time.After(time.Duration(seconds) * time.Second):
			_ = jnode.Stop()
		}
	}()
}

// Pid provided for the running node process.
func (jnode *Jnode) Pid() int {
	if jnode.cmd.Process == nil {
		return 0
	}
	return jnode.cmd.Process.Pid
}

// AddSecretFile to node config
func (jnode *Jnode) AddSecretFile(secretFile string) {
	jnode.Secrets = append(jnode.Secrets, secretFile)
}

// BinName set the executable name/full path if not the default one.
func BinName(name string) {
	jnodeName = name
}
