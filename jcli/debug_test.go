package jcli_test

import (
	"fmt"

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
