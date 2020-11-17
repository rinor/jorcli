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
	// p256k1_vcommsk1x6gt4ggttcjkplygz7gg6z9k3d7erssyfh7j48guu8pjz0f7rhasg408yn
}

func ExampleVotesCommitteeCommunicationKeyToPublic_stdin() {
	var (
		stdinSk      = []byte("p256k1_vcommsk1x6gt4ggttcjkplygz7gg6z9k3d7erssyfh7j48guu8pjz0f7rhasg408yn")
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
	// p256k1_vcommpk1qj0n8j609pq3y92hn9kxge9jsta5vhjqezrmnuaa8m50q632p2ulz7f98mtj20gzehmlygvh4jl6uz5kjrdp3ls75anevx6la0k294v4fx9ckr
}

func ExampleVotesCommitteeMemberKeyGenerate() {
	var (
		seed      = ""
		crs       = "049b8327d929a0e45285c04d19c9fffbee065c266b701972922d807228120e43f34ad68ac77f6ec0205fe39f7c5b6055dad973a03464a3a743302de0feaf6ec6d9"
		threshold = uint8(2)
		keys      = []string{
			"p256k1_vcommpk1qj0n8j609pq3y92hn9kxge9jsta5vhjqezrmnuaa8m50q632p2ulz7f98mtj20gzehmlygvh4jl6uz5kjrdp3ls75anevx6la0k294v4fx9ckr", // our comm key
			"p256k1_vcommpk1qj0n8j609pq3y92hn9kxge9jsta5vhjqezrmnuaa8m50q632p2ulz7f98mtj20gzehmlygvh4jl6uz5kjrdp3ls75anevx6la0k294v4fx9ckr", // other comm key
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
			"p256k1_vcommpk1qj0n8j609pq3y92hn9kxge9jsta5vhjqezrmnuaa8m50q632p2ulz7f98mtj20gzehmlygvh4jl6uz5kjrdp3ls75anevx6la0k294v4fx9ckr",
			"p256k1_vcommpk1qj0n8j609pq3y92hn9kxge9jsta5vhjqezrmnuaa8m50q632p2ulz7f98mtj20gzehmlygvh4jl6uz5kjrdp3ls75anevx6la0k294v4fx9ckr",
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
	// p256k1_membersk1ftte8aend7y4tndlxcrn6xsf4nrm7yn04fu5d9eytkfnyq8eg59sfy4kuf
}

func ExampleVotesCommitteeMemberKeyToPublic_stdin() {
	var (
		stdinSk      = []byte("p256k1_membersk1ftte8aend7y4tndlxcrn6xsf4nrm7yn04fu5d9eytkfnyq8eg59sfy4kuf")
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
	// p256k1_memberpk1qn8fk57s4gux6a23unz6y56dx6pahrxpxnqegs8ht8hgxx3fc2k5cnwq0yuf6m9lyc2srxr3nks3xufuwdxjxrdztxuce6kcym0038fhcxmsn6
}

func ExampleVotesEncryptingKey() {
	var (
		keys = []string{
			"p256k1_memberpk1qn8fk57s4gux6a23unz6y56dx6pahrxpxnqegs8ht8hgxx3fc2k5cnwq0yuf6m9lyc2srxr3nks3xufuwdxjxrdztxuce6kcym0038fhcxmsn6",
			"p256k1_memberpk1qn8fk57s4gux6a23unz6y56dx6pahrxpxnqegs8ht8hgxx3fc2k5cnwq0yuf6m9lyc2srxr3nks3xufuwdxjxrdztxuce6kcym0038fhcxmsn6",
		}

		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "encvote.pk" - will also save output to that file
	)

	privateKey, err := jcli.VotesEncryptingKey(keys, outputFileSk)

	if err != nil {
		fmt.Printf("VotesEncryptingKey: %s - %s", err, string(privateKey))
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// p256k1_votepk1qsm95agaa6j6cs07tyu40t0sm3clzssnhxjusvjqezc5wjtlv3gjlrvr0ta0g64s6j52dddu4vz9x3nfs7mlf9kf6qp7rwtzjj05thp4tpqhe3
}
