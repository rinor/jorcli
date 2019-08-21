package jcli

import (
	"fmt"
	"io/ioutil"
)

// GenesisInit - Create a default Genesis file with appropriate documentation to help creating the YAML file.
// jcli genesis init
func GenesisInit() ([]byte, error) {
	return execStd(nil, "jcli", "genesis", "init")
}

// GenesisHash - print the block hash (aka the block id) of the block0.
// STDIN | jcli genesis hash [--input <FILE_INPUT>]
func GenesisHash(
	block0_bin []byte,
	input_file string,
) ([]byte, error) {
	if len(block0_bin) == 0 && input_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "block0_bin", "input_file")
	}

	arg := []string{"genesis", "hash"}
	if input_file != "" {
		arg = append(arg, "--input", input_file)
		block0_bin = nil // reset STDIN - not needed since input_file has priority over STDIN
	}

	return execStd(block0_bin, "jcli", arg...)
}

// GenesisEncode - create the block 0 file (the genesis block of the blockchain) from a given yaml file.
// STDIN | jcli genesis encode [--input <FILE_INPUT>] [--output <FILE_OUTPUT>]
func GenesisEncode(
	block0_txt []byte,
	input_file string,
	output_file string,
) ([]byte, error) {
	if len(block0_txt) == 0 && input_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "block0_txt", "input_file")
	}

	arg := []string{"genesis", "encode"}
	if input_file != "" {
		arg = append(arg, "--input", input_file)
		if len(block0_txt) > 0 {
			block0_txt = nil // reset STDIN - not needed since input_file has priority over STDIN
		}
	}
	if output_file != "" {
		arg = append(arg, "--output", output_file)
	}

	out, err := execStd(block0_txt, "jcli", arg...)
	if err != nil || output_file == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file)
}

// GenesisDecode - Decode the block 0 and print the corresponding YAML file.
// STDIN | jcli genesis decode [--input <FILE_INPUT>] [--output <FILE_OUTPUT>]
func GenesisDecode(
	block0_bin []byte,
	input_file string,
	output_file string,
) ([]byte, error) {
	if len(block0_bin) == 0 && input_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "block0_bin", "input_file")
	}

	arg := []string{"genesis", "decode"}
	if input_file != "" {
		arg = append(arg, "--input", input_file)
		block0_bin = nil // reset STDIN - not needed since input_file has priority over STDIN
	}
	if output_file != "" {
		arg = append(arg, "--output", output_file)
	}

	out, err := execStd(block0_bin, "jcli", arg...)
	if err != nil || output_file == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file)
}
