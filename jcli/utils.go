package jcli

import (
	"fmt"
	"strings"
)

// UtilsBech32ConvertFromStdin - convert a bech32 with hrp n into a bech32 with prefix m.
// jcli does not use STDIN for this function, but keep it for convenience
func UtilsBech32ConvertFromStdin(
	bech32 []byte,
	newPrefix string,
) ([]byte, error) {
	if len(bech32) == 0 {
		return nil, fmt.Errorf("%s : EMPTY", "bech32")
	}

	return UtilsBech32Convert(strings.TrimSuffix(string(bech32), "\n"), newPrefix)
}

// UtilsBech32Convert - convert a bech32 with hrp n into a bech32 with prefix m.
//
// jcli utils bech32-convert <FROM_BECH32> <NEW_PREFIX>
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

	return execStd(nil, "jcli", arg...)
}
