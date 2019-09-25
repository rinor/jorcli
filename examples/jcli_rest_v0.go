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
		restApiAddress = "http://127.0.0.1:8001/api" // leader leader
		// restApiAddress = "http://127.0.0.1:8002/api" // passive node

		// addresses generated from "node_bootstrap_and_run" example
		// we are using seed option, hence we know those values here
		faucetAddr = "jnode_ta1skep7gv070w5hj43jk8uplca76cct09rmv4phgdvqv0e2386qh65jguztau"
		fixedAddr  = "jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c"
	)

	//////////////////////
	//  jcli rest usage //
	//////////////////////

	restSettings, err := jcli.RestSettings(restApiAddress, "json")
	log.Printf("RestSettings: %s - %v\n", b2s(restSettings), err)

	restNodeStats, err := jcli.RestNodeStats(restApiAddress, "json")
	log.Printf("RestNodeStats: %s - %v\n", b2s(restNodeStats), err)

	restTip, err := jcli.RestTip(restApiAddress)
	log.Printf("RestTip: %s - %v\n", b2s(restTip), err)

	restAccFc, err := jcli.RestAccount(faucetAddr, restApiAddress, "json")
	log.Printf("RestAccount Faucet: %s - %v\n", b2s(restAccFc), err)

	restAccFx, err := jcli.RestAccount(fixedAddr, restApiAddress, "json")
	log.Printf("RestAccount Fixed: %s - %v\n", b2s(restAccFx), err)

	restLeaders, err := jcli.RestLeaders(restApiAddress, "json")
	log.Printf("RestLeaders: %s - %v\n", b2s(restLeaders), err)

	restStakePools, err := jcli.RestStakePools(restApiAddress, "json")
	log.Printf("RestStakePools: %s - %v\n", b2s(restStakePools), err)

	restStake, err := jcli.RestStake(restApiAddress, "json")
	log.Printf("RestStake: %s - %v\n", b2s(restStake), err)

	restLeadersLogs, err := jcli.RestLeadersLogs(restApiAddress, "json")
	log.Printf("RestLeadersLogs: %s - %v\n", b2s(restLeadersLogs), err)

	restMessageLogs, err := jcli.RestMessageLogs(restApiAddress, "json")
	log.Printf("RestMessageLogs: %s - %v\n", b2s(restMessageLogs), err)

	//////////////////////////////////////////////////////////
	// remove comment to shutdown the node using jcli rest  //
	//////////////////////////////////////////////////////////
	// rsd, err := jcli.RestShutdown(restApiAddress, "")
	// log.Printf("RestShutdown: %s - %v\n", b2s(rsd), err)
}
