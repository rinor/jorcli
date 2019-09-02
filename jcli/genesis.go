package jcli

import (
	"fmt"
	"io/ioutil"
)

// GenesisInit - Create a default Genesis file with appropriate documentation to help creating the YAML file.
//
//  jcli genesis init | STDOUT
func GenesisInit() ([]byte, error) {
	return execStd(nil, "jcli", "genesis", "init")
}

// GenesisHash - print the block hash (aka the block id) of the block0.
//
//  [STDIN] | jcli genesis hash [--input <FILE_INPUT>] | STDOUT
func GenesisHash(
	block0Bin []byte,
	inputFile string,
) ([]byte, error) {
	if len(block0Bin) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "block0Bin", "inputFile")
	}

	arg := []string{"genesis", "hash"}
	if inputFile != "" {
		arg = append(arg, "--input", inputFile)
		block0Bin = nil
	}

	return execStd(block0Bin, "jcli", arg...)
}

// GenesisEncode - create the block 0 file (the genesis block of the blockchain) from a given yaml file.
//
//  [STDIN] | jcli genesis encode [--input <FILE_INPUT>] [--output <FILE_OUTPUT>] | [STDOUT]
func GenesisEncode(
	block0Txt []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(block0Txt) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "block0Txt", "inputFile")
	}

	arg := []string{"genesis", "encode"}
	if inputFile != "" {
		arg = append(arg, "--input", inputFile)
		block0Txt = nil
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := execStd(block0Txt, "jcli", arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// GenesisDecode - Decode the block 0 and print the corresponding YAML file.
//
//  [STDIN] | jcli genesis decode [--input <FILE_INPUT>] [--output <FILE_OUTPUT>] | [STDOUT]
func GenesisDecode(
	block0Bin []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(block0Bin) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "block0Bin", "inputFile")
	}

	arg := []string{"genesis", "decode"}
	if inputFile != "" {
		arg = append(arg, "--input", inputFile)
		block0Bin = nil
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := execStd(block0Bin, "jcli", arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}
