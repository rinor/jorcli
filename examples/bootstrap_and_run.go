//$(which go) run $0 $@; exit $?

package main

import (
	"encoding/hex"
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

func wait(secs int) {
	time.Sleep(time.Duration(secs) * time.Second)
}

func main() {
	var (
		restProto   = "http"      // proto
		restAddress = "127.0.0.1" // ip
		restPort    = 8443        // port
		// "http://127.0.0.1:8443/api"
		restApiAddress = restProto + "://" + restAddress + ":" + strconv.Itoa(restPort) + "/api" // proto://ip:port/api
	)
	_ = restApiAddress

	var (
		consensus      = "genesis_praos" // bft or genesis_praos
		discrimination = "testing"       // "" (empty defaults to "production")
		addressPrefix  = "jnode_ta"      // "" (empty defaults to "ca")
	)

	// set binary name/path if not default
	/*
		jcli.BinName("jcli")         // defaults to "jcli"
		jnode.BinName("jormungandr") // defaults to "jormungandr"
	*/

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
	block0cfg.BlockchainConfiguration.SlotsPerEpoch = 6
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
	fatalOn(err)

	block0Yaml, err := block0cfg.ToYaml()
	fatalOn(err)
	// need this file for starting the node
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
	// need this file for starting the node
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

	nodeCfgYaml, err := nodeCfg.ToYaml()
	fatalOn(err)
	// need this file for starting the node
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
	node.Config = nodeCfgFile
	node.AddSecretFile(secretCfgFile) // or node.Secrets = append(node.Secrets, secretCfgFile)

	node.Stdout, err = os.Create(filepath.Join(workingDir, "stdout.log"))
	fatalOn(err)
	node.Stderr, err = os.Create(filepath.Join(workingDir, "stderr.log"))
	fatalOn(err)

	// Run the node (Start + Wait)
	err = node.Run()
	if err != nil {
		log.Printf("node.Run FAILED: %v", err)
	}

	// node.Stop() // Stop the node now
	node.StopAfter(300) // Stop the node after x seconds

	// can also use use jcli rest to stop the node
	// rsd, err := jcli.RestShutdown(restApiAddress, "")
	// log.Printf("RestShutdown: %s - %v", b2s(rsd), err)

	//////////////////////
	//  jcli rest usage //
	//////////////////////

	waitSec := 5
	wait(waitSec)
	restSettings, err := jcli.RestSettings(restApiAddress, "json")
	log.Printf("RestSettings: %s - %v", b2s(restSettings), err)

	wait(waitSec)
	restNodeStats, err := jcli.RestNodeStats(restApiAddress, "json")
	log.Printf("RestNodeStats: %s - %v", b2s(restNodeStats), err)

	wait(waitSec)
	restTip, err := jcli.RestTip(restApiAddress)
	log.Printf("RestTip: %s - %v", b2s(restTip), err)

	wait(waitSec)
	restAccFc, err := jcli.RestAccount(b2s(faucetAddr), restApiAddress, "json")
	log.Printf("RestAccount Faucet: %s - %v", b2s(restAccFc), err)

	wait(waitSec)
	restAccFx, err := jcli.RestAccount(b2s(fixedAddr), restApiAddress, "json")
	log.Printf("RestAccount Fixed: %s - %v", b2s(restAccFx), err)

	wait(waitSec)
	restLeaders, err := jcli.RestLeaders(restApiAddress, "json")
	log.Printf("RestLeaders: %s - %v", b2s(restLeaders), err)

	wait(waitSec)
	restStakePools, err := jcli.RestStakePools(restApiAddress, "json")
	log.Printf("RestStakePools: %s - %v", b2s(restStakePools), err)

	wait(waitSec)
	restStake, err := jcli.RestStake(restApiAddress, "json")
	log.Printf("RestStake: %s - %v", b2s(restStake), err)

	wait(waitSec)
	restLeadersLogs, err := jcli.RestLeadersLogs(restApiAddress, "json")
	log.Printf("RestLeadersLogs: %s - %v", b2s(restLeadersLogs), err)

	wait(waitSec)
	restMessageLogs, err := jcli.RestMessageLogs(restApiAddress, "json")
	log.Printf("RestMessageLogs: %s - %v", b2s(restMessageLogs), err)

	// Wait for the node to stop.
	log.Println("Waiting...")
	node.Wait() // This blocks here. use "jcli rest v0 shutdown get -h "http://127.0.0.1:8443/api""

	// All done. Node has stopped.
	log.Println("...Done")
}
