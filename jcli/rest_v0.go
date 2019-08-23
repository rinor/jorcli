package jcli

import (
	"fmt"
	"strconv"
)

/* ******************** ACCOUNT ******************** */

// RestAccount - Get account state.
//
// jcli rest v0 account get <account-id> --host <host> [--output-format <format>]
func RestAccount(account_id string, host string, output_format string) ([]byte, error) {
	if account_id == "" {
		return nil, fmt.Errorf("parameter missing : %s", "account_id")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "account", "get", account_id, "--host", host}
	if output_format != "" {
		arg = append(arg, "--output-format", output_format)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** BLOCK ******************** */

// RestBlock - Get block data.
//
// jcli rest v0 block <block-id> get --host <host>
func RestBlock(block_id string, host string) ([]byte, error) {
	if block_id == "" {
		return nil, fmt.Errorf("parameter missing : %s", "block_id")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "block", block_id, "get", "--host", host}

	return execStd(nil, "jcli", arg...)
}

// RestBlockNextID - Get block descendant ID.
//
// jcli rest v0 block <block-id> next-id get [--count <count>] --host <host>
func RestBlockNextID(block_id string, count_ids uint, host string) ([]byte, error) {
	if block_id == "" {
		return nil, fmt.Errorf("parameter missing : %s", "block_id")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	// NOTE: don't like this uint at all, but keep it for now
	countDefault, countMin, countMax := uint(1), uint(1), uint(100) // count_ids must be in this range 1-100.

	if count_ids < countMin || count_ids > countMax {
		count_ids = countDefault
		// return nil, fmt.Errorf("%s: value must be between %d - %d", "count_ids", countMin, countMax)
	}

	arg := []string{
		"rest", "v0", "block", block_id, "next-id", "get",
		"--count", strconv.FormatUint(uint64(count_ids), 10),
		"--host", host,
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** LEADERS ******************** */

// RestLeadersDelete - Delete leader.
//
// jcli rest v0 leaders delete <id> --host <host>
func RestLeadersDelete(leader_id uint32, host string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{
		"rest", "v0", "leaders", "delete",
		strconv.FormatUint(uint64(leader_id), 10), // FIXME: leader_id > 0
		"--host", host,
	}

	return execStd(nil, "jcli", arg...)
}

// RestLeaders - Get list of leader IDs.
//
// jcli rest v0 leaders get --host <host> [--output-format <format>]
func RestLeaders(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "leaders", "get", "--host", host}
	if output_format != "" {
		arg = append(arg, "--output-format", output_format)
	}

	return execStd(nil, "jcli", arg...)
}

// RestLeadersLogs - Get leadership logs.
//
// jcli rest v0 leaders logs get --host <host> [--output-format <format>]
func RestLeadersLogs(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "leaders", "logs", "get", "--host", host}
	if output_format != "" {
		arg = append(arg, "--output-format", output_format)
	}

	return execStd(nil, "jcli", arg...)
}

// RestLeadersPost - Register new leader and get its ID.
//
// STDIN | jcli rest v0 leaders post --host <host> [--file <input_file>]
func RestLeadersPost(stdin_sk []byte, host string, input_file_sk string) ([]byte, error) {
	if len(stdin_sk) == 0 && input_file_sk == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_sk", "input_file_sk")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "leaders", "post", "--host", host}
	if input_file_sk != "" {
		arg = append(arg, "--file", input_file_sk)
		stdin_sk = nil
	}

	return execStd(stdin_sk, "jcli", arg...)
}

/* ******************** MESSAGE ******************** */

// RestMessageLogs - get the node's logs on the message pool.
// This will provide information on pending transaction, rejected transaction
// and or when a transaction has been added in a block
//
// jcli rest v0 message logs --host <host> [--output-format <format>]
func RestMessageLogs(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "message", "logs", "--host", host}
	if output_format != "" {
		arg = append(arg, "--output-format", output_format)
	}

	return execStd(nil, "jcli", arg...)
}

// RestMessagePost - Post message and prints id for posted message.
//
// STDIN | jcli rest v0 message post --host <host> [--file <input_file>]
func RestMessagePost(stdin_msg []byte, host string, input_file_msg string) ([]byte, error) {
	if len(stdin_msg) == 0 && input_file_msg == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdin_msg", "input_file_msg")
	}
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "message", "post", "--host", host}
	if input_file_msg != "" {
		arg = append(arg, "--file", input_file_msg)
		stdin_msg = nil
	}

	return execStd(stdin_msg, "jcli", arg...)
}

/* ******************** NODE ******************** */

// RestNodeStats - Get node information.
//
// jcli rest v0 node stats get --host <host> --output-format <format>
func RestNodeStats(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "node", "stats", "get", "--host", host}
	if output_format != "" {
		arg = append(arg, "--output-format", output_format)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** SETTINGS ******************** */

// RestSettings - Get node settings.
//
// jcli rest v0 settings get --host <host> --output-format <format>
func RestSettings(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "settings", "get", "--host", host}
	if output_format != "" {
		arg = append(arg, "--output-format", output_format)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** SHUTDOWN ******************** */

// RestShutdown - Shutdown node.
//
// jcli rest v0 shutdown get --host <host>
func RestShutdown(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "shutdown", "get", "--host", host}

	return execStd(nil, "jcli", arg...)
}

/* ******************** STAKE-POOLS ******************** */

// RestStakePools - Get stake pool IDs
//
// jcli rest v0 stake-pools get --host <host> --output-format <format>
func RestStakePools(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "stake-pools", "get", "--host", host}
	if output_format != "" {
		arg = append(arg, "--output-format", output_format)
	}

	return execStd(nil, "jcli", arg...)
}

/* ******************** TIP ******************** */

// RestTip - Get tip.
//
// jcli rest v0 tip get --host <host>
func RestTip(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "top", "get", "--host", host}

	return execStd(nil, "jcli", arg...)
}

/* ******************** UTXO ******************** */

// RestUTxOs - Get all UTXOs.
//
// jcli rest v0 utxo get --host <host> --output-format <format>
func RestUTxOs(host string, output_format string) ([]byte, error) {
	if host == "" {
		return nil, fmt.Errorf("parameter missing : %s", "host")
	}

	arg := []string{"rest", "v0", "utxo", "get", "--host", host}
	if output_format != "" {
		arg = append(arg, "--output-format", output_format)
	}

	return execStd(nil, "jcli", arg...)
}
