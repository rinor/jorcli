# Examples

1) [node_bootstrap_and_run.go](#node-bootstrap-and-run)
2) [jcli_rest_v0.go](#jcli-rest-v0)
3) [node_passive_run.go](#node-passive-run)

## Node Bootstrap And Run

Bootstrap configuration and start a **leader** node.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)


## JCLI Rest v0

Query the node using rest capabilies of `jcli`.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)

## Node Passive Run

Start a passive node with:

- `--genesis-block-hash` from [node_bootstrap_and_run.go](#node-bootstrap-and-run)

- `--trusted-peer` the node from [node_bootstrap_and_run.go](#node-bootstrap-and-run)

and connect to **leader** node.

It shows the usage of:

- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)