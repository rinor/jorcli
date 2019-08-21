package jcli

import (
	"fmt"
	"strings"
)

// UtilsBech32ConvertFromStdin - convert a bech32 with hrp n into a bech32 with prefix m.
// jcli does not use STDIN for this function, but keep it for convenience
func UtilsBech32ConvertFromStdin(
	stdin_bech32 []byte,
	new_prefix string,
) ([]byte, error) {
	if len(stdin_bech32) == 0 {
		return nil, fmt.Errorf("%s : EMPTY", "stdin_bech32")
	}
	return UtilsBech32Convert(strings.TrimSuffix(string(stdin_bech32), "\n"), new_prefix)
}

// UtilsBech32Convert - convert a bech32 with hrp n into a bech32 with prefix m.
// jcli utils bech32-convert <FROM_BECH32> <NEW_PREFIX>
func UtilsBech32Convert(
	bech32 string,
	new_prefix string,
) ([]byte, error) {
	if bech32 == "" {
		return nil, fmt.Errorf("parameter missing : %s", "bech32")
	}
	if new_prefix == "" {
		return nil, fmt.Errorf("parameter missing : %s", "new_prefix")
	}
	arg := []string{"utils", "bech32-convert", bech32, new_prefix}

	return execStd(nil, "jcli", arg...)
}
