// Package jnode provides jormungandr binary and block0 config helpers.
package jnode

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
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
	ConfigFile       string   // --config <file>
	Storage          string   // --storage <storage>
	SecretFiles      []string // --secret <secret>...
	TrustedPeers     []string // --trusted-peer <trusted_peer>...
	RestListen       string   // --rest-listen
	EnableExplorer   bool     // --enable-explorer
	// Log
	LogFormat string // --log-format <log_format>
	LogLevel  string // --log-level <log_level>
	LogOutput string // --log-output <log_output>
	// Extra
	WorkingDir string
	Stdout     io.Writer
	Stderr     io.Writer
	// internal usage
	cmd  *exec.Cmd
	done chan struct{}
}

// NewJnode returns a Jnode with some defaults.
func NewJnode() *Jnode {
	return &Jnode{
		WorkingDir: os.TempDir(),
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		done:       make(chan struct{}),
	}
}

func (jnode *Jnode) buildCmdArg() []string {
	var arg []string

	if jnode.ConfigFile != "" {
		arg = append(arg, "--config", jnode.ConfigFile)
	}

	if jnode.GenesisBlock != "" {
		arg = append(arg, "--genesis-block", jnode.GenesisBlock)
	}

	if jnode.GenesisBlockHash != "" {
		arg = append(arg, "--genesis-block-hash", jnode.GenesisBlockHash)
	}

	if jnode.Storage != "" {
		arg = append(arg, "--storage", jnode.Storage)
	}

	if len(jnode.SecretFiles) > 0 {
		for i := range jnode.SecretFiles {
			arg = append(arg, "--secret", jnode.SecretFiles[i])
		}
	}

	if len(jnode.TrustedPeers) > 0 {
		for i := range jnode.TrustedPeers {
			arg = append(arg, "--trusted-peer", jnode.TrustedPeers[i])
		}
	}

	if jnode.LogFormat != "" {
		arg = append(arg, "--log-format", jnode.LogFormat)
	}
	if jnode.LogLevel != "" {
		arg = append(arg, "--log-level", jnode.LogLevel)
	}
	if jnode.LogOutput != "" {
		arg = append(arg, "--log-output", jnode.LogOutput)
	}

	if jnode.RestListen != "" {
		arg = append(arg, "--rest-listen", jnode.RestListen)
	}

	if jnode.EnableExplorer {
		arg = append(arg, "--enable-explorer")
	}

	return arg
}

// Run starts the node.
func (jnode *Jnode) Run() error {
	jnode.cmd = exec.Command(jnodeName, jnode.buildCmdArg()...)

	jnode.cmd.Dir = jnode.WorkingDir
	jnode.cmd.Stdout = jnode.Stdout
	jnode.cmd.Stderr = jnode.Stderr

	err := jnode.cmd.Start()
	if err != nil {
		return err
	}

	// FIXME: find an effective way to catch errors of stderr
	// since cmd.Start does not care about them.
	// Ex: config errors will cause Start() to report no errors
	//     when in fact the node reports errors on stderr and stops.

	jnode.handleSigs()
	go jnode.cmdWait()

	return nil
}

// cmdWait for the node to terminate,
// and close the stop channel.
func (jnode *Jnode) cmdWait() {
	err := jnode.cmd.Wait()
	if err != nil {
		log.Printf("cmd.Wait() - %v", err) // FIXME: handle shutdown
	}
	select {
	case <-jnode.done:
	default:
		close(jnode.done)
	}
}

// handleSigs SIGINT + SIGTERM
func (jnode *Jnode) handleSigs() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		_ = jnode.Stop()
	}()
}

// Wait for the node to stop.
func (jnode *Jnode) Wait() {
	<-jnode.done
}

// Stop the node if running.
func (jnode *Jnode) Stop() error {
	if jnode.cmd.Process == nil {
		return fmt.Errorf("%s : exec: not started", "jnode.Stop")
	}
	return jnode.cmd.Process.Kill()
}

// StopAfter seconds.
func (jnode *Jnode) StopAfter(d time.Duration) error {
	if jnode.cmd.Process == nil {
		return fmt.Errorf("%s : exec: not started", "jnode.StopAfter")
	}

	go func() {
		select {
		case <-jnode.done:
		case <-time.After(d):
			_ = jnode.Stop()
		}
	}()
	return nil
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
	jnode.SecretFiles = append(jnode.SecretFiles, secretFile)
}

// AddTrustedPeer to node config
func (jnode *Jnode) AddTrustedPeer(address string, id string) {
	jnode.TrustedPeers = append(jnode.TrustedPeers, address+"@"+id)
}

// BinName set the executable name/full path if not the default one.
func BinName(name string) {
	jnodeName = name
}
