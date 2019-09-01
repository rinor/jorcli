package jcli

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// TransactionNew - create a new staging transaction. The transaction is initially empty.
//
// STDIN | jcli transaction new [--staging <staging-file>] | [STDOUT]
func TransactionNew(
	stdinStaging []byte,
	stagingFile string,
) ([]byte, error) {
	arg := []string{"transaction", "new"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionAddInput - add UTxO input to the transaction.
//
// STDIN | jcli transaction add-input <transaction-id> <index> <value> [--staging <staging-file>]
func TransactionAddInput(
	stdinStaging []byte,
	stagingFile string,
	fragmentID string,
	outputIndex uint8,
	value uint64,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}
	if fragmentID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "fragmentID")
	}

	arg := []string{
		"transaction", "add-input",
		fragmentID,
		strconv.FormatUint(uint64(outputIndex), 10),
		strconv.FormatUint(value, 10),
	}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionAddAccount - add Account input to the transaction.
//
// STDIN | jcli transaction add-account <account> <value> [--staging <staging-file>]
func TransactionAddAccount(
	stdinStaging []byte,
	stagingFile string,
	account string,
	value uint64,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}
	if account == "" {
		return nil, fmt.Errorf("parameter missing : %s", "account")
	}

	arg := []string{
		"transaction", "add-account",
		account,
		strconv.FormatUint(value, 10),
	}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionAddOutput - add output to the transaction.
//
// STDIN | jcli transaction add-output <address> <value> [--staging <staging-file>]
func TransactionAddOutput(
	stdinStaging []byte,
	stagingFile string,
	address string,
	value uint64,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}
	if address == "" {
		return nil, fmt.Errorf("parameter missing : %s", "address")
	}

	arg := []string{
		"transaction", "add-output",
		address,
		strconv.FormatUint(value, 10),
	}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionAddWitness - add output to the finalized transaction.
//
// STDIN | jcli transaction add-witness <witness> [--staging <staging-file>]
func TransactionAddWitness(
	stdinStaging []byte,
	stagingFile string,
	witnessFile string, // FIXME: UPSTREAM add witness description since it is empty
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}
	if witnessFile == "" {
		return nil, fmt.Errorf("parameter missing : %s", "witnessFile")
	}

	arg := []string{"transaction", "add-witness", witnessFile}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionAddCertificate - set a certificate to the Transaction.
// If there is already an extra certificate in the transaction it will be replaced with the new one.
//
// STDIN | jcli transaction add-certificate <value> [--staging <staging-file>]
func TransactionAddCertificate(
	stdinStaging []byte,
	stagingFile string,
	certificateBech32 string, // FIXME: UPSTREAM add value description since is ambiguous
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}
	if certificateBech32 == "" {
		return nil, fmt.Errorf("parameter missing : %s", "certificateBech32")
	}

	arg := []string{"transaction", "add-certificate", certificateBech32}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionFinalize - Lock a transaction including provided fees and start adding witnesses.
//
// STDIN | jcli transaction finalize [--fee-certificate <certificate>]
//                                   [--fee-coefficient <coefficient>]
//                                   [--fee-constant <constant>]
//                                   [change]
//                                   [--staging <staging-file>]
func TransactionFinalize(
	stdinStaging []byte,
	stagingFile string,
	feeCertificate uint64,
	feeCoefficient uint64,
	feeConstant uint64,
	changeAddress string, // FIXME: UPSTREAM add change description since is ambiguous
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "finalize"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}
	arg = append(arg,
		"--fee-certificate", strconv.FormatUint(feeCertificate, 10),
		"--fee-coefficient", strconv.FormatUint(feeCoefficient, 10),
		"--fee-constant", strconv.FormatUint(feeConstant, 10),
	)
	if changeAddress != "" {
		arg = append(arg, changeAddress)
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionSeal - Finalize the transaction.
//
// STDIN | jcli transaction seal [--staging <staging-file>]
func TransactionSeal(
	stdinStaging []byte,
	stagingFile string,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "seal"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionId - get the Transaction ID from the given transaction
// (if the transaction is edited, the returned value will change).
//
// STDIN | jcli transaction id [--staging <staging-file>]
func TransactionId(
	stdinStaging []byte,
	stagingFile string,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "id"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	return execStd(stdinStaging, "jcli", arg...)
}

// TransactionToMessage - get the message format out of a sealed transaction.
//
// STDIN | jcli transaction to-message [--staging <staging-file>]
func TransactionToMessage(
	stdinStaging []byte,
	stagingFile string,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "to-message"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionMakeWitness - create witnesses.
//
// STDIN | jcli transaction make-witness <transaction-id> --genesis-block-hash <genesis-block-hash>
//                                       --type <witness-type (utxo,legacy-utxo,account)>
//                                       [--account-spending-counter <account-spending-counter> (mandatory if --type=account)]
//                                       [<output file>] [<secret file>]
func TransactionMakeWitness(
	stdinKey []byte,
	dataForWitness string, // FIXME: UPSTREAM (the real transaction id is fragmentID, but here we need trancasctionID -> data-for-witness)
	block0Hash string,
	typeWitness string, accountSpendingCounter uint32,
	outputFile string,
	inputFileKey string,
) ([]byte, error) {
	if len(stdinKey) == 0 && inputFileKey == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinKey", "inputFileKey")
	}
	if block0Hash == "" {
		return nil, fmt.Errorf("parameter missing : %s", "block0Hash")
	}
	if typeWitness == "" {
		return nil, fmt.Errorf("parameter missing : %s", "typeWitness")
	}

	arg := []string{
		"transaction", "make-witness", dataForWitness,
		"--genesis-block-hash", block0Hash,
		"--type", typeWitness,
	}
	if typeWitness == "account" {
		arg = append(arg, "--account-spending-counter", strconv.FormatUint(uint64(accountSpendingCounter), 10))
	}
	if outputFile != "" {
		arg = append(arg, outputFile) // TODO: UPSTREAM unify with "--output" as other file output commands
	}
	if inputFileKey != "" && outputFile != "" {
		arg = append(arg, inputFileKey) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinKey = nil                  // reset STDIN - not needed since input_file has priority over STDIN
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	// convert input_file to stdin
	if inputFileKey != "" && outputFile == "" {
		var err error // prevent variable shadowing of stdinKey
		stdinKey, err = ioutil.ReadFile(inputFileKey)
		if err != nil {
			return nil, err
		}
	}

	out, err := execStd(stdinKey, "jcli", arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
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
	stdinStaging []byte, stagingFile string,
	feeCertificate uint64,
	feeCoefficient uint64,
	feeConstant uint64,
	outputFile string, // TODO: UPSTREAM unify with "--output" as other file output commands
	format string,
	onlyUTxOs bool, // TODO: UPSTREAM convert to --only-utxos
	onlyAccounts bool, // TODO: UPSTREAM convert to --only-accounts
	onlyOutputs bool, // TODO: UPSTREAM convert to --only-outputs
	formatUTxOInput string, // TODO: UPSTREAM convert to --format-utxo-input <data>
	formatAccountInput string, // TODO: UPSTREAM convert to --format-account-input <data>
	formatOutput string, // TODO: UPSTREAM convert to --format-output <data>
	prefix string,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "info"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}
	arg = append(arg,
		"--fee-certificate", strconv.FormatUint(feeCertificate, 10),
		"--fee-coefficient", strconv.FormatUint(feeCoefficient, 10),
		"--fee-constant", strconv.FormatUint(feeConstant, 10),
	)
	if format != "" {
		arg = append(arg, "--format", format)
	}
	if prefix != "" {
		arg = append(arg, "--prefix", prefix)
	}
	if outputFile != "" {
		arg = append(arg, outputFile)
	}
	arg = append(arg,
		strconv.FormatBool(onlyUTxOs),
		strconv.FormatBool(onlyAccounts),
		strconv.FormatBool(onlyOutputs),
	)
	if formatUTxOInput != "" {
		arg = append(arg, formatUTxOInput)
	}
	if formatAccountInput != "" {
		arg = append(arg, formatAccountInput)
	}
	if formatOutput != "" {
		arg = append(arg, formatOutput)
	}

	out, err := execStd(stdinStaging, "jcli", arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

/////////////////////////////////////////////////////////////////////
// This is not yet implemented in jcli                             //
// Check https://github.com/input-output-hk/jormungandr/issues/674 //
/////////////////////////////////////////////////////////////////////

// TransactionDataForWitness - Sign data hash
//
// STDIN | jcli transaction data-for-witness [--staging <staging-file>]
func TransactionDataForWitness(
	stdinStaging []byte,
	stagingFile string,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}
	/*
		arg := []string{"transaction", "data-for-witness"}
	*/
	arg := []string{"transaction", "id"} // FIXME: restore data-for-witness once implemented
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil // reset STDIN - not needed since stagingFile has priority over STDIN
	}

	return execStd(stdinStaging, "jcli", arg...)
}
