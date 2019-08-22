package jcli

import (
	"fmt"
	"io/ioutil"
)

// KeyGenerate - generate a private key without seed.
//
// jcli key generate --type=<key_type> [OUTPUT_FILE]
func KeyGenerate(
	key_type string,
	output_file_sk string,
) ([]byte, error) {
	if key_type == "" {
		return nil, fmt.Errorf("parameter missing : %s", "key_type")
	}

	return KeyGenerateFromSeed("", key_type, output_file_sk)
}

// KeyGenerateFromSeed - generate a private key using a SEED value.
//
// jcli key generate --type=<key_type> [--seed=<SEED>] [OUTPUT_FILE]
func KeyGenerateFromSeed(
	seed string,
	key_type string,
	output_file_sk string,
) ([]byte, error) {
	if key_type == "" {
		return nil, fmt.Errorf("parameter missing : %s", "key_type")
	}
	arg := []string{"key", "generate", "--type", key_type}
	if seed != "" {
		arg = append(arg, "--seed", seed)
	}
	if output_file_sk != "" {
		arg = append(arg, output_file_sk)
	}

	out, err := execStd(nil, "jcli", arg...)
	if err != nil || output_file_sk == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file_sk)
}

// KeyToPublic - get the public key out of a given private key.
//
// STDIN | jcli key to-public [--input=input_file] [OUTPUT_FILE] (input file has priority over STDIN)
func KeyToPublic(
	stdin_sk []byte,
	input_file_sk string,
	output_file_pk string,
) ([]byte, error) {
	if len(stdin_sk) == 0 && input_file_sk == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_sk", "input_file_sk")
	}

	arg := []string{"key", "to-public"}
	if input_file_sk != "" {
		arg = append(arg, "--input", input_file_sk)
		stdin_sk = nil // reset STDIN - not needed since input_file has priority over STDIN
	}
	if output_file_pk != "" {
		arg = append(arg, output_file_pk) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(stdin_sk, "jcli", arg...)
	if err != nil || output_file_pk == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file_pk)
}

// KeyToBytes - get the bytes out of a private key. [TODO: encodes also public key but corverts it wrong]
//
// STDIN | jcli key to-bytes [OUTPUT_FILE] [INPUT_FILE]
func KeyToBytes(
	stdin_key []byte,
	output_file_bytes string,
	input_file_key string,
) ([]byte, error) {
	if len(stdin_key) == 0 && input_file_key == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_key", "input_file_key")
	}

	arg := []string{"key", "to-bytes"}
	if output_file_bytes != "" {
		arg = append(arg, output_file_bytes)
	}
	if input_file_key != "" && output_file_bytes != "" {
		arg = append(arg, input_file_key) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdin_key = nil                   // reset STDIN since not needed
	}

	// TODO: Remove this once UPSTREAM fixed (--input and --output)
	// convert input_file to stdin
	if input_file_key != "" && output_file_bytes == "" {
		var err error // prevent variable shadowing of stdin_key
		stdin_key, err = ioutil.ReadFile(input_file_key)
		if err != nil {
			return nil, err
		}
	}

	out, err := execStd(stdin_key, "jcli", arg...)
	if err != nil || output_file_bytes == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file_bytes)
}

// KeyFromBytes - retrive a private key from the given bytes. [TODO: UPSTREAM encodes also public key but corverts it wrong]
//
// STDIN | jcli key from-bytes --type=<key_type> [INPUT_BYTES] [OUTPUT_FILE]
func KeyFromBytes(
	stdin_key []byte,
	key_type string,
	input_file_bytes string,
	output_file_key string,
) ([]byte, error) {
	if key_type == "" {
		return nil, fmt.Errorf("parameter missing : %s", "key_type")
	}
	if len(stdin_key) == 0 && input_file_bytes == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_key", "input_file_bytes")
	}

	arg := []string{"key", "from-bytes", "--type", key_type}
	if input_file_bytes != "" {
		arg = append(arg, input_file_bytes) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdin_key = nil                     // reset STDIN since not needed
	}
	if output_file_key != "" && input_file_bytes != "" {
		arg = append(arg, output_file_key) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(stdin_key, "jcli", arg...)
	if err != nil /* || output_file_key == "" */ {
		return out, err
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	// convert stdout to output_file
	if output_file_key != "" && input_file_bytes == "" {
		if err = ioutil.WriteFile(output_file_key, out, 0644); err != nil {
			return nil, err
		}
	}
	if output_file_key == "" {
		return out, nil
	}

	return ioutil.ReadFile(output_file_key)
}
