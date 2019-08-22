package jcli

import "fmt"

// DebugMessage - Decode hex-encoded message an display its content.
//
// STDIN | jcli debug message [--input <input>]
func DebugMessage(
	stdin_hex []byte,
	input_file string,
) ([]byte, error) {
	if len(stdin_hex) == 0 && input_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_hex", "input_file")
	}
	arg := []string{"debug", "message"}

	if input_file != "" {
		arg = append(arg, "--input", input_file)
		stdin_hex = nil
	}

	return execStd(stdin_hex, "jcli", arg...)
}
