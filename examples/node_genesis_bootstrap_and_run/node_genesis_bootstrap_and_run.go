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

// get a fixed date if possible.
// needed for testing only to have a known genesis hash.
func block0Date() int64 {
	block0Date, err := time.Parse(time.RFC3339, "2017-09-29T00:00:00.000Z")
	if err != nil {
		return time.Now().Unix()
	}
	return block0Date.Unix()
}

/* seeds used [0-4] */
func main() {
	var (
		err error

		// Rest
		restAddr    = "127.0.0.1" // rest ip
		restPort    = 8001        // rest port
		restAddress = restAddr + ":" + strconv.Itoa(restPort)

		// P2P
		p2pIPver = "ip4" // ipv4 or ipv6
		p2pProto = "tcp" // tcp

		// P2P Public
		p2pPubAddr       = "127.0.0.1" // PublicAddres
		p2pPubPort       = 9001        // node P2P Public Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPubPort)

		// P2P Listen
		p2pListenAddr    = "127.0.0.1" // ListenAddress
		p2pListenPort    = 9001        // node P2P Public Port
		p2pListenAddress = "/" + p2pIPver + "/" + p2pListenAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pListenPort)

		// General
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

	////////////////
	// BFT LEADER //
	////////////////

	leaderSK, err := jcli.KeyGenerate(seed(2), "Ed25519", "")
	fatalOn(err, b2s(leaderSK))
	leaderPK, err := jcli.KeyToPublic(leaderSK, "", "")
	fatalOn(err, b2s(leaderPK))

	/////////////////////////
	// STAKE POOL Creation //
	/////////////////////////

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

	// note we will use the Faucet and Fixed as owners of this pool
	stakePoolOwners := []string{
		b2s(faucetPK),
		b2s(fixedPK),
	}
	stakePoolManagementThreshold := uint16(len(stakePoolOwners)) // uint16(2) -  (since we have 2 owners)
	stakePoolSerial := uint64(1010101010)
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

	///////////////////////////
	// STAKE POOL Delegation //
	///////////////////////////

	stakePoolID, err := jcli.CertificateGetStakePoolID(stakePoolCertSigned, "", "")
	fatalOn(err, b2s(stakePoolID))

	// FAUCET delegation (is also one the pool owners)
	stakeDelegationFaucetCert, err := jcli.CertificateNewStakeDelegation(b2s(stakePoolID), b2s(faucetPK), "")
	fatalOn(err, b2s(stakeDelegationFaucetCert))
	stakeDelegationFaucetCertSigned, err := jcli.CertificateSign(stakeDelegationFaucetCert, faucetFileSK, "", "")
	fatalOn(err, b2s(stakeDelegationFaucetCertSigned))

	// FIXED delegation (is also one the pool owners)
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
	block0cfg.BlockchainConfiguration.Block0Date = block0Date()
	block0cfg.BlockchainConfiguration.Block0Consensus = consensus
	block0cfg.BlockchainConfiguration.Discrimination = block0Discrimination

	block0cfg.BlockchainConfiguration.SlotDuration = 2
	block0cfg.BlockchainConfiguration.SlotsPerEpoch = 150

	block0cfg.BlockchainConfiguration.LinearFees.Certificate = 10000
	block0cfg.BlockchainConfiguration.LinearFees.Coefficient = 50
	block0cfg.BlockchainConfiguration.LinearFees.Constant = 1000

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
	fatalOn(err)

	//////////////////////////////////////////////////////////////////
	// START - Add BULK generated addresses to genesis block0       //
	for i := range srcFaucets {
		err = block0cfg.AddInitialFund(srcFaucets[i], 25_000_000_000)
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

	block0Hash, err := jcli.GenesisHash(block0Bin, "")
	fatalOn(err, b2s(block0Hash))

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

	nodeCfg.Storage = "jnode_storage"

	nodeCfg.Rest.Enabled = true       // default is "false" (rest disabled)
	nodeCfg.Rest.Listen = restAddress // 127.0.0.1:8443 is default value

	nodeCfg.Explorer.Enabled = false // default is "false" (explorer disabled)

	nodeCfg.P2P.PublicAddress = p2pPublicAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.ListenAddress = p2pListenAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.AllowPrivateAddresses = true     // for private addresses

	nodeCfg.Log.Level = "info" // default is "trace"

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

	log.Printf("Genesis Hash: %s", block0Hash)
	log.Printf("StakePool ID: %s", stakePoolID)
	log.Println("Leader Node - Running...")
	node.Wait()                          // Wait for the node to stop.
	log.Println("...Leader Node - Done") // All done. Node has stopped.
}
