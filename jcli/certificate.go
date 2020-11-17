package jcli

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// CertificateShowStakePoolID - get the stake pool id from the given stake pool registration certificate.
//
//  [STDIN] | jcli certificate show stake-pool-id [--input <FILE_INPUT>] [--output <FILE_OUTPUT>] | [STDOUT]
func CertificateShowStakePoolID(
	stdinCertSigned []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinCertSigned) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCertSigned", "inputFile")
	}

	arg := []string{"certificate", "show", "stake-pool-id"}
	if inputFile != "" {
		arg = append(arg, "--input", inputFile)
		stdinCertSigned = nil
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(stdinCertSigned, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewOwnerStakeDelegation - build an owner stake delegation certificate.
//
//  jcli certificate new owner-stake-delegation <STAKE_POOL_ID:weight>... [--output <output>] | [STDOUT]
func CertificateNewOwnerStakeDelegation(
	weightedPoolID []string,
	outputFile string,
) ([]byte, error) {
	if len(weightedPoolID) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "weightedPoolID")
	}

	maxPools := 8 // The maximum number of delegation pools
	if len(weightedPoolID) > maxPools {
		return nil, fmt.Errorf("%s expected between %d - %d, got %d", "weightedPoolID", 1, maxPools, len(weightedPoolID))
	}

	arg := []string{
		"certificate", "new", "owner-stake-delegation",
	}
	arg = append(arg, weightedPoolID...) // FIXME: should check data validity!

	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewStakeDelegation - build a stake delegation certificate.
//
//  jcli certificate new stake-delegation <STAKE_KEY> <STAKE_POOL_ID:weight>... [--output <output>] | [STDOUT]
func CertificateNewStakeDelegation(
	stakeKey string,
	weightedPoolID []string,
	outputFile string,
) ([]byte, error) {
	if stakeKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "stakeKey")
	}
	if len(weightedPoolID) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "weightedPoolID")
	}

	maxPools := 8 // The maximum number of delegation pools
	if len(weightedPoolID) > maxPools {
		return nil, fmt.Errorf("%s expected between %d - %d, got %d", "weightedPoolID", 1, maxPools, len(weightedPoolID))
	}

	arg := []string{
		"certificate", "new", "stake-delegation",
		stakeKey,
	}
	arg = append(arg, weightedPoolID...) // FIXME: should check data validity!

	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewStakePoolRegistration - build a stake pool registration certificate with single/multiple owners.
//
//  jcli certificate new stake-pool-registration
//                                              --kes-key <KES_KEY>
//                                              --vrf-key <VRF_KEY>
//                                              --start-validity <SECONDS-SINCE-START>
//                                              --management-threshold <THRESHOLD>
//                                              --owner <OWNER_PUBLIC_KEY> --owner... (max owners TBD)
//                                              [--operator <OPERATOR_PUBLIC_KEY> --operator... (max operators TBD)]
//                                              [--tax-fixed <TAX_VALUE>]
//                                              [--tax-ratio <TAX_RATIO>]
//                                              [--tax-limit <TAX_LIMIT>]
//                                              [--reward-account <REWARD_ACCOUNT>]
//                                              [output] | STDOUT
func CertificateNewStakePoolRegistration(
	kesKey string,
	vrfKey string,
	startValidity uint64,
	managementThreshold uint8,
	owner []string,
	operator []string,
	taxFixed uint64,
	taxRatio string,
	taxLimit uint64,
	rewardAccount string,
	outputFile string,
) ([]byte, error) {
	if kesKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "kesKey")
	}
	if vrfKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "vrfKey")
	}
	if len(owner) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "owner")
	}
	// TODO: Confirm/Fix the limits
	/*

		maxOwners := 31   // 5 bits for the owners for a maximum of 31 elements
		maxOperators := 3 // 2 bits for the operators for a maximum of 3 elements
		if len(owner) > maxOwners {
			return nil, fmt.Errorf("%s expected between %d - %d, got %d", "owner", 1, maxOwners, len(owner))
		}
		if len(operator) > maxOperators {
			return nil, fmt.Errorf("%s expected between %d - %d, got %d", "operator", 0, maxOperators, len(operator))
		}

		// managementThreshold <= #owners and > 0
		if managementThreshold < 1 || int(managementThreshold) > len(owner) {
			return nil, fmt.Errorf("%s expected between %d - %d, got %d", "managementThreshold", 1, len(owner), managementThreshold)
		}
	*/
	arg := []string{
		"certificate", "new", "stake-pool-registration",
		"--kes-key", kesKey,
		"--vrf-key", vrfKey,
		"--start-validity", strconv.FormatUint(startValidity, 10),
		"--management-threshold", strconv.FormatUint(uint64(managementThreshold), 10),
	}
	for _, ownerPublicKey := range owner {
		arg = append(arg, "--owner", ownerPublicKey) // FIXME: should check data validity!
	}
	for _, operatorPublicKey := range operator {
		arg = append(arg, "--operator", operatorPublicKey) // FIXME: should check data validity!
	}

	if taxFixed > 0 {
		arg = append(arg, "--tax-fixed", strconv.FormatUint(taxFixed, 10))
	}
	if taxRatio != "" {
		arg = append(arg, "--tax-ration", taxRatio)
	}
	if taxLimit > 0 {
		arg = append(arg, "--tax-limit", strconv.FormatUint(taxLimit, 10))
	}
	if rewardAccount != "" {
		arg = append(arg, "--reward-account", rewardAccount)
	}
	if outputFile != "" {
		arg = append(arg, outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewStakePoolRetirement - retire the given stake pool ID From the blockchain.
// By doing so all remaining stake delegated to this stake pool will become pending and will need to be re-delegated.
//
// jcli certificate new stake-pool-retirement --pool-id <POOL_ID> --retirement-time <SECONDS-SINCE-START> [output]  | STDOUT
func CertificateNewStakePoolRetirement(
	poolID string,
	retirementTime uint64,
	outputFile string,
) ([]byte, error) {
	if poolID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "poolID")
	}

	arg := []string{
		"certificate", "new", "stake-pool-retirement",
		"--pool-id", poolID,
		"--retirement-time", strconv.FormatUint(retirementTime, 10),
	}
	if outputFile != "" {
		arg = append(arg, outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewVotePlan - create a vote plan certificate form given config data/file.
//
//  STDIN | jcli certificate new vote-plan [<FILE_INPUT>] [--output <FILE_OUTPUT>] | [STDOUT]
func CertificateNewVotePlan(
	stdinConfig []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinConfig) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinConfig", "inputFile")
	}

	arg := []string{
		"certificate", "new", "vote-plan",
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(stdinConfig, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateShowVotePlanID - get the vote plan id from the given vote plan certificate.
//
//  [STDIN] | jcli certificate show vote-plan-id [--input <FILE_INPUT>] [--output <FILE_OUTPUT>] | [STDOUT]
func CertificateShowVotePlanID(
	stdinCert []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinCert) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCert", "inputFile")
	}

	arg := []string{"certificate", "show", "vote-plan-id"}
	if inputFile != "" {
		arg = append(arg, "--input", inputFile)
		stdinCert = nil
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(stdinCert, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewVoteCastPublic - create a vote cast certificate for public voteplan.
//
// jcli certificate new vote-cast public --vote-plan-id <vote-plan-id> --proposal-index <proposal-index> --choice <choice> [--output <FILE_OUTPUT>] | STDOUT
func CertificateNewVoteCastPublic(
	votePlanID string,
	proposalIndex uint8,
	choice uint8,
	// optionsSize uint8,
	outputFile string,
) ([]byte, error) {
	if votePlanID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "votePlanID")
	}

	arg := []string{
		"certificate", "new", "vote-cast", "public",
		"--vote-plan-id", votePlanID,
		"--proposal-index", strconv.FormatUint(uint64(proposalIndex), 10),
		"--choice", strconv.FormatUint(uint64(choice), 10),
		// "--options-size", strconv.FormatUint(uint64(optionsSize), 10),
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewVoteCastPrivate - create a vote cast certificate for private voteplan.
//
//  [STDIN] | jcli certificate new vote-cast private  --vote-plan-id <vote-plan-id> --proposal-index <proposal-index> --choice <choice> [--key-path <encrypting-key-path>] [--output <FILE_OUTPUT>] | STDOUT
func CertificateNewVoteCastPrivate(
	stdinEncKey []byte,
	votePlanID string,
	proposalIndex uint8,
	choice uint8,
	optionsSize uint8,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinEncKey) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinEncKey", "inputFile")
	}
	if votePlanID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "votePlanID")
	}

	arg := []string{
		"certificate", "new", "vote-cast", "private",
		"--vote-plan-id", votePlanID,
		"--proposal-index", strconv.FormatUint(uint64(proposalIndex), 10),
		"--choice", strconv.FormatUint(uint64(choice), 10),
		"--options-size", strconv.FormatUint(uint64(optionsSize), 10),
	}
	if inputFile != "" {
		arg = append(arg, "--key-path", inputFile)
		stdinEncKey = nil
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(stdinEncKey, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewVoteTally - create a vote tally certificate.
//
// jcli certificate new vote-tally --vote-plan-id <id> [--output <FILE_OUTPUT>] | STDOUT
func CertificateNewVoteTally(
	votePlanID string,
	outputFile string,
) ([]byte, error) {
	if votePlanID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "votePlanID")
	}

	arg := []string{
		"certificate", "new", "vote-tally",
		"--vote-plan-id", votePlanID,
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewEncryptedVoteTally - create an encrypted vote tally certificate.
//
// jcli certificate new encrypted-vote-tally --vote-plan-id <id> [--output <FILE_OUTPUT>] | STDOUT
func CertificateNewEncryptedVoteTally(
	votePlanID string,
	outputFile string,
) ([]byte, error) {
	if votePlanID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "votePlanID")
	}

	arg := []string{
		"certificate", "new", "encrypted-vote-tally",
		"--vote-plan-id", votePlanID,
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateSign - Sign certificate,
// you can call this command multiple time to add multiple signatures if this is required.
//
//  [STDIN] | jcli certificate sign --key=<signing-key file>... [--certificate=<input file>] [--output=<output file>] | [STDOUT]
func CertificateSign(
	stdinCert []byte,
	signingKeyFile []string,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinCert) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCert", "inputFile")
	}
	if len(signingKeyFile) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "signingKeyFile")
	}

	arg := []string{"certificate", "sign"}
	for _, signKeyFile := range signingKeyFile {
		arg = append(arg, "--key", signKeyFile) // FIXME: should check data validity!
	}
	if inputFile != "" {
		arg = append(arg, "--certificate", inputFile)
		stdinCert = nil
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(stdinCert, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificatePrint - Print certificate.
//
//  [STDIN] | jcli certificate print [<input file>] | STDOUT
func CertificatePrint(
	stdinCert []byte,
	inputFile string,
) ([]byte, error) {
	if len(stdinCert) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCert", "inputFile")
	}

	arg := []string{"certificate", "print"}
	if inputFile != "" {
		arg = append(arg, inputFile) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinCert = nil
	}

	return jcli(stdinCert, arg...)
}
