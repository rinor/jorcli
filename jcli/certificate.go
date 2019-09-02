package jcli

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// CertificateGetStakePoolID - get the stake pool id from the given stake pool registration certificate.
//
// STDIN | jcli certificate get-stake-pool-id [<FILE_INPUT>] [<FILE_OUTPUT>]
func CertificateGetStakePoolID(
	stdinCert []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinCert) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCert", "inputFile")
	}

	arg := []string{"certificate", "get-stake-pool-id"}
	if inputFile != "" {
		arg = append(arg, inputFile) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinCert = nil              // reset STDIN - not needed since inputFile has priority over STDIN
	}
	if outputFile != "" && inputFile != "" {
		arg = append(arg, outputFile) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(stdinCert, "jcli", arg...)
	if err != nil /* || outputFile == "" */ {
		return out, err
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	// convert stdout to outputFile
	if outputFile != "" && inputFile == "" {
		if err = ioutil.WriteFile(outputFile, out, 0644); err != nil {
			return nil, err
		}
	}
	if outputFile == "" {
		return out, nil
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewStakeDelegation - build a stake delegation certificate.
//
// jcli certificate new stake-delegation <STAKE_POOL_ID> <STAKE_KEY> [output]
func CertificateNewStakeDelegation(
	stakePoolID string,
	stakeKey string,
	outputFile string,
) ([]byte, error) {
	if stakePoolID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "stake_pstakePoolIDool_id")
	}
	if stakeKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "stakeKey")
	}

	arg := []string{"certificate", "new", "stake-delegation", stakePoolID, stakeKey}
	if outputFile != "" {
		arg = append(arg, outputFile) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(nil, "jcli", arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// BUG(rinor): The certificate 'serial' is declared as uint64 when actually it should be uint128.

// CertificateNewStakePoolRegistration - build a stake pool registration certificate with single/multiple owners.
//
// jcli certificate new stake-pool-registration --kes-key <KES_KEY>
//                                              --vrf-key <VRF_KEY>
//                                              --start-validity <SECONDS-SINCE-START>
//                                              --management-threshold <THRESHOLD> (<= #owners and > 0)
//                                              --serial <SERIAL>
//                                              [--owner <PUBLIC_KEY> --owner <PUBLIC_KEY> ...]
//                                              [output]
func CertificateNewStakePoolRegistration(
	kesKey string,
	vrfKey string,
	startValidity uint64,
	managementThreshold uint8,
	serial uint64,
	owner []string,
	outputFile string,
) ([]byte, error) {
	if kesKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "kesKey")
	}
	if vrfKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "vrfKey")
	}
	if managementThreshold < 1 || int(managementThreshold) > len(owner) {
		return nil, fmt.Errorf("%s expected between %d - %d, got %d", "managementThreshold", 1, len(owner), managementThreshold)
	}

	arg := []string{
		"certificate", "new", "stake-pool-registration",
		"--kes-key", kesKey,
		"--vrf-key", vrfKey,
		"--start-validity", strconv.FormatUint(startValidity, 10),
		"--management-threshold", strconv.FormatUint(uint64(managementThreshold), 10),
		"--serial", strconv.FormatUint(serial, 10),
	}
	for _, ownerPublicKey := range owner {
		arg = append(arg, "--owner", ownerPublicKey) // FIXME: should check data validity!
	}
	if outputFile != "" {
		arg = append(arg, outputFile) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(nil, "jcli", arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateSign - Sign certificate,
// you can call this command multiple time to add multiple signatures if this is required.
//
// STDIN | jcli certificate sign <signing-key file> [<input file>] [<output file>]
func CertificateSign(
	stdinCert []byte,
	signingKeyFile string,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinCert) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCert", "inputFile")
	}
	if signingKeyFile == "" {
		return nil, fmt.Errorf("parameter missing : %s", "signingKeyFile")
	}

	arg := []string{"certificate", "sign", signingKeyFile}
	if inputFile != "" {
		arg = append(arg, inputFile) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinCert = nil              // reset STDIN - not needed since inputFile has priority over STDIN
	}
	if outputFile != "" && inputFile != "" {
		arg = append(arg, outputFile) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(stdinCert, "jcli", arg...)
	if err != nil /* || outputFile == "" */ {
		return out, err
	}

	// TODO: Remove this once UPSTREAM fixed (--input and --output)
	// convert stdout to outputFile
	if outputFile != "" && inputFile == "" {
		if err = ioutil.WriteFile(outputFile, out, 0644); err != nil {
			return nil, err
		}
	}
	if outputFile == "" {
		return out, nil
	}

	return ioutil.ReadFile(outputFile)
}

// CertificatePrint - Print certificate.
//
// STDIN | jcli certificate print [<input file>]
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
		stdinCert = nil              // reset STDIN - not needed since inputFile has priority over STDIN
	}

	return execStd(stdinCert, "jcli", arg...)
}
