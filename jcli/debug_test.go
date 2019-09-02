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
	//     AuthenticatedTransaction {
	//         transaction: Transaction {
	//             inputs: [
	//                 Input {
	//                     index_or_account: 255,
	//                     value: Value(
	//                         100,
	//                     ),
	//                     input_ptr: [
	//                         120,
	//                         107,
	//                         24,
	//                         43,
	//                         20,
	//                         68,
	//                         111,
	//                         118,
	//                         219,
	//                         226,
	//                         45,
	//                         181,
	//                         215,
	//                         56,
	//                         148,
	//                         155,
	//                         25,
	//                         236,
	//                         78,
	//                         206,
	//                         102,
	//                         160,
	//                         36,
	//                         116,
	//                         204,
	//                         238,
	//                         43,
	//                         46,
	//                         62,
	//                         91,
	//                         87,
	//                         94,
	//                     ],
	//                 },
	//             ],
	//             outputs: [
	//                 Output {
	//                     address: Address(
	//                         Test,
	//                         Account(
	//                             a670453a1990552003ccca84c6f9974f69561709cb2fabd3d1594e89cc9c0b46,
	//                         ),
	//                     ),
	//                     value: Value(
	//                         50,
	//                     ),
	//                 },
	//                 Output {
	//                     address: Address(
	//                         Test,
	//                         Account(
	//                             786b182b14446f76dbe22db5d738949b19ec4ece66a02474ccee2b2e3e5b575e,
	//                         ),
	//                     ),
	//                     value: Value(
	//                         43,
	//                     ),
	//                 },
	//             ],
	//             extra: NoExtra,
	//         },
	//         witnesses: [
	//             Account(
	//                 ffd45e56ee8a4709f33f9fea91a4bf3af32608c896e4a606b47eccc6551bf6b9871728c59b316063c6041b4faf979ca8d842570cc5d8d4db88e2dbd3d3a8170d,
	//             ),
	//         ],
	//     },
	// )
}
