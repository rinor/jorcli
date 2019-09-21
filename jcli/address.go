package jcli

import (
	"fmt"
)

// AddressInfo - display the content and info of a bech32 formatted address.
//
//  jcli address info <ADDRESS (in bech32 format)> | STDOUT
func AddressInfo(
	addressBech32 string,
) ([]byte, error) {
	if addressBech32 == "" {
		return nil, fmt.Errorf("parameter missing : %s", "addressBech32")
	}

	arg := []string{"address", "info", addressBech32}

	return jcli(nil, arg...)
}

// AddressAccount - create an address from the the single public key.
//
//  jcli address account <PUBLIC_KEY> [--prefix <address_prefix>] [--testing] | STDOUT
func AddressAccount(
	publicKey string,
	prefix string,
	discrimination string,
) ([]byte, error) {
	if publicKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "publicKey")
	}

	arg := []string{"address", "account", publicKey}
	if prefix != "" {
		arg = append(arg, "--prefix", prefix)
	}
	if discrimination != "" {
		arg = append(arg, "--"+discrimination)
	}

	return jcli(nil, arg...)
}

// AddressSingle - create an address from the single public key. This address does not have delegation.
//
//  jcli address single <PUBLIC_KEY> [DELEGATION_KEY] [--prefix <address_prefix>] [--testing] | STDOUT
func AddressSingle(
	publicKey string,
	groupPublicKey string,
	prefix string,
	discrimination string,
) ([]byte, error) {
	if publicKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "publicKey")
	}

	arg := []string{"address", "single", publicKey}
	if groupPublicKey != "" {
		arg = append(arg, groupPublicKey)
	}
	if prefix != "" {
		arg = append(arg, "--prefix", prefix)
	}
	if discrimination != "" {
		arg = append(arg, "--"+discrimination)
	}

	return jcli(nil, arg...)
}
