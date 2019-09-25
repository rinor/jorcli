//$(which go) run $0 $@; exit $?

package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

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

// seed generated from an int. For the same int the same seed is returned.
// Useful for reproducible batch key generation,
// for example the index of a slice/array can be a param.
func seed(i int) string {
	in := []byte(strconv.Itoa(i))
	out := make([]byte, 32-len(in), 32)
	out = append(out, in...)

	return hex.EncodeToString(out)
}

// b2s converts []byte to string with all leading
// and trailing white space removed, as defined by Unicode.
func b2s(b []byte) string {
	return strings.TrimSpace(string(b))
}

// buildAccountAddr returns a new account address
func buildAccountAddr(seed string, addressPrefix string, discrimination string) (string, error) {
	var (
		err error

		faucetSK   []byte
		faucetPK   []byte
		faucetAddr []byte
	)
	// private key
	faucetSK, err = jcli.KeyGenerate(seed, "Ed25519Extended", "")
	if err != nil {
		return "", fmt.Errorf("KeyGenerate: %s - %s", err, faucetSK)
	}
	// public key
	faucetPK, err = jcli.KeyToPublic(faucetSK, "", "")
	if err != nil {
		return "", fmt.Errorf("KeyToPublic: %s - %s", err, faucetPK)
	}
	// account address
	faucetAddr, err = jcli.AddressAccount(b2s(faucetPK), addressPrefix, discrimination)
	if err != nil {
		return "", fmt.Errorf("AddressAccount: %s - %s", err, faucetAddr)
	}
	return b2s(faucetAddr), err
}

/* seeds used [0-1,10-11] */
func main() {
	var (
		err error

		// Rest
		restAddr       = "127.0.0.1" // rest ip
		restPort       = 8003        // rest port
		restAddress    = restAddr + ":" + strconv.Itoa(restPort)
		restApiAddress = "http://" + restAddr + ":" + strconv.Itoa(restPort) + "/api" // self

		// P2P
		p2pIPver         = "ip4"       // ipv4 or ipv6
		p2pProto         = "tcp"       // tcp
		p2pPubAddr       = "127.0.0.1" // PublicAddres
		p2pPort          = 9003        // node P2P Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPort)

		discrimination = "testing"  // "" (empty defaults to "production")
		addressPrefix  = "jnode_ta" // "" (empty defaults to "ca")

		// Trusted peers
		trustedPeerLeader  = "/ip4/127.0.0.1/tcp/9001" // Leader (genesis) node
		trustedPeerPassive = "/ip4/127.0.0.1/tcp/9002" // Passive node

		// Genesis Block0 Hash retrieved from example (1)
		block0Hash = "1162376908bb94488eb2e2d4cc4572b192034a3eb603a3019a0a471683d10333"
	)

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
	workingDir, err := ioutil.TempDir("", "jnode_")
	fatalOn(err, "workingDir")
	log.Printf("Working Directory: %s", workingDir)

	////////////
	// FAUCET //
	////////////

	// will need this one file later for certificate signing
	faucetFileSK := workingDir + string(os.PathSeparator) + "faucet_key.sk"

	faucetSK, err := jcli.KeyGenerate(seed(0), "Ed25519Extended", faucetFileSK)
	fatalOn(err, b2s(faucetSK))
	faucetPK, err := jcli.KeyToPublic(faucetSK, "", "")
	fatalOn(err, b2s(faucetPK))
	faucetAddr, err := jcli.AddressAccount(b2s(faucetPK), addressPrefix, discrimination)
	fatalOn(err, b2s(faucetAddr))

	///////////
	// FIXED //
	///////////

	// will need this one file later for certificate signing
	fixedFileSK := workingDir + string(os.PathSeparator) + "fixed_key.sk"

	fixedSK, err := jcli.KeyGenerate(seed(1), "Ed25519Extended", fixedFileSK)
	fatalOn(err, b2s(fixedSK))
	fixedPK, err := jcli.KeyToPublic(fixedSK, "", "")
	fatalOn(err, b2s(fixedPK))
	fixedAddr, err := jcli.AddressAccount(b2s(fixedPK), addressPrefix, discrimination)
	fatalOn(err, b2s(fixedAddr))

	/////////////////////////
	// STAKE POOL Creation //
	/////////////////////////

	// VRF
	poolVrfSK, err := jcli.KeyGenerate(seed(10), "Curve25519_2HashDH", "")
	fatalOn(err, b2s(poolVrfSK))
	poolVrfPK, err := jcli.KeyToPublic(poolVrfSK, "", "")
	fatalOn(err, b2s(poolVrfPK))

	// KES
	poolKesSK, err := jcli.KeyGenerate(seed(11), "SumEd25519_12", "")
	fatalOn(err, b2s(poolKesSK))
	poolKesPK, err := jcli.KeyToPublic(poolKesSK, "", "")
	fatalOn(err, b2s(poolKesPK))

	// note we will use the Faucet and Fixed as owners of this pool
	stakePoolOwners := []string{
		b2s(faucetPK),
		b2s(fixedPK),
	}
	stakePoolManagementThreshold := uint16(len(stakePoolOwners)) // uint16(2) -  (since we have 2 owners)
	stakePoolSerial := uint64(2020202020)
	stakePoolStartValidity := uint64(0)

	stakePoolCert, err := jcli.CertificateNewStakePoolRegistration(
		b2s(poolKesPK),
		b2s(poolVrfPK),
		stakePoolStartValidity,
		stakePoolManagementThreshold,
		stakePoolSerial,
		stakePoolOwners,
		"",
	)
	fatalOn(err, b2s(stakePoolCert))

	// Sign the certificate with FAUCET private key
	stakePoolCertSigned, err := jcli.CertificateSign(stakePoolCert, faucetFileSK, "", "")
	fatalOn(err, b2s(stakePoolCertSigned))

	// Sign the certificate also with FIXED private key
	stakePoolCertSigned, err = jcli.CertificateSign(stakePoolCertSigned, fixedFileSK, "", "")
	fatalOn(err, b2s(stakePoolCertSigned))

	stakePoolID, err := jcli.CertificateGetStakePoolID(stakePoolCertSigned, "", "")
	fatalOn(err, b2s(stakePoolID))

	//////////////////////
	//  secrets config  //
	//////////////////////

	secretCfg := jnode.NewSecretConfig()

	secretCfg.Genesis.SigKey = b2s(poolKesSK)
	secretCfg.Genesis.VrfKey = b2s(poolVrfSK)
	secretCfg.Genesis.NodeID = b2s(stakePoolID)

	secretCfgYaml, err := secretCfg.ToYaml()
	fatalOn(err)
	// need this file for starting the node (--secret)
	secretCfgFile := workingDir + string(os.PathSeparator) + "pool-secret.yaml"
	err = ioutil.WriteFile(secretCfgFile, secretCfgYaml, 0644)
	fatalOn(err)

	// fmt.Printf("%s", secretCfgYaml)

	///////////////////
	//  node config  //
	///////////////////

	nodeCfg := jnode.NewNodeConfig()

	nodeCfg.Storage = ""                         // memory storage ("jnode_storage" default)
	nodeCfg.Rest.Listen = restAddress            // 127.0.0.1:8443 is default value
	nodeCfg.P2P.PublicAddress = p2pPublicAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.Log.Level = "debug"                  // default is "trace"

	// needed for testing on private ip addresses
	nodeCfg.P2P.AllowPrivateAddresses = true // default false

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
	node.GenesisBlockHash = block0Hash      // add block0 hash
	node.AddTrustedPeer(trustedPeerLeader)  // add leader from example (1) as trusted
	node.AddTrustedPeer(trustedPeerPassive) // add passive from example (3) as trusted

	node.AddSecretFile(secretCfgFile)
	// or node.SecretFiles = append(node.SecretFiles, secretCfgFile)

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

	/*****************************************************************

		At this point the StakePool is configured and running,
		but the node behaves like a passive one since:
		1) the StakePool is not yet registered on the network
		2) Even if it was registered the StakePool has no stake yet.

	*******************************************************************/

	log.Printf("Genesis Hash: %s", block0Hash)
	log.Printf("StakePool ID: %s", stakePoolID)
	log.Println("StakePool Node - Running...")
	node.Wait()                             // Wait for the node to stop.
	log.Println("...StakePool Node - Done") // All done. Node has stopped.
}
