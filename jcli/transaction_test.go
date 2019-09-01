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
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
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

func ExampleTransactionFinalize_staging() {
	var (
		stdinStaging   []byte
		stagingFile    = "tx.staging"
		feeCertificate = uint64(3)
		feeCoefficient = uint64(2)
		feeConstant    = uint64(1)
		changeAddress  = "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"
	)

	tx, err := jcli.TransactionFinalize(stdinStaging, stagingFile, feeCertificate, feeCoefficient, feeConstant, changeAddress)

	if err != nil {
		fmt.Printf("TransactionFinalize: %s", err)
	} else {
		fmt.Printf("%v", tx)
	}
}

func TestTransactionFinalize(t *testing.T) {
	var (
		stdinStaging   = loadBytes(t, "tx-03_add_output_staging.golden")
		stagingFile    = ""
		feeCertificate = uint64(3)
		feeCoefficient = uint64(2)
		feeConstant    = uint64(1)
		changeAddress  = "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"
		expectedTx     = loadBytes(t, "tx-04_finalize_staging.golden")
	)

	tx, err := jcli.TransactionFinalize(stdinStaging, stagingFile, feeCertificate, feeCoefficient, feeConstant, changeAddress)
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
		fmt.Printf("%v", txDataForWitness)
	}
}

func TestTransactionDataForWitness(t *testing.T) {
	var (
		stdinStaging             = loadBytes(t, "tx-04_finalize_staging.golden")
		stagingFile              = ""
		expectedtxDataForWitness = loadBytes(t, "tx-05_data-for-witness.golden")
	)
	txDataForWitness, err := jcli.TransactionDataForWitness(stdinStaging, stagingFile)
	ok(t, err)
	equals(t, expectedtxDataForWitness, txDataForWitness)
}
