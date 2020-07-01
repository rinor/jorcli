package jcli_test

import (
	"fmt"
	"strings"

	"github.com/rinor/jorcli/jcli"
)

func ExampleCertificateNewStakePoolRegistration() {
	var (
		kesKey              string
		vrfKey              string
		startValidity       = uint64(0)
		managementThreshold = uint8(1)
		owner               = []string{"ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"}
		operator            []string // no operator in this case
		outputFile          = ""     // "" - output to STDOUT only, "stakePool.cert" - will also save output to that file

		taxFixed = uint64(0)
		taxRatio = ""
		taxLimit = uint64(0)

		rewardAccount = ""

		// seed used only for testing and reproducibility
		seed = "0000000000000000000000000000000000000000000000000000000000000000"
	)

	// kes25519-12-sk1qqqqqqqtawlq47yfd829sz0hjpdfs7ywhjkpy2f3z60ptpavadx9egx09zfg75v6qj9xqvea45rfpaxvu0unyt8ecvx2k7yxnlugfx980akf0qapznjg7gpfljj9pp8zqy8p2t76xksa79wxgre5t3qmtu6yme9zh2te3zfn8ctjdawxtx4umds8ljthth6f203qyah5azhpk82tsl7n3wsu97r8qssp5u36z5tdxv3rukje7hqkgmvys5vk7x4f6dds3zwcfqp7gmqxmnjv20ne5swzprjvsngtkl244v4wmw6daent3av5q0k4rqktl73qg6063ref9wa83za2e9r5t6aaqfz67hq0x5283j7f7p0fn6mtsahf5s4pukwcqdfkynzqy3fx36xj40ze6avxped0hcm5mjkdv0hznhqtrxer8u24pk92uc5nq679evw9ddpdzr8alzu5mgmth3fz3epm0k6m693tydsgutw7xjcu675pwdxavz707recqhjgc6w0264fks5r4unt862m25g5pfndnsxye53v8n5f42cmefcdvfmy930yxm8vgz69d34a8dn3v8qhcqthpukvedv63yy62mrt8anr5zwppf8g326n5h38mnzmm7nf4644x5szpgll3uhy32ejmdjmnhn8kvyt2dq8q2673yzznjgzcrruha4arzaeevzgts5xuenc9pcrdnuvku89qdyzlek72kdtszk6dmcsju3v9z5egfxsrwa3eu6tuejhqmrnjrhfmrk98ugnnz6ugayrwtwcrp7y7q98jcd2hprvavxmtesnpxn2df08vy3tk8ypwfh2a43let52tc6q3e9usylyskwyaswkt469am9f5x4v30vslutz7czayx7j4n0g52jyyx7fn57xj43z9xfarmezjew4z4q6v5fnnxh5fy6tu73xe8tremhdexuzkrlzvnq0c97jnm0d8a2uk7ewu5zqf8xaf45tag0487wf98y63ewlfdyuf53awy08heppl7nax3deahunfmr4c2sp9l8dryz7qzmc8djmxm5u6ta45dhlfx40tdwwdhm83xepn8epkw7cfw7hc27tehspcry4nhgjzkd6wdl4wn0ljnsvnwpz37ztfdy0p7yqxkzd8kvvjgwm2748mym45g2wq782ceglj0edd8lfh9g8hjhfsptdlkq2kyldk2udqukk6yzlysf0y08c7lcf9r63ngzg5cpn8kksdy85en3ljv3vl8pse2mcs60l3pyc5alke97a33vfeppn8t75d4xcp2h8u03l3gjwqj74djz09q6g4np9m0u79fupje0kd9vupa9lcs5lrthplyvet85vca96eflj7zz84806yqur4lq6se3wes8wjp9kn5xtuslk66jvx4e9vfu6tp87fw9fyekedmv9pv2579stuf4w846ql54ewdpjcd8j0gwf4ekql0t95zu80y3gc3cmzc2lxvcpt724am87pr4a3m6j6f5tc9zw52zc376ducpc6586a6vakdca6xkzm4s4m59tsvmnxya289e32v60dmkcz3pk03jxe4mrnj6wl2udfgf49eylstc6sjmr5zve0a338l8xye0n4yad5n53ka8t43kuwpycvqe4qje7nx4w4c3zpxrhxkcdvdexvl3w344cf864kekr4rgrds5luv36nn2slj73v2wxwgm4rsvf8ggazy3ytzx6dwvc8986wretny8mua6maf3fdtrdyp0c787d8tjg30xkvwfyll7s4pj0yn8s3lgyg5qsc6v5w0sa0gs0l5s7r7mnn0q6rvthkvkag226yw0xfmake7kjsrsv4klu733fpwrk34lq3zphrzamcge60z425x34l3g7qj49l4d53rfrl6emvr9dxagflldu32550c4uz4wrvxq0fa463cazxvhnl9
	kesPrivateKey, err := jcli.KeyGenerate(seed, "SumEd25519_12", "")
	if err != nil {
		fmt.Printf("kesPrivateKey FAILED: %s\n", err)
		return
	}

	// kes25519-12-pk1j4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4s4zrqja
	kesPublicKey, err := jcli.KeyToPublic(kesPrivateKey, "", "")
	if err != nil {
		fmt.Printf("kesPublicKey FAILED: %s\n", err)
		return
	}

	// vrf_sk1fffu87autxtsae0cttup8p6allqn4yz29ef6uln9lgx75mnzeyqshc9lsp
	vrfPrivateKey, err := jcli.KeyGenerate(seed, "Curve25519_2HashDH", "")
	if err != nil {
		fmt.Printf("vrfPrivateKey FAILED: %s\n", err)
		return
	}

	// vrf_pk1n3n2xwwgx38eyt7rypkttkhgzjjefsqh0hfjxhp9fkwypxn9hqyq6870lk
	vrfPublicKey, err := jcli.KeyToPublic(vrfPrivateKey, "", "")
	if err != nil {
		fmt.Printf("vrfPublicKey FAILED: %s\n", err)
		return
	}

	// convert to string from []byte and remove newline/space
	kesKey = strings.TrimSpace(string(kesPublicKey))
	vrfKey = strings.TrimSpace(string(vrfPublicKey))

	stakePoolCert, err := jcli.CertificateNewStakePoolRegistration(kesKey, vrfKey, startValidity, managementThreshold, owner, operator, taxFixed, taxRatio, taxLimit, rewardAccount, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewStakePoolRegistration: %s", err)
	} else {
		fmt.Printf("%s", stakePoolCert)
	}
	// Output:
	//
	// cert1qvqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqpmj470
}

func ExampleCertificateSign_registration_stdin() {
	var (
		stdinCert      = []byte("cert1qvqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqpmj470")
		signingKeyFile = []string{"testdata/private_key_txt.golden"} // ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
		inputFile      = ""                                          // "" - input from STDIN (stdinCert []byte), "stakePool.cert" - will load the certificate from that file
		outputFile     = ""                                          // "" - output to STDOUT ([]byte) only, "stakePool.signed_cert" - will also save output to that file
	)

	stakePoolSignedCert, err := jcli.CertificateSign(stdinCert, signingKeyFile, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateSign: %s", err)
	} else {
		fmt.Printf("%s", stakePoolSignedCert)
	}
	// Output:
	//
	// signedcert1qvqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqpqprrm7e7sj4ljteluexy39nvrpylyxngaflylzku6uevcq0ygej047hmc82yd89pjrf925tqjm039dkdrksasyfvdl4yjzt4n2n2mqctxtncl7
}

func ExampleCertificateGetStakePoolID_stdin() {
	var (
		stdinCertSigned = []byte("cert1qvqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqpmj470")
		inputFile       = "" // "" - input from STDIN (stdinCertSigned []byte), "stakePool.signed_cert" - will load the certificate from that file
		outputFile      = "" // "" - output to STDOUT ([]byte) only, "stakePool.id" - will also save output to that file
	)

	stakePoolID, err := jcli.CertificateGetStakePoolID(stdinCertSigned, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateGetStakePoolID: %s", err)
	} else {
		fmt.Printf("%s", stakePoolID)
	}
	// Output:
	//
	// 6f93df128a8fa9f7c2f48c8f2590d21f51c0ac05b331536604117ff7218d0bd8
}

func ExampleCertificateNewOwnerStakeDelegation() {
	var (
		stakePoolID = "6f93df128a8fa9f7c2f48c8f2590d21f51c0ac05b331536604117ff7218d0bd8"
		outputFile  = "" // "" - output to STDOUT ([]byte) only, "stakePoolOwnerDelegation.new" - will also save output to that file
	)

	stakeNewDeleg, err := jcli.CertificateNewOwnerStakeDelegation([]string{stakePoolID}, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewStakeDelegation: %s", err)
	} else {
		fmt.Printf("%s", stakeNewDeleg)
	}
	// Output:
	//
	// cert1qgqkly7lz29gl20hct6gere9jrfp75wq4szmxv2nvczpzllhyxxshkqxlmcrj
}

func ExampleCertificateNewStakeDelegation() {
	var (
		stakePoolID = "6f93df128a8fa9f7c2f48c8f2590d21f51c0ac05b331536604117ff7218d0bd8"
		stakeKey    = "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v" // Public Key
		outputFile  = ""                                                                      // "" - output to STDOUT ([]byte) only, "stakePoolDelegation.new" - will also save output to that file
	)

	stakeNewDeleg, err := jcli.CertificateNewStakeDelegation(stakeKey, []string{stakePoolID}, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewStakeDelegation: %s", err)
	} else {
		fmt.Printf("%s", stakeNewDeleg)
	}
	// Output:
	//
	// cert1q9uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4uqt0j0039z5048mu9ayv3ujep5sl28q2cpdnx9fkvpq30lmjrrgtmqn0cwmu
}

func ExampleCertificateSign_delegation_stdin() {
	var (
		stdinCert      = []byte("cert1q9uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4uqt0j0039z5048mu9ayv3ujep5sl28q2cpdnx9fkvpq30lmjrrgtmqn0cwmu")
		signingKeyFile = []string{"testdata/private_key_txt.golden"} // ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
		inputFile      = ""                                          // "" - input from STDIN (stdinCert []byte), "stakePoolDelegation.cert" - will load the certificate from that file
		outputFile     = ""                                          // "" - output to STDOUT ([]byte) only, "stakePoolDelegation.signed_cert" - will also save output to that file
	)

	stakePoolDelegationSignedCert, err := jcli.CertificateSign(stdinCert, signingKeyFile, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateSign: %s", err)
	} else {
		fmt.Printf("%s", stakePoolDelegationSignedCert)
	}
	// Output:
	//
	// signedcert1q9uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4uqt0j0039z5048mu9ayv3ujep5sl28q2cpdnx9fkvpq30lmjrrgtmqqctzczvu6e3v65m40n40c3y2pnu4vhd888dygkrtnfm0ts92fe50jy0h0ugh6wlvgy4xvr3lz4uuqzg2xgu6vv8tr24jrwhg0l09klp5wvwzl5
}

func ExampleCertificateNewStakePoolRetirement() {
	var (
		poolID         = "6f93df128a8fa9f7c2f48c8f2590d21f51c0ac05b331536604117ff7218d0bd8"
		retirementTime = uint64(0)
		outputFile     = ""
	)

	stakePoolRetirementCert, err := jcli.CertificateNewStakePoolRetirement(poolID, retirementTime, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewStakePoolRetirement: %s", err)
	} else {
		fmt.Printf("%s", stakePoolRetirementCert)
	}
	// Output:
	//
	// cert1q3he8hcj3286na7z7jxg7fvs6g04rs9vqkenz5mxqsghlaep359asqqqqqqqqqqqqqlpzcht
}

func ExampleCertificateNewVotePlan_stdin() {
	var (
		stdinConfig = []byte(`
{
  "payload_type": "public",
  "vote_start": {
    "epoch": 0,
    "slot_id": 100
  },
  "vote_end": {
    "epoch": 0,
    "slot_id": 200
  },
  "committee_end": {
    "epoch": 0,
    "slot_id": 300
  },
  "proposals": [
    {
      "external_id": "adb92757155d09e7f92c9f100866a92dddd35abd2a789a44ae19ab9a1dbc3280",
      "options": 3,
      "action": "off_chain"
    },
    {
      "external_id": "6778d37161c3962fe62c9fa8a31a55bccf6ec2d1ea254a467d8cd994709fc404",
      "options": 3,
      "action": "off_chain"
    }
  ]
}
`)
		inputFile  = ""
		outputFile = ""
	)

	votePlanCert, err := jcli.CertificateNewVotePlan(stdinConfig, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewVotePlan: %s", err)
	} else {
		fmt.Printf("%s", votePlanCert)
	}
	// Output:
	//
	// cert1qcqqqqqqqqqqqeqqqqqqqqqqqryqqqqqqqqqqqfvqyp2mwf82u246z08lykf7yqgv65jmhwnt27j57y6gjhpn2u6rk7r9qqrqpnh35m3v8pevtlx9j063gc62k7v7mkz684z2jjx0kxdn9rsnlzqgqcqx4dhc9
}

func ExampleCertificateGetVotePlanID_stdin() {
	var (
		stdinCert  = []byte("cert1qcqqqqqqqqqqqeqqqqqqqqqqqryqqqqqqqqqqqfvqyp2mwf82u246z08lykf7yqgv65jmhwnt27j57y6gjhpn2u6rk7r9qqrqpnh35m3v8pevtlx9j063gc62k7v7mkz684z2jjx0kxdn9rsnlzqgqcqx4dhc9")
		inputFile  = "" // "" - input from STDIN (stdinCert []byte), "votePlan.cert" - will load the certificate from that file
		outputFile = "" // "" - output to STDOUT ([]byte) only, "votePlan.id" - will also save output to that file
	)

	votePlanID, err := jcli.CertificateGetVotePlanID(stdinCert, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateGetVotePlanID: %s", err)
	} else {
		fmt.Printf("%s", votePlanID)
	}
	// Output:
	//
	// 7bfc5132cfd4aa459491199f069aa9dc19e30fd372e1873b62cb0b6700ac0ec2

}

func ExampleCertificateNewVoteCast() {
	var (
		votePlanID    = "7bfc5132cfd4aa459491199f069aa9dc19e30fd372e1873b62cb0b6700ac0ec2"
		proposalIndex = uint8(0)
		choice        = uint8(1)
		privacy       = "public"
		outputFile    = ""
	)

	voteCastCert, err := jcli.CertificateNewVoteCast(votePlanID, proposalIndex, choice, privacy, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewVoteCast: %s", err)
	} else {
		fmt.Printf("%s", voteCastCert)
	}
	// Output:
	//
	// cert1qaalc5fjel2253v5jyve7p5648wpncc06dewrpemvt9skecq4s8vyqqpqy7ct20d
}

func ExampleCertificateNewVoteTally() {
	var (
		votePlanID = "7bfc5132cfd4aa459491199f069aa9dc19e30fd372e1873b62cb0b6700ac0ec2"
		outputFile = ""
	)

	voteTallyCert, err := jcli.CertificateNewVoteTally(votePlanID, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewVoteTally: %s", err)
	} else {
		fmt.Printf("%s", voteTallyCert)
	}
	// Output:
	//
	// cert1ppalc5fjel2253v5jyve7p5648wpncc06dewrpemvt9skecq4s8vyqgzcncz3
}

func ExampleCertificateSign_retirement_stdin() {
	var (
		stdinCert      = []byte("cert1q3he8hcj3286na7z7jxg7fvs6g04rs9vqkenz5mxqsghlaep359asqqqqqqqqqqqqqlpzcht")
		signingKeyFile = []string{"testdata/private_key_txt.golden"} // ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
		inputFile      = ""                                          // "" - input from STDIN (stdinCert []byte), "stakePoolRetirement.cert" - will load the certificate from that file
		outputFile     = ""                                          // "" - output to STDOUT ([]byte) only, "stakePoolRetirement.signed_cert" - will also save output to that file
	)

	stakePoolRetirementSignedCert, err := jcli.CertificateSign(stdinCert, signingKeyFile, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateSign: %s", err)
	} else {
		fmt.Printf("%s", stakePoolRetirementSignedCert)
	}
	// Output:
	//
	// signedcert1q3he8hcj3286na7z7jxg7fvs6g04rs9vqkenz5mxqsghlaep359asqqqqqqqqqqqqqqsq49x23cxqe3wtpawjnjhfu29gezruvd8uyemd2wdm3g6w55l0pswg76pchp70r5dk343fxfctphzweuy04sd79wdmqcznepfzxkf25pqkxc6x0
}

func ExampleCertificatePrint_registrationSigned_stdin() {
	var (
		stdinCert = []byte("signedcert1qvqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqpqprrm7e7sj4ljteluexy39nvrpylyxngaflylzku6uevcq0ygej047hmc82yd89pjrf925tqjm039dkdrksasyfvdl4yjzt4n2n2mqctxtncl7")
		inputFile = "" // "" - input from STDIN (stdinCert []byte), "stakePool.signed_cert" - will load the certificate from that file
	)

	certPrint, err := jcli.CertificatePrint(stdinCert, inputFile)

	if err != nil {
		fmt.Printf("CertificatePrint: %s", err)
	} else {
		fmt.Printf("%s", certPrint)
	}
	// Output:
	//
	// Certificate(PoolRegistration(PoolRegistration { serial: 0, start_validity: TimeOffsetSeconds(DurationSeconds(0)), permissions: PoolPermissions(1), owners: [786b182b14446f76dbe22db5d738949b19ec4ece66a02474ccee2b2e3e5b575e], operators: [], rewards: TaxType { fixed: Value(0), ratio: Ratio { numerator: 0, denominator: 1 }, max_limit: None }, reward_account: None, keys: GenesisPraosLeader { kes_public_key: 954798c34404bd28528e7c4241ad1c63989dab19130e0278d84b069a11114e2b, vrf_public_key: 9c66a339c8344f922fc3206cb5dae814a594c0177dd3235c254d9c409a65b808 } }))
}

func ExampleCertificatePrint_delegationSigned_stdin() {
	var (
		stdinCert = []byte("signedcert1q9uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4uqt0j0039z5048mu9ayv3ujep5sl28q2cpdnx9fkvpq30lmjrrgtmqqctzczvu6e3v65m40n40c3y2pnu4vhd888dygkrtnfm0ts92fe50jy0h0ugh6wlvgy4xvr3lz4uuqzg2xgu6vv8tr24jrwhg0l09klp5wvwzl5")
		inputFile = "" // "" - input from STDIN (stdinCert []byte), "stakePoolDelegation.signed_cert" - will load the certificate from that file
	)

	certPrint, err := jcli.CertificatePrint(stdinCert, inputFile)

	if err != nil {
		fmt.Printf("CertificatePrint: %s", err)
	} else {
		fmt.Printf("%s", certPrint)
	}
	// Output:
	//
	// Certificate(StakeDelegation(StakeDelegation { account_id: UnspecifiedAccountIdentifier([120, 107, 24, 43, 20, 68, 111, 118, 219, 226, 45, 181, 215, 56, 148, 155, 25, 236, 78, 206, 102, 160, 36, 116, 204, 238, 43, 46, 62, 91, 87, 94]), delegation: Full($ hash_ty(0x6f93df128a8fa9f7c2f48c8f2590d21f51c0ac05b331536604117ff7218d0bd8)) }))
}

func ExampleCertificatePrint_retirementSigned_stdin() {
	var (
		stdinCert = []byte("signedcert1q3he8hcj3286na7z7jxg7fvs6g04rs9vqkenz5mxqsghlaep359asqqqqqqqqqqqqqqsq49x23cxqe3wtpawjnjhfu29gezruvd8uyemd2wdm3g6w55l0pswg76pchp70r5dk343fxfctphzweuy04sd79wdmqcznepfzxkf25pqkxc6x0")
		inputFile = "" // "" - input from STDIN (stdinCert []byte), "stakePoolRetirement.signed_cert" - will load the certificate from that file
	)

	certPrint, err := jcli.CertificatePrint(stdinCert, inputFile)

	if err != nil {
		fmt.Printf("CertificatePrint: %s", err)
	} else {
		fmt.Printf("%s", certPrint)
	}
	// Output:
	//
	// Certificate(PoolRetirement(PoolRetirement { pool_id: $ hash_ty(0x6f93df128a8fa9f7c2f48c8f2590d21f51c0ac05b331536604117ff7218d0bd8), retirement_time: TimeOffsetSeconds(DurationSeconds(0)) }))
}
