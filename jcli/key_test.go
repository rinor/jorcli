package jcli_test

import (
	"fmt"
	"testing"

	"github.com/rinor/jorcli/jcli"
)

func ExampleKeyGenerate_noseed() {
	var (
		seed         = ""
		keyType      = "ed25519extended"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "privateKey.sk" - will also save output to that file

	)

	privateKey, err := jcli.KeyGenerate(seed, keyType, outputFileSk)

	if err != nil {
		fmt.Printf("KeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
}

func ExampleKeyGenerate_seed() {
	var (
		seed         = "0000000000000000000000000000000000000000000000000000000000000000"
		keyType      = "ed25519extended"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "privateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.KeyGenerate(seed, keyType, outputFileSk)

	if err != nil {
		fmt.Printf("KeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
}

func TestKeyToPublic_file(t *testing.T) {
	var (
		stdinSk           []byte
		inputFileSk       = filePath(t, "private_key_txt.golden")
		outputFilePk      = ""
		expectedPublicKey = loadBytes(t, "public_key_txt.golden")
	)

	publicKey, err := jcli.KeyToPublic(stdinSk, inputFileSk, outputFilePk)
	ok(t, err)
	equals(t, expectedPublicKey, publicKey) // Prod: bytes.Equal(expectedPublicKey, publicKey)
}

func ExampleKeyToPublic_stdin() {
	var (
		stdinSk      = []byte("ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4")
		inputFileSk  = "" // "" - input from STDIN (stdinSk []byte), "privateKey.sk" - will load the private key from that file
		outputFilePk = "" // "" - output to STDOUT ([]byte) only, "publicKey.pk" - will also save the public key to that file
	)

	publicKey, err := jcli.KeyToPublic(stdinSk, inputFileSk, outputFilePk)

	if err != nil {
		fmt.Printf("KeyToPublic: %s", err)
	} else {
		fmt.Printf("%s", string(publicKey))
	}
	// Output:
	//
	// ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
}

func TestKeyToBytes_file(t *testing.T) {
	var (
		stdinSk                 []byte
		outputFile              = ""
		inputFileSk             = filePath(t, "private_key_txt.golden")
		expectedPrivateKeyBytes = loadBytes(t, "private_key_bytes.golden")
	)

	privateKeyBytes, err := jcli.KeyToBytes(stdinSk, outputFile, inputFileSk)
	ok(t, err)
	equals(t, expectedPrivateKeyBytes, privateKeyBytes) // Prod: bytes.Equal(expectedPublicKey, actualPublicKey)
}

func ExampleKeyToBytes_stdin() {
	var (
		stdinSk     = []byte("ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4")
		outputFile  = ""
		inputFileSk = ""
	)

	privateKeyBytes, err := jcli.KeyToBytes(stdinSk, outputFile, inputFileSk)

	if err != nil {
		fmt.Printf("KeyToBytes: %s", err)
	} else {
		fmt.Printf("%s", string(privateKeyBytes))
	}
	// Output:
	//
	// 70b8e0ada0f13d90405d6ae55386bd28bdd219b8a08ded1aa836efcc8b770d47da41597c5157488d7724e03fb8d84a376a43b8f41518a11cc387b669b2ee6586
}

func TestKeyFromBytes_file(t *testing.T) {
	var (
		stdinSk            []byte
		keyType            = "ed25519extended"
		inputFile          = filePath(t, "private_key_bytes.golden")
		outputFileSk       = ""
		expectedPrivateKey = loadBytes(t, "private_key_txt.golden")
	)

	privateKey, err := jcli.KeyFromBytes(stdinSk, keyType, inputFile, outputFileSk)
	ok(t, err)
	equals(t, expectedPrivateKey, privateKey) // Prod: bytes.Equal(expectedPublicKey, privateKey)
}

func ExampleKeyFromBytes_stdin() {
	var (
		stdinSk      = []byte("70b8e0ada0f13d90405d6ae55386bd28bdd219b8a08ded1aa836efcc8b770d47da41597c5157488d7724e03fb8d84a376a43b8f41518a11cc387b669b2ee6586")
		keyType      = "ed25519extended"
		inputFile    = ""
		outputFileSk = ""
	)

	privateKey, err := jcli.KeyFromBytes(stdinSk, keyType, inputFile, outputFileSk)

	if err != nil {
		fmt.Printf("KeyFromBytes: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
}
