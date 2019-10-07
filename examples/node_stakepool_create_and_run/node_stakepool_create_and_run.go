//$(which go) run $0 $@; exit $?

package main

import (
	"encoding/hex"
	"encoding/json"
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

func fatalStop(node *jnode.Jnode, err error, str ...string) {
	if err != nil {
		_ = node.Stop()
		node.Wait()
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

/* seeds used [0-1,10-11] */
func main() {
	var (
		err error

		// Rest
		restAddr       = "127.0.0.1" // rest ip
		restPort       = 8003        // rest port
		restAddress    = restAddr + ":" + strconv.Itoa(restPort)
		restAddressAPI = "http://" + restAddr + ":" + strconv.Itoa(restPort) + "/api" // self
		// Rest trusted nodes
		restLeaderAPI  = "http://127.0.0.1:8001/api" // leader node
		restPassiveAPI = "http://127.0.0.1:8002/api" // passive node

		// P2P
		p2pIPver = "ip4" // ipv4 or ipv6
		p2pProto = "tcp" // tcp

		// P2P Public
		p2pPubAddr       = "127.0.0.1" // PublicAddres
		p2pPubPort       = 9003        // node P2P Public Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPubPort)

		// P2P Listen
		p2pListenAddr    = "127.0.0.1" // ListenAddress
		p2pListenPort    = 9003        // node P2P Public Port
		p2pListenAddress = "/" + p2pIPver + "/" + p2pListenAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pListenPort)

		// General
		discrimination = "testing"  // "" (empty defaults to "production")
		addressPrefix  = "jnode_ta" // "" (empty defaults to "ca")

		// Trusted peers
		trustedPeerLeader  = "/ip4/127.0.0.1/tcp/9001" // Leader (genesis) node
		trustedPeerPassive = "/ip4/127.0.0.1/tcp/9002" // Passive node

		// Genesis Block0 Hash retrieved from example (1)
		block0Hash = "ea7d7d70182c7c9b3820a509d1e87c9a8ec2ad1acaf09645b5c84bed1a938224"
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
	stakePoolManagementThreshold := uint16(len(stakePoolOwners))
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
		1) the StakePool is not yet registered on the network.
		2) Even if it was registered the StakePool has no stake yet.

	*******************************************************************/

	/////////////////////////////
	// STAKE POOL Registration //
	/////////////////////////////

	// give some time for the rest interface to come online
	// FIXME: check if rest server has come online
	log.Println("Waiting for rest interface...")
	time.Sleep(5 * time.Second)

	// FIXME: The correct behaviour is to wait for the node to sync.
	//
	// We could check the tip from some trusted nodes,
	// but only from those having rest available,
	// until a new api is available to check for overall netwok status.
	//
	// In the networked testnet we can only check
	// if the local rest server has come online,
	// since the public rest interfaces are probably disabled.
	//

	// Genesis leader node tip
	leaderTip, err := jcli.RestTip(restLeaderAPI)
	log.Printf("LeaderTip: %s - %v\n", b2s(leaderTip), err)
	fatalStop(node, err)
	time.Sleep(1 * time.Second)

	// passive node tip
	passiveTip, err := jcli.RestTip(restPassiveAPI)
	log.Printf("PassiveTip: %s - %v\n", b2s(passiveTip), err)
	fatalStop(node, err)
	time.Sleep(1 * time.Second)

	// stake pool (self) node tip
	selfTip, err := jcli.RestTip(restAddressAPI)
	log.Printf("SelfTip: %s - %v\n", b2s(selfTip), err)
	fatalStop(node, err)
	time.Sleep(1 * time.Second)

	// Since the pool has 2 owners, lets make both of them pay :)
	//
	// the total ammount to pay for this transaction is 11100, because:
	// LinearFees.Certificate = 10000 (and the tx contains a certificate)
	// LinearFees.Coefficient =    50
	// LinearFees.Constant    =  1000
	// -------------------------------
	// TOTAL (lovelace)       = 11050 + 50 = 11100 (split it in half for each owner)
	//
	// In the testnet you can get those values by querying
	// the settings using Rest API.
	// TODO: restSettings, err := jcli.RestSettings(restAddrAPI, "json")
	//
	var (
		feeCertificate = uint64(10000)
		feeCoefficient = uint64(50)
		feeConstant    = uint64(1000)

		txHalfCost = uint64(5550) // 11100/2

		// spending counter data
		jsonData      map[string]interface{}
		faucetCounter uint32
		fixedCounter  uint32
	)

	//////////////////////////////
	// 1 - Create a transaction //
	//////////////////////////////

	txStaging, err := jcli.TransactionNew(nil, "")
	fatalStop(node, err, "TransactionNew", b2s(txStaging))

	///////////////////////////////////////////
	// 2 - add accounts to transaction input //
	///////////////////////////////////////////

	// 2.a - Add the FAUCET Account address to the transaction
	txStaging, err = jcli.TransactionAddAccount(txStaging, "", b2s(faucetAddr), txHalfCost)
	fatalStop(node, err, "TransactionAddAccount FAUCET", b2s(txStaging))

	// 2.b -  Add the FIXED Account address to the transaction
	txStaging, err = jcli.TransactionAddAccount(txStaging, "", b2s(fixedAddr), txHalfCost)
	fatalStop(node, err, "TransactionAddAccount FIXED", b2s(txStaging))

	////////////////////////////////////////////////
	// 3 - Add the certificate to the transaction //
	////////////////////////////////////////////////

	txStaging, err = jcli.TransactionAddCertificate(txStaging, "", b2s(stakePoolCertSigned))
	fatalStop(node, err, "TransactionAddAccount FIXED", b2s(txStaging))

	// TODO: check if transaction is balanced
	// otherwise finalize will fail

	//////////////////////////////////
	// 4 - Finalize the transaction //
	//////////////////////////////////

	txStaging, err = jcli.TransactionFinalize(txStaging, "", feeCertificate, feeCoefficient, feeConstant, b2s(fixedAddr))
	fatalStop(node, err, "TransactionFinalize", b2s(txStaging))

	////////////////////////////
	// 5 - Make the witnesses //
	////////////////////////////

	// 5.a - Get transaction data for witness (right now the same as TransactionID)
	txDataForWitness, err := jcli.TransactionDataForWitness(txStaging, "")
	fatalStop(node, err, "TransactionDataForWitness", b2s(txDataForWitness))

	// 5.b - FAUCET witness

	// Get faucet account spending counter
	faucetState, err := jcli.RestAccount(b2s(faucetAddr), restAddressAPI, "json")
	fatalStop(node, err, b2s(faucetState))
	err = json.Unmarshal(faucetState, &jsonData)
	fatalStop(node, err)
	jsonCounter, ok := jsonData["counter"].(float64)
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "faucetCounter"))
	}
	faucetCounter = uint32(jsonCounter)

	// save the witness data to this file
	faucetWitnessFile := workingDir + string(os.PathSeparator) + "faucet.witness"
	faucetWitness, err := jcli.TransactionMakeWitness(
		faucetSK,
		b2s(txDataForWitness),
		block0Hash,
		"account", faucetCounter,
		faucetWitnessFile,
		"",
	)
	fatalStop(node, err, "TransactionMakeWitness FAUCET", b2s(faucetWitness))

	// 5.c - FIXED witness

	// Get fixed account spending counter
	fixedState, err := jcli.RestAccount(b2s(fixedAddr), restAddressAPI, "json")
	fatalStop(node, err, b2s(fixedState))
	err = json.Unmarshal(fixedState, &jsonData)
	fatalStop(node, err)
	jsonCounter, ok = jsonData["counter"].(float64)
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "fixedCounter"))
	}
	fixedCounter = uint32(jsonCounter)

	// save the witness data to this file
	fixedWitnessFile := workingDir + string(os.PathSeparator) + "fixed.witness"
	fixedWitness, err := jcli.TransactionMakeWitness(
		fixedSK,
		b2s(txDataForWitness),
		block0Hash,
		"account", fixedCounter,
		fixedWitnessFile,
		"",
	)
	fatalStop(node, err, "TransactionMakeWitness FIXED", b2s(fixedWitness))

	//////////////////////////////////////////////
	// 6 - Add the witnesses to the transaction //
	//////////////////////////////////////////////

	// 6.a - Add FAUCET witness
	txStaging, err = jcli.TransactionAddWitness(txStaging, "", faucetWitnessFile)
	fatalStop(node, err, "TransactionAddWitness FAUCET", b2s(txStaging))

	// 6.b - Add FIXED witness
	txStaging, err = jcli.TransactionAddWitness(txStaging, "", fixedWitnessFile)
	fatalStop(node, err, "TransactionAddWitness FIXED", b2s(txStaging))

	/////////////////////////////
	// 7. Seal the transaction //
	/////////////////////////////

	txStaging, err = jcli.TransactionSeal(txStaging, "")
	fatalStop(node, err, "TransactionSeal", b2s(txStaging))

	///////////////////////////////////////////
	// 8. Convert the transaction to message //
	///////////////////////////////////////////

	txMessage, err := jcli.TransactionToMessage(txStaging, "")
	fatalStop(node, err, "TransactionToMessage", b2s(txMessage))

	///////////////////////////////////////////////
	// 9. Send the transaction to the blockchain //
	///////////////////////////////////////////////

	fragmentID, err := jcli.RestMessagePost(txMessage, restAddressAPI, "")
	fatalStop(node, err, "RestMessagePost", b2s(fragmentID))

	// Display transaction info
	txInfo, err := jcli.TransactionInfo(
		txStaging, "",
		feeCertificate,
		feeCoefficient,
		feeConstant,
		addressPrefix,
		"",
		"",
	)
	log.Printf("TransactionInfo:\n%s : %v\n", b2s(txInfo), err)

	/*****************************************************************

		At this point the StakePool is configured and running,
		but the node behaves like a passive one since:
		1) StakePool is registered on the network, but has no stake yet.

	*******************************************************************/

	///////////////////////////
	// STAKE POOL Delegation //
	///////////////////////////

	log.Printf("Genesis Hash: %s", block0Hash)
	log.Printf("StakePool ID: %s", stakePoolID)
	log.Println("StakePool Node - Running...")
	node.Wait()                             // Wait for the node to stop.
	log.Println("...StakePool Node - Done") // All done. Node has stopped.
}
