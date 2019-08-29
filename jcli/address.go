package jcli

import (
	"fmt"
	"strings"
)

// AddressInfoFromStdin - display the content and info of a bech32 formatted address.
//
// jcli does not use STDIN for this function, but keep it for convenience
func AddressInfoFromStdin(
	addressBech32 []byte,
) ([]byte, error) {
	if len(addressBech32) == 0 {
		return nil, fmt.Errorf("%s : EMPTY", "stdin_bech32")
	}

	return AddressInfo(strings.TrimSuffix(string(addressBech32), "\n"))
}

// AddressInfo - display the content and info of a bech32 formatted address.
//
// jcli address info <ADDRESS (in bech32 format)>
func AddressInfo(
	addressBech32 string,
) ([]byte, error) {
	if addressBech32 == "" {
		return nil, fmt.Errorf("parameter missing : %s", "addressBech32")
	}

	arg := []string{"address", "info", addressBech32}

	return execStd(nil, "jcli", arg...)
}

// AddressAccountFromStdin - create an address from the the single public key.
//
// jcli does not use STDIN for this function, but keep it for convenience
func AddressAccountFromStdin(
	publicKey []byte,
	prefix string,
	discrimination string,
) ([]byte, error) {
	if len(publicKey) == 0 {
		return nil, fmt.Errorf("%s : EMPTY", "publicKey")
	}

	return AddressAccount(strings.TrimSuffix(string(publicKey), "\n"), prefix, discrimination)
}

// AddressAccount - create an address from the the single public key.
//
// jcli address account <PUBLIC_KEY> [--prefix <address_prefix>] [--testing]
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

	return execStd(nil, "jcli", arg...)
}

// AddressSingleFromStdin - create an address from the single public key. This address does not have delegation.
//
// jcli does not use STDIN for this function, but keep it for convenience
func AddressSingleFromStdin(
	publicKey []byte,
	groupPublicKey string,
	prefix string,
	discrimination string,
) ([]byte, error) {
	if len(publicKey) == 0 {
		return nil, fmt.Errorf("%s : EMPTY", "publicKey")
	}

	return AddressSingle(strings.TrimSuffix(string(publicKey), "\n"), groupPublicKey, prefix, discrimination)
}

// AddressSingle - create an address from the single public key. This address does not have delegation.
//
// jcli address single <PUBLIC_KEY> [DELEGATION_KEY] [--prefix <address_prefix>] [--testing]
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

	return execStd(nil, "jcli", arg...)
}
