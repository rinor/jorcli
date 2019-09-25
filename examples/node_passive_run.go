//$(which go) run $0 $@; exit $?

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/rinor/jorcli/jnode"
)

// fatalOn be careful with it in production,
// since it uses os.Exit(1) which affects the control flow.
// use pattern:
// if err != nil {
// 	....
// }
func fatalOn(err error, str ...string) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Fatalf("%s:%d %s -> %s", fn, line, str, err.Error())
	}
}

// b2s converts []byte to string with all leading
// and trailing white space removed, as defined by Unicode.
func b2s(b []byte) string {
	return strings.TrimSpace(string(b))
}

func main() {

	var (
		err error

		// Rest
		restAddr    = "127.0.0.1" // rest ip
		restPort    = 8002        // rest port
		restAddress = restAddr + ":" + strconv.Itoa(restPort)

		// P2P
		p2pIPver         = "ip4"       // ipv4 or ipv6
		p2pProto         = "tcp"       // tcp
		p2pPubAddr       = "127.0.0.1" // PublicAddres
		p2pPort          = 9002        // node P2P Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPort)

		// Trusted peers
		trustedPeer = "/ip4/127.0.0.1/tcp/9001"

		// Genesis Block0 Hash retrieved from example (1)
		block0Hash = "9facc7df455ee673f409ca062da0104f15b8b729a0faf694f457d3f3d390e6a8"
	)

	// set binary name/path if not default,
	jnode.BinName("jormungandr") // default is "jormungandr"

	// get jormungandr version
	jormungandrVersion, err := jnode.VersionFull()
	fatalOn(err, b2s(jormungandrVersion))
	log.Printf("Using: %s", jormungandrVersion)

	// create a new temporary directory inside your systems temp dir
	workingDir, err := ioutil.TempDir("", "jnode_")
	fatalOn(err, "workingDir")
	log.Printf("Working Directory: %s", workingDir)

	///////////////////
	//  node config  //
	///////////////////

	nodeCfg := jnode.NewNodeConfig()

	nodeCfg.Storage = ""                         // memory storage ("jnode_storage" default)
	nodeCfg.Rest.Listen = restAddress            // 127.0.0.1:8443 is default value
	nodeCfg.P2P.PublicAddress = p2pPublicAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.Log.Level = "debug"                  // default is "trace"

	// config not yet available on upstream,
	// it will be needed for testing on private ip addresses
	// nodeCfg.P2P.AllowPrivateAddresses = true // default false

	nodeCfgYaml, err := nodeCfg.ToYaml()
	fatalOn(err)
	// need this file for starting the node (--config)
	nodeCfgFile := workingDir + string(os.PathSeparator) + "node-config.yaml"
	err = ioutil.WriteFile(nodeCfgFile, nodeCfgYaml, 0644)
	fatalOn(err)

	// fmt.Printf("%s", nodeCfgYaml)

	//////////////////////
	// running the node //
	//////////////////////

	node := jnode.NewJnode()
	node.WorkingDir = workingDir
	node.ConfigFile = nodeCfgFile
	node.AddTrustedPeer(trustedPeer)   // add leader from example (1) as trusted
	node.GenesisBlockHash = block0Hash // add block0 hash

	node.Stdout, err = os.Create(filepath.Join(workingDir, "stdout.log"))
	fatalOn(err)
	node.Stderr, err = os.Create(filepath.Join(workingDir, "stderr.log"))
	fatalOn(err)

	// Run the node (Start + Wait)
	err = node.Run()
	if err != nil {
		log.Fatalf("node.Run FAILED: %v", err)
	}

	// _ = node.Stop() // Stop the node now
	_ = node.StopAfter(60 * time.Minute) // Stop the node after time.Duration

	log.Printf("Genesis Hash: %s", block0Hash)
	log.Println("Passive Node - Running...")
	node.Wait()                           // Wait for the node to stop.
	log.Println("...Passive Node - Done") // All done. Node has stopped.

}
