package jcli_test

import (
	"fmt"

	"github.com/rinor/jorcli/jcli"
)

func ExampleAddressAccount() {
	var (
		publicKey      = "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"
		addressPrefix  = "ta"
		discrimination = "testing"
	)

	ac, err := jcli.AddressAccount(publicKey, addressPrefix, discrimination)

	if err != nil {
		fmt.Printf("AddressAccount: %s", err)
	} else {
		fmt.Printf("%s", string(ac))
	}
	// Output:
	//
	// ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap
}

func ExampleAddressInfo() {
	var (
		addressBech32 = "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"
	)

	adi, err := jcli.AddressInfo(addressBech32)

	if err != nil {
		fmt.Printf("AddressInfo: %s", err)
	} else {
		fmt.Printf("%s", string(adi))
	}
	// Output:
	//
	// discrimination: testing
	// account: ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
}

func ExampleAddressInfo_single() {
	var (
		addressBech32 = "ta1s3uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4u7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467x80kcm"
	)

	adi, err := jcli.AddressInfo(addressBech32)

	if err != nil {
		fmt.Printf("AddressInfo: %s", err)
	} else {
		fmt.Printf("%s", string(adi))
	}
	// Output:
	//
	// discrimination: testing
	// public key: ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
	// group key:  ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
}

func ExampleAddressSingle() {
	var (
		publicKey      = "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"
		groupPublicKey = "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"
		addressPrefix  = "ta"
		discrimination = "testing"
	)

	ac, err := jcli.AddressSingle(publicKey, groupPublicKey, addressPrefix, discrimination)

	if err != nil {
		fmt.Printf("ExampleAddressSingle: %s", err)
	} else {
		fmt.Printf("%s", string(ac))
	}
	// Output:
	//
	// ta1s3uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4u7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467x80kcm
}
