package jcli

import (
	"fmt"
	"io/ioutil"
)

// KeyGenerate - generate a private key using a SEED value.
//
// jcli key generate --type=<key_type> [--seed=<SEED>] [OUTPUT_FILE]
func KeyGenerate(
	seed string,
	keyType string,
	outputFileSk string,
) ([]byte, error) {
	if keyType == "" {
		return nil, fmt.Errorf("parameter missing : %s", "keyType")
	}
	arg := []string{"key", "generate", "--type", keyType}
	if seed != "" {
		arg = append(arg, "--seed", seed)
	}
	if outputFileSk != "" {
		arg = append(arg, outputFileSk)
	}

	out, err := execStd(nil, "jcli", arg...)
	if err != nil || outputFileSk == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFileSk)
}

// KeyToPublic - get the public key out of a given private key.
//
// STDIN | jcli key to-public [--input=input_file] [OUTPUT_FILE] (input file has priority over STDIN)
func KeyToPublic(
	stdinSk []byte,
	inputFileSk string,
	outputFilePk string,
) ([]byte, error) {
	if len(stdinSk) == 0 && inputFileSk == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinSk", "inputFileSk")
	}

	arg := []string{"key", "to-public"}
	if inputFileSk != "" {
		arg = append(arg, "--input", inputFileSk)
		stdinSk = nil // reset STDIN - not needed since input_file has priority over STDIN
	}
	if outputFilePk != "" {
		arg = append(arg, outputFilePk) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(stdinSk, "jcli", arg...)
	if err != nil || outputFilePk == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFilePk)
}

// KeyToBytes - get the bytes out of a private key. [TODO: encodes also public key but corverts it wrong]
//
// STDIN | jcli key to-bytes [OUTPUT_FILE] [INPUT_FILE]
func KeyToBytes(
	stdinSk []byte,
	outputFile string,
	inputFileSk string,
) ([]byte, error) {
	if len(stdinSk) == 0 && inputFileSk == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinSk", "inputFileSk")
	}

	arg := []string{"key", "to-bytes"}
	if outputFile != "" {
		arg = append(arg, outputFile)
	}
	if inputFileSk != "" && outputFile != "" {
		arg = append(arg, inputFileSk) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinSk = nil                  // reset STDIN since not needed
	}

	// TODO: Remove this once UPSTREAM fixed (--input and --output)
	// convert input_file to stdin
	if inputFileSk != "" && outputFile == "" {
		var err error // prevent variable shadowing of stdinSk
		stdinSk, err = ioutil.ReadFile(inputFileSk)
		if err != nil {
			return nil, err
		}
	}

	out, err := execStd(stdinSk, "jcli", arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// KeyFromBytes - retrive a private key from the given bytes. [TODO: UPSTREAM encodes also public key but corverts it wrong]
//
// STDIN | jcli key from-bytes --type=<key_type> [INPUT_BYTES] [OUTPUT_FILE]
func KeyFromBytes(
	stdinSk []byte,
	keyType string,
	inputFile string,
	outputFileSk string,
) ([]byte, error) {
	if keyType == "" {
		return nil, fmt.Errorf("parameter missing : %s", "keyType")
	}
	if len(stdinSk) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinSk", "inputFile")
	}

	arg := []string{"key", "from-bytes", "--type", keyType}
	if inputFile != "" {
		arg = append(arg, inputFile) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinSk = nil                // reset STDIN since not needed
	}
	if outputFileSk != "" && inputFile != "" {
		arg = append(arg, outputFileSk) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(stdinSk, "jcli", arg...)
	if err != nil /* || outputFileSk == "" */ {
		return out, err
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	// convert stdout to output_file
	if outputFileSk != "" && inputFile == "" {
		if err = ioutil.WriteFile(outputFileSk, out, 0644); err != nil {
			return nil, err
		}
	}
	if outputFileSk == "" {
		return out, nil
	}

	return ioutil.ReadFile(outputFileSk)
}
