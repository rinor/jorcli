//$(which go) run $0 $@; exit $?

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/rinor/jorcli/jcli"
)

type extendedKey struct {
	Discrimination string `json:"Discrimination,omitempty"`
	AddressPrefix  string `json:"AddressPrefix,omitempty"`

	KeySeed string `json:"KeySeed,omitempty"`
	KeyType string `json:"KeyType"`

	PrivateKey string `json:"PrivateKey"`
	PublicKey  string `json:"PublicKey"`

	Account string `json:"Account,omitempty"`
	UTxO    string `json:"UTxO,omitempty"`
}

func newKey(seed string, keyType string) (*extendedKey, error) {
	var (
		err error
		sk  []byte
		pk  []byte
	)
	// private key
	sk, err = jcli.KeyGenerate(seed, keyType, "")
	if err != nil {
		return nil, fmt.Errorf("KeyGenerate: %s - %s", err, sk)
	}
	// public key
	pk, err = jcli.KeyToPublic(sk, "", "")
	if err != nil {
		return nil, fmt.Errorf("KeyToPublic: %s - %s", err, pk)
	}

	return &extendedKey{
		KeySeed:    seed,
		KeyType:    keyType,
		PrivateKey: b2s(sk),
		PublicKey:  b2s(pk),
	}, nil
}

func (ea *extendedKey) buildAccount(prefix string, discrimination string) error {
	var (
		err error
		acc []byte
	)
	// account address
	acc, err = jcli.AddressAccount(ea.PublicKey, prefix, discrimination)
	if err != nil {
		return fmt.Errorf("AddressAccount: %s - %s", err, acc)
	}
	ea.Account = b2s(acc)
	ea.Discrimination = discrimination
	ea.AddressPrefix = prefix

	return nil
}

/*
{
  "counter": 1,
  "delegation": {
    "pools": [
      [
        "20ba61c4d0b044962ada2536ab703eeaf95ddf7b90d9900a737988b80abb9415",
        1
      ]
    ]
  },
  "value": 9999999999988950
}
*/
type accountRest struct {
	// Number of transactions performed with this account
	Counter uint32 `json:"counter"`
	// Current balance of this account
	Value uint64 `json:"value"`
	// Hex-encoded stake pool ID this account is delegating to
	Delegation struct {
		Pools []interface{} `json:"pools,omitempty"`
	} `json:"delegation,omitempty"`
}

/*
{
  "block0Hash": "999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d",
  "block0Time": "2017-09-29T00:00:00+00:00",
  "consensusVersion": "genesis",
  "currSlotStartTime": "2019-10-15T13:00:52+00:00",
  "fees": {
    "certificate": 10000,
    "coefficient": 50,
    "constant": 1000
  },
  "maxTxsPerBlock": 255,
  "slotDuration": 2,
  "slotsPerEpoch": 150
}
*/
type settingsRest struct {
	// Hex-encoded hash of block0
	Block0Hash string `json:"block0Hash"`
	// When block0 was created
	Block0Time time.Time `json:"block0Time"`
	// Version of consensus, which is currently used
	ConsensusVersion string `json:"consensusVersion"`
	// When current slot was opened, not set if none is currently open
	CurrSlotStartTime time.Time `json:"currSlotStartTime,omitempty"`
	Fees              struct {
		// Fee per certificate used in witness
		Certificate uint64 `json:"certificate"`
		// Fee per every input and output of transaction
		Coefficient uint64 `json:"coefficient"`
		// Base fee per transaction
		Constant uint64 `json:"constant"`
	} `json:"fees"`
	// Maximum number of transactions in block
	MaxTxsPerBlock uint32 `json:"maxTxsPerBlock"`
	// Slot duration in seconds
	SlotDuration uint8 `json:"slotDuration"`
	// number of slots per epoch
	SlotsPerEpoch uint32 `json:"slotsPerEpoch"`
}

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

/* seeds used [1-2], [50] ,[60], [100-199] */
const (
	// faucetSeed    = 1  // seed for faucet
	// fixedSeed     = 2  // seed for fixed
	// gepSeed       = 50 // seed the owner of an extra pool in genesis block
	delegatorSeed = 60 // seed for new stake delegator example (3)

	seedStartBulk  = 100 // seed key generation start
	totSrcAddrBulk = 100 // total number of account addresses

	pathSep = string(os.PathSeparator)
)

func main() {
	var (
		err     error
		jcliBin string
	)
	// Check for jcli binary. PATH or Local folder
	_, err = exec.LookPath("jcli")
	if err != nil {
		if jcliBin, err = exec.LookPath("." + pathSep + "jcli"); err != nil {
			fatalOn(fmt.Errorf("%s - Not Found", "jcli"))
		}
		jcli.BinName(jcliBin) // use local folder binary
	}

	var (
		restAddrAPI = "http://127.0.0.11:8001/api" // genesis node

		discrimination = "testing"  // "" (empty defaults to "production")
		addressPrefix  = "jnode_ta" // "" (empty defaults to "ca")

		keyType = "Ed25519Extended"
	)

	// rebuild DELEGATOR data, since will send lovelaces to  that address
	delegator, err := newKey(seed(delegatorSeed), keyType)
	fatalOn(err)
	err = delegator.buildAccount(addressPrefix, discrimination)
	fatalOn(err)
	err = writeJsonFile(delegator, "delegator_"+delegator.Account+".json")
	fatalOn(err)

	// get blockchain settings
	restSettings, err := jcli.RestSettings(restAddrAPI, "json")
	fatalOn(err, b2s(restSettings))
	settings := settingsRest{}
	err = json.Unmarshal(restSettings, &settings)
	fatalOn(err)

	var (
		//             1000     + (      1    +        1   ) * 50          [+ 10000      ]
		// total fees: constant + (num_inputs + num_outputs) * coefficient [+ certificate]
		feesTotal = settings.Fees.Constant + 2*settings.Fees.Coefficient // 1100
		ammount   = uint64(8900)
	)

	type bulk struct {
		faucet   *extendedKey
		message  []byte
		fragment []byte
	}
	var bulkData [totSrcAddrBulk]bulk

	// rebuild bulk addresses (already present in genesis block)
	for i := 0; i < totSrcAddrBulk; i++ {
		keys, err := newKey(seed(seedStartBulk+i), keyType)
		fatalOn(err)
		err = keys.buildAccount(addressPrefix, discrimination)
		fatalOn(err)
		err = writeJsonFile(keys, keys.Account+".json")
		fatalOn(err)

		bulkData[i].faucet = keys
	}

	// get spending counter and balance of each account address
	// and prepare the transaction messages, but do not send them yet
	for i := range bulkData {
		srcAccount := bulkData[i].faucet.Account
		srcPrivateKey := bulkData[i].faucet.PrivateKey

		restAccFc, err := jcli.RestAccount(srcAccount, restAddrAPI, "json")
		fatalOn(err, b2s(restAccFc))
		newAccountRest := accountRest{}
		err = json.Unmarshal(restAccFc, &newAccountRest)
		fatalOn(err)
		log.Printf("%s - %d - %d\n", srcAccount, newAccountRest.Counter, newAccountRest.Value)

		// prepare transaction
		txStaging, err := jcli.TransactionNew(nil, "")
		fatalOn(err, b2s(txStaging))

		txStaging, err = jcli.TransactionAddAccount(txStaging, "", srcAccount, ammount+feesTotal)
		fatalOn(err, b2s(txStaging))

		txStaging, err = jcli.TransactionAddOutput(txStaging, "", delegator.Account, ammount)
		fatalOn(err, b2s(txStaging))

		// get balance value from transaction info
		txInfo, err := jcli.TransactionInfo(
			txStaging, "",
			settings.Fees.Certificate,
			settings.Fees.Coefficient,
			settings.Fees.Constant,
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
		fatalOn(err, b2s(txInfo))
		txBalance := b2s(txInfo)
		if txBalance == "" {
			fatalOn(fmt.Errorf("TransactionInfo, BALANCE has no data [balance=%s]", txBalance))
		}
		txBalanceAmmount, err := strconv.Atoi(txBalance) // BUG: if balance outside (int) range ...
		fatalOn(err, txBalance)
		switch {
		case txBalanceAmmount < 0:
			fatalOn(fmt.Errorf("TransactionInfo, NOT BALANCED [balance=%s], Finalize will fail", txBalance))
		case txBalanceAmmount > 0:
			fatalOn(fmt.Errorf("TransactionInfo, NOT BALANCED [balance=%s], Will be rejected", txBalance))
		default:
			// Transaction is balanced :)
		}

		txStaging, err = jcli.TransactionFinalize(
			txStaging, "",
			settings.Fees.Certificate,
			settings.Fees.Coefficient,
			settings.Fees.Constant,
			srcAccount,
		)
		fatalOn(err, b2s(txStaging))

		// Get transaction data for witness (right now the same as TransactionID)
		txDataForWitness, err := jcli.TransactionDataForWitness(txStaging, "")
		fatalOn(err, b2s(txDataForWitness))

		// save the witness data to this file
		witnessFile := srcAccount + ".witness"
		witness, err := jcli.TransactionMakeWitness(
			[]byte(srcPrivateKey),
			b2s(txDataForWitness),
			settings.Block0Hash,
			"account", newAccountRest.Counter,
			witnessFile,
			"",
		)
		fatalOn(err, b2s(witness))

		txStaging, err = jcli.TransactionAddWitness(txStaging, "", witnessFile)
		fatalOn(err, b2s(txStaging))
		_ = os.Remove(witnessFile)

		txStaging, err = jcli.TransactionSeal(txStaging, "")
		fatalOn(err, b2s(txStaging))

		txMessage, err := jcli.TransactionToMessage(txStaging, "")
		fatalOn(err, b2s(txMessage))

		bulkData[i].message = txMessage
	}

	// send transaction message and get fragment_id back.
	for i := range bulkData {
		fragmentID, err := jcli.RestMessagePost(bulkData[i].message, restAddrAPI, "")
		fatalOn(err, b2s(fragmentID))

		bulkData[i].fragment = fragmentID

		log.Printf("%s - %s\n", bulkData[i].faucet.Account, b2s(fragmentID))
	}

}

/* KIT */

// writeJsonFile to accounts/filename after marshalling data to json
func writeJsonFile(v interface{}, fileName string) error {
	finalFile := "accounts" + pathSep + fileName
	addrJson, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	if fileName == "" {
		fmt.Printf("%s\n", addrJson)
		return nil
	}
	return ioutil.WriteFile(finalFile, addrJson, 0644)
}

// b2s converts []byte to string with all leading
// and trailing white space removed, as defined by Unicode.
func b2s(b []byte) string {
	return strings.TrimSpace(string(b))
}
