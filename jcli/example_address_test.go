package jcli_test

import (
	"fmt"
	"strings"

	"github.com/rinor/jorcli/jcli"
)

func ExampleAddressAccount() {
	public_key := "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"
	address_prefix := "ta"
	discrimination := "testing"

	ac, err := jcli.AddressAccount(public_key, address_prefix, discrimination)
	if err != nil {
		fmt.Printf("AddressAccount: %s - %s", err, ac)
		return
	}
	fmt.Printf("%s", strings.TrimSuffix(string(ac), "\n"))
	// Output:
	//
	// ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap
}

func ExampleAddressAccountFromStdin() {
	public_key := []byte("ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v")
	address_prefix := "ta"
	discrimination := "testing"

	ac, err := jcli.AddressAccountFromStdin(public_key, address_prefix, discrimination)
	if err != nil {
		fmt.Printf("AddressAccountFromStdin: %s - %s", err, ac)
		return
	}
	fmt.Printf("%s", strings.TrimSuffix(string(ac), "\n"))
	// Output:
	//
	// ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap
}

func ExampleAddressInfo() {
	address_bech32 := "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"

	adi, err := jcli.AddressInfo(address_bech32)
	if err != nil {
		fmt.Printf("AddressInfo: %s - %s", err, adi)
		return
	}
	fmt.Printf("%s", strings.TrimSuffix(string(adi), "\n"))
	// Output:
	//
	// discrimination: testing
	// account: ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
}

func ExampleAddressInfo_single() {
	address_bech32 := "ta1s3uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4u7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467x80kcm"

	adi, err := jcli.AddressInfo(address_bech32)
	if err != nil {
		fmt.Printf("AddressInfo: %s - %s", err, adi)
		return
	}
	fmt.Printf("%s", strings.TrimSuffix(string(adi), "\n"))
	// Output:
	//
	// discrimination: testing
	// public key: ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
	// group key:  ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
}

func ExampleAddressInfoFromStdin() {
	address_bech32 := []byte("ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap")

	adi, err := jcli.AddressInfoFromStdin(address_bech32)
	if err != nil {
		fmt.Printf("AddressInfoFromStdin: %s - %s", err, adi)
		return
	}
	fmt.Printf("%s", strings.TrimSuffix(string(adi), "\n"))
	// Output:
	//
	// discrimination: testing
	// account: ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
}

func ExampleAddressSingle() {
	public_key := "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"
	group_public_key := "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"
	address_prefix := "ta"
	discrimination := "testing"

	ac, err := jcli.AddressSingle(public_key, group_public_key, address_prefix, discrimination)
	if err != nil {
		fmt.Printf("AddressSingleFromStdin: %s - %s", err, ac)
		return
	}
	fmt.Printf("%s", strings.TrimSuffix(string(ac), "\n"))
	// Output:
	//
	// ta1s3uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4u7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467x80kcm
}

func ExampleAddressSingleFromStdin() {
	public_key := []byte("ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v")
	group_public_key := "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"
	address_prefix := "ta"
	discrimination := "testing"

	ac, err := jcli.AddressSingleFromStdin(public_key, group_public_key, address_prefix, discrimination)
	if err != nil {
		fmt.Printf("AddressSingleFromStdin: %s - %s", err, ac)
		return
	}
	fmt.Printf("%s", strings.TrimSuffix(string(ac), "\n"))
	// Output:
	//
	// ta1s3uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4u7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467x80kcm
}
