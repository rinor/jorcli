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

func main() {

	var (
		err error

		// Rest
		restAddr    = "127.0.0.22" // rest ip
		restPort    = 8001         // rest port
		restAddress = restAddr + ":" + strconv.Itoa(restPort)

		// P2P
		p2pIPver = "ip4" // ipv4 or ipv6
		p2pProto = "tcp" // tcp

		// P2P Public
		p2pPubAddr       = "127.0.0.22" // PublicAddres
		p2pPubPort       = 9001         // node P2P Public Port
		p2pPublicAddress = "/" + p2pIPver + "/" + p2pPubAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pPubPort)

		// P2P Listen
		p2pListenAddr    = "127.0.0.22" // ListenAddress
		p2pListenPort    = 9001         // node P2P Public Port
		p2pListenAddress = "/" + p2pIPver + "/" + p2pListenAddr + "/" + p2pProto + "/" + strconv.Itoa(p2pListenPort)

		// General
		discrimination = "testing"  // "" (empty defaults to "production")
		addressPrefix  = "jnode_ta" // "" (empty defaults to "ca")

		// Trusted peers
		trustedPeerLeader = "/ip4/127.0.0.11/tcp/9001" // Leader (genesis) node (example 1)

		// Genesis Block0 Hash retrieved from example (1)
		block0Hash = "9a0245551cded1536defeb494a979624656b33bebeac9a95130a92fad347ade6"

		// genesis accounts data
		genesisExtraSeed = 50 // seed the owner of an extra pool in genesis block
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
	log.Println()
	log.Printf("Working Directory: %s", workingDir)

	/**********************************************************************************************************/

	///////////////////////////////////////
	// GENESIS Extra STAKE POOL Creation //
	///////////////////////////////////////

	//////////////////////////////
	// Genesis Extra Pool Owner //
	//////////////////////////////

	// will need this one file later for certificate signing
	gepoFileSK := workingDir + string(os.PathSeparator) + "gepo_key.sk"

	gepoSK, err := jcli.KeyGenerate(seed(genesisExtraSeed), "Ed25519Extended", gepoFileSK)
	fatalOn(err, b2s(gepoSK))
	gepoPK, err := jcli.KeyToPublic(gepoSK, "", "")
	fatalOn(err, b2s(gepoPK))
	gepoAddr, err := jcli.AddressAccount(b2s(gepoPK), addressPrefix, discrimination)
	fatalOn(err, b2s(gepoAddr))

	// gep VRF
	gepPoolVrfSK, err := jcli.KeyGenerate(seed(genesisExtraSeed+1), "Curve25519_2HashDH", "")
	fatalOn(err, b2s(gepPoolVrfSK))
	gepPoolVrfPK, err := jcli.KeyToPublic(gepPoolVrfSK, "", "")
	fatalOn(err, b2s(gepPoolVrfPK))

	// gep KES
	gepPoolKesSK, err := jcli.KeyGenerate(seed(genesisExtraSeed+2), "SumEd25519_12", "")
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
	gepStakePoolCertSigned, err := jcli.CertificateSign(gepStakePoolCert, gepoFileSK, "", "")
	fatalOn(err, b2s(gepStakePoolCertSigned))

	/////////////////////////////////
	// Extra STAKE POOL Delegation //
	/////////////////////////////////

	gepStakePoolID, err := jcli.CertificateGetStakePoolID(gepStakePoolCertSigned, "", "")
	fatalOn(err, b2s(gepStakePoolID))

	// Genesis Extra Pool Owner delegation
	stakeDelegationGepoCert, err := jcli.CertificateNewStakeDelegation(b2s(gepStakePoolID), b2s(gepoPK), "")
	fatalOn(err, b2s(stakeDelegationGepoCert))
	stakeDelegationGepoCertSigned, err := jcli.CertificateSign(stakeDelegationGepoCert, gepoFileSK, "", "")
	fatalOn(err, b2s(stakeDelegationGepoCertSigned))

	/**********************************************************************************************************/

	//////////////////////
	//  secrets config  //
	//////////////////////

	secretCfg := jnode.NewSecretConfig()

	secretCfg.Genesis.SigKey = b2s(gepPoolKesSK)
	secretCfg.Genesis.VrfKey = b2s(gepPoolVrfSK)
	secretCfg.Genesis.NodeID = b2s(gepStakePoolID)

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
	node.GenesisBlockHash = block0Hash // add block0 hash

	node.AddTrustedPeer(trustedPeerLeader) // add leader from example (1) as trusted

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
	log.Printf("StakePool ID       : %s", gepStakePoolID)
	log.Printf("StakePool Owner    : %s", gepoAddr)
	log.Printf("StakePool Delegator: %s", gepoAddr)
	log.Println()

	log.Println("Genesis StakePool Node - Running...")
	node.Wait()                                     // Wait for the node to stop.
	log.Println("...Genesis StakePool Node - Done") // All done. Node has stopped.
}
