# Examples

1) [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
2) [node_passive_run](#passive-node-run)
3) [node_stakepool_create_and_run](#stakepool-node-create-and-run)
10) [jcli_rest_v0](#jcli-rest-v0)

## Genesis Node Bootstrap And Run

Bootstrap configuration and start a **leader** node.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info

TODO

## Passive Node Run

Start a passive node with:

- `--genesis-block-hash` from [node_genesis_bootstrap_and_run](#node-genesis-bootstrap-and-run)
- `--trusted-peer` the node from [node_genesis_bootstrap_and_run](#node-genesis-bootstrap-and-run)

and connect to **leader** node.

It shows the usage of:

- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info

TODO

## StakePool Node Create And Run

Creates a StakePool node with appropriate configuration:

- `--secret` self generate pool secrets
- `--trusted-peer` the node from [node_genesis_bootstrap_and_run](#node-genesis-bootstrap-and-run)
- `--trusted-peer` the node from [node_passive_run](#node-passive-run)

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info
At this point the StakePool is configured and running,
but the node behaves like a passive one since:

1. the StakePool is not yet registered on the network
2. Even if it was registered the StakePool has no stake yet.

## JCLI Rest v0

Query the node using rest capabilies of `jcli`.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)