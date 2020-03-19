package jcli_test

import (
	"fmt"
	"testing"

	"github.com/rinor/jorcli/jcli"
)

func ExampleDebugMessage() {
	var (
		stdinHex  []byte
		inputFile = "testdata/tx-09_to_message.golden"
	)

	dbgMesg, err := jcli.DebugMessage(stdinHex, inputFile)

	if err != nil {
		fmt.Printf("DebugMessage: %s", err)
	} else {
		fmt.Printf("%s", dbgMesg)
	}
	// Output:
	//
	// Transaction(
	//     Transaction {
	//         payload: [
	//             1,
	//             2,
	//         ],
	//         nb_inputs: 1,
	//         nb_outputs: 2,
	//         nb_witnesses: 1,
	//         total_input_value: Ok(
	//             Value(
	//                 100,
	//             ),
	//         ),
	//         total_output_value: Ok(
	//             Value(
	//                 93,
	//             ),
	//         ),
	//     },
	// )
}

func ExampleDebugBlock() {
	var (
		stdinHex  []byte
		inputFile = "testdata/block_PoolRegistration_hex.golden"
	)

	dbgMesg, err := jcli.DebugBlock(stdinHex, inputFile)

	if err != nil {
		fmt.Printf("DebugMessage: %s", err)
	} else {
		fmt.Printf("%s", dbgMesg)
	}
}

func TestDebugBlock_PoolRegistration(t *testing.T) {
	var (
		stdinHex       []byte
		inputFile      = filePath(t, "block_PoolRegistration_hex.golden")
		expectedOutput = loadBytes(t, "block_PoolRegistration_txt.golden")
	)

	output, err := jcli.DebugBlock(stdinHex, inputFile)
	ok(t, err)
	equals(t, expectedOutput, output) // Prod: bytes.Equal(expectedOutput, output)
}
