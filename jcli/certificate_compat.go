package jcli

// CertificateGetStakePoolID - deprecated. Use CertificateGetStakePoolID.
//
//  [STDIN] | jcli certificate show stake-pool-id [--input <FILE_INPUT>] [--output <FILE_OUTPUT>] | [STDOUT]
func CertificateGetStakePoolID(
	stdinCertSigned []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	return CertificateShowStakePoolID(stdinCertSigned, inputFile, outputFile)
}

// CertificateGetVotePlanID - deprecated. Use CertificateShowVotePlanID.
//
//  [STDIN] | jcli certificate show vote-plan-id [--input <FILE_INPUT>] [--output <FILE_OUTPUT>] | [STDOUT]
func CertificateGetVotePlanID(
	stdinCert []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	return CertificateShowVotePlanID(stdinCert, inputFile, outputFile)
}
