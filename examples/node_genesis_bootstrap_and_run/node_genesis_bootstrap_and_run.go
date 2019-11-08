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

// nodePID builds a node public_id from a seed int
// For the same int the same value is returned.
func nodePID(i int) string {
	in := []byte(strconv.Itoa(i))
	out := make([]byte, 24-len(in), 24)
	out = append(out, in...)

	return hex.EncodeToString(out)
}

// buildAccountAddr returns a new account address
func buildAccountAddr(seed string, addressPrefix string, discrimination string) (string, error) {
	var (
		err error

		secretKey   []byte
		publicKey   []byte
		accountAddr []byte
	)
	// private key
	secretKey, err = jcli.KeyGenerate(seed, "Ed25519Extended", "")
	if err != nil {
		return "", fmt.Errorf("KeyGenerate: %s - %s", err, secretKey)
	}
	// public key
	publicKey, err = jcli.KeyToPublic(secretKey, "", "")
	if err != nil {
		return "", fmt.Errorf("KeyToPublic: %s - %s", err, publicKey)
	}
	// account address
	accountAddr, err = jcli.AddressAccount(b2s(publicKey), addressPrefix, discrimination)
	if err != nil {
		return "", fmt.Errorf("AddressAccount: %s - %s", err, accountAddr)
	}
	return b2s(accountAddr), err
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

/* seeds used [0-5], [50-52] ,[60], [100-1099] */
const (
	seedPublicID = 0 // seed for p2p public_id
	faucetSeed   = 1 // seed for faucet
	fixedSeed    = 2 // seed for fixed
	leaderSeed   = 3 // seed for bft leader
	localVrfSeed = 4 // seed for local pool VRF
	localKesSeed = 5 // seed for local pool KES

	gepSeed    = 50 // seed the owner of an extra pool in genesis block
	gepVrfSeed = 51 // seed for extra pool VRF
	gepKesSeed = 52 // seed for extra pool KES

	delegatorSeed = 60 // seed for new stake delegator example (3)

	seedStartBulk  = 100  // seed key generation start
	totSrcAddrBulk = 1000 // total number of account addresses
)

func main() {
	var (
		err error

		// Rest
		restAddr    = "127.0.0.11" // rest ip
		restPort    = 8001         // rest port
		restAddress = restAddr + ":" + strconv.Itoa(restPort)

		// P2P
		p2pIPver = "ip4" // ipv4 or ipv6
		p2pProto = "tcp" // tcp

		// P2P Public
		p2pPubAddr       = "127.0.0.11" // PublicAddres
		p2pPubPort       = 9001         // node P2P Public Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPubPort)

		// P2P Listen
		p2pListenAddr    = "127.0.0.11" // ListenAddress
		p2pListenPort    = 9001         // node P2P Public Port
		p2pListenAddress = "/" + p2pIPver + "/" + p2pListenAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pListenPort)

		// General
		consensus      = "genesis_praos" // bft or genesis_praos
		discrimination = "testing"       // "" (empty defaults to "production")
		addressPrefix  = "jnode_ta"      // "" (empty defaults to "ca")

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
	workingDir, err := ioutil.TempDir("", "jnode_genesis_")
	fatalOn(err, "workingDir")
	log.Println()
	log.Printf("Working Directory: %s", workingDir)

	/////////////////////////////////////////////////////////////////////////////////////
	// START - BULK generate                                                           //
	//                                                                                 //
	// this will be used to generate bulk addresses and include them in genesis block0 //
	// so we can use them as source for bulk transactions from other examples.         //

	var (
		srcFaucets [totSrcAddrBulk]string
	)
	// build bulk addresses
	for i := 0; i < totSrcAddrBulk; i++ {
		srcFaucets[i], err = buildAccountAddr(
			seed(seedStartBulk+i),
			addressPrefix,
			discrimination,
		)
		fatalOn(err)
	}

	//                                                                                //
	// DONE - BULK generate                                                           //
	////////////////////////////////////////////////////////////////////////////////////

	///////////////
	// DELEGATOR //
	///////////////

	// generate also an account for delegating in example(3)
	delegatorAddr, err := buildAccountAddr(
		seed(delegatorSeed),
		addressPrefix,
		discrimination,
	)
	fatalOn(err)

	////////////
	// FAUCET //
	////////////

	// will need this one file later for certificate signing
	faucetFileSK := workingDir + string(os.PathSeparator) + "faucet_key.sk"

	faucetSK, err := jcli.KeyGenerate(seed(faucetSeed), "Ed25519Extended", faucetFileSK)
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

	fixedSK, err := jcli.KeyGenerate(seed(fixedSeed), "Ed25519Extended", fixedFileSK)
	fatalOn(err, b2s(fixedSK))
	fixedPK, err := jcli.KeyToPublic(fixedSK, "", "")
	fatalOn(err, b2s(fixedPK))
	fixedAddr, err := jcli.AddressAccount(b2s(fixedPK), addressPrefix, discrimination)
	fatalOn(err, b2s(fixedAddr))

	////////////////
	// BFT LEADER //
	////////////////

	leaderSK, err := jcli.KeyGenerate(seed(leaderSeed), "Ed25519", "")
	fatalOn(err, b2s(leaderSK))
	leaderPK, err := jcli.KeyToPublic(leaderSK, "", "")
	fatalOn(err, b2s(leaderPK))

	///////////////////////////////////////
	// GENESIS Local STAKE POOL Creation //
	///////////////////////////////////////

	// VRF
	poolVrfSK, err := jcli.KeyGenerate(seed(localVrfSeed), "Curve25519_2HashDH", "")
	fatalOn(err, b2s(poolVrfSK))
	poolVrfPK, err := jcli.KeyToPublic(poolVrfSK, "", "")
	fatalOn(err, b2s(poolVrfPK))

	// KES
	poolKesSK, err := jcli.KeyGenerate(seed(localKesSeed), "SumEd25519_12", "")
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

	// Sign the certificate with FAUCET private key and also with FIXED private key
	stakePoolCertSigned, err := jcli.CertificateSign(stakePoolCert, []string{faucetFileSK, fixedFileSK}, "", "")
	fatalOn(err, b2s(stakePoolCertSigned))

	/////////////////////////////////
	// Local STAKE POOL Delegation //
	/////////////////////////////////

	// We can get poolID from signed or unsigned certificate
	stakePoolID, err := jcli.CertificateGetStakePoolID(stakePoolCert, "", "")
	fatalOn(err, b2s(stakePoolID))

	// FAUCET delegation (is also one the pool owners)
	stakeDelegationFaucetCert, err := jcli.CertificateNewStakeDelegation(b2s(stakePoolID), b2s(faucetPK), "")
	fatalOn(err, b2s(stakeDelegationFaucetCert))
	stakeDelegationFaucetCertSigned, err := jcli.CertificateSign(stakeDelegationFaucetCert, []string{faucetFileSK}, "", "")
	fatalOn(err, b2s(stakeDelegationFaucetCertSigned))

	// FIXED delegation (is also one the pool owners)
	stakeDelegationFixedCert, err := jcli.CertificateNewStakeDelegation(b2s(stakePoolID), b2s(fixedPK), "")
	fatalOn(err, b2s(stakeDelegationFixedCert))
	stakeDelegationFixedCertSigned, err := jcli.CertificateSign(stakeDelegationFixedCert, []string{fixedFileSK}, "", "")
	fatalOn(err, b2s(stakeDelegationFixedCertSigned))

	/**********************************************************************************************************/

	///////////////////////////////////////
	// GENESIS Extra STAKE POOL Creation //
	///////////////////////////////////////

	//////////////////////////////
	// Genesis Extra Pool Owner //
	//////////////////////////////

	// will need this one file later for certificate signing
	gepoFileSK := workingDir + string(os.PathSeparator) + "gepo_key.sk"

	gepoSK, err := jcli.KeyGenerate(seed(gepSeed), "Ed25519Extended", gepoFileSK)
	fatalOn(err, b2s(gepoSK))
	gepoPK, err := jcli.KeyToPublic(gepoSK, "", "")
	fatalOn(err, b2s(gepoPK))
	gepoAddr, err := jcli.AddressAccount(b2s(gepoPK), addressPrefix, discrimination)
	fatalOn(err, b2s(gepoAddr))

	// gep VRF
	gepPoolVrfSK, err := jcli.KeyGenerate(seed(gepVrfSeed), "Curve25519_2HashDH", "")
	fatalOn(err, b2s(gepPoolVrfSK))
	gepPoolVrfPK, err := jcli.KeyToPublic(gepPoolVrfSK, "", "")
	fatalOn(err, b2s(gepPoolVrfPK))

	// gep KES
	gepPoolKesSK, err := jcli.KeyGenerate(seed(gepKesSeed), "SumEd25519_12", "")
	fatalOn(err, b2s(gepPoolKesSK))
	gepPoolKesPK, err := jcli.KeyToPublic(gepPoolKesSK, "", "")
	fatalOn(err, b2s(gepPoolKesPK))

	// Genesis Extra Pool Owner of this pool
	gepStakePoolOwners := []string{
		b2s(gepoPK),
	}
	gepStakePoolManagementThreshold := uint16(len(gepStakePoolOwners)) // uint16(1) -  (since we have 1 owner)
	gepStakePoolSerial := uint64(2020202020)
	gepStakePoolStartValidity := uint64(0)

	gepStakePoolCert, err := jcli.CertificateNewStakePoolRegistration(
		b2s(gepPoolKesPK),
		b2s(gepPoolVrfPK),
		gepStakePoolStartValidity,
		gepStakePoolManagementThreshold,
		gepStakePoolSerial,
		gepStakePoolOwners,
		"",
	)
	fatalOn(err, b2s(gepStakePoolCert))

	// Sign the certificate with Genesis Extra Pool Owner private key
	gepStakePoolCertSigned, err := jcli.CertificateSign(gepStakePoolCert, []string{gepoFileSK}, "", "")
	fatalOn(err, b2s(gepStakePoolCertSigned))

	/////////////////////////////////
	// Extra STAKE POOL Delegation //
	/////////////////////////////////

	gepStakePoolID, err := jcli.CertificateGetStakePoolID(gepStakePoolCertSigned, "", "")
	fatalOn(err, b2s(gepStakePoolID))

	// Genesis Extra Pool Owner delegation
	stakeDelegationGepoCert, err := jcli.CertificateNewStakeDelegation(b2s(gepStakePoolID), b2s(gepoPK), "")
	fatalOn(err, b2s(stakeDelegationGepoCert))
	stakeDelegationGepoCertSigned, err := jcli.CertificateSign(stakeDelegationGepoCert, []string{gepoFileSK}, "", "")
	fatalOn(err, b2s(stakeDelegationGepoCertSigned))

	/**********************************************************************************************************/

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
	block0cfg.BlockchainConfiguration.SlotsPerEpoch = 450
	// block0cfg.BlockchainConfiguration.KesUpdateSpeed = 300

	block0cfg.BlockchainConfiguration.LinearFees.Certificate = 10000
	block0cfg.BlockchainConfiguration.LinearFees.Coefficient = 50
	block0cfg.BlockchainConfiguration.LinearFees.Constant = 1000

	// Bft Leader not used, just to satisfy config
	err = block0cfg.AddConsensusLeader(b2s(leaderPK))
	fatalOn(err)

	// Note: Funds involved in delegation (Faucet, Fixed, Gepo)
	// need to be added before the stake delegation certificates,
	// otherwise genesis encode will fail.
	//
	// check https://github.com/input-output-hk/jormungandr/issues/917
	//
	// FIXME: since initial is a slice, need to respect that order,
	// but need to handle this correctly to remove order dependency from code.

	// funds
	err = block0cfg.AddInitialFund(b2s(faucetAddr), 10_000_000_000_000_000)
	fatalOn(err)
	err = block0cfg.AddInitialFund(b2s(fixedAddr), 10_000_000_000_000_000)
	fatalOn(err)
	err = block0cfg.AddInitialFund(b2s(gepoAddr), 10_000_000_000_000_000)
	fatalOn(err)
	err = block0cfg.AddInitialFund(delegatorAddr, 10_000_000_000_000_000)
	fatalOn(err)
	//////////////////////////////////////////////////////////////////
	// START - Add BULK generated addresses to genesis block0       //
	for i := range srcFaucets {
		err = block0cfg.AddInitialFund(srcFaucets[i], 5_000_000_000_000)
		fatalOn(err)
	}
	// DONE - Add BULK generated addresses to genesis block0        //
	//////////////////////////////////////////////////////////////////

	// genesis main stake pool data
	err = block0cfg.AddInitialCertificate(b2s(stakePoolCertSigned))
	fatalOn(err)
	err = block0cfg.AddInitialCertificate(b2s(stakeDelegationFaucetCertSigned))
	fatalOn(err)
	err = block0cfg.AddInitialCertificate(b2s(stakeDelegationFixedCertSigned))
	fatalOn(err)
	// genesis extra stake pool data
	err = block0cfg.AddInitialCertificate(b2s(gepStakePoolCertSigned))
	fatalOn(err)
	err = block0cfg.AddInitialCertificate(b2s(stakeDelegationGepoCertSigned))
	fatalOn(err)

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

	//secretCfg.Bft.SigningKey = b2s(leaderSK) // Keep it

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

	// p2p node public_id
	nodePublicID := nodePID(seedPublicID)

	nodeCfg := jnode.NewNodeConfig()

	nodeCfg.Storage = "jnode_storage"

	nodeCfg.Rest.Enabled = true       // default is "false" (rest disabled)
	nodeCfg.Rest.Listen = restAddress // 127.0.0.1:8443 is default value

	nodeCfg.Explorer.Enabled = false // default is "false" (explorer disabled)

	nodeCfg.P2P.PublicAddress = p2pPublicAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.ListenAddress = p2pListenAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.PublicID = nodePublicID          // jörmungandr will generate a random key, if not set
	nodeCfg.P2P.AllowPrivateAddresses = true     // for private addresses

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
	node.GenesisBlock = block0BinFile
	node.ConfigFile = nodeCfgFile

	// node.AddTrustedPeer(trustedPeerGenesisStake)

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

	log.Println()
	log.Printf("Genesis Hash: %s", block0Hash)
	log.Println()
	log.Printf("LOCAL StakePool ID       : %s", stakePoolID)
	log.Printf("LOCAL StakePool Owner    : %s", faucetAddr)
	log.Printf("LOCAL StakePool Owner    : %s", fixedAddr)
	log.Printf("LOCAL StakePool Delegator: %s", faucetAddr)
	log.Printf("LOCAL StakePool Delegator: %s", fixedAddr)
	log.Println()
	log.Printf("EXTRA StakePool ID       : %s", gepStakePoolID)
	log.Printf("EXTRA StakePool Owner    : %s", gepoAddr)
	log.Printf("EXTRA StakePool Delegator: %s", gepoAddr)
	log.Println()
	log.Printf("NodePublicID for trusted: %s", nodePublicID)
	log.Println()

	log.Println("Genesis Node - Running...")
	node.Wait()                           // Wait for the node to stop.
	log.Println("...Genesis Node - Done") // All done. Node has stopped.
}
