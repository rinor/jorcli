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

func main() {
	var (
		err error

		restAddress = "127.0.0.1" // ip
		restPort    = 8443        // port

		consensus      = "genesis_praos" // bft or genesis_praos
		discrimination = "testing"       // "" (empty defaults to "production")
		addressPrefix  = "jnode_ta"      // "" (empty defaults to "ca")
	)

	// set binary name/path if not default,
	// provided as example since the ones set here,
	// are also the default values.
	jcli.BinName("jcli")         // default is "jcli"
	jnode.BinName("jormungandr") // default is "jormungandr"

	/////////////////////////////////////////////////////////////////////////////////////
	// START - BULK generate                                                           //
	//                                                                                 //
	// this will be used to generate bulk addresses and include them in genesis block0 //
	// so we can use them as source for bulk transactions from other examples.         //
	const ( //
		seedStart  = 101 // seed key generation start
		totSrcAddr = 100 // total number of account addresses
	)
	var srcFaucets [totSrcAddr]string
	// build bulk addresses
	for i := 0; i < totSrcAddr; i++ {
		srcFaucets[i], err = buildAccountAddr(
			seed(seedStart+i),
			addressPrefix,
			discrimination,
		)
		fatalOn(err)
	}
	//                                                                                //
	// DONE - BULK generate                                                           //
	////////////////////////////////////////////////////////////////////////////////////

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

	////////////
	// LEADER //
	////////////

	leaderSK, err := jcli.KeyGenerate(seed(2), "Ed25519", "")
	fatalOn(err, b2s(leaderSK))
	leaderPK, err := jcli.KeyToPublic(leaderSK, "", "")
	fatalOn(err, b2s(leaderPK))

	////////////////
	// STAKE POOL //
	////////////////

	// VRF
	poolVrfSK, err := jcli.KeyGenerate(seed(3), "Curve25519_2HashDH", "")
	fatalOn(err, b2s(poolVrfSK))
	poolVrfPK, err := jcli.KeyToPublic(poolVrfSK, "", "")
	fatalOn(err, b2s(poolVrfPK))

	// KES
	poolKesSK, err := jcli.KeyGenerate(seed(4), "SumEd25519_12", "")
	fatalOn(err, b2s(poolKesSK))
	poolKesPK, err := jcli.KeyToPublic(poolKesSK, "", "")
	fatalOn(err, b2s(poolKesPK))

	// note we will use the faucet as the owner to this pool
	var stakePoolOwners []string
	stakePoolOwners = append(stakePoolOwners, b2s(faucetPK))

	stakePoolCert, err := jcli.CertificateNewStakePoolRegistration(
		b2s(poolKesPK),
		b2s(poolVrfPK),
		uint64(0),
		uint16(1),
		uint64(1010101010),
		stakePoolOwners,
		"",
	)
	fatalOn(err, b2s(stakePoolCert))

	// note we are using faucet as the owner of the pool
	stakePoolCertSigned, err := jcli.CertificateSign(stakePoolCert, faucetFileSK, "", "")
	fatalOn(err, b2s(stakePoolCertSigned))

	stakePoolID, err := jcli.CertificateGetStakePoolID(stakePoolCertSigned, "", "")
	fatalOn(err, b2s(stakePoolID))

	stakeDelegationFaucetCert, err := jcli.CertificateNewStakeDelegation(b2s(stakePoolID), b2s(faucetPK), "")
	fatalOn(err, b2s(stakeDelegationFaucetCert))
	stakeDelegationFaucetCertSigned, err := jcli.CertificateSign(stakeDelegationFaucetCert, faucetFileSK, "", "")
	fatalOn(err, b2s(stakeDelegationFaucetCertSigned))

	stakeDelegationFixedCert, err := jcli.CertificateNewStakeDelegation(b2s(stakePoolID), b2s(fixedPK), "")
	fatalOn(err, b2s(stakeDelegationFixedCert))
	stakeDelegationFixedCertSigned, err := jcli.CertificateSign(stakeDelegationFixedCert, fixedFileSK, "", "")
	fatalOn(err, b2s(stakeDelegationFixedCertSigned))

	/////////////////////
	//  block0 config  //
	/////////////////////

	block0cfg := jnode.NewBlock0Config()

	block0Discrimination := "production"
	if discrimination == "testing" {
		block0Discrimination = "test"
	}

	// set/change config params
	block0cfg.BlockchainConfiguration.Block0Consensus = consensus
	block0cfg.BlockchainConfiguration.Discrimination = block0Discrimination
	block0cfg.BlockchainConfiguration.SlotDuration = 10
	block0cfg.BlockchainConfiguration.SlotsPerEpoch = 60
	block0cfg.BlockchainConfiguration.LinearFees.Constant = 10

	err = block0cfg.AddConsensusLeader(b2s(leaderPK))
	fatalOn(err)

	err = block0cfg.AddInitialFund(b2s(faucetAddr), 1_000_000_000_000)
	fatalOn(err)
	err = block0cfg.AddInitialFund(b2s(fixedAddr), 1_000_000_000_000)
	fatalOn(err)

	err = block0cfg.AddInitialCertificate(b2s(stakePoolCertSigned))
	fatalOn(err)
	err = block0cfg.AddInitialCertificate(b2s(stakeDelegationFaucetCertSigned))
	fatalOn(err)
	err = block0cfg.AddInitialCertificate(b2s(stakeDelegationFixedCertSigned))

	//////////////////////////////////////////////////////////////////
	// START - Add BULK generated addresses to genesis block0       //
	for i := range srcFaucets {
		err = block0cfg.AddInitialFund(srcFaucets[i], 1_000_000_000)
		fatalOn(err)
	}
	// DONE - Add BULK generated addresses to genesis block0        //
	//////////////////////////////////////////////////////////////////

	block0Yaml, err := block0cfg.ToYaml()
	fatalOn(err)
	// need this file for starting the node (--genesis-block)
	block0BinFile := workingDir + string(os.PathSeparator) + "block-0.bin"

	// block0BinFile will be created by jcli
	block0Bin, err := jcli.GenesisEncode(block0Yaml, "", block0BinFile)
	fatalOn(err, b2s(block0Bin))

	/*
		// Or we can create block0BinFile by our self
		block0Bin, err := jcli.GenesisEncode(block0Yaml, "", "")
		fatalOn(err, b2s(block0Bin))
		err = ioutil.WriteFile(block0BinFile, block0Bin, 0644)
		fatalOn(err)
	*/

	// fmt.Printf("%s", block0Yaml)

	//////////////////////
	//  secrets config  //
	//////////////////////

	secretCfg := jnode.NewSecretConfig()

	secretCfg.Genesis.SigKey = b2s(poolKesSK)
	secretCfg.Genesis.VrfKey = b2s(poolVrfSK)
	secretCfg.Genesis.NodeID = b2s(stakePoolID)
	secretCfg.Bft.SigningKey = b2s(leaderSK)

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

	nodeCfg.Storage = ""                                             // memory storage ("jnode_storage" default)
	nodeCfg.Rest.Listen = restAddress + ":" + strconv.Itoa(restPort) // 127.0.0.1:8443 is also default value
	nodeCfg.P2P.PublicAddress = "/ip4/" + restAddress + "/tcp/8299"  // /ip4/127.0.0.1/tcp/8299 is also default value
	nodeCfg.Log.Level = "debug"                                      // default is "trace"

	// config not yet available on upstream, it will be needed for testing on private ip addresses
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
	node.GenesisBlock = block0BinFile
	node.ConfigFile = nodeCfgFile
	node.AddSecretFile(secretCfgFile) // or node.SecretFiles = append(node.SecretFiles, secretCfgFile)

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

	log.Println("Running...")
	node.Wait()            // Wait for the node to stop.
	log.Println("...Done") // All done. Node has stopped.
}
