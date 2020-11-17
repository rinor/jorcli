package jcli

// VotesEncryptingVoteKey - deprecated. Use VotesEncryptingKey.
//
//  jcli votes encrypting-key --keys=<member-keys>... [OUTPUT_FILE] | [STDOUT]
func VotesEncryptingVoteKey(
	keys []string,
	outputFileSk string,
) ([]byte, error) {
	return VotesEncryptingKey(keys, outputFileSk)
}
