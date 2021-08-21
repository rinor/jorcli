package jcli_test

import (
	"fmt"
	"testing"

	"github.com/rinor/jorcli/jcli"
)

func ExampleTransactionNew_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"
	)

	tx, err := jcli.TransactionNew(stdinStaging, stagingFile)

	if err != nil {
		fmt.Printf("TransactionNew: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func ExampleTransactionNew_stdout() {
	var (
		stdinStaging []byte
		stagingFile  = ""
	)

	tx, err := jcli.TransactionNew(stdinStaging, stagingFile)

	if err != nil {
		fmt.Printf("TransactionNew: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
	// Output:
	//
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func TestTransactionNew(t *testing.T) {
	var (
		stdinStaging []byte
		stagingFile  = ""
		expectedTx   = loadBytes(t, "tx-01_new_staging.golden")
	)

	tx, err := jcli.TransactionNew(stdinStaging, stagingFile)
	ok(t, err)
	equals(t, expectedTx, tx)
}

func ExampleTransactionAddAccount_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"
		account      = "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"
		value        = uint64(100)
	)

	tx, err := jcli.TransactionAddAccount(stdinStaging, stagingFile, account, value)

	if err != nil {
		fmt.Printf("TransactionAddAccount: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestTransactionAddAccount(t *testing.T) {
	var (
		stdinStaging = loadBytes(t, "tx-01_new_staging.golden")
		stagingFile  = ""
		account      = "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"
		value        = uint64(100)
		expectedTx   = loadBytes(t, "tx-02_add_account_staging.golden")
	)

	tx, err := jcli.TransactionAddAccount(stdinStaging, stagingFile, account, value)
	ok(t, err)
	equals(t, expectedTx, tx)
}

func ExampleTransactionAddOutput_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"
		address      = "ta1skn8q3f6rxg92gqren9gf3heja8kj4shp89jl27n69v5azwvns95vlgw0pz"
		value        = uint64(50)
	)

	tx, err := jcli.TransactionAddOutput(stdinStaging, stagingFile, address, value)

	if err != nil {
		fmt.Printf("TransactionAddOutput: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestTransactionAddOutput(t *testing.T) {
	var (
		stdinStaging = loadBytes(t, "tx-02_add_account_staging.golden")
		stagingFile  = ""
		address      = "ta1skn8q3f6rxg92gqren9gf3heja8kj4shp89jl27n69v5azwvns95vlgw0pz"
		value        = uint64(50)
		expectedTx   = loadBytes(t, "tx-03_add_output_staging.golden")
	)

	tx, err := jcli.TransactionAddOutput(stdinStaging, stagingFile, address, value)
	ok(t, err)
	equals(t, expectedTx, tx)
}

func ExampleTransactionSetExpiryDate_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"
		blockDate    = "3.14"
	)

	tx, err := jcli.TransactionSetExpiryDate(stdinStaging, stagingFile, blockDate)

	if err != nil {
		fmt.Printf("TransactionSetExpiryDate: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestTransactionSetExpiryDate(t *testing.T) {
	var (
		stdinStaging = loadBytes(t, "tx-03_add_output_staging.golden")
		stagingFile  = ""
		blockDate    = "3.14"
		expectedTx   = loadBytes(t, "tx-04_set_expiry_date_staging.golden")
	)

	tx, err := jcli.TransactionSetExpiryDate(stdinStaging, stagingFile, blockDate)
	ok(t, err)
	equals(t, expectedTx, tx)
}

func ExampleTransactionFinalize_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"

		feeCertificate = uint64(3)
		feeCoefficient = uint64(2)
		feeConstant    = uint64(1)

		feeCertificatePoolRegistration     = uint64(3)
		feeCertificateStakeDelegation      = uint64(3)
		feeCertificateOwnerStakeDelegation = uint64(3)

		feeCertificateVoteCast = uint64(1)
		feeCertificateVotePlan = uint64(2)

		changeAddress = "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"
	)

	tx, err := jcli.TransactionFinalize(
		stdinStaging, stagingFile,
		feeCertificate,
		feeCoefficient,
		feeConstant,
		feeCertificatePoolRegistration,
		feeCertificateStakeDelegation,
		feeCertificateOwnerStakeDelegation,
		feeCertificateVoteCast,
		feeCertificateVotePlan,
		changeAddress,
	)

	if err != nil {
		fmt.Printf("TransactionFinalize: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestTransactionFinalize(t *testing.T) {
	var (
		stdinStaging = loadBytes(t, "tx-04_set_expiry_date_staging.golden")
		stagingFile  = ""

		feeCertificate = uint64(3)
		feeCoefficient = uint64(2)
		feeConstant    = uint64(1)

		feeCertificatePoolRegistration     = uint64(3)
		feeCertificateStakeDelegation      = uint64(3)
		feeCertificateOwnerStakeDelegation = uint64(3)

		feeCertificateVoteCast = uint64(1)
		feeCertificateVotePlan = uint64(2)

		changeAddress = "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"
		expectedTx    = loadBytes(t, "tx-05_finalize_staging.golden")
	)

	tx, err := jcli.TransactionFinalize(
		stdinStaging, stagingFile,
		feeCertificate,
		feeCoefficient,
		feeConstant,
		feeCertificatePoolRegistration,
		feeCertificateStakeDelegation,
		feeCertificateOwnerStakeDelegation,
		feeCertificateVoteCast,
		feeCertificateVotePlan,
		changeAddress,
	)
	ok(t, err)
	equals(t, expectedTx, tx)
}

func ExampleTransactionDataForWitness_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"
	)
	txDataForWitness, err := jcli.TransactionDataForWitness(stdinStaging, stagingFile)

	if err != nil {
		fmt.Printf("TransactionDataForWitness: %s", err)
	} else {
		fmt.Printf("%s", string(txDataForWitness))
	}
	// Ex: 260d1f43854062f558036da376196a35ecf482515dbe88ba4d8109bbdf34c52c
}

func TestTransactionDataForWitness(t *testing.T) {
	var (
		stdinStaging             = loadBytes(t, "tx-05_finalize_staging.golden")
		stagingFile              = ""
		expectedTxDataForWitness = loadBytes(t, "tx-06_data-for-witness.golden")
	)
	txDataForWitness, err := jcli.TransactionDataForWitness(stdinStaging, stagingFile)
	ok(t, err)
	equals(t, expectedTxDataForWitness, txDataForWitness)
}

func ExampleTransactionMakeWitness_stdin() {
	var (
		stdinKey               = []byte("ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4")
		dataForWitness         = "aa27a2258f0df9b25b5ad3969e826bf0877f3d7d912e9ff7b8dcdb3ee6cb3819"
		block0Hash             = "8571f91f59857a5bf033bbe9024d8c360790b3f19f70a72969e9a1f2902b5a71"
		typeWitness            = "account"
		accountSpendingCounter = uint32(0)
		outputFile             = "witness.out"
		inputFileKey           = "" // "witness.secret" - instead of stdinKey use a file that contains Private Key
	)
	witness, err := jcli.TransactionMakeWitness(stdinKey, dataForWitness, block0Hash, typeWitness, accountSpendingCounter, outputFile, inputFileKey)

	if err != nil {
		fmt.Printf("TransactionMakeWitness: %s", err)
	} else {
		fmt.Printf("%v", witness)
	}
}

func TestTransactionMakeWitness(t *testing.T) {
	var (
		stdinKey               []byte                                                               //= loadBytes(t, "private_key_txt.golden")
		dataForWitness         = "aa27a2258f0df9b25b5ad3969e826bf0877f3d7d912e9ff7b8dcdb3ee6cb3819" // strings.TrimSpace(string(loadBytes(t, "tx-06_data-for-witness.golden")))
		block0Hash             = "8571f91f59857a5bf033bbe9024d8c360790b3f19f70a72969e9a1f2902b5a71"
		typeWitness            = "account"
		accountSpendingCounter = uint32(0)
		outputFile             = ""
		inputFileKey           = filePath(t, "private_key_txt.golden")
		expectedWitness        = loadBytes(t, "tx-07_witness_out.golden")
	)

	witness, err := jcli.TransactionMakeWitness(stdinKey, dataForWitness, block0Hash, typeWitness, accountSpendingCounter, outputFile, inputFileKey)
	ok(t, err)
	equals(t, expectedWitness, witness)
}

func ExampleTransactionAddWitness_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"
		witnessFile  = "witness.out"
	)

	tx, err := jcli.TransactionAddWitness(stdinStaging, stagingFile, witnessFile)

	if err != nil {
		fmt.Printf("TransactionAddWitness: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestTransactionAddWitness(t *testing.T) {
	var (
		stdinStaging = loadBytes(t, "tx-05_finalize_staging.golden")
		stagingFile  = ""
		witnessFile  = filePath(t, "tx-07_witness_out.golden")
		expectedTx   = loadBytes(t, "tx-08_add_witness_staging.golden")
	)

	tx, err := jcli.TransactionAddWitness(stdinStaging, stagingFile, witnessFile)
	ok(t, err)
	equals(t, expectedTx, tx)
}

func ExampleTransactionSeal_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"
	)

	tx, err := jcli.TransactionSeal(stdinStaging, stagingFile)

	if err != nil {
		fmt.Printf("TransactionSeal: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestTransactionSeal(t *testing.T) {
	var (
		stdinStaging = loadBytes(t, "tx-08_add_witness_staging.golden")
		stagingFile  = ""
		expectedTx   = loadBytes(t, "tx-09_seal_staging.golden")
	)

	tx, err := jcli.TransactionSeal(stdinStaging, stagingFile)
	ok(t, err)
	equals(t, expectedTx, tx)
}

func ExampleTransactionToMessage_staging() {
	var (
		stdinStaging []byte
		stagingFile  = "tx.staging"
	)

	message, err := jcli.TransactionToMessage(stdinStaging, stagingFile)

	if err != nil {
		fmt.Printf("TransactionToMessage: %s", err)
	} else {
		fmt.Printf("%v", message)
	}
}

func TestTransactionToMessage(t *testing.T) {
	var (
		stdinStaging = loadBytes(t, "tx-09_seal_staging.golden")
		stagingFile  = ""
		expectedMsg  = loadBytes(t, "tx-10_to_message.golden")
	)

	msg, err := jcli.TransactionToMessage(stdinStaging, stagingFile)
	ok(t, err)
	equals(t, expectedMsg, msg)
}

func ExampleTransactionAddCertificate_staging() {
	var (
		stdinStaging      []byte
		stagingFile       = "tx.staging"
		certificateBech32 = "cert1qvqqqqqqqqqqqqqqqqqqq0xsn2eqqqqqqqqqqqqqqyqhs6cc9v2ygmmkm03zmdwh8z2fkx0vfm8xdgpywnxwu2ew8ed4wh5uv63nnjp5f7fzlseqdj6a46q55k2vq9ma6v34cf2dn3qf5edcpz250xxrgszt62zj3e7yysddr33e38dtryfsuqncmp9sdxs3z98zk45mr2u"
	)

	tx, err := jcli.TransactionAddCertificate(stdinStaging, stagingFile, certificateBech32)

	if err != nil {
		fmt.Printf("TransactionAddCertificate: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestTransactioninfo_staging(t *testing.T) {
	var (
		stdinStaging = loadBytes(t, "tx-08_add_witness_staging.golden")
		stagingFile  = ""
		expectedInfo = loadBytes(t, "tx-08_info.golden")

		feeCertificate = uint64(3)
		feeCoefficient = uint64(2)
		feeConstant    = uint64(1)

		feeCertificatePoolRegistration     = uint64(3)
		feeCertificateStakeDelegation      = uint64(3)
		feeCertificateOwnerStakeDelegation = uint64(3)

		feeCertificateVoteCast = uint64(1)
		feeCertificateVotePlan = uint64(2)

		prefix = "ta"

		outputFormat = ""
		outputFile   = ""
	)

	txInfo, err := jcli.TransactionInfo(
		stdinStaging,
		stagingFile,
		feeCertificate,
		feeCoefficient,
		feeConstant,
		feeCertificatePoolRegistration,
		feeCertificateStakeDelegation,
		feeCertificateOwnerStakeDelegation,
		feeCertificateVoteCast,
		feeCertificateVotePlan,
		prefix,
		outputFormat,
		outputFile,
	)
	ok(t, err)
	equals(t, expectedInfo, txInfo)
}

func TestTransactionFragmentID(t *testing.T) {
	var (
		stdinStaging         = loadBytes(t, "tx-09_seal_staging.golden")
		stagingFile          = ""
		expectedTxFragmentID = []byte("4b1ad6d369d9ebe2d0bd1813b6d0d9df3b43dbbc41d3b24eec65e6ae4da8addc\n")
	)
	txFragmentID, err := jcli.TransactionFragmentID(stdinStaging, stagingFile)
	ok(t, err)
	equals(t, expectedTxFragmentID, txFragmentID)
}
