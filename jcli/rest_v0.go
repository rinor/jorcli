package jcli

import (
	"fmt"
	"strconv"
)

/* ******************** ACCOUNT ******************** */

// RestAccount - Get account state.
//
//  jcli rest v0 account get <account-id> --host <host> [--output-format <format>] | STDOUT
func RestAccount(
	accountID string,
	host string,
	outputFormat string,
) ([]byte, error) {
	if accountID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "accountID")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "account", "get", accountID, "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** BLOCK ******************** */

// RestBlock - Get block data.
//
//  jcli rest v0 block <block-id> get --host <host> | STDOUT
func RestBlock(
	blockID string,
	host string,
) ([]byte, error) {
	if blockID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "blockID")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "block", blockID, "get", "--host", host}

	return execStd(nil, "jcli", arg...)
}

// RestBlockNextID - Get block descendant ID.
//
//  jcli rest v0 block <block-id> next-id get [--count <count>] --host <host> | STDOUT
func RestBlockNextID(
	blockID string,
	countIds uint,
	host string,
) ([]byte, error) {
	if blockID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "blockID")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
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
		"--host", host,
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** LEADERS ******************** */

// RestLeadersDelete - Delete leader.
//
//  jcli rest v0 leaders delete <id> --host <host> | STDOUT
func RestLeadersDelete(
	leaderID uint32,
	host string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{
		"rest", "v0", "leaders", "delete",
		strconv.FormatUint(uint64(leaderID), 10), // FIXME: leaderID > 0
		"--host", host,
	}

	return execStd(nil, "jcli", arg...)
}

// RestLeaders - Get list of leader IDs.
//
//  jcli rest v0 leaders get --host <host> [--output-format <format>] | STDOUT
func RestLeaders(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "leaders", "get", "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}

// RestLeadersLogs - Get leadership logs.
//
//  jcli rest v0 leaders logs get --host <host> [--output-format <format>] | STDOUT
func RestLeadersLogs(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "leaders", "logs", "get", "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}

// RestLeadersPost - Register new leader and get its ID.
//
//  [STDIN] | jcli rest v0 leaders post --host <host> [--file <input_file>] | STDOUT
func RestLeadersPost(
	stdinSk []byte,
	host string,
	inputFileSk string,
) ([]byte, error) {
	if len(stdinSk) == 0 && inputFileSk == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinSk", "inputFileSk")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "leaders", "post", "--host", host}
	if inputFileSk != "" {
		arg = append(arg, "--file", inputFileSk)
		stdinSk = nil
	}

	return execStd(stdinSk, "jcli", arg...)
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
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "message", "logs", "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}

// RestMessagePost - Post message and prints id for posted message.
//
//  [STDIN] | jcli rest v0 message post --host <host> [--file <input_file>] | STDOUT
func RestMessagePost(
	stdinMsg []byte,
	host string,
	inputFileMsg string,
) ([]byte, error) {
	if len(stdinMsg) == 0 && inputFileMsg == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinMsg", "inputFileMsg")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "message", "post", "--host", host}
	if inputFileMsg != "" {
		arg = append(arg, "--file", inputFileMsg)
		stdinMsg = nil
	}

	return execStd(stdinMsg, "jcli", arg...)
}

/* ******************** NODE ******************** */

// RestNodeStats - Get node information.
//
//  jcli rest v0 node stats get --host <host> --output-format <format> | STDOUT
func RestNodeStats(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "node", "stats", "get", "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** SETTINGS ******************** */

// RestSettings - Get node settings.
//
//  jcli rest v0 settings get --host <host> --output-format <format> | STDOUT
func RestSettings(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "settings", "get", "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** SHUTDOWN ******************** */

// RestShutdown - Shutdown node.
//
//  jcli rest v0 shutdown get --host <host> | STDOUT
func RestShutdown(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "shutdown", "get", "--host", host}

	return execStd(nil, "jcli", arg...)
}

/* ******************** STAKE ******************** */

// RestStake - Get stake distribution
//
//  jcli rest v0 stake get --host <host> --output-format <format> | STDOUT
func RestStake(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "stake", "get", "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** STAKE-POOLS ******************** */

// RestStakePools - Get stake pool IDs
//
//  jcli rest v0 stake-pools get --host <host> --output-format <format> | STDOUT
func RestStakePools(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "stake-pools", "get", "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** TIP ******************** */

// RestTip - Get tip.
//
//  jcli rest v0 tip get --host <host> | STDOUT
func RestTip(
	host string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "top", "get", "--host", host}

	return execStd(nil, "jcli", arg...)
}

/* ******************** UTXO ******************** */

// RestUTxO - Get all UTXOs.
//
//  jcli rest v0 utxo get --host <host> --output-format <format> | STDOUT
func RestUTxO(
	host string,
	outputFormat string,
) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "utxo", "get", "--host", host}
	if outputFormat != "" {
		arg = append(arg, "--output-format", outputFormat)
	}

	return execStd(nil, "jcli", arg...)
}
