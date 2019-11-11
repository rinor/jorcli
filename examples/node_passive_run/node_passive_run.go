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

	"github.com/rinor/jorcli/jcli"
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
		restAddr    = "127.0.0.44" // rest ip
		restPort    = 8001         // rest port
		restAddress = restAddr + ":" + strconv.Itoa(restPort)

		// P2P
		p2pIPver = "ip4" // ipv4 or ipv6
		p2pProto = "tcp" // tcp

		// P2P Public
		p2pPubAddr       = "127.0.0.44" // PublicAddres
		p2pPubPort       = 9001         // node P2P Public Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPubPort)

		// P2P Listen
		p2pListenAddr    = "127.0.0.44" // ListenAddress
		p2pListenPort    = 9001         // node P2P Public Port
		p2pListenAddress = "/" + p2pIPver + "/" + p2pListenAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pListenPort)

		// Trusted peers
		/*
			leaderAddr = "/ip4/127.0.0.11/tcp/9001"                         // Leader (genesis) node (example 1)
			leaderID   = "000000000000000000000000000000000000000000000030" // Leader public_id

			gepAddr = "/ip4/127.0.0.22/tcp/9001"                         // Genesis stake pool node (example 2)
			gepID   = "000000000000000000000000000000000000000000003130" // Genesis stake pool public_id
		*/
		delegatorAddr = "/ip4/127.0.0.33/tcp/9001"                         // stake pool node (example 3)
		delegatorID   = "333333333333333333333333333333333333333333333333" // delegator pool public_id

		// Genesis Block0 Hash retrieved from example (1)
		block0Hash = "116f3e765a825a68dc1ac0a3f8993447dccef5641b0450e31dbe0a2cf1c79cad"

		// Node config log
		nodeCfgLogLevel = "debug"
	)

	// Set RUST_BACKTRACE=full env
	err = os.Setenv("RUST_BACKTRACE", "full")
	fatalOn(err, "Failed to set env (RUST_BACKTRACE=full)")

	// set binary name/path if not default,
	// provided as example since the ones set here,
	// are also the default values.
	jcli.BinName("jcli")         // default is "jcli"
	jnode.BinName("jormungandr") // default is "jormungandr"

	// get jcli version
	jcliVersion, err := jcli.VersionFull()
	fatalOn(err, b2s(jcliVersion))
	log.Printf("Using: %s", jcliVersion)

	// get jormungandr version
	jormungandrVersion, err := jnode.VersionFull()
	fatalOn(err, b2s(jormungandrVersion))
	log.Printf("Using: %s", jormungandrVersion)

	// create a new temporary directory inside your systems temp dir
	workingDir, err := ioutil.TempDir("", "jnode_passive_")
	fatalOn(err, "workingDir")
	log.Println()
	log.Printf("Working Directory: %s", workingDir)

	///////////////////
	//  node config  //
	///////////////////

	// p2p node public_id
	nodePublicID := "444444444444444444444444444444444444444444444444"

	nodeCfg := jnode.NewNodeConfig()

	nodeCfg.Storage = "jnode_storage"

	nodeCfg.Rest.Enabled = true       // default is "false" (rest disabled)
	nodeCfg.Rest.Listen = restAddress // 127.0.0.1:8443 is default value

	nodeCfg.Explorer.Enabled = true // default is "false" (explorer disabled)

	nodeCfg.P2P.PublicAddress = p2pPublicAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.ListenAddress = p2pListenAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.PublicID = nodePublicID          // jörmungandr will generate a random key, if not set
	nodeCfg.P2P.AllowPrivateAddresses = true     // for private addresses
	nodeCfg.P2P.Policy.QuarantineDuration = "5m" // default to "30m"

	// add trusted peer to config file
	// nodeCfg.AddTrustedPeer(leaderAddr, leaderID)
	// nodeCfg.AddTrustedPeer(gepAddr, gepID)
	nodeCfg.AddTrustedPeer(delegatorAddr, delegatorID)

	nodeCfg.Log.Level = nodeCfgLogLevel // default is "trace"

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
	node.GenesisBlockHash = block0Hash // add block0 hash

	// add trusted peer cmd args (not needed if using config)
	// node.AddTrustedPeer(leaderAddr, leaderID)       // add leader from example (1) as trusted
	// node.AddTrustedPeer(gepAddr, gepID)             // add genesis stake pool from example (2) as trusted
	node.AddTrustedPeer(delegatorAddr, delegatorID) // add delegator stake pool from example (3) as trusted

	node.Stdout, err = os.Create(filepath.Join(workingDir, "stdout.log"))
	fatalOn(err)
	node.Stderr, err = os.Create(filepath.Join(workingDir, "stderr.log"))
	fatalOn(err)

	// Run the node (Start + Wait)
	err = node.Run()
	if err != nil {
		log.Fatalf("node.Run FAILED: %v", err)
	}

	log.Println()
	log.Printf("Genesis Hash: %s", block0Hash)
	log.Println()
	log.Printf("NodePublicID for trusted: %s", nodePublicID)
	log.Println()

	log.Println("Passive/Explorer Node - Running...")
	node.Wait()                                    // Wait for the node to stop.
	log.Println("...Passive/Explorer Node - Done") // All done. Node has stopped.
}
