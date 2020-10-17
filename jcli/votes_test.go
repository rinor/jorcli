package jcli_test

import (
	"fmt"

	"github.com/rinor/jorcli/jcli"
)

func ExampleVotesCRSGenerate() {
	var (
		seed         = ""
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "crs.key" - will also save output to that file
	)

	privateKey, err := jcli.VotesCRSGenerate(seed, outputFileSk)

	if err != nil {
		fmt.Printf("VotesCRSGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
}

func ExampleVotesCRSGenerate_seed() {
	var (
		seed         = "0000000000000000000000000000000000000000000000000000000000000000"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "crs.key" - will also save output to that file
	)

	privateKey, err := jcli.VotesCRSGenerate(seed, outputFileSk)

	if err != nil {
		fmt.Printf("VotesCRSGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// 049b8327d929a0e45285c04d19c9fffbee065c266b701972922d807228120e43f34ad68ac77f6ec0205fe39f7c5b6055dad973a03464a3a743302de0feaf6ec6d9
}

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
	// 3690baa10b5e2560fc8817908d08b68b7d91c2044dfd2a9d1ce1c3213d3e1dfb
}

func ExampleVotesCommitteeCommunicationKeyToPublic_stdin() {
	var (
		stdinSk      = []byte("3690baa10b5e2560fc8817908d08b68b7d91c2044dfd2a9d1ce1c3213d3e1dfb")
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
	// 049f33cb4f2841121557996c6464b282fb465e40c887b9f3bd3ee8f06a2a0ab9f179253ed7253d02cdf7f22197acbfae0a9690da18fe1ea767961b5febeca2d595
}

func ExampleVotesCommitteeMemberKeyGenerate() {
	var (
		seed      = ""
		crs       = "049b8327d929a0e45285c04d19c9fffbee065c266b701972922d807228120e43f34ad68ac77f6ec0205fe39f7c5b6055dad973a03464a3a743302de0feaf6ec6d9"
		threshold = uint8(2)
		keys      = []string{
			"049f33cb4f2841121557996c6464b282fb465e40c887b9f3bd3ee8f06a2a0ab9f179253ed7253d02cdf7f22197acbfae0a9690da18fe1ea767961b5febeca2d595", // our comm key
			"0409da52c697004d1e4fd743706fa6032266ee256522677ad1e419ce5e07f5451cfd4b6e45d2b27ef263af997ebec93679e57cbe046eae8cddac4692de5ed54527", // other comm key
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
			"049f33cb4f2841121557996c6464b282fb465e40c887b9f3bd3ee8f06a2a0ab9f179253ed7253d02cdf7f22197acbfae0a9690da18fe1ea767961b5febeca2d595",
			"0409da52c697004d1e4fd743706fa6032266ee256522677ad1e419ce5e07f5451cfd4b6e45d2b27ef263af997ebec93679e57cbe046eae8cddac4692de5ed54527",
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
	// 4ad793f7336f8955cdbf36073d1a09acc7bf126faa794697245d933200f9450b
}

func ExampleVotesCommitteeMemberKeyToPublic_stdin() {
	var (
		stdinSk      = []byte("4ad793f7336f8955cdbf36073d1a09acc7bf126faa794697245d933200f9450b")
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
	// 04ce9b53d0aa386d7551e4c5a2534d3683db8cc134c19440f759ee831a29c2ad4c4dc079389d6cbf26150198719da113713c734d230da259b98cead826def89d37
}

func ExampleVotesEncryptingVoteKey() {
	var (
		keys = []string{
			"04ce9b53d0aa386d7551e4c5a2534d3683db8cc134c19440f759ee831a29c2ad4c4dc079389d6cbf26150198719da113713c734d230da259b98cead826def89d37",
			"047d2f10e2b4fb4cc50ab13b1bc41cfbb105808fcc8819e15532d5cf55fcced2de8c253e092ca6e47573faa3afd9b4d908fa7004ff02c05cc5e341504206dc104d",
		}

		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "memberPrivateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.VotesEncryptingVoteKey(keys, outputFileSk)

	if err != nil {
		fmt.Printf("VotesEncryptingVoteKey: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// 04e37aa182ea49206252dbcaec1f3d432bb91ee330b61eeae4bc1554ad626e5e3a41c734e1e9279218db4418564ef14c0858fafd15419dfca57111561cdf5014ba
}
