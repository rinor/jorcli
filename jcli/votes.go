package jcli

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

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

// VotesElectionKey - Build an encryption vote key.
//
//  jcli votes election-key --keys=<member-keys>... [OUTPUT_FILE] | [STDOUT]
func VotesElectionKey(
	keys []string,
	outputFileSk string,
) ([]byte, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "keys")
	}

	arg := []string{
		"votes", "election-key",
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

// VotesTallyDecryptionShares - .
//
// [STDIN] | jcli votes tally decryption-shares
//                                               --key=<key_file>
//                                               [--vote-plan=<vote_plan_file>]
//                                               [--vote-plan-id=<vote-plan-id>] | [STDOUT]
func VotesTallyDecryptionShares(
	stdinVP []byte,
	inputFileSk string,
	inputFileVP string,
	vpID string,
) ([]byte, error) {
	if inputFileSk == "" {
		return nil, fmt.Errorf("parameter missing : %s", "inputFileSk")
	}
	if len(stdinVP) == 0 && inputFileVP == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinVP", "inputFileVP")
	}

	arg := []string{
		"votes", "tally", "decryption-shares",
		"--key", inputFileSk,
	}
	if inputFileVP != "" {
		arg = append(arg, "--vote-plan", inputFileVP)
		stdinVP = nil
	}
	if vpID != "" {
		arg = append(arg, "--vote-plan-id", vpID)
	}

	return jcli(stdinVP, arg...)
}

// VotesTallyMergeShares - .
//
// jcli votes tally merge-shares <shares_file> | [STDOUT]
func VotesTallyMergeShares(
	inputFileShares string,
) ([]byte, error) {
	if inputFileShares == "" {
		return nil, fmt.Errorf("parameter missing : %s", "inputFileShares")
	}
	arg := []string{
		"votes", "tally", "merge-shares",
		inputFileShares,
	}

	return jcli(nil, arg...)
}

// VotesTallyDecryptResults - .
//
// jcli votes tally decrypt-results
//                                  [--shares=<shares_file>]
//                                  [--threshold=<threshold>]
//                                  [--vote-plan=<vote_plan_file>]
//                                  [--vote-plan-id=<vote-plan-id>]
//                                  [--output-format=<format>] | [STDOUT]
func VotesTallyDecryptResults(
	stdin []byte,
	inputFileShares string,
	threshold uint8,
	inputFileVP string,
	vpID string,
	outputFormat string,
) ([]byte, error) {
	if len(stdin) == 0 && (inputFileShares == "" || inputFileVP == "") {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin", "inputFileShares || inputFileVP")
	}
	if len(stdin) != 0 && (inputFileShares == "" && inputFileVP == "") {
		return nil, fmt.Errorf("%s : NOT EMPTY but both parameters missing : %s", "stdin", "inputFileShares && inputFileVP")
	}

	arg := []string{
		"votes", "tally", "decrypt-result",
	}
	if inputFileShares != "" {
		arg = append(arg, "--shares", inputFileShares)
	}
	if threshold > 0 {
		arg = append(arg, "--threshold", strconv.FormatUint(uint64(threshold), 10))
	}
	if inputFileVP != "" {
		arg = append(arg, "--vote-plan", inputFileVP)
	}
	if vpID != "" {
		arg = append(arg, "--vote-plan-id", vpID)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	if inputFileShares != "" && inputFileVP != "" {
		stdin = nil
	}

	return jcli(stdin, arg...)
}
