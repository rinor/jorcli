# Examples

1) [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
2) [node_passive_run](#passive-node-run)
3) [node_stakepool_create_and_run](#stakepool-node-create-and-run)
4) [jcli_rest_v0](#jcli-rest-v0)

## Genesis Node Bootstrap And Run

Bootstrap configuration and start a **leader** node.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info Genesis Node

TODO

## Passive Node Run

Start a passive node with:

- `--genesis-block-hash` from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--trusted-peer` the node from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)

and connect to **leader** node.

It shows the usage of:

- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info Passive Node

TODO

## StakePool Node Create And Run

Creates a StakePool node with appropriate configuration:

- `--secret` self generate pool secrets
- `--trusted-peer` the node from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--trusted-peer` the node from [node_passive_run](#passive-node-run)

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info StakePool Node

At this point the StakePool is configured and running,
but the node behaves like a passive one since:

1. StakePool is registered on the network, but has no stake yet.

## JCLI Rest v0

Query the node using rest capabilies of `jcli`.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
