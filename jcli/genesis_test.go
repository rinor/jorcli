package jcli_test

import (
	"testing"

	"github.com/rinor/jorcli/jcli"
)

// TODO: Improve test
func TestGenesisInit(t *testing.T) {
	_, err := jcli.GenesisInit()
	ok(t, err)
}

func TestGenesisHash_file(t *testing.T) {
	var (
		block0Bin          []byte
		inputFile          = filePath(t, "genesis_block0_bin.golden")
		expectedBlock0Hash = []byte("8571f91f59857a5bf033bbe9024d8c360790b3f19f70a72969e9a1f2902b5a71\n")
	)

	genesisBlock0Hash, err := jcli.GenesisHash(block0Bin, inputFile)
	ok(t, err)
	equals(t, expectedBlock0Hash, genesisBlock0Hash) // Prod: bytes.Equal(expectedBlock0Hash, genesisBlock0Hash)
}

func TestGenesisHash_stdin(t *testing.T) {
	var (
		block0Bin          = loadBytes(t, "genesis_block0_bin.golden")
		inputFile          = ""
		expectedBlock0Hash = []byte("8571f91f59857a5bf033bbe9024d8c360790b3f19f70a72969e9a1f2902b5a71\n")
	)

	genesisBlock0Hash, err := jcli.GenesisHash(block0Bin, inputFile)
	ok(t, err)
	equals(t, expectedBlock0Hash, genesisBlock0Hash) // Prod: bytes.Equal(expectedBlock0Hash, genesisBlock0Hash)
}

func TestGenesisEncode_file(t *testing.T) {
	var (
		block0Txt         []byte
		inputFile         = filePath(t, "genesis_block0_txt.golden")
		outputFile        = ""
		expectedBlock0Bin = loadBytes(t, "genesis_block0_bin.golden")
	)

	genesisBlock0Bin, err := jcli.GenesisEncode(block0Txt, inputFile, outputFile)
	ok(t, err)
	equals(t, expectedBlock0Bin, genesisBlock0Bin) // Prod: bytes.Equal(expectedBlock0Bin, genesisBlock0Bin)
}

func TestGenesisEncode_stdin(t *testing.T) {
	var (
		block0Txt         = loadBytes(t, "genesis_block0_txt.golden")
		inputFile         = ""
		outputFile        = ""
		expectedBlock0Bin = loadBytes(t, "genesis_block0_bin.golden")
	)

	genesisBlock0Bin, err := jcli.GenesisEncode(block0Txt, inputFile, outputFile)
	ok(t, err)
	equals(t, expectedBlock0Bin, genesisBlock0Bin) // Prod: bytes.Equal(expectedBlock0Bin, genesisBlock0Bin)
}

func TestGenesisDecode_file(t *testing.T) {
	var (
		block0Bin         []byte
		inputFile         = filePath(t, "genesis_block0_bin.golden")
		outputFile        = ""
		expectedBlock0Txt = loadBytes(t, "genesis_block0_txt.golden")
	)

	genesisBlock0Txt, err := jcli.GenesisDecode(block0Bin, inputFile, outputFile)
	ok(t, err)
	equals(t, expectedBlock0Txt, genesisBlock0Txt) // Prod: bytes.Equal(expectedBlock0Txt, genesisBlock0Txt)
}

func TestGenesisDecode_stdin(t *testing.T) {
	var (
		block0Bin         = loadBytes(t, "genesis_block0_bin.golden")
		inputFile         = ""
		outputFile        = ""
		expectedBlock0Txt = loadBytes(t, "genesis_block0_txt.golden")
	)

	genesisBlock0Txt, err := jcli.GenesisDecode(block0Bin, inputFile, outputFile)
	ok(t, err)
	equals(t, expectedBlock0Txt, genesisBlock0Txt) // Prod: bytes.Equal(expectedBlock0Txt, genesisBlock0Txt)
}
