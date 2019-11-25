package jcli

import (
	"fmt"
	"os"
	"strconv"
)

const envJormungandrRestAPIURL = "JORMUNGANDR_RESTAPI_URL"

/* ******************** ACCOUNT ******************** */

// RestAccount - Get account state.
//
//  jcli rest v0 account get <account-id> --host <host> [--output-format <format>] | STDOUT
func RestAccount(
	accountID string,
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}
	if accountID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "accountID")
	}

	arg := []string{"rest", "v0", "account", "get", accountID}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

/* ******************** BLOCK ******************** */

// RestBlock - Get block data.
//
//  jcli rest v0 block <block-id> get --host <host> | STDOUT
func RestBlock(
	blockID string,
	host string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}
	if blockID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "blockID")
	}

	arg := []string{"rest", "v0", "block", blockID, "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}

	return jcli(nil, arg...)
}

// RestBlockNextID - Get block descendant ID.
//
//  jcli rest v0 block <block-id> next-id get [--count <count>] --host <host> | STDOUT
func RestBlockNextID(
	blockID string,
	countIds uint,
	host string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}
	if blockID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "blockID")
	}

	// NOTE: don't like this uint at all, but keep it for now
	countDefault, countMin, countMax := uint(1), uint(1), uint(100) // countIds must be in this range 1-100.

	if countIds < countMin || countIds > countMax {
		countIds = countDefault
		// return nil, fmt.Errorf("%s: value must be between %d - %d", "countIds", countMin, countMax)
	}

	arg := []string{
		"rest", "v0", "block", blockID, "next-id", "get",
		"--count", strconv.FormatUint(uint64(countIds), 10),
	}
	if host != "" {
		arg = append(arg, "--host", host)
	}

	return jcli(nil, arg...)
}

/* ******************** LEADERS ******************** */

// RestLeadersDelete - Delete leader.
//
//  jcli rest v0 leaders delete <id> --host <host> | STDOUT
func RestLeadersDelete(
	leaderID uint32,
	host string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{
		"rest", "v0", "leaders", "delete",
		strconv.FormatUint(uint64(leaderID), 10), // FIXME: leaderID > 0
	}
	if host != "" {
		arg = append(arg, "--host", host)
	}

	return jcli(nil, arg...)
}

// RestLeaders - Get list of leader IDs.
//
//  jcli rest v0 leaders get --host <host> [--output-format <format>] | STDOUT
func RestLeaders(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "leaders", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

// RestLeadersLogs - Get leadership logs.
//
//  jcli rest v0 leaders logs get --host <host> [--output-format <format>] | STDOUT
func RestLeadersLogs(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "leaders", "logs", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

// RestLeadersPost - Register new leader and get its ID.
//
//  [STDIN] | jcli rest v0 leaders post --host <host> [--file <input_file>] | STDOUT
func RestLeadersPost(
	stdinSk []byte,
	host string,
	inputFileSk string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}
	if len(stdinSk) == 0 && inputFileSk == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinSk", "inputFileSk")
	}

	arg := []string{"rest", "v0", "leaders", "post"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if inputFileSk != "" {
		arg = append(arg, "--file", inputFileSk)
		stdinSk = nil
	}

	return jcli(stdinSk, arg...)
}

/* ******************** MESSAGE ******************** */

// RestMessageLogs - get the node's logs on the message pool.
// This will provide information on pending transaction, rejected transaction
// and or when a transaction has been added in a block
//
//  jcli rest v0 message logs --host <host> [--output-format <format>] | STDOUT
func RestMessageLogs(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "message", "logs"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

// RestMessagePost - Post message and prints id for posted message.
//
//  [STDIN] | jcli rest v0 message post --host <host> [--file <input_file>] | STDOUT
func RestMessagePost(
	stdinMsg []byte,
	host string,
	inputFileMsg string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}
	if len(stdinMsg) == 0 && inputFileMsg == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinMsg", "inputFileMsg")
	}

	arg := []string{"rest", "v0", "message", "post"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if inputFileMsg != "" {
		arg = append(arg, "--file", inputFileMsg)
		stdinMsg = nil
	}

	return jcli(stdinMsg, arg...)
}

/* ******************** NODE ******************** */

// RestNodeStats - Get node information.
//
//  jcli rest v0 node stats get --host <host> --output-format <format> | STDOUT
func RestNodeStats(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "node", "stats", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

/* ******************** SETTINGS ******************** */

// RestSettings - Get node settings.
//
//  jcli rest v0 settings get --host <host> --output-format <format> | STDOUT
func RestSettings(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "settings", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

/* ******************** SHUTDOWN ******************** */

// RestShutdown - Shutdown node.
//
//  jcli rest v0 shutdown get --host <host> | STDOUT
func RestShutdown(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "shutdown", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}

	return jcli(nil, arg...)
}

/* ******************** STAKE ******************** */

// RestStake - Get stake distribution
//
//  jcli rest v0 stake get --host <host> --output-format <format> | STDOUT
func RestStake(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "stake", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

/* ******************** STAKE-POOLS ******************** */

// RestStakePools - Get stake pool IDs
//
//  jcli rest v0 stake-pools get --host <host> --output-format <format> | STDOUT
func RestStakePools(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "stake-pools", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

//
// RestStakePool - Get stake pool details
//
//  jcli rest v0 stake-pool get <pool-id> --host <host> --output-format <format> | STDOUT
func RestStakePool(
	poolID string,
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}
	if poolID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "poolID")
	}

	arg := []string{"rest", "v0", "stake-pool", "get", poolID}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

/* ******************** TIP ******************** */

// RestTip - Get tip.
//
//  jcli rest v0 tip get --host <host> | STDOUT
func RestTip(
	host string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "tip", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}

	return jcli(nil, arg...)
}

/* ******************** UTXO ******************** */

// RestUTxO - UTXO information.
//
//  jcli rest v0 utxo <fragment-id> <output-index> get --host <host> --output-format <format> | STDOUT
func RestUTxO(
	fragmentID string,
	outputIndex uint8,
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}
	if fragmentID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "fragmentID")
	}

	arg := []string{
		"rest", "v0",
		"utxo", "fragmentID", strconv.FormatUint(uint64(outputIndex), 10),
		"get",
	}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}

/* ******************** NETWORK ******************** */

// RestNetworkStats - Get network information.
//
//  jcli rest v0 network stats get --host <host> --output-format <format> | STDOUT
func RestNetworkStats(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" && os.Getenv(envJormungandrRestAPIURL) == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "network", "stats", "get"}
	if host != "" {
		arg = append(arg, "--host", host)
	}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return jcli(nil, arg...)
}
