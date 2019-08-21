package jcli

import (
	"fmt"
	"io/ioutil"
)

// CertificateGetStakePoolID - get the stake pool id from the given stake pool registration certificate.
// STDIN | jcli certificate get-stake-pool-id [<FILE_INPUT>] [<FILE_OUTPUT>]
func CertificateGetStakePoolID(
	stdin_cert []byte,
	input_file string,
	output_file string,
) ([]byte, error) {
	if len(stdin_cert) == 0 && input_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_cert", "input_file")
	}

	arg := []string{"certificate", "get-stake-pool-id"}
	if input_file != "" {
		arg = append(arg, input_file) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdin_cert = nil              // reset STDIN - not needed since input_file has priority over STDIN
	}
	if output_file != "" && input_file != "" {
		arg = append(arg, output_file) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(stdin_cert, "jcli", arg...)
	if err != nil /* || output_file == "" */ {
		return out, err
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	// convert stdout to output_file
	if output_file != "" && input_file == "" {
		if err := ioutil.WriteFile(output_file, out, 0644); err != nil {
			return nil, err
		}
	}

	if output_file == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file)
}

// CertificateNewStakeDelegation - build a stake delegation certificate.
// jcli certificate new stake-delegation <STAKE_POOL_ID> <STAKE_KEY> [output]
func CertificateNewStakeDelegation(
	stake_pool_id string,
	stake_key string,
	output_file string,
) ([]byte, error) {
	if stake_pool_id == "" {
		return nil, fmt.Errorf("parameter missing : %s", "stake_pool_id")
	}
	if stake_key == "" {
		return nil, fmt.Errorf("parameter missing : %s", "stake_key")
	}

	arg := []string{"certificate", "new", "stake-delegation", stake_pool_id, stake_key}
	if output_file != "" {
		arg = append(arg, output_file) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(nil, "jcli", arg...)
	if err != nil || output_file == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file)
}

// CertificateNewStakePoolRegistrationSingleOwner - build a stake pool registration certificate with single owner.
// jcli certificate new stake-pool-registration --kes-key <KES_KEY>
//                                              --vrf-key <VRF_KEY>
//                                              --serial <SERIAL>
//                                              [--owner <PUBLIC_KEY>]
//                                              [output]
func CertificateNewStakePoolRegistrationSingleOwner(
	kes_key string,
	vrf_key string,
	serial string,
	owner string,
	output_file string,
) ([]byte, error) {
	return CertificateNewStakePoolRegistration(kes_key, vrf_key, serial, []string{owner}, output_file)
}

// CertificateNewStakePoolRegistration - build a stake pool registration certificate.
// jcli certificate new stake-pool-registration --kes-key <KES_KEY>
//                                              --vrf-key <VRF_KEY>
//                                              --serial <SERIAL>
//                                              [--owner <PUBLIC_KEY> --owner <PUBLIC_KEY> ...]
//                                              [output]
func CertificateNewStakePoolRegistration(
	kes_key string,
	vrf_key string,
	serial string,
	owner []string,
	output_file string,
) ([]byte, error) {
	if kes_key == "" {
		return nil, fmt.Errorf("parameter missing : %s", "kes_key")
	}
	if vrf_key == "" {
		return nil, fmt.Errorf("parameter missing : %s", "vrf_key")
	}
	if serial == "" {
		return nil, fmt.Errorf("parameter missing : %s", "serial")
	}

	arg := []string{
		"certificate", "new", "stake-pool-registration",
		"--kes-key", kes_key,
		"--vrf-key", vrf_key, "--serial", serial,
	}
	for _, owner_pk := range owner {
		arg = append(arg, "--owner", owner_pk) // FIXME: should check data validity!
	}
	if output_file != "" {
		arg = append(arg, output_file) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(nil, "jcli", arg...)
	if err != nil || output_file == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file)
}

// CertificateSign - Sign certificate, you can call this command multiple time to add multiple signatures if this is required.
// STDIN | jcli certificate sign <signing-key file> [<input file>] [<output file>]
func CertificateSign(
	stdin_cert []byte,
	signing_key_file string,
	input_file string,
	output_file string,
) ([]byte, error) {
	if len(stdin_cert) == 0 && input_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_cert", "input_file")
	}
	if signing_key_file == "" {
		return nil, fmt.Errorf("parameter missing : %s", "signing_key_file")
	}

	arg := []string{"certificate", "sign", signing_key_file}
	if input_file != "" {
		arg = append(arg, input_file) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdin_cert = nil              // reset STDIN - not needed since input_file has priority over STDIN
	}
	if output_file != "" && input_file != "" {
		arg = append(arg, output_file) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := execStd(stdin_cert, "jcli", arg...)
	if err != nil /* || output_file == "" */ {
		return out, err
	}

	// TODO: Remove this once UPSTREAM fixed (--input and --output)
	// convert stdout to output_file
	if output_file != "" && input_file == "" {
		err = ioutil.WriteFile(output_file, out, 0644)
		if err != nil {
			return nil, err
		}
	}
	if output_file == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file)
}
