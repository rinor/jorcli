package jcli

import (
	"fmt"
	"strings"
)

// AddressInfoFromStdin - display the content and info of a bech32 formatted address.
//
// jcli does not use STDIN for this function, but keep it for convenience
func AddressInfoFromStdin(
	address_bech32 []byte,
) ([]byte, error) {
	if len(address_bech32) == 0 {
		return nil, fmt.Errorf("%s : EMPTY", "stdin_bech32")
	}

	return AddressInfo(strings.TrimSuffix(string(address_bech32), "\n"))
}

// AddressInfo - display the content and info of a bech32 formatted address.
//
// jcli address info <ADDRESS (in bech32 format)>
func AddressInfo(
	address_bech32 string,
) ([]byte, error) {
	if address_bech32 == "" {
		return nil, fmt.Errorf("parameter missing : %s", "address_bech32")
	}

	arg := []string{"address", "info", address_bech32}

	return execStd(nil, "jcli", arg...)
}

// AddressAccountFromStdin - create an address from the the single public key.
//
// jcli does not use STDIN for this function, but keep it for convenience
func AddressAccountFromStdin(
	public_key []byte,
	prefix string,
	discrimination string,
) ([]byte, error) {
	if len(public_key) == 0 {
		return nil, fmt.Errorf("%s : EMPTY", "public_key")
	}

	return AddressAccount(strings.TrimSuffix(string(public_key), "\n"), prefix, discrimination)
}

// AddressAccount - create an address from the the single public key.
//
// jcli address account <PUBLIC_KEY> [--prefix <address_prefix>] [--testing]
func AddressAccount(
	public_key string,
	prefix string,
	discrimination string,
) ([]byte, error) {
	if public_key == "" {
		return nil, fmt.Errorf("parameter missing : %s", "public_key")
	}

	arg := []string{"address", "account", public_key}
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
	public_key []byte,
	group_pk string,
	discrimination string,
) ([]byte, error) {
	if len(public_key) == 0 {
		return nil, fmt.Errorf("%s : EMPTY", "public_key")
	}

	return AddressSingle(strings.TrimSuffix(string(public_key), "\n"), group_pk, discrimination)
}

// AddressSingle - create an address from the single public key. This address does not have delegation.
//
// jcli address single <PUBLIC_KEY> [DELEGATION_KEY] [--testing]
func AddressSingle(
	public_key string,
	group_public_key string,
	discrimination string,
) ([]byte, error) {
	if public_key == "" {
		return nil, fmt.Errorf("parameter missing : %s", "public_key")
	}

	arg := []string{"address", "single", public_key}
	if group_public_key != "" {
		arg = append(arg, group_public_key)
	}
	if discrimination != "" {
		arg = append(arg, "--"+discrimination)
	}

	return execStd(nil, "jcli", arg...)
}
