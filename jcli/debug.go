package jcli

import "fmt"

// DebugMessage - Decode hex-encoded message an display its content.
//
//  [STDIN] | jcli debug message [--input <input>] | STDOUT
func DebugMessage(
	stdinHex []byte,
	inputFile string,
) ([]byte, error) {
	if len(stdinHex) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinHex", "inputFile")
	}

	arg := []string{"debug", "message"}
	if inputFile != "" {
		arg = append(arg, "--input", inputFile)
		stdinHex = nil
	}

	return execStd(stdinHex, "jcli", arg...)
}
