package jcli

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// VotesCrsGenerate - generate the Common Reference String.
//
//  jcli votes crs generate [--seed=<SEED>] [OUTPUT_FILE] | [STDOUT]
func VotesCRSGenerate(
	seed string,
	outputFileSk string,
) ([]byte, error) {
	arg := []string{
		"votes", "crs", "generate",
	}
	if seed != "" {
		arg = append(arg, "--seed", seed)
	}
	if outputFileSk != "" {
		arg = append(arg, outputFileSk)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFileSk == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFileSk)
}

// VotesCommitteeCommunicationKeyGenerate - generate a committee communication private key.
//
//  jcli votes committee communication-key generate [--seed=<SEED>] [OUTPUT_FILE] | [STDOUT]
func VotesCommitteeCommunicationKeyGenerate(
	seed string,
	outputFileSk string,
) ([]byte, error) {
	arg := []string{
		"votes", "committee", "communication-key", "generate",
	}
	if seed != "" {
		arg = append(arg, "--seed", seed)
	}
	if outputFileSk != "" {
		arg = append(arg, outputFileSk)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFileSk == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFileSk)
}

// VotesCommitteeCommunicationKeyToPublic - get the public key out of a given committee communication private key.
//
//  [STDIN] | jcli votes committee communication-key to-public [--input=input_file] [OUTPUT_FILE] | [STDOUT]
func VotesCommitteeCommunicationKeyToPublic(
	stdinSk []byte,
	inputFileSk string,
	outputFilePk string,
) ([]byte, error) {
	if len(stdinSk) == 0 && inputFileSk == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinSk", "inputFileSk")
	}

	arg := []string{
		"votes", "committee", "communication-key", "to-public",
	}
	if inputFileSk != "" {
		arg = append(arg, "--input", inputFileSk)
		stdinSk = nil
	}
	if outputFilePk != "" {
		arg = append(arg, outputFilePk)
	}

	out, err := jcli(stdinSk, arg...)
	if err != nil || outputFilePk == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFilePk)
}

// VotesCommitteeMemberKeyGenerate - generate a committee member private key.
//
//  jcli votes committee member-key generate
//                                          --crs=<CRS>
//                                          --keys=<COMMUNICATION_KEYS>...
//                                          --threshold=<THRESHOLD>
//                                          --index=<INDEX>
//                                          [--seed=<SEED>]
//                                          [OUTPUT_FILE] | [STDOUT]
func VotesCommitteeMemberKeyGenerate(
	crs string,
	threshold uint8,
	keys []string,
	index uint8,
	seed string,
	outputFileSk string,
) ([]byte, error) {
	if crs == "" {
		return nil, fmt.Errorf("parameter missing : %s", "crs")
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "keys")
	}
	if threshold == 0 || int(threshold) > len(keys) {
		return nil, fmt.Errorf("%s expected between %d - %d, got %d", "threshold", 1, len(keys), threshold)
	}
	if int(index) >= len(keys) {
		return nil, fmt.Errorf("%s expected between %d - %d, got %d", "index", 0, len(keys)-1, index)
	}

	arg := []string{
		"votes", "committee", "member-key", "generate",
		"--crs", crs,
		"--threshold", strconv.FormatUint(uint64(threshold), 10),
		"--index", strconv.FormatUint(uint64(index), 10),
	}
	for _, key := range keys {
		arg = append(arg, "--keys", key) // FIXME: should check data validity!
	}
	if seed != "" {
		arg = append(arg, "--seed", seed)
	}
	if outputFileSk != "" {
		arg = append(arg, outputFileSk)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFileSk == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFileSk)
}

// VotesCommitteeMemberKeyToPublic - get the public key out of a given committee member private key.
//
//  [STDIN] |jcli votes committee member-key to-public [--input=input_file] [OUTPUT_FILE] | [STDOUT]
func VotesCommitteeMemberKeyToPublic(
	stdinSk []byte,
	inputFileSk string,
	outputFilePk string,
) ([]byte, error) {
	if len(stdinSk) == 0 && inputFileSk == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinSk", "inputFileSk")
	}

	arg := []string{
		"votes", "committee", "member-key", "to-public",
	}
	if inputFileSk != "" {
		arg = append(arg, "--input", inputFileSk)
		stdinSk = nil
	}
	if outputFilePk != "" {
		arg = append(arg, outputFilePk)
	}

	out, err := jcli(stdinSk, arg...)
	if err != nil || outputFilePk == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFilePk)
}

// VotesEncryptingKey - Build an encryption vote key.
//
//  jcli votes encrypting-key --keys=<member-keys>... [OUTPUT_FILE] | [STDOUT]
func VotesEncryptingKey(
	keys []string,
	outputFileSk string,
) ([]byte, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "keys")
	}

	arg := []string{
		"votes", "encrypting-key",
	}
	for _, key := range keys {
		arg = append(arg, "--keys", key) // FIXME: should check data validity!
	}
	if outputFileSk != "" {
		arg = append(arg, outputFileSk)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFileSk == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFileSk)
}
