package jcli

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// TransactionNew - create a new staging transaction. The transaction is initially empty.
//
//  [STDIN] | jcli transaction new [--staging <staging-file>] | [STDOUT]
func TransactionNew(
	stdinStaging []byte,
	stagingFile string,
) ([]byte, error) {
	arg := []string{"transaction", "new"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TODO: TransactionAddInput tests and examples

// TransactionAddInput - add UTxO input to the transaction.
//
//  [STDIN] | jcli transaction add-input <fragment-id> <index> <value> [--staging <staging-file>] | [STDOUT]
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
		stdinStaging = nil
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionAddAccount - add Account input to the transaction.
//
//  [STDIN] | jcli transaction add-account <account> <value> [--staging <staging-file>] | [STDOUT]
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
		stdinStaging = nil
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionAddOutput - add output to the transaction.
//
//  [STDIN] | jcli transaction add-output <address> <value> [--staging <staging-file>] | [STDOUT]
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
		stdinStaging = nil
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionAddWitness - add output to the finalized transaction.
//
//  [STDIN] | jcli transaction add-witness <witness> [--staging <staging-file>] | [STDOUT]
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
		stdinStaging = nil
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TODO: TransactionAddCertificate tests. Example done

// TransactionAddCertificate - set a certificate to the Transaction.
// If there is already an extra certificate in the transaction it will be replaced with the new one.
//
//  [STDIN] | jcli transaction add-certificate <value> [--staging <staging-file>] | [STDOUT]
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
		stdinStaging = nil
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionFinalize - Lock a transaction including provided fees and start adding witnesses.
//
//  [STDIN] | jcli transaction finalize
//                                      [--fee-certificate <certificate>]
//                                      [--fee-coefficient <coefficient>]
//                                      [--fee-constant <constant>]
//                                      [--fee-pool-registration <certificate-pool-registration>]
//                                      [--fee-stake-delegation <certificate-stake-delegation>]
//                                      [--fee-owner-stake-delegation <certificate-owner-stake-delegation>]
//                                      [--fee-vote-cast <certificate-vote-cast>]
//                                      [--fee-vote-plan <certificate-vote-plan>]
//                                      [change]
//                                      [--staging <staging-file>] | [STDOUT]
func TransactionFinalize(
	stdinStaging []byte,
	stagingFile string,
	feeCertificate uint64,
	feeCoefficient uint64,
	feeConstant uint64,
	feeCertificatePoolRegistration uint64,
	feeCertificateStakeDelegation uint64,
	feeCertificateOwnerStakeDelegation uint64,
	feeCertificateVoteCast uint64,
	feeCertificateVotePlan uint64,
	changeAddress string, // FIXME: UPSTREAM add change description since is ambiguous
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "finalize"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil
	}
	arg = append(arg,
		"--fee-certificate", strconv.FormatUint(feeCertificate, 10),
		"--fee-coefficient", strconv.FormatUint(feeCoefficient, 10),
		"--fee-constant", strconv.FormatUint(feeConstant, 10),
	)
	if feeCertificatePoolRegistration > 0 {
		arg = append(arg,
			"--fee-pool-registration", strconv.FormatUint(feeCertificatePoolRegistration, 10),
		)
	}
	if feeCertificateStakeDelegation > 0 {
		arg = append(arg,
			"--fee-stake-delegation", strconv.FormatUint(feeCertificateStakeDelegation, 10),
		)
	}
	if feeCertificateOwnerStakeDelegation > 0 {
		arg = append(arg,
			"--fee-owner-stake-delegation", strconv.FormatUint(feeCertificateOwnerStakeDelegation, 10),
		)
	}
	if feeCertificateVoteCast > 0 {
		arg = append(arg,
			"--fee-vote-cast", strconv.FormatUint(feeCertificateVoteCast, 10),
		)
	}
	if feeCertificateVotePlan > 0 {
		arg = append(arg,
			"--fee-vote-plan", strconv.FormatUint(feeCertificateVotePlan, 10),
		)
	}
	if changeAddress != "" {
		arg = append(arg, changeAddress)
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionSeal - Finalize the transaction.
//
//  [STDIN] | jcli transaction seal [--staging <staging-file>] | [STDOUT]
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
		stdinStaging = nil
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionToMessage - get the message format out of a sealed transaction.
//
//  [STDIN] | jcli transaction to-message [--staging <staging-file>] | [STDOUT]
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
		stdinStaging = nil
	}

	return jcli(stdinStaging, arg...)
}

// TransactionToMessageFile - get the message format out of a sealed transaction
// and write it to the given outputFile.
// This is like TransactionToMessage with outFile added since jcli has only STDOUT
func TransactionToMessageFile(
	stdinStaging []byte,
	stagingFile string,
	outputFile string,
) ([]byte, error) {
	if outputFile == "" {
		return nil, fmt.Errorf("parameter missing : %s", "outputFile")
	}

	message, err := TransactionToMessage(stdinStaging, stagingFile)
	if err != nil {
		return nil, err
	}
	if err = ioutil.WriteFile(outputFile, message, 0644); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(outputFile)
}

// TransactionMakeWitness - create witnesses.
//
//  [STDIN] | jcli transaction make-witness
//                                          <transaction-id>
//                                          --genesis-block-hash <genesis-block-hash>
//                                          --type <witness-type (utxo,legacy-utxo,account)>
//                                          [--account-spending-counter <account-spending-counter> (mandatory if --type=account)]
//                                          [<output file>]
//                                          [<secret file>] | [STDOUT]
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
		stdinKey = nil
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	//
	// convert input_file to stdin
	if inputFileKey != "" && outputFile == "" {
		var (
			err error // [xXx] - prevent variable shadowing of stdinKey
		)
		stdinKey, err = ioutil.ReadFile(inputFileKey)
		if err != nil {
			return nil, err
		}
	}

	out, err := jcli(stdinKey, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// TransactionInfo - display the info regarding a given transaction.
//
//  [STDIN] | jcli transaction info
//                                  [--staging <staging-file>]
//                                  [--fee-certificate <certificate>]
//                                  [--fee-coefficient <coefficient>]
//                                  [--fee-constant <constant>]
//                                  [--fee-pool-registration <certificate-pool-registration>]
//                                  [--fee-stake-delegation <certificate-stake-delegation>]
//                                  [--fee-owner-stake-delegation <certificate-owner-stake-delegation>]
//                                  [--fee-vote-cast <certificate-vote-cast>]
//                                  [--fee-vote-plan <certificate-vote-plan>]
//                                  [--output-format <format>]
//                                  [--output <output file>]
//                                  [--prefix <address prefix>] | [STDOUT]
func TransactionInfo(
	stdinStaging []byte, stagingFile string,
	feeCertificate uint64,
	feeCoefficient uint64,
	feeConstant uint64,
	feeCertificatePoolRegistration uint64,
	feeCertificateStakeDelegation uint64,
	feeCertificateOwnerStakeDelegation uint64,
	feeCertificateVoteCast uint64,
	feeCertificateVotePlan uint64,
	prefix string,
	outputFormat string,
	outputFile string,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "info"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil
	}
	arg = append(arg,
		"--fee-certificate", strconv.FormatUint(feeCertificate, 10),
		"--fee-coefficient", strconv.FormatUint(feeCoefficient, 10),
		"--fee-constant", strconv.FormatUint(feeConstant, 10),
	)
	if feeCertificatePoolRegistration > 0 {
		arg = append(arg,
			"--fee-pool-registration", strconv.FormatUint(feeCertificatePoolRegistration, 10),
		)
	}
	if feeCertificateStakeDelegation > 0 {
		arg = append(arg,
			"--fee-stake-delegation", strconv.FormatUint(feeCertificateStakeDelegation, 10),
		)
	}
	if feeCertificateOwnerStakeDelegation > 0 {
		arg = append(arg,
			"--fee-owner-stake-delegation", strconv.FormatUint(feeCertificateOwnerStakeDelegation, 10),
		)
	}
	if feeCertificateVoteCast > 0 {
		arg = append(arg,
			"--fee-vote-cast", strconv.FormatUint(feeCertificateVoteCast, 10),
		)
	}
	if feeCertificateVotePlan > 0 {
		arg = append(arg,
			"--fee-vote-plan", strconv.FormatUint(feeCertificateVotePlan, 10),
		)
	}
	if prefix != "" {
		arg = append(arg, "--prefix", prefix)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// TODO: tests and examples

// TransactionAuth - make auth
//
//  [STDIN] | jcli transaction auth --key <signing-keys>... [--staging <staging-file>] | [STDOUT]
func TransactionAuth(
	stdinStaging []byte,
	stagingFile string,
	signingKeyFile []string,
) ([]byte, error) {

	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}
	if len(signingKeyFile) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "key")
	}

	arg := []string{"transaction", "auth"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil
	}
	for _, authKeyFile := range signingKeyFile {
		arg = append(arg, "--key", authKeyFile) // FIXME: should check data validity!
	}

	out, err := jcli(stdinStaging, arg...)
	if err != nil || stagingFile == "" {
		return out, err
	}

	return ioutil.ReadFile(stagingFile)
}

// TransactionDataForWitness - Sign data hash
//
//  [STDIN] | jcli transaction data-for-witness [--staging <staging-file>] | [STDOUT]
func TransactionDataForWitness(
	stdinStaging []byte,
	stagingFile string,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "data-for-witness"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil
	}

	return jcli(stdinStaging, arg...)
}

// TransactionFragmentID - get the Fragment ID from the given SEALED transaction
//
//  [STDIN] | jcli transaction fragment-id [--staging <staging-file>] | [STDOUT]
func TransactionFragmentID(
	stdinStaging []byte,
	stagingFile string,
) ([]byte, error) {
	if len(stdinStaging) == 0 && stagingFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinStaging", "stagingFile")
	}

	arg := []string{"transaction", "fragment-id"}
	if stagingFile != "" {
		arg = append(arg, "--staging", stagingFile)
		stdinStaging = nil
	}

	return jcli(stdinStaging, arg...)
}
