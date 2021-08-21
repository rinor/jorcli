package jcli_test

import (
	"fmt"
	"testing"

	"github.com/rinor/jorcli/jcli"
)

func ExampleKeyGenerate_noseed() {
	var (
		seed         = ""
		keyType      = "ed25519extended"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "privateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.KeyGenerate(seed, keyType, outputFileSk)

	if err != nil {
		fmt.Printf("KeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
}

func ExampleKeyGenerate_ed25519Extended_seed() {
	var (
		seed         = "0000000000000000000000000000000000000000000000000000000000000000"
		keyType      = "ed25519extended"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "privateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.KeyGenerate(seed, keyType, outputFileSk)

	if err != nil {
		fmt.Printf("KeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
}

func TestKeyToPublic_file(t *testing.T) {
	var (
		stdinSk           []byte
		inputFileSk       = filePath(t, "private_key_txt.golden")
		outputFilePk      = ""
		expectedPublicKey = loadBytes(t, "public_key_txt.golden")
	)

	publicKey, err := jcli.KeyToPublic(stdinSk, inputFileSk, outputFilePk)
	ok(t, err)
	equals(t, expectedPublicKey, publicKey) // Prod: bytes.Equal(expectedPublicKey, publicKey)
}

func ExampleKeyToPublic_stdin() {
	var (
		stdinSk      = []byte("ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4")
		inputFileSk  = "" // "" - input from STDIN (stdinSk []byte), "privateKey.sk" - will load the private key from that file
		outputFilePk = "" // "" - output to STDOUT ([]byte) only, "publicKey.pk" - will also save the public key to that file
	)

	publicKey, err := jcli.KeyToPublic(stdinSk, inputFileSk, outputFilePk)

	if err != nil {
		fmt.Printf("KeyToPublic: %s", err)
	} else {
		fmt.Printf("%s", string(publicKey))
	}
	// Output:
	//
	// ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v
}

func TestKeyToBytes_file(t *testing.T) {
	var (
		stdinSk                 []byte
		outputFile              = ""
		inputFileSk             = filePath(t, "private_key_txt.golden")
		expectedPrivateKeyBytes = loadBytes(t, "private_key_bytes.golden")
	)

	privateKeyBytes, err := jcli.KeyToBytes(stdinSk, outputFile, inputFileSk)
	ok(t, err)
	equals(t, expectedPrivateKeyBytes, privateKeyBytes) // Prod: bytes.Equal(expectedPrivateKeyBytes, privateKeyBytes)
}

func ExampleKeyToBytes_stdin() {
	var (
		stdinSk     = []byte("ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4")
		outputFile  = ""
		inputFileSk = ""
	)

	privateKeyBytes, err := jcli.KeyToBytes(stdinSk, outputFile, inputFileSk)

	if err != nil {
		fmt.Printf("KeyToBytes: %s", err)
	} else {
		fmt.Printf("%s", string(privateKeyBytes))
	}
	// Output:
	//
	// 70b8e0ada0f13d90405d6ae55386bd28bdd219b8a08ded1aa836efcc8b770d47da41597c5157488d7724e03fb8d84a376a43b8f41518a11cc387b669b2ee6586
}

func TestKeyFromBytes_file(t *testing.T) {
	var (
		stdinSk            []byte
		keyType            = "ed25519extended"
		inputFile          = filePath(t, "private_key_bytes.golden")
		outputFileSk       = ""
		expectedPrivateKey = loadBytes(t, "private_key_txt.golden")
	)

	privateKey, err := jcli.KeyFromBytes(stdinSk, keyType, inputFile, outputFileSk)
	ok(t, err)
	equals(t, expectedPrivateKey, privateKey) // Prod: bytes.Equal(expectedPrivateKey, privateKey)
}

func ExampleKeyFromBytes_stdin() {
	var (
		stdinSk      = []byte("70b8e0ada0f13d90405d6ae55386bd28bdd219b8a08ded1aa836efcc8b770d47da41597c5157488d7724e03fb8d84a376a43b8f41518a11cc387b669b2ee6586")
		keyType      = "ed25519extended"
		inputFile    = ""
		outputFileSk = ""
	)

	privateKey, err := jcli.KeyFromBytes(stdinSk, keyType, inputFile, outputFileSk)

	if err != nil {
		fmt.Printf("KeyFromBytes: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
}

func TestKeySign_file(t *testing.T) {
	var (
		stdinData     []byte
		inputFileSk   = filePath(t, "private_key_txt.golden")
		inputFileData = filePath(t, "key_sign_txt.golden")
		outputFileSig = ""
		expectedSig   = loadBytes(t, "key_sign_signature.golden")
	)

	sig, err := jcli.KeySign(stdinData, inputFileSk, inputFileData, outputFileSig)
	ok(t, err)
	equals(t, expectedSig, sig) // Prod: bytes.Equal(expectedSig, sig)
}

func TestKeySign_stdin(t *testing.T) {
	var (
		stdinData     = loadBytes(t, "key_sign_txt.golden")
		inputFileSk   = filePath(t, "private_key_txt.golden")
		inputFileData = ""
		outputFileSig = ""
		expectedSig   = loadBytes(t, "key_sign_signature.golden")
	)

	sig, err := jcli.KeySign(stdinData, inputFileSk, inputFileData, outputFileSig)
	ok(t, err)
	equals(t, expectedSig, sig) // Prod: bytes.Equal(expectedSig, sig)
}

func TestKeyVerify_file(t *testing.T) {
	var (
		stdinData      []byte
		inputFilePk    = filePath(t, "public_key_txt.golden")
		inputFileData  = filePath(t, "key_sign_txt.golden")
		inputFileSig   = filePath(t, "key_sign_signature.golden")
		expectedVerify = []byte("Success\n")
	)

	verify, err := jcli.KeyVerify(stdinData, inputFilePk, inputFileSig, inputFileData)
	ok(t, err)
	equals(t, expectedVerify, verify) // Prod: bytes.Equal(expectedVerify, verify)
}

func TestKeyVerify_stdin(t *testing.T) {
	var (
		stdinData      = loadBytes(t, "key_sign_txt.golden")
		inputFilePk    = filePath(t, "public_key_txt.golden")
		inputFileData  = ""
		inputFileSig   = filePath(t, "key_sign_signature.golden")
		expectedVerify = []byte("Success\n")
	)

	verify, err := jcli.KeyVerify(stdinData, inputFilePk, inputFileSig, inputFileData)
	ok(t, err)
	equals(t, expectedVerify, verify) // Prod: bytes.Equal(expectedVerify, verify)
}

func ExampleKeyDerive_stdin() {
	var (
		stdinParentKey     = []byte("xprv1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtp5lqlnmu4238paf3w5h03ej6zqdev8jngzgudjkjykx2vlr9mn6a5ec2azq")
		inputFileParentKey = "" // "" - input from STDIN (stdinParentKey []byte), "parent.key" - will load the private key from that file
		indexChildKey      = uint32(0)
		outputFileChildKey = "" // "" - output to STDOUT ([]byte) only, "child.key" - will also save the derived key to that file
	)

	childKey, err := jcli.KeyDerive(stdinParentKey, inputFileParentKey, indexChildKey, outputFileChildKey)

	if err != nil {
		fmt.Printf("KeyDerive: %s", err)
	} else {
		fmt.Printf("%s", string(childKey))
	}
	// Output:
	//
	// xprv10z6xdkgmvpqqvfgw4y37fy0sqwv66gkc04hee4w7phf0drnhp4rkvk5q0w0fjjj378r9296536av9g7ypcl8w7exeef3ud4p24yhl5870s27hrfs7vglwfhus8a32wkqknlnjj456h5gjwmtslyqka70accut0ax
}

func ExampleKeyGenerate_ed25519_seed() {
	var (
		seed         = "0000000000000000000000000000000000000000000000000000000000000000"
		keyType      = "ed25519"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "privateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.KeyGenerate(seed, keyType, outputFileSk)

	if err != nil {
		fmt.Printf("KeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// ed25519_sk1w6uwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhphrs8yjml0
}

func ExampleKeyGenerate_ed25519bip32_seed() {
	var (
		seed         = "0000000000000000000000000000000000000000000000000000000000000000"
		keyType      = "ed25519bip32"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "privateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.KeyGenerate(seed, keyType, outputFileSk)

	if err != nil {
		fmt.Printf("KeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// xprv1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtp5lqlnmu4238paf3w5h03ej6zqdev8jngzgudjkjykx2vlr9mn6a5ec2azq
}

func ExampleKeyGenerate_sumed2551912_seed() {
	var (
		seed         = "0000000000000000000000000000000000000000000000000000000000000000"
		keyType      = "sumed25519_12"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "privateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.KeyGenerate(seed, keyType, outputFileSk)

	if err != nil {
		fmt.Printf("KeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// kes25519-12-sk1qqqqqqqtawlq47yfd829sz0hjpdfs7ywhjkpy2f3z60ptpavadx9egx09zfg75v6qj9xqvea45rfpaxvu0unyt8ecvx2k7yxnlugfx980akf0qapznjg7gpfljj9pp8zqy8p2t76xksa79wxgre5t3qmtu6yme9zh2te3zfn8ctjdawxtx4umds8ljthth6f203qyah5azhpk82tsl7n3wsu97r8qssp5u36z5tdxv3rukje7hqkgmvys5vk7x4f6dds3zwcfqp7gmqxmnjv20ne5swzprjvsngtkl244v4wmw6daent3av5q0k4rqktl73qg6063ref9wa83za2e9r5t6aaqfz67hq0x5283j7f7p0fn6mtsahf5s4pukwcqdfkynzqy3fx36xj40ze6avxped0hcm5mjkdv0hznhqtrxer8u24pk92uc5nq679evw9ddpdzr8alzu5mgmth3fz3epm0k6m693tydsgutw7xjcu675pwdxavz707recqhjgc6w0264fks5r4unt862m25g5pfndnsxye53v8n5f42cmefcdvfmy930yxm8vgz69d34a8dn3v8qhcqthpukvedv63yy62mrt8anr5zwppf8g326n5h38mnzmm7nf4644x5szpgll3uhy32ejmdjmnhn8kvyt2dq8q2673yzznjgzcrruha4arzaeevzgts5xuenc9pcrdnuvku89qdyzlek72kdtszk6dmcsju3v9z5egfxsrwa3eu6tuejhqmrnjrhfmrk98ugnnz6ugayrwtwcrp7y7q98jcd2hprvavxmtesnpxn2df08vy3tk8ypwfh2a43let52tc6q3e9usylyskwyaswkt469am9f5x4v30vslutz7czayx7j4n0g52jyyx7fn57xj43z9xfarmezjew4z4q6v5fnnxh5fy6tu73xe8tremhdexuzkrlzvnq0c97jnm0d8a2uk7ewu5zqf8xaf45tag0487wf98y63ewlfdyuf53awy08heppl7nax3deahunfmr4c2sp9l8dryz7qzmc8djmxm5u6ta45dhlfx40tdwwdhm83xepn8epkw7cfw7hc27tehspcry4nhgjzkd6wdl4wn0ljnsvnwpz37ztfdy0p7yqxkzd8kvvjgwm2748mym45g2wq782ceglj0edd8lfh9g8hjhfsptdlkq2kyldk2udqukk6yzlysf0y08c7lcf9r63ngzg5cpn8kksdy85en3ljv3vl8pse2mcs60l3pyc5alke97a33vfeppn8t75d4xcp2h8u03l3gjwqj74djz09q6g4np9m0u79fupje0kd9vupa9lcs5lrthplyvet85vca96eflj7zz84806yqur4lq6se3wes8wjp9kn5xtuslk66jvx4e9vfu6tp87fw9fyekedmv9pv2579stuf4w846ql54ewdpjcd8j0gwf4ekql0t95zu80y3gc3cmzc2lxvcpt724am87pr4a3m6j6f5tc9zw52zc376ducpc6586a6vakdca6xkzm4s4m59tsvmnxya289e32v60dmkcz3pk03jxe4mrnj6wl2udfgf49eylstc6sjmr5zve0a338l8xye0n4yad5n53ka8t43kuwpycvqe4qje7nx4w4c3zpxrhxkcdvdexvl3w344cf864kekr4rgrds5luv36nn2slj73v2wxwgm4rsvf8ggazy3ytzx6dwvc8986wretny8mua6maf3fdtrdyp0c787d8tjg30xkvwfyll7s4pj0yn8s3lgyg5qsc6v5w0sa0gs0l5s7r7mnn0q6rvthkvkag226yw0xfmake7kjsrsv4klu733fpwrk34lq3zphrzamcge60z425x34l3g7qj49l4d53rfrl6emvr9dxagflldu32550c4uz4wrvxq0fa463cazxvhnl9
}

func ExampleKeyGenerate_ristrettoGroup2HashDh_seed() {
	var (
		seed         = "0000000000000000000000000000000000000000000000000000000000000000"
		keyType      = "RistrettoGroup2HashDh"
		outputFileSk = "" // "" - output to STDOUT ([]byte) only, "privateKey.sk" - will also save output to that file
	)

	privateKey, err := jcli.KeyGenerate(seed, keyType, outputFileSk)

	if err != nil {
		fmt.Printf("KeyGenerate: %s", err)
	} else {
		fmt.Printf("%s", string(privateKey))
	}
	// Output:
	//
	// vrf_sk1fffu87autxtsae0cttup8p6allqn4yz29ef6uln9lgx75mnzeyqshc9lsp
}
