package jcli

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// TransactionNew - create a new staging transaction. The transaction is initially empty.
//
// STDIN | jcli transaction new [--staging <staging-file>]
func TransactionNew(
	stdin_staging []byte,
	staging_file string,
) ([]byte, error) {
	arg := []string{"transaction", "new"}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionAddInput - add UTxO input to the transaction.
//
// STDIN | jcli transaction add-input <transaction-id> <index> <value> [--staging <staging-file>]
func TransactionAddInput(
	stdin_staging []byte,
	staging_file string,
	fragment_id string,
	output_index uint8,
	value uint64,
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}
	if fragment_id == "" {
		return nil, fmt.Errorf("parameter missing : %s", "fragment_id")
	}

	arg := []string{"transaction", "add-input", fragment_id, strconv.FormatUint(uint64(output_index), 10), strconv.FormatUint(value, 10)}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}
	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionAddAccount - add Account input to the transaction.
//
// STDIN | jcli transaction add-account <account> <value> [--staging <staging-file>]
func TransactionAddAccount(
	stdin_staging []byte,
	staging_file string,
	account string,
	value uint64,
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}
	if account == "" {
		return nil, fmt.Errorf("parameter missing : %s", "account")
	}

	arg := []string{"transaction", "add-account", account, strconv.FormatUint(value, 10)}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionAddOutput - add output to the transaction.
//
// STDIN | jcli transaction add-output <address> <value> [--staging <staging-file>]
func TransactionAddOutput(
	stdin_staging []byte,
	staging_file string,
	address string,
	value uint64,
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}
	if address == "" {
		return nil, fmt.Errorf("parameter missing : %s", "address")
	}

	arg := []string{"transaction", "add-output", address, strconv.FormatUint(value, 10)}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionAddWitness - add output to the finalized transaction.
//
// STDIN | jcli transaction add-witness <witness> [--staging <staging-file>]
func TransactionAddWitness(
	stdin_staging []byte,
	staging_file string,
	witness_file string, // FIXME: UPSTREAM add witness description since it is empty
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}
	if witness_file == "" {
		return nil, fmt.Errorf("parameter missing : %s", "witness_file")
	}

	arg := []string{"transaction", "add-witness", witness_file}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionAddCertificate - set a certificate to the Transaction.
// If there is already an extra certificate in the transaction it will be replaced with the new one.
//
// STDIN | jcli transaction add-certificate <value> [--staging <staging-file>]
func TransactionAddCertificate(
	stdin_staging []byte,
	staging_file string,
	certificate_bech32 string, // FIXME: UPSTREAM add value description since is ambiguous
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}
	if certificate_bech32 == "" {
		return nil, fmt.Errorf("parameter missing : %s", "certificate_bech32")
	}

	arg := []string{"transaction", "add-certificate", certificate_bech32}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionFinalize - Lock a transaction including provided fees and start adding witnesses.
//
// STDIN | jcli transaction finalize [--fee-certificate <certificate>]
//                                   [--fee-coefficient <coefficient>]
//                                   [--fee-constant <constant>]
//                                   [change]
//                                   [--staging <staging-file>]
func TransactionFinalize(
	stdin_staging []byte,
	staging_file string,
	fee_certificate uint64,
	fee_coefficient uint64,
	fee_constant uint64,
	change_address string, // FIXME: UPSTREAM add change description since is ambiguous
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}

	arg := []string{"transaction", "finalize"}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}
	arg = append(arg,
		"--fee-certificate", strconv.FormatUint(fee_certificate, 10),
		"--fee-coefficient", strconv.FormatUint(fee_coefficient, 10),
		"--fee-constant", strconv.FormatUint(fee_constant, 10),
	)
	if change_address != "" {
		arg = append(arg, change_address)
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionSeal - Finalize the transaction.
//
// STDIN | jcli transaction seal [--staging <staging-file>]
func TransactionSeal(
	stdin_staging []byte,
	staging_file string,
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}

	arg := []string{"transaction", "seal"}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionId - get the Transaction ID from the given transaction
// (if the transaction is edited, the returned value will change).
//
// STDIN | jcli transaction id [--staging <staging-file>]
func TransactionId(
	stdin_staging []byte,
	staging_file string,
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}

	arg := []string{"transaction", "id"}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionToMessage - get the message format out of a sealed transaction.
//
// STDIN | jcli transaction to-message [--staging <staging-file>]
func TransactionToMessage(
	stdin_staging []byte,
	staging_file string,
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}

	arg := []string{"transaction", "to-message"}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || staging_file == "" {
		return out, err
	}

	return ioutil.ReadFile(staging_file)
}

// TransactionMakeWitness - create witnesses.
//
// STDIN | jcli transaction make-witness <transaction-id> --genesis-block-hash <genesis-block-hash>
//                                       --type <witness-type (utxo,legacy-utxo,account)>
//                                       [--account-spending-counter <account-spending-counter> (mandatory if --type=account)]
//                                       [<output file>] [<secret file>]
func TransactionMakeWitness(
	stdin_key []byte,
	transaction_id string, // FIXME: UPSTREAM (the real transaction id is fragmentID, but here we need trancasctionID -> data-for-witness)
	block0_hash string,
	type_witness string, account_spending_counter uint32,
	output_file string,
	input_file_key string,
) ([]byte, error) {
	if len(stdin_key) == 0 && input_file_key == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_key", "input_file_key")
	}
	if block0_hash == "" {
		return nil, fmt.Errorf("parameter missing : %s", "block0_hash")
	}
	if type_witness == "" {
		return nil, fmt.Errorf("parameter missing : %s", "type_witness")
	}

	arg := []string{"transaction", "make-witness", "--genesis-block-hash", block0_hash, "--type", type_witness}
	if type_witness == "account" {
		arg = append(arg, "--account-spending-counter", strconv.FormatUint(uint64(account_spending_counter), 10))
	}
	if output_file != "" {
		arg = append(arg, output_file) // TODO: UPSTREAM unify with "--output" as other file output commands
	}
	if input_file_key != "" && output_file != "" {
		arg = append(arg, input_file_key) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdin_key = nil                   // reset STDIN - not needed since input_file has priority over STDIN
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	// convert input_file to stdin
	if input_file_key != "" && output_file == "" {
		var err error // prevent variable shadowing of stdin_key
		stdin_key, err = ioutil.ReadFile(input_file_key)
		if err != nil {
			return nil, err
		}
	}

	out, err := execStd(stdin_key, "jcli", arg...)
	if err != nil || output_file == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file)
}

// TransactionInfo - display the info regarding a given transaction.
//
// STDIN | jcli transaction info [--staging <staging-file>]
//                               [--fee-certificate <certificate>]
//                               [--fee-coefficient <coefficient>]
//                               [--fee-constant <constant>]
//                               [--format <format>]
//                               [<output file>]
//                               [<only-utxos>]
//                               [<only-accounts>]
//                               [<only-outputs>]
//                               [<format-utxo-input>]
//                               [<format-account-input>]
//                               [<format-output>]
//                               [--prefix <address prefix>]
func TransactionInfo(
	stdin_staging []byte, staging_file string,
	fee_certificate uint64,
	fee_coefficient uint64,
	fee_constant uint64,
	output_file string,
	format string,
	only_utxos bool,
	only_accounts bool,
	only_outputs bool,
	format_utxo_input string,
	format_account_input string,
	format_output string,
	prefix string,
) ([]byte, error) {
	if len(stdin_staging) == 0 && staging_file == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
	}

	arg := []string{"transaction", "info"}
	if staging_file != "" {
		arg = append(arg, "--staging", staging_file)
		stdin_staging = nil
	}
	arg = append(arg,
		"--fee-certificate", strconv.FormatUint(fee_certificate, 10),
		"--fee-coefficient", strconv.FormatUint(fee_coefficient, 10),
		"--fee-constant", strconv.FormatUint(fee_constant, 10),
	)
	if format != "" {
		arg = append(arg, "--format", format)
	}
	if prefix != "" {
		arg = append(arg, "--prefix", prefix)
	}
	if output_file != "" {
		arg = append(arg, output_file)
	}
	arg = append(arg, strconv.FormatBool(only_utxos), strconv.FormatBool(only_accounts), strconv.FormatBool(only_outputs))
	if format_utxo_input != "" {
		arg = append(arg, format_utxo_input)
	}
	if format_account_input != "" {
		arg = append(arg, format_account_input)
	}
	if format_output != "" {
		arg = append(arg, format_output)
	}

	out, err := execStd(stdin_staging, "jcli", arg...)
	if err != nil || output_file == "" {
		return out, err
	}

	return ioutil.ReadFile(output_file)
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// // This is not yet implemented check https://github.com/input-output-hk/jormungandr/issues/674
// //
// // TransactionDataForWitness -
// // STDIN | jcli transaction data-for-witness [--staging <staging-file>]
// func TransactionDataForWitness(
// 	stdin_staging []byte,
// 	staging_file string,
// ) ([]byte, error) {
// 	if len(stdin_staging) == 0 && staging_file == "" {
// 		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_staging", "staging_file")
// 	}
// 	arg := []string{"transaction", "data-for-witness"}

// 	if staging_file != "" {
// 		arg = append(arg, "--staging", staging_file)
// 		stdin_staging = nil
// 	}

// 	out, err := execStd(stdin_staging, "jcli", arg...)
// 	if err != nil || staging_file == "" {
// 		return out, err
// 	}

// 	return ioutil.ReadFile(staging_file)
// }

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
