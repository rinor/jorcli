package jcli

import (
	"fmt"
)

// UtilsBech32Convert - convert a bech32 with hrp n into a bech32 with prefix m.
//
//  jcli utils bech32-convert <FROM_BECH32> <NEW_PREFIX> | STDOUT
func UtilsBech32Convert(
	bech32 string,
	newPrefix string,
) ([]byte, error) {
	if bech32 == "" {
		return nil, fmt.Errorf("parameter missing : %s", "bech32")
	}
	if newPrefix == "" {
		return nil, fmt.Errorf("parameter missing : %s", "newPrefix")
	}

	arg := []string{"utils", "bech32-convert", bech32, newPrefix}

	return jcli(nil, arg...)
}
