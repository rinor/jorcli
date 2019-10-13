# Examples

1) [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
2) [node_genesis_stakepool_run](#genesis-stakepool-node-run)
3) [node_stakepool_create_and_run](#stakepool-node-create-and-run)
4) [node_passive_run](#passive-node-run)
5) [jcli_rest_v0](#jcli-rest-v0)

## Genesis Node Bootstrap And Run

Bootstrap configuration and start a **leader** node with and active local stake pool.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info Genesis Node

```log
2019/10/13 17:57:25 Using: jcli 0.5.6+lock (master-dbd5e2e8, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/13 17:57:25 Using: jormungandr 0.5.6+lock (master-dbd5e2e8, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/13 17:57:25
2019/10/13 17:57:25 Working Directory: /tmp/jnode_651461336
2019/10/13 17:57:32
2019/10/13 17:57:32 Genesis Hash: 999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d
2019/10/13 17:57:32
2019/10/13 17:57:32 LOCAL StakePool ID       : a93cf67dac50f84f2b74f3cccad1c21a2df2c364037e5dc1dd8017c1d320fc9d
2019/10/13 17:57:32 LOCAL StakePool Owner    : jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/10/13 17:57:32 LOCAL StakePool Owner    : jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/10/13 17:57:32 LOCAL StakePool Delegator: jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/10/13 17:57:32 LOCAL StakePool Delegator: jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/10/13 17:57:32
2019/10/13 17:57:32 EXTRA StakePool ID       : 2f3471d99a42e3c75362ea3a217f190143675300ca9893edd86519eff118c9aa
2019/10/13 17:57:32 EXTRA StakePool Owner    : jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/10/13 17:57:32 EXTRA StakePool Delegator: jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/10/13 17:57:32
2019/10/13 17:57:32 NodeID: ed25519_pk1thawa4wxfhn9hh9xll04npw9pv0djgnvcun90nw9szupfw95lvns94qgpu
2019/10/13 17:57:32
2019/10/13 17:57:32 Genesis Node - Running...
```

## Genesis Stakepool Node Run

Configure and start the extra stake pool already encoded in genesis node.

- `--genesis-block-hash` from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--secret` self generate pool secrets (related to the extra pool encoded in genesis block0)
- `--trusted-peer` the node from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info Genesis Stakepol Node

```log
2019/10/13 18:01:40 Using: jormungandr 0.5.6+lock (master-dbd5e2e8, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/13 18:01:40
2019/10/13 18:01:40 Working Directory: /tmp/jnode_559982986
2019/10/13 18:01:42
2019/10/13 18:01:42 Genesis Hash: 999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d
2019/10/13 18:01:42
2019/10/13 18:01:42 StakePool ID       : 2f3471d99a42e3c75362ea3a217f190143675300ca9893edd86519eff118c9aa
2019/10/13 18:01:42 StakePool Owner    : jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/10/13 18:01:42 StakePool Delegator: jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/10/13 18:01:42
2019/10/13 18:01:42 NodeID: ed25519_pk1z5u62jwftwrepu53nj655cdzjrhv4dlry9d7c602j6dagfpwp34q5gjcmr
2019/10/13 18:01:42
2019/10/13 18:01:42 Genesis StakePool Node - Running...
```

## StakePool Node Create And Run

Creates a StakePool node with appropriate configuration,
register it to the network and run it:

- `--genesis-block-hash` from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--secret` self generate pool secrets
- `--trusted-peer` the node from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--trusted-peer` the node from [node_genesis_stakepool_run](#genesis-stakepool-node-run)

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info StakePool Node

```log
2019/10/13 18:13:09 Using: jcli 0.5.6+lock (master-dbd5e2e8, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/13 18:13:09 Using: jormungandr 0.5.6+lock (master-dbd5e2e8, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/13 18:13:09
2019/10/13 18:13:09 Working Directory: /tmp/jnode_857293357
2019/10/13 18:13:09
2019/10/13 18:13:11 Waiting for rest interface...
2019/10/13 18:13:16 SelfTip: 4d31256f4544586b63006ba7e13ac4a9529c9b64b5758416f339f949ab832caa
2019/10/13 18:13:17 Wait for pool registration certificate transaction [de44b891ef786855c9d2c90650b0af6685c48c7467501496f2927a8a1457c7e1] status change...
2019/10/13 18:13:31 FragmentID: de44b891ef786855c9d2c90650b0af6685c48c7467501496f2927a8a1457c7e1 - InABlock [214490.105]
2019/10/13 18:13:32 Wait for delegation certificate transaction [e030916d4ca6f3e9ba9b16d624aa261d00e8a47e1adce99724de73176c665139] status change...
2019/10/13 18:14:21 FragmentID: e030916d4ca6f3e9ba9b16d624aa261d00e8a47e1adce99724de73176c665139 - InABlock [214490.130]
2019/10/13 18:14:21
2019/10/13 18:14:21 Genesis Hash: 999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d
2019/10/13 18:14:21
2019/10/13 18:14:21 StakePool ID       : 20ba61c4d0b044962ada2536ab703eeaf95ddf7b90d9900a737988b80abb9415
2019/10/13 18:14:21 StakePool Owner    : jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/10/13 18:14:21 StakePool Owner    : jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/10/13 18:14:21 StakePool Delegator: jnode_ta1s5a8e4qye5rwttc9qrek0e30htttmpvvuf967mdp35pcx80t6e2psskthdh
2019/10/13 18:14:21
2019/10/13 18:14:21 NodeID: ed25519_pk19qzyd6xxed7rc3nxj0qgnsuyxkpqvlcue44l7l3f5kkr9dj378ss2wnm22
2019/10/13 18:14:21
2019/10/13 18:14:21 StakePool Node - Running...
```

## Passive Node Run

Start a passive node with:

- `--genesis-block-hash` from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--trusted-peer` the node from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--trusted-peer` the node from [node_genesis_stakepool_run](#genesis-stakepool-node-run)
- `--trusted-peer` the node from [node_stakepool_create_and_run](#stakepool-node-create-and-run)

and connect to **leader** node.

It shows the usage of:

- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info Passive Node

```log
2019/10/13 18:19:31 Using: jormungandr 0.5.6+lock (master-dbd5e2e8, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/13 18:19:31
2019/10/13 18:19:31 Working Directory: /tmp/jnode_889162869
2019/10/13 18:19:31
2019/10/13 18:19:31 Genesis Hash: 999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d
2019/10/13 18:19:31
2019/10/13 18:19:31 NodeID: ed25519_pk1heevnwdnhve9035nr5llwxx0wfc59uhx6c82442eanz0ntudyfqqcfhtwf
2019/10/13 18:19:31
2019/10/13 18:19:31 Passive Node - Running...
```

## JCLI Rest v0

Query the node using rest capabilies of `jcli`.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
