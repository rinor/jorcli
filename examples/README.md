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
2019/10/14 11:53:20 Using: jcli 0.6.0-rc1 (master-d715f410, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/14 11:53:20 Using: jormungandr 0.6.0-rc1 (master-d715f410, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/14 11:53:20
2019/10/14 11:53:20 Working Directory: /tmp/jnode_073168556
2019/10/14 11:53:27
2019/10/14 11:53:27 Genesis Hash: 999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d
2019/10/14 11:53:27
2019/10/14 11:53:27 LOCAL StakePool ID       : a93cf67dac50f84f2b74f3cccad1c21a2df2c364037e5dc1dd8017c1d320fc9d
2019/10/14 11:53:27 LOCAL StakePool Owner    : jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/10/14 11:53:27 LOCAL StakePool Owner    : jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/10/14 11:53:27 LOCAL StakePool Delegator: jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/10/14 11:53:27 LOCAL StakePool Delegator: jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/10/14 11:53:27
2019/10/14 11:53:27 EXTRA StakePool ID       : 2f3471d99a42e3c75362ea3a217f190143675300ca9893edd86519eff118c9aa
2019/10/14 11:53:27 EXTRA StakePool Owner    : jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/10/14 11:53:27 EXTRA StakePool Delegator: jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/10/14 11:53:27
2019/10/14 11:53:27 NodePublicID for trusted: ed25519_pk1thawa4wxfhn9hh9xll04npw9pv0djgnvcun90nw9szupfw95lvns94qgpu
2019/10/14 11:53:27 NodePublicID in logs    : 5dfaeed5c64de65bdca6ffdf5985c50b1ed9226cc72657cdc580b814b8b4fb27
2019/10/14 11:53:27
2019/10/14 11:53:27 Genesis Node - Running...
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
2019/10/14 11:53:28 Using: jcli 0.6.0-rc1 (master-d715f410, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/14 11:53:28 Using: jormungandr 0.6.0-rc1 (master-d715f410, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/14 11:53:28
2019/10/14 11:53:28 Working Directory: /tmp/jnode_853412662
2019/10/14 11:53:30
2019/10/14 11:53:30 Genesis Hash: 999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d
2019/10/14 11:53:30
2019/10/14 11:53:30 StakePool ID       : 2f3471d99a42e3c75362ea3a217f190143675300ca9893edd86519eff118c9aa
2019/10/14 11:53:30 StakePool Owner    : jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/10/14 11:53:30 StakePool Delegator: jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/10/14 11:53:30
2019/10/14 11:53:30 NodePublicID for trusted: ed25519_pk1z5u62jwftwrepu53nj655cdzjrhv4dlry9d7c602j6dagfpwp34q5gjcmr
2019/10/14 11:53:30 NodePublicID in logs    : 1539a549c95b8790f2919cb54a61a290eecab7e3215bec69ea969bd4242e0c6a
2019/10/14 11:53:30
2019/10/14 11:53:30 Genesis StakePool Node - Running...
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
2019/10/14 11:53:34 Using: jcli 0.6.0-rc1 (master-d715f410, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/14 11:53:34 Using: jormungandr 0.6.0-rc1 (master-d715f410, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/14 11:53:34
2019/10/14 11:53:34 Working Directory: /tmp/jnode_675424923
2019/10/14 11:53:34
2019/10/14 11:53:36 Waiting for rest interface...
2019/10/14 11:53:41 SelfTip: 7ddc01834479fd7dcbea8002761f7bba43f06402e486a66dad4335204d498c3f
2019/10/14 11:53:42 Wait for pool registration certificate transaction [de44b891ef786855c9d2c90650b0af6685c48c7467501496f2927a8a1457c7e1] status change...
2019/10/14 11:53:57 FragmentID: de44b891ef786855c9d2c90650b0af6685c48c7467501496f2927a8a1457c7e1 - InABlock [214702.118 (838c740b62d0efcf5fc7207d69820a85dc2d8e6c99a10d83962ecdeab21db3fb)]
2019/10/14 11:53:57 Wait for delegation certificate transaction [e030916d4ca6f3e9ba9b16d624aa261d00e8a47e1adce99724de73176c665139] status change...
2019/10/14 11:53:59 FragmentID: e030916d4ca6f3e9ba9b16d624aa261d00e8a47e1adce99724de73176c665139 - InABlock [214702.119 (9ad641bf41f228e5950926d711a4c81039901486eed355059ef95454b112ded7)]
2019/10/14 11:53:59
2019/10/14 11:53:59 Genesis Hash: 999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d
2019/10/14 11:53:59
2019/10/14 11:53:59 StakePool ID       : 20ba61c4d0b044962ada2536ab703eeaf95ddf7b90d9900a737988b80abb9415
2019/10/14 11:53:59 StakePool Owner    : jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/10/14 11:53:59 StakePool Owner    : jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/10/14 11:53:59 StakePool Delegator: jnode_ta1s5a8e4qye5rwttc9qrek0e30htttmpvvuf967mdp35pcx80t6e2psskthdh
2019/10/14 11:53:59
2019/10/14 11:53:59 NodePublicID for trusted: ed25519_pk19qzyd6xxed7rc3nxj0qgnsuyxkpqvlcue44l7l3f5kkr9dj378ss2wnm22
2019/10/14 11:53:59 NodePublicID in logs    : 280446e8c6cb7c3c466693c089c3843582067f1ccd6bff7e29a5ac32b651f1e1
2019/10/14 11:53:59
2019/10/14 11:53:59 Delegator StakePool Node - Running...
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
2019/10/14 11:54:01 Using: jcli 0.6.0-rc1 (master-d715f410, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/14 11:54:01 Using: jormungandr 0.6.0-rc1 (master-d715f410, debug, linux [x86_64]) - [rustc 1.38.0 (625451e37 2019-09-23)]
2019/10/14 11:54:01
2019/10/14 11:54:01 Working Directory: /tmp/jnode_186119459
2019/10/14 11:54:01
2019/10/14 11:54:01 Genesis Hash: 999772edda51c486687218bd00a94e09659becf09db5257b03487157a08dac4d
2019/10/14 11:54:01
2019/10/14 11:54:01 NodePublicID for trusted: ed25519_pk1heevnwdnhve9035nr5llwxx0wfc59uhx6c82442eanz0ntudyfqqcfhtwf
2019/10/14 11:54:01 NodePublicID in logs    : be72c9b9b3bb3257c6931d3ff718cf727142f2e6d60eaad559ecc4f9af8d2240
2019/10/14 11:54:01
2019/10/14 11:54:01 Passive/Explorer Node - Running...
```

## JCLI Rest v0

Query the node using rest capabilies of `jcli`.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
