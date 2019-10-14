//$(which go) run $0 $@; exit $?

package main

import (
	"log"
	"strings"

	"github.com/rinor/jorcli/jcli"
)

// b2s converts []byte to string with all leading
// and trailing white space removed, as defined by Unicode.
func b2s(b []byte) string {
	return strings.TrimSpace(string(b))
}

func main() {
	var (
		restAddrAPI = "http://127.0.0.11:8001/api" // genesis node
		// restAddrAPI = "http://127.0.0.22:8001/api" // extra genesis pool node
		// restAddrAPI = "http://127.0.0.33:8001/api" // stakepool node
		// restAddrAPI = "http://127.0.0.44:8001/api" // passive/explorer node node

		// addresses generated from "node_bootstrap_and_run" example
		// we are using seed option, hence we know those values here
		faucetAddr = "jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c"
		fixedAddr  = "jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z"
	)

	//////////////////////
	//  jcli rest usage //
	//////////////////////

	restSettings, err := jcli.RestSettings(restAddrAPI, "json")
	log.Printf("RestSettings: %s - %v\n", b2s(restSettings), err)

	restNetworkStats, err := jcli.RestNetworkStats(restAddrAPI, "json")
	log.Printf("RestNetworkStats: %s - %v\n", b2s(restNetworkStats), err)

	restNodeStats, err := jcli.RestNodeStats(restAddrAPI, "json")
	log.Printf("RestNodeStats: %s - %v\n", b2s(restNodeStats), err)

	restTip, err := jcli.RestTip(restAddrAPI)
	log.Printf("RestTip: %s - %v\n", b2s(restTip), err)

	restAccFc, err := jcli.RestAccount(faucetAddr, restAddrAPI, "json")
	log.Printf("RestAccount Faucet: %s - %v\n", b2s(restAccFc), err)

	restAccFx, err := jcli.RestAccount(fixedAddr, restAddrAPI, "json")
	log.Printf("RestAccount Fixed: %s - %v\n", b2s(restAccFx), err)

	restStakePools, err := jcli.RestStakePools(restAddrAPI, "json")
	log.Printf("RestStakePools: %s - %v\n", b2s(restStakePools), err)

	restStake, err := jcli.RestStake(restAddrAPI, "json")
	log.Printf("RestStake: %s - %v\n", b2s(restStake), err)

	restLeaders, err := jcli.RestLeaders(restAddrAPI, "json")
	log.Printf("RestLeaders: %s - %v\n", b2s(restLeaders), err)

	restLeadersLogs, err := jcli.RestLeadersLogs(restAddrAPI, "json")
	log.Printf("RestLeadersLogs: %s - %v\n", b2s(restLeadersLogs), err)

	restMessageLogs, err := jcli.RestMessageLogs(restAddrAPI, "json")
	log.Printf("RestMessageLogs: %s - %v\n", b2s(restMessageLogs), err)

	//////////////////////////////////////////////////////////
	// remove comment to shutdown the node using jcli rest  //
	//////////////////////////////////////////////////////////
	// rsd, err := jcli.RestShutdown(restAddrAPI, "")
	// log.Printf("RestShutdown: %s - %v\n", b2s(rsd), err)
}
