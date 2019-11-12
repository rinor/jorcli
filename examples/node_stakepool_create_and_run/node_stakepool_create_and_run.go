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
	"reflect"
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

/* seeds used [1-2,6-7], [60] */
const (
	// genesis accounts data
	faucetSeed = 1 // seed for faucet
	fixedSeed  = 2 // seed for fixed
	// pool secrets seed
	vrfSeed = 6
	kesSeed = 7

	delegatorSeed = 60 // seed for delegator

)

/* seeds used [0-1,10-11] */
func main() {
	var (
		err error

		// Rest
		restAddr       = "127.0.0.33" // rest ip
		restPort       = 8001         // rest port
		restAddress    = restAddr + ":" + strconv.Itoa(restPort)
		restAddressAPI = "http://" + restAddr + ":" + strconv.Itoa(restPort) + "/api" // self

		// P2P
		p2pIPver = "ip4" // ipv4 or ipv6
		p2pProto = "tcp" // tcp

		// P2P Public
		p2pPubAddr       = "127.0.0.33" // PublicAddres
		p2pPubPort       = 9001         // node P2P Public Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPubPort)

		// P2P Listen
		p2pListenAddr    = "127.0.0.33" // ListenAddress
		p2pListenPort    = 9001         // node P2P Public Port
		p2pListenAddress = "/" + p2pIPver + "/" + p2pListenAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pListenPort)

		// General
		discrimination = "testing" // "" (empty defaults to "production")
		addressPrefix  = ""        // "" (empty defaults to "ca")

		// Trusted peers
		/*
			leaderAddr = "/ip4/127.0.0.11/tcp/9001"                         // Leader (genesis) node (example 1)
			leaderID   = "000000000000000000000000000000000000000000000030" // Leader public_id
		*/
		gepAddr = "/ip4/127.0.0.22/tcp/9001"                         // Genesis stake pool node (example 2)
		gepID   = "222222222222222222222222222222222222222222222222" // Genesis stake pool public_id

		// Genesis Block0 Hash retrieved from example (1)
		block0Hash = "4fc09867e8e377811b86f5b940eef8dfe3e7388bbb9dacce0cb95d5105de6562"

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
	workingDir, err := ioutil.TempDir("", "jnode_dynstake_")
	fatalOn(err, "workingDir")
	log.Println()
	log.Printf("Working Directory: %s", workingDir)
	log.Println()

	///////////////
	// DELEGATOR //
	///////////////

	// will need this one file later for delegation certificate `transaction auth`
	// since delegator is not the pool owner
	delegatorFileSK := workingDir + string(os.PathSeparator) + "delegator_key.sk"

	delegatorSK, err := jcli.KeyGenerate(seed(delegatorSeed), "Ed25519Extended", delegatorFileSK)
	fatalOn(err, b2s(delegatorSK))
	delegatorPK, err := jcli.KeyToPublic(delegatorSK, "", "")
	fatalOn(err, b2s(delegatorPK))
	delegatorAddr, err := jcli.AddressAccount(b2s(delegatorPK), addressPrefix, discrimination)
	fatalOn(err, b2s(delegatorAddr))

	////////////
	// FAUCET //
	////////////

	// will need this one file later for pool certificate `transaction auth`
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

	// will need this one file later for pool certificate `transaction auth`
	fixedFileSK := workingDir + string(os.PathSeparator) + "fixed_key.sk"

	fixedSK, err := jcli.KeyGenerate(seed(fixedSeed), "Ed25519Extended", fixedFileSK)
	fatalOn(err, b2s(fixedSK))
	fixedPK, err := jcli.KeyToPublic(fixedSK, "", "")
	fatalOn(err, b2s(fixedPK))
	fixedAddr, err := jcli.AddressAccount(b2s(fixedPK), addressPrefix, discrimination)
	fatalOn(err, b2s(fixedAddr))

	/////////////////////////
	// STAKE POOL Creation //
	/////////////////////////

	// VRF
	poolVrfSK, err := jcli.KeyGenerate(seed(vrfSeed), "Curve25519_2HashDH", "")
	fatalOn(err, b2s(poolVrfSK))
	poolVrfPK, err := jcli.KeyToPublic(poolVrfSK, "", "")
	fatalOn(err, b2s(poolVrfPK))

	// KES
	poolKesSK, err := jcli.KeyGenerate(seed(kesSeed), "SumEd25519_12", "")
	fatalOn(err, b2s(poolKesSK))
	poolKesPK, err := jcli.KeyToPublic(poolKesSK, "", "")
	fatalOn(err, b2s(poolKesPK))

	// note we will use the Faucet and Fixed as owners of this pool
	stakePoolOwners := []string{
		b2s(faucetPK),
		b2s(fixedPK),
	}
	stakePoolManagementThreshold := uint8(len(stakePoolOwners))
	stakePoolSerial := uint64(3030303030)
	stakePoolStartValidity := uint64(0)

	stakePoolCert, err := jcli.CertificateNewStakePoolRegistration(
		b2s(poolKesPK),
		b2s(poolVrfPK),
		stakePoolStartValidity,
		stakePoolManagementThreshold,
		stakePoolSerial,
		stakePoolOwners,
		nil,
		"",
	)
	fatalOn(err, b2s(stakePoolCert))

	// We can get poolID from signed or unsigned certificate
	stakePoolID, err := jcli.CertificateGetStakePoolID(stakePoolCert, "", "")
	fatalOn(err, b2s(stakePoolID))

	// We do not need to sign the certificate anymore since:
	//
	// certificate sign - is now only used for the block0 creation certificates,
	// as for certificates using the transaction system, we now need - `transaction auth`.
	//
	// jcli transaction auth - used after sealing
	// similar to `jcli certificate sign` but apply to an almost built transaction.
	// Needed for:
	//     StakeDelegation      - YES (jcli certificate new stake-delegation)
	//     PoolRegistration     - YES (jcli certificate new stake-pool-registration)
	//     PoolRetirement       - YES (N/A)
	//     PoolUpdate           - YES (N/A)
	//     OwnerStakeDelegation - NO  (N/A - jcli certificate new stake-delegation)
	//
	/*
		// Sign the certificate with FAUCET private key and also with FIXED private key
		stakePoolCertSigned, err := jcli.CertificateSign(stakePoolCert, []string{faucetFileSK, fixedFileSK}, "", "")
		fatalOn(err, b2s(stakePoolCertSigned))
	*/

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

	// p2p node public_id
	nodePublicID := "333333333333333333333333333333333333333333333333"

	nodeCfg := jnode.NewNodeConfig()

	nodeCfg.Storage = "jnode_storage"

	nodeCfg.Rest.Enabled = true       // default is "false" (rest disabled)
	nodeCfg.Rest.Listen = restAddress // 127.0.0.1:8443 is default value

	nodeCfg.Explorer.Enabled = false // default is "false" (explorer disabled)

	nodeCfg.P2P.PublicAddress = p2pPublicAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.ListenAddress = p2pListenAddress // /ip4/127.0.0.1/tcp/8299 is default value
	nodeCfg.P2P.PublicID = nodePublicID          // jörmungandr will generate a random key, if not set
	nodeCfg.P2P.AllowPrivateAddresses = true     // for private addresses
	nodeCfg.P2P.Policy.QuarantineDuration = "5m" // default to "30m"

	// add trusted peer to config file
	// nodeCfg.AddTrustedPeer(leaderAddr, leaderID)
	nodeCfg.AddTrustedPeer(gepAddr, gepID)

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
	// node.AddTrustedPeer(leaderAddr, leaderID) // add leader from example (1) as trusted
	node.AddTrustedPeer(gepAddr, gepID) // add genesis stake pool from example (2) as trusted

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

	/*****************************************************************

		At this point the StakePool is configured and running,
		but the node behaves like a passive one since:
		1) the StakePool is not yet registered on the network.
		2) Even if it was registered the StakePool has no stake yet.

	*******************************************************************/

	/////////////////////////////
	// STAKE POOL Registration //
	/////////////////////////////

	var (
		// generic interface used for json data.
		jsonData map[string]interface{}

		// spending counter data
		faucetCounter uint32
		fixedCounter  uint32

		// fees data
		feeCertificate uint64
		feeCoefficient uint64
		feeConstant    uint64

		nodeReady bool
		nodeState string
	)

	// give some time for the rest interface to come online
	log.Println("Waiting for rest interface...")
	stateCheckFreq := 1 // seconds
	for !nodeReady {
		time.Sleep(time.Duration(stateCheckFreq) * time.Second)

		nodeStats, err := jcli.RestNodeStats(restAddressAPI, "json")
		fatalStop(node, err, b2s(nodeStats))

		err = json.Unmarshal(nodeStats, &jsonData)
		fatalStop(node, err)

		// StartingRestServer, PreparingStorage, PreparingBlock0, Bootstrapping, StartingWorkers, Running
		// {
		//   "state": "PreparingBlock0"
		// }

		ok := false // one of those things (prevent shadowing of "nodeState")
		nodeState, ok = jsonData["state"].(string)
		if !ok {
			fatalStop(node, fmt.Errorf("%s - NOT FOUND", "jsonNodeState"))
		}

		// unnecessary, but keep it
		switch nodeState {
		case "StartingRestServer":
			stateCheckFreq = 1
		case "PreparingStorage":
			stateCheckFreq = 1
		case "PreparingBlock0":
			stateCheckFreq = 1
		case "Bootstrapping":
			stateCheckFreq = 5
		case "StartingWorkers":
			stateCheckFreq = 1
		case "Running":
			nodeReady = !nodeReady
		}
	}
	log.Printf("...Node state [%s]\n", nodeState)

	// stake pool (self) node tip
	selfTip, err := jcli.RestTip(restAddressAPI)
	fatalStop(node, err, b2s(selfTip))
	log.Printf("SelfTip: %s\n", b2s(selfTip))

	// Since the pool has 2 owners, lets make both of them pay :)
	//
	// the total ammount to pay for this transaction is 11100, because:
	// total fees: constant + (num_inputs + num_outputs) * coefficient [+ certificate]
	//
	// LinearFees.Certificate = 10_000 (and the tx contains a certificate)
	// LinearFees.Coefficient =     50 (we have 2 inputs so total is 100)
	// LinearFees.Constant    =  1_000
	// -------------------------------
	// TOTAL (lovelace)       =  1_000 + (2 + 0)*50 + 10_000 = 11_100

	// get blockchain setting from rest
	blockchainSettings, err := jcli.RestSettings(restAddressAPI, "json")
	fatalStop(node, err, "RestSettings", b2s(blockchainSettings))

	// will use a generic way to parse json data without using structs,
	// just to show how is done.
	// One can also build a struct representing the settings or
	// only the fees using jnode.LinearFees for example.

	err = json.Unmarshal(blockchainSettings, &jsonData)
	fatalStop(node, err)

	// fees
	jsonFees, ok := jsonData["fees"].(map[string]interface{})
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "jsonFees"))
	}

	// fee certificate
	jsonFeeCertificate, ok := jsonFees["certificate"].(float64)
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "jsonFeeCertificate"))
	}
	feeCertificate = uint64(jsonFeeCertificate)

	// fee coefficient
	jsonFeeCoefficient, ok := jsonFees["coefficient"].(float64)
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "jsonFeeCoefficient"))
	}
	feeCoefficient = uint64(jsonFeeCoefficient)

	// fee constant
	jsonFeeConstant, ok := jsonFees["constant"].(float64)
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "jsonFeeConstant"))
	}
	feeConstant = uint64(jsonFeeConstant)

	// we will have 2 inputs and 0 outputs.
	// total fees: constant + (num_inputs + num_outputs) * coefficient + certificate
	totalFees := feeConstant + 2*feeCoefficient + feeCertificate

	//////////////////////////////
	// 1 - Create a transaction //
	//////////////////////////////

	txStaging, err := jcli.TransactionNew(nil, "")
	fatalStop(node, err, "TransactionNew", b2s(txStaging))

	///////////////////////////////////////////
	// 2 - add accounts to transaction input //
	///////////////////////////////////////////

	// 2.a - Add the FAUCET Account address to the transaction
	txStaging, err = jcli.TransactionAddAccount(txStaging, "", b2s(faucetAddr), totalFees/2)
	fatalStop(node, err, "TransactionAddAccount FAUCET", b2s(txStaging))

	// 2.b -  Add the FIXED Account address to the transaction
	txStaging, err = jcli.TransactionAddAccount(txStaging, "", b2s(fixedAddr), totalFees-(totalFees/2))
	fatalStop(node, err, "TransactionAddAccount FIXED", b2s(txStaging))

	////////////////////////////////////////////////
	// 3 - Add the certificate to the transaction //
	////////////////////////////////////////////////

	// This needs an unsigned certificate now, otherwise will fail
	txStaging, err = jcli.TransactionAddCertificate(txStaging, "", b2s(stakePoolCert))
	fatalStop(node, err, "TransactionAddCertificate", b2s(txStaging))

	// Check if transaction is balanced otherwise
	// finalize will fail (if balance != 0)

	// get balance value from transaction info
	txInfo, err := jcli.TransactionInfo(
		txStaging, "",
		feeCertificate,
		feeCoefficient,
		feeConstant,
		addressPrefix,
		"",
		"{balance}",
		"",
		"",
		"",
		false,
		false,
		false,
	)
	fatalStop(node, err, "TransactionInfo BALANCE", b2s(txInfo))

	// get string value.
	// when math ops required, convert it to number.
	txBalance := b2s(txInfo)

	// jic, since shouldn't happen (unless jcli cmd changed)
	if txBalance == "" {
		fatalStop(node, fmt.Errorf("TransactionInfo, BALANCE has no data [balance=%s]", txBalance))
	}

	// BUG: if balance outside (int) range ...
	txBalanceAmmount, err := strconv.Atoi(txBalance)
	fatalStop(node, err, "strconv.Atoi(txBalance)", txBalance)
	if txBalanceAmmount != 0 {
		fatalStop(node, fmt.Errorf("TransactionInfo, NOT BALANCED [balance=%s], Finalize will fail", txBalance))
	}

	//////////////////////////////
	// 4 - Lock the transaction //
	//////////////////////////////

	txStaging, err = jcli.TransactionFinalize(txStaging, "", feeCertificate, feeCoefficient, feeConstant, b2s(fixedAddr))
	fatalStop(node, err, "TransactionFinalize", b2s(txStaging))

	////////////////////////////
	// 5 - Make the witnesses //
	////////////////////////////

	// 5.a - Get transaction data for witness
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

	//////////////////////////////
	// 7 - Seal the transaction //
	//////////////////////////////

	txStaging, err = jcli.TransactionSeal(txStaging, "")
	fatalStop(node, err, "TransactionSeal", b2s(txStaging))

	//////////////////////////////
	// 8 - Auth the transaction //
	//////////////////////////////

	// Since the transaction contains a stake pool registration certificate
	// we need to auth it with owners secret keys FAUCET and FIXED
	txStaging, err = jcli.TransactionAuth(txStaging, "", []string{faucetFileSK, fixedFileSK})
	fatalStop(node, err, "TransactionAuth", b2s(txStaging))

	////////////////////////////////////////////
	// 9 - Convert the transaction to message //
	////////////////////////////////////////////

	txMessage, err := jcli.TransactionToMessage(txStaging, "")
	fatalStop(node, err, "TransactionToMessage", b2s(txMessage))

	/////////////////////////////////////////////////
	// 10 - Send the transaction to the blockchain //
	/////////////////////////////////////////////////

	fragmentID, err := jcli.RestMessagePost(txMessage, restAddressAPI, "")
	fatalStop(node, err, "RestMessagePost", b2s(fragmentID))

	///////////////////////////////////////////////
	// 11 - Check certificate transaction status //
	///////////////////////////////////////////////

	// There are different ways to set a checkpoint when to check for status changes.
	// One could check for tip changes and then check for status changes,
	// or use a loop with timeout based on slot_duration.
	// In this example a loop is used.
	//
	// If the node has explorer enabled, one could also use graphql queries to get the status.
	// TODO: implement this example.
	//
	// NOTE:
	// - the tip can change also during sync if the node is behind,
	//   so it does not guaranties the transaction inclusion
	//
	// - slot leader election in genesis_praos is unpredictable

	var (
		logFragmentID  = b2s(fragmentID)
		fragmentStatus string
		fragmentInfo   string

		// generic interface used for json data.
		jsonDataList []map[string]interface{}
	)

	log.Printf("Wait for pool registration certificate transaction [%s] status change...\n", logFragmentID)

	// 150 derived from slots_per_epoch
	for x, done := 0, false; !done && x < 150; x++ {
		fragmentLogs, err := jcli.RestMessageLogs(restAddressAPI, "json")
		fatalStop(node, err, "RestMessageLogs", b2s(fragmentLogs))

		err = json.Unmarshal(fragmentLogs, &jsonDataList)
		fatalStop(node, err)

		// This may be resourse intensive depending on the number of message logs.
		for i := range jsonDataList {
			logID, ok := jsonDataList[i]["fragment_id"].(string)
			if !ok {
				fatalStop(node, fmt.Errorf("%s - NOT FOUND", "fragment_id"))
			}

			// we are interested in a specific fragment_id
			if logFragmentID != logID {
				continue
			}

			status, ok := jsonDataList[i]["status"]
			if !ok {
				fatalStop(node, fmt.Errorf("%s - NOT FOUND", "status"))
			}

			switch reflect.TypeOf(status).Kind() {

			case reflect.String:
				/**************************************
				   "status": "Pending"
				**************************************/
				fragmentStatus = status.(string)

			case reflect.Map:
				/**************************************
				   "status": {
				     "InABlock": {
					   "block": "3aa748dd766dbe0bc1c77ab5200c85cda8743464ba07fdf98944bfa63339e571",
				       "date": "114237.32"
				     }
				   }
				**************************************/
				blockDate, accepted := status.(map[string]interface{})["InABlock"]
				if accepted {
					fragmentStatus = "InABlock"
					info, ok := blockDate.(map[string]interface{})["date"]
					if ok {
						fragmentInfo = info.(string)
					}
					info, ok = blockDate.(map[string]interface{})["block"]
					if ok {
						fragmentInfo = fragmentInfo + " (" + info.(string) + ")"
					}
					done = true
					break
				}

				/**************************************
				   "status": {
				     "Rejected": {
				       "reason": "some reason info"
				     }
				   }
				**************************************/
				reason, rejected := status.(map[string]interface{})["Rejected"]
				if rejected {
					fragmentStatus = "Rejected"
					info, ok := reason.(map[string]interface{})["reason"]
					if ok {
						fragmentInfo = info.(string)
					}
					done = true
					// break
				}
			} /* switch */

		} /* for */

		if !done {
			time.Sleep(2 * time.Second) // 2 is derived from slot_duration
		}

	} /* for */

	log.Printf("FragmentID: %s - %s [%s]\n", logFragmentID, fragmentStatus, fragmentInfo)
	switch fragmentStatus {
	case "":
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", logFragmentID))
	case "Pending":
		fatalStop(node, fmt.Errorf("%s - %s", logFragmentID, fragmentStatus))
	case "Rejected":
		fatalStop(node, fmt.Errorf("%s - %s [%s]", logFragmentID, fragmentStatus, fragmentInfo))
	case "InABlock":
		// transaction included in a block
	default:
		fatalStop(node, fmt.Errorf("unknown status for %s - %s [%s]", logFragmentID, fragmentStatus, fragmentInfo))
	}

	/////////////////////////////////////////
	// 12 - Check the stake pool is listed //
	/////////////////////////////////////////

	var (
		// generic interface used for json data.
		jsonStakePoolList []string

		ledgerPoolID   = b2s(stakePoolID)
		stakePoolFound = false
	)
	jsonStakePools, err := jcli.RestStakePools(restAddressAPI, "json")
	fatalStop(node, err, "RestStakePools", b2s(jsonStakePools))

	err = json.Unmarshal(jsonStakePools, &jsonStakePoolList)
	fatalStop(node, err)

	for i := range jsonStakePoolList {
		// we are interested in a specific poolID
		if ledgerPoolID != jsonStakePoolList[i] {
			continue
		}
		stakePoolFound = true
		break
	}

	if !stakePoolFound {
		fatalStop(node, fmt.Errorf("StakePool %s - Not found", ledgerPoolID))
	}

	/*****************************************************************

		At this point the StakePool is configured and running,
		but the node behaves like a passive one since:
		1) StakePool is registered on the network,
		   but has NO STAKE yet.

	*******************************************************************/

	///////////////////////////
	// STAKE POOL Delegation //
	///////////////////////////

	// DELEGATOR account will stake to this new pool
	// FAUCET and FIXED are the owners of this pool
	// so we still need to auth the transaction.

	/////////////////////////////////////////////////
	// 1 - Create the stake delegation certificate //
	/////////////////////////////////////////////////

	delegationCert, err := jcli.CertificateNewStakeDelegation(b2s(delegatorPK), []string{b2s(stakePoolID)}, "")
	fatalStop(node, err, "CertificateNewStakeDelegation", b2s(delegationCert))

	// This is transaction based certificate, no need to sign
	/*
		///////////////////////////////////////////////
		// 2 - Sign the stake delegation certificate //
		///////////////////////////////////////////////

		delegationCertSigned, err := jcli.CertificateSign(delegationCert, []string{delegatorFileSK}, "", "")
		fatalStop(node, err, "CertificateSign", b2s(delegationCertSigned))
	*/

	///////////////////////////////////////
	// 2 - Create a transaction and send //
	///////////////////////////////////////

	// The flow is almost the same as the
	// transaction flow in Stake Pool Registration
	// so the comments will be less verbose

	var (
		// we will have 1 input and 0 outputs.
		// total fees: constant + (num_inputs + num_outputs) * coefficient + certificate
		totalDelegationFees = feeConstant + 1*feeCoefficient + feeCertificate
	)

	dtxStaging, err := jcli.TransactionNew(nil, "")
	fatalStop(node, err, "TransactionNew", b2s(dtxStaging))

	dtxStaging, err = jcli.TransactionAddAccount(dtxStaging, "", b2s(delegatorAddr), totalDelegationFees)
	fatalStop(node, err, "TransactionAddAccount FIXED", b2s(dtxStaging))

	dtxStaging, err = jcli.TransactionAddCertificate(dtxStaging, "", b2s(delegationCert))
	fatalStop(node, err, "TransactionAddAccount FIXED", b2s(dtxStaging))

	dtxInfo, err := jcli.TransactionInfo(
		dtxStaging, "",
		feeCertificate,
		feeCoefficient,
		feeConstant,
		addressPrefix,
		"",
		"{balance}",
		"",
		"",
		"",
		false,
		false,
		false,
	)
	fatalStop(node, err, "TransactionInfo BALANCE", b2s(dtxInfo))

	dtxBalance := b2s(dtxInfo)
	if dtxBalance == "" {
		fatalStop(node, fmt.Errorf("TransactionInfo, BALANCE has no data [balance=%s]", dtxBalance))
	}

	dtxBalanceAmmount, err := strconv.Atoi(dtxBalance)
	fatalStop(node, err, "strconv.Atoi(dtxBalance)", dtxBalance)
	if dtxBalanceAmmount != 0 {
		fatalStop(node, fmt.Errorf("TransactionInfo, NOT BALANCED [balance=%s], Finalize will fail", dtxBalance))
	}

	dtxStaging, err = jcli.TransactionFinalize(dtxStaging, "", feeCertificate, feeCoefficient, feeConstant, b2s(delegatorAddr))
	fatalStop(node, err, "TransactionFinalize", b2s(dtxStaging))

	dtxDataForWitness, err := jcli.TransactionDataForWitness(dtxStaging, "")
	fatalStop(node, err, "TransactionDataForWitness", b2s(dtxDataForWitness))

	delegatorState, err := jcli.RestAccount(b2s(delegatorAddr), restAddressAPI, "json")
	fatalStop(node, err, b2s(delegatorState))

	err = json.Unmarshal(delegatorState, &jsonData)
	fatalStop(node, err)

	jsonCounter, ok = jsonData["counter"].(float64)
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "delegatorCounter"))
	}

	delegatorCounter := uint32(jsonCounter)

	// save the witness data to this file
	delegatorWitnessFile := workingDir + string(os.PathSeparator) + "delegator.witness"
	dfixedWitness, err := jcli.TransactionMakeWitness(
		delegatorSK,
		b2s(dtxDataForWitness),
		block0Hash,
		"account", delegatorCounter,
		delegatorWitnessFile,
		"",
	)
	fatalStop(node, err, "TransactionMakeWitness FIXED", b2s(dfixedWitness))

	dtxStaging, err = jcli.TransactionAddWitness(dtxStaging, "", delegatorWitnessFile)
	fatalStop(node, err, "TransactionAddWitness FIXED", b2s(dtxStaging))

	dtxStaging, err = jcli.TransactionSeal(dtxStaging, "")
	fatalStop(node, err, "TransactionSeal", b2s(dtxStaging))

	// The transaction contains a stake pool delegation certificate,
	// but since the delegator is NOT one of the pools owners
	// we need to auth it with delegator secret keys.
	dtxStaging, err = jcli.TransactionAuth(dtxStaging, "", []string{delegatorFileSK})
	fatalStop(node, err, "TransactionAuth", b2s(dtxStaging))

	dtxMessage, err := jcli.TransactionToMessage(dtxStaging, "")
	fatalStop(node, err, "TransactionToMessage", b2s(dtxMessage))

	dfragmentID, err := jcli.RestMessagePost(dtxMessage, restAddressAPI, "")
	fatalStop(node, err, "RestMessagePost", b2s(dfragmentID))

	var (
		dlogFragmentID  = b2s(dfragmentID)
		dfragmentStatus string
		dfragmentInfo   string
	)

	log.Printf("Wait for delegation certificate transaction [%s] status change...\n", dlogFragmentID)

	// 150 derived from slots_per_epoch
	for x, done := 0, false; !done && x < 150; x++ {
		fragmentLogs, err := jcli.RestMessageLogs(restAddressAPI, "json")
		fatalStop(node, err, "RestMessageLogs", b2s(fragmentLogs))

		err = json.Unmarshal(fragmentLogs, &jsonDataList)
		fatalStop(node, err)

		// This may be resourse intensive depending on the number of message logs.
		for i := range jsonDataList {
			logID, ok := jsonDataList[i]["fragment_id"].(string)
			if !ok {
				fatalStop(node, fmt.Errorf("%s - NOT FOUND", "fragment_id"))
			}

			// we are interested in a specific fragment_id
			if dlogFragmentID != logID {
				continue
			}

			status, ok := jsonDataList[i]["status"]
			if !ok {
				fatalStop(node, fmt.Errorf("%s - NOT FOUND", "status"))
			}

			switch reflect.TypeOf(status).Kind() {

			case reflect.String:
				dfragmentStatus = status.(string)
			case reflect.Map:
				blockDate, accepted := status.(map[string]interface{})["InABlock"]
				if accepted {
					dfragmentStatus = "InABlock"
					info, ok := blockDate.(map[string]interface{})["date"]
					if ok {
						dfragmentInfo = info.(string)
					}
					info, ok = blockDate.(map[string]interface{})["block"]
					if ok {
						dfragmentInfo = dfragmentInfo + " (" + info.(string) + ")"
					}
					done = true
					break
				}
				reason, rejected := status.(map[string]interface{})["Rejected"]
				if rejected {
					dfragmentStatus = "Rejected"
					info, ok := reason.(map[string]interface{})["reason"]
					if ok {
						dfragmentInfo = info.(string)
					}
					done = true
					// break
				}
			} /* switch */

		} /* for */

		if !done {
			time.Sleep(2 * time.Second) // 2 is derived from slot_duration
		}

	} /* for */

	log.Printf("FragmentID: %s - %s [%s]\n", dlogFragmentID, dfragmentStatus, dfragmentInfo)
	switch dfragmentStatus {
	case "":
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", dlogFragmentID))
	case "Pending":
		fatalStop(node, fmt.Errorf("%s - %s", dlogFragmentID, dfragmentStatus))
	case "Rejected":
		fatalStop(node, fmt.Errorf("%s - %s [%s]", dlogFragmentID, dfragmentStatus, dfragmentInfo))
	case "InABlock":
		// transaction included in a block
	default:
		fatalStop(node, fmt.Errorf("unknown status for %s - %s [%s]", dlogFragmentID, dfragmentStatus, dfragmentInfo))
	}

	/////////////////////////////////
	// Check the delegation status //
	/////////////////////////////////

	jsonDelegatorData, err := jcli.RestAccount(b2s(delegatorAddr), restAddressAPI, "json")
	fatalStop(node, err, "RestAccount", b2s(jsonDelegatorData))

	err = json.Unmarshal(jsonDelegatorData, &jsonData)
	fatalStop(node, err)

	accDelegation, ok := jsonData["delegation"].(map[string]interface{})
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "delegation"))
	}

	accPools, ok := accDelegation["pools"].([]interface{})
	if !ok {
		fatalStop(node, fmt.Errorf("%s - NOT FOUND", "pools"))
	}

	stakeDelegationFound := false
	for i := range accPools {
		for _, poolData := range accPools[i].([]interface{}) {
			poolID, ok := poolData.(string)
			if !ok {
				// probably this is the counter (float64)
				continue
			}
			if poolID != b2s(stakePoolID) {
				continue
			}
			stakeDelegationFound = true
			break
		}
		if stakeDelegationFound {
			break
		}
	}

	if !stakeDelegationFound {
		fatalStop(node, fmt.Errorf("%s delegation - NOT FOUND", stakePoolID))
	}

	/*****************************************************************

		At this point the StakePool is ready and
		in the next 2 epochs will become slot leader contendent

	*******************************************************************/

	log.Println()
	log.Printf("Genesis Hash: %s", block0Hash)
	log.Println()
	log.Printf("StakePool ID       : %s", stakePoolID)
	log.Printf("StakePool Owner    : %s", faucetAddr)
	log.Printf("StakePool Owner    : %s", fixedAddr)
	log.Printf("StakePool Delegator: %s", delegatorAddr)
	log.Println()
	log.Printf("NodePublicID for trusted: %s", nodePublicID)
	log.Println()

	log.Println("Delegator StakePool Node - Running...")
	node.Wait()                                       // Wait for the node to stop.
	log.Println("...Delegator StakePool Node - Done") // All done. Node has stopped.
}
