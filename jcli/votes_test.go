package jcli_test

import (
	"fmt"

	"github.com/rinor/jorcli/jcli"
)

func ExampleVotesCommitteeCommunicationKeyGenerate() {
	var (
		seed         = ""
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "communicationPrivateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.VotesCommitteeCommunicationKeyGenerate(seed, outputFileSk)

	if err != nil {
		fmt.Printf("VotesCommitteeCommunicationKeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
}

func ExampleVotesCommitteeCommunicationKeyGenerate_seed() {
	var (
		seed         = "0000000000000000000000000000000000000000000000000000000000000011"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "communicationPrivateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.VotesCommitteeCommunicationKeyGenerate(seed, outputFileSk)

	if err != nil {
		fmt.Printf("VotesCommitteeCommunicationKeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// p256k1_vcommsk1hqg47t39a4ahvjwr99hkn4h3hxjvacezgjks7xm3tls9p4zxlvqqpjuqr8
}

func ExampleVotesCommitteeCommunicationKeyToPublic_stdin() {
	var (
		stdinSk      = []byte("p256k1_vcommsk1hqg47t39a4ahvjwr99hkn4h3hxjvacezgjks7xm3tls9p4zxlvqqpjuqr8")
		inputFileSk  = "" // "" - input from STDIN (stdinSk []byte), "communicationPrivateKey.sk" - will load the private key from that file
		outputFilePk = "" // "" - output to STDOUT ([]byte) only, "communicationPublicKey.pk" - will also save the public key to that file
	)

	publicKey, err := jcli.VotesCommitteeCommunicationKeyToPublic(stdinSk, inputFileSk, outputFilePk)

	if err != nil {
		fmt.Printf("VotesCommitteeCommunicationKeyToPublic: %s", err)
	} else {
		fmt.Printf("%s", string(publicKey))
	}
	// Output:
	//
	// p256k1_vcommpk16tkulvdqjmvcwafnh77x6j4xywu52srzx23jl9xyytxtuf845q6snw8k27
}

func ExampleVotesCommitteeMemberKeyGenerate() {
	var (
		seed      = ""
		crs       = "049b8327d929a0e45285c04d19c9fffbee065c266b701972922d807228120e43f34ad68ac77f6ec0205fe39f7c5b6055dad973a03464a3a743302de0feaf6ec6d9"
		threshold = uint8(2)
		keys      = []string{
			"p256k1_vcommpk16tkulvdqjmvcwafnh77x6j4xywu52srzx23jl9xyytxtuf845q6snw8k27", // our comm key
			"p256k1_vcommpk16tkulvdqjmvcwafnh77x6j4xywu52srzx23jl9xyytxtuf845q6snw8k27", // other comm key
		}
		index        = uint8(0) // index of our own key in the list of keys
		outputFileSk = ""       // "" - output to STDOUT ([]byte) only, "memberPrivateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.VotesCommitteeMemberKeyGenerate(crs, threshold, keys, index, seed, outputFileSk)

	if err != nil {
		fmt.Printf("VotesCommitteeMemberKeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
}

func ExampleVotesCommitteeMemberKeyGenerate_seed() {
	var (
		seed      = "0000000000000000000000000000000000000000000000000000000000000022"
		crs       = "049b8327d929a0e45285c04d19c9fffbee065c266b701972922d807228120e43f34ad68ac77f6ec0205fe39f7c5b6055dad973a03464a3a743302de0feaf6ec6d9"
		threshold = uint8(2)
		keys      = []string{
			"p256k1_vcommpk16tkulvdqjmvcwafnh77x6j4xywu52srzx23jl9xyytxtuf845q6snw8k27",
			"p256k1_vcommpk16tkulvdqjmvcwafnh77x6j4xywu52srzx23jl9xyytxtuf845q6snw8k27",
		}
		index        = uint8(0)
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "memberPrivateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.VotesCommitteeMemberKeyGenerate(crs, threshold, keys, index, seed, outputFileSk)

	if err != nil {
		fmt.Printf("VotesCommitteeMemberKeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// p256k1_membersk1kfy7mqqurem2z5y5leg5uhxkvp4pczqghl8h9dj4yptqjx9rzspqq0gs7p
}

func ExampleVotesCommitteeMemberKeyToPublic_stdin() {
	var (
		stdinSk      = []byte("p256k1_membersk1kfy7mqqurem2z5y5leg5uhxkvp4pczqghl8h9dj4yptqjx9rzspqq0gs7p")
		inputFileSk  = "" // "" - input from STDIN (stdinSk []byte), "memberPrivateKey.sk" - will load the private key from that file
		outputFilePk = "" // "" - output to STDOUT ([]byte) only, "memberPublicKey.pk" - will also save the public key to that file
	)

	publicKey, err := jcli.VotesCommitteeMemberKeyToPublic(stdinSk, inputFileSk, outputFilePk)

	if err != nil {
		fmt.Printf("VotesCommitteeMemberKeyToPublic: %s", err)
	} else {
		fmt.Printf("%s", string(publicKey))
	}
	// Output:
	//
	// p256k1_memberpk1s3aznnq2qtmktdt9gxsv6mt67et0nwymx9hufz0026pp0x8td5nspsge8e
}

func ExampleVotesElectionKey() {
	var (
		keys = []string{
			"p256k1_memberpk1s3aznnq2qtmktdt9gxsv6mt67et0nwymx9hufz0026pp0x8td5nspsge8e",
			"p256k1_memberpk1s3aznnq2qtmktdt9gxsv6mt67et0nwymx9hufz0026pp0x8td5nspsge8e",
		}

		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "encvote.pk" - will also save output to that file
	)

	privateKey, err := jcli.VotesElectionKey(keys, outputFileSk)

	if err != nil {
		fmt.Printf("VotesElectionKey: %s - %s", err, string(privateKey))
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// p256k1_votepk1d6rdft5nr3pgtn8t4xseat8nv2nlqeda73khy5084fz4wp7efe2skdy6kp
}
