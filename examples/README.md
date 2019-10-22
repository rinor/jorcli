# Examples

1) [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
2) [node_genesis_stakepool_run](#genesis-stakepool-node-run)
3) [node_stakepool_create_and_run](#stakepool-node-create-and-run)
4) [node_passive_run](#passive-node-run)
5) [jcli_rest_v0](#jcli-rest-v0)
6) [jcli_bulk_send](#jcli-bulk-send)

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

## JCLI Bulk send

Bulk send 100 transaction to 1 address (delegator address).
This is a sequential version in order not to stress too much the rest interface.
It will be enhanced and converted in a concurrent one.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)

Example log showing first the **account address** - **spending counter** - **account balance**

```log
2019/10/15 20:51:35 jnode_ta1skjryzc2x6a4dwpxmkygznf3czvpcmuequ9gqca0hcufqp87vw5a5y3lgal - 5 - 49999999950000
2019/10/15 20:51:35 jnode_ta1s4pmyagmqdanyfjtdrz9jj8jrf84dyqhnf2w67man3xlspt842cquhc9xek - 5 - 49999999950000
2019/10/15 20:51:35 jnode_ta1sk9ydx0xy3q4pz477yfc2gx33vlgepa0zjqdujmsm83n4zx8xwn4yf5tzdd - 5 - 49999999950000
2019/10/15 20:51:35 jnode_ta1s53w0sa28fw0ykreyw7nyclgl3rjprw7mkccpk084k8rmg7un474gsg862x - 5 - 49999999950000
2019/10/15 20:51:35 jnode_ta1sh2xfzwxrqlxal6vnmhqtzkm3k4nvum8rlk3fd2ksudpk0las39466an2tz - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1s47g6peqfzfxmhzhc07xqrlhqmq3f0q24v6nr8a42lqduegtj9kvcfkm3hn - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1s4gnnljx4wa7rfxr0swxqln4wg9s4vj8kt3xsa5djmya984zea5h2t8vm3k - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1s5au5mf0axuppq9qsfe90kwne0jjtjdkh93pmfplap56yue24kvz6cljd03 - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1s47zzz8nrwww9z4w2yl9d3yjckxap6dyd5u27pryuyk8vpje0jkzjzkjnka - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1s4cca439qvt7dazppu9uuvl8xpwf45dtxqmh2q3wgxr8tx4hjg2wyhdthuu - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1s5xlflm5ytzvzhr6p6m9gj2qlw03g20qwdgwxnsv2ll4kj0jthvgzgfulva - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1shhagrmnc2d2gc89me7el5d57706e8vmvm35wd277u9tq7tnqu9j79h3hvd - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1s4dd2fkla3deusgweywzge9t0cetm064kllf9gsfvtfpc9nfs7sgcgf4e3v - 5 - 49999999950000
2019/10/15 20:51:36 jnode_ta1skkf06ftnlexqwvx0fzgs82vqnrnm9s84adhr89wu5q0ly3at5xuxtlepl3 - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1s4mtn4jmefk00qfmv7j3t6vesqunarnsmht9herk7yuf8x4mzs8cztt0n58 - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1shldwftcpng0ah3kzrc7ufzl39uafsjkrmk8yu7cum05r9fuslj65d7m8ap - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1shykqlluq8ass3kgj5zalhrmwlf4d2qd4lam5rhsgcw2urjy8rwpc6qh6f6 - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1s49dckh3gqtfzqrfe5t4x93j6kzhq4ssuc9zfdk5fdwfd6gat8c4j2xp2x7 - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1sk4l6zqcc4tgsfp0akeknpkd35jrrsz97827z35mg36wrc6fumhxw5tay3z - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1skxqlkz49rrnpwd3heapa76rnrp2adent77w4asrq4sayw3wwpp9jfcn5k8 - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1sh8jfxjgtmdym04x3tr0tzugxw0a7hqsj0ydr40xt2rz69hsxy99g0m7e0s - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1s4yvyn93sjmr7e7rlr8r4qsq76mfrr5mpumqq8vvzur33uergnu8uudrs99 - 5 - 49999999950000
2019/10/15 20:51:37 jnode_ta1s5hpjd5uy5npu6d6pc7kycz3pawkzd9m6kau3kq20mq2easv5pfyj802gv4 - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1skuwg6kam9wlckh7kzrlul4fq022nfr52rdzd0094550wkmr2ydk7z3dwz0 - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1s550tzlpy3xhlezde72cwn70xg94924k04luw5nkxfgwhdkf0e9g29x06ff - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1skfys5mtl9w5chu7dlk7f0p3xtqrvauhw866ykg6jmevtjkdmumj69sl4nq - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1s5jvr7hehcme4hlgp0jl30s0jhl0hhl3lm7h406s50u65t7c3myvkklfdj9 - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1s5gmt78jszdcjthq5snrawgyf4rycdjcfkvfkjnxx437pzkvxkxrwkhh537 - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1sh8fjw43kxwzmu6stas4rw4fvsgwd9mtnrdhykdas6833nu05ex6swx55ja - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1skzlvder8cv6fsnuxqm6dkrgra97py2jprkxcgdzlrmu5m6yglvnxs6qqwm - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1sk66alarl3xqklwee22xe07kpxxs94nh96r37ye4dl6uwh96zf0m724980u - 5 - 49999999950000
2019/10/15 20:51:38 jnode_ta1sh4rm259rqpuazujn7z7l3d5hcwymgum97g2c4453cnvpmg8athk5dwxkm7 - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1s4y5wz08tcfv8xdmdqkf5knxnm7qk9snft3g045txswvycc3vq2m75e7utq - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1s55zzdkw3r95hdsym29vvlacan40we90pqu5vrempuz0x48cw7tqqswzjmt - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1s43ymn68j9yq6gmdmt745tfdj8pee8ct02qmd9zmd2d3chc6m3dqwrxhkhx - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1skng9thtxeyxkn8hg4038q423e2wkhzym4c3t87hsk5tghtfc7hlyjehx33 - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1s5jp5xhmqqv2uyzdnx2ge4tsdkxs7gqvy956trtjl2c0fhnkudl6gp46xuf - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1sk59zz4klx6j8md78246rqhk8q0gnnp6gj4qje94yr2c24ggrswgwqm0scu - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1s45xm9nryk8nxdfhvznjuxdav842muqr8c0czntj5500qxkckjylkwa4xa9 - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1s42v0gfknf648du70rx9qu68t4zfrklnl0ajpwh4mn8u4n3txn5a25wafxc - 5 - 49999999950000
2019/10/15 20:51:39 jnode_ta1s5sxsls6warl8q62juflm9ny58jzjyq9tddparqv9xj55p6ymp4lq8vu6y0 - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1sh4pygj8p3d0yasksql7yjwueltrndurcv4ngqar5qxmjmytd22w7zljjnj - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1s53qm6edw5trx4zeqak8mcdx0706fnmjnxp68me8y055kq5z4p9t67xwhjf - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1skylrf9ulphq6x2hq2vmcflf8kc335uyg80ll72v7pd7md8wwna5g8dquqx - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1shjz9fnn423jchm8ukf2dtcwmh65czg05d0g3xymepanpsnugk3zxjqgvm0 - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1sk3dc8cxuvad3wg9fkeap5m4fz9v3epfrmk9w0yx0c9cv7k2pvq25zxzwsp - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1s5aze7naejjz487l6vz5zuhqf2fn48qve8fxzv5f53mfkjlapaj9cq4fqxk - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1s4pp6gntvdg39nju3q9jhy9l20lqajf2l7f0czryfzpl7kkcc3pey3545yh - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1s5zssp6n2x3yy72jym8r7ypjkxw6np7dgnf6e5yyslml9ddm7qj5s4yfjdw - 5 - 49999999950000
2019/10/15 20:51:40 jnode_ta1s40579t28us2kgq5vznjzmaudt2pa9tjq4kg2q38pqesdw9mfy9ws7xke69 - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1s5nm64z0sk9r8jp9y20nv0eegmq6gm2hpghpnp3z9lzzdk4uerhfx55cua7 - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1s5ec5mymsuvdaeljcg7c7s7x9wcg4dmykf35lu0n0qk4mpmcgrnp706f8e4 - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1s4ahxnflfda06xcdw909npk92c5synlnk7pectq7rs8l8mr6p6wqsuptjp5 - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1shv0a8dx6lpj4fr8xup9drf2a4cmpzvlwh5enwjcqaj6v3gel76rwygtyg8 - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1s59zyusjregyhhwchmdf0lly3ld53tj0s49lrcya238pdmmmrt50jcwfatx - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1s5x78szsanpd8smym2rdjvhwlhuz8ccncvj0vyu0mm8zxkc6h6kmzrlz3mj - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1skmm0369a8efkf2rntvs5ymvwwfrw9zv6n9zkd8mhv96wscj5fuqgytd2dy - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1s43s0kp0l85lld0nutvq6xfkkcmmmw3wh46x7r4y57h97u50l5ft5ec5y2g - 5 - 49999999950000
2019/10/15 20:51:41 jnode_ta1s4u3c0sflmltp9lnvy6s76qgvxjgtn0aw3sysyep4p5nkg2h3djywt64zay - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1s4e9jas42rcqs3ufcnlp3mvp7nyt3mrf9qee0hca83pdxewlu63nu9qdh5u - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1s420wwu2hfgwj5tvwava78tugq06evkfw8ngu8wyn2xuwttt78qqg6ayxr5 - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1s4267em4razxtdl83538kl5f350h5f353dygp0kgvmk5a82p726v5jt8tex - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1shn9axum2cphz2f66r2sr6uqwjq4mzh2xgkxnyw9ksnt3u47klzy7m9zggk - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1s5tq3s7r3ezmuhh9v9ha5050gltzf7mllx2aejqhck5txhmqzhqqs9y52wf - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1skx3tk6upatfg4s063yzp84jd5g6dp4sdlkzu47gjrfct9nfugvlsgg3hqu - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1skezrvu4m2hklr2dlfthswl2sc649ua8yhrfnz0uhfus9xsufpjryfj6sqw - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1s55g3dunvppflvnu2f6a08akmhwnpyzcpxras0qvnhdwnjk0unzm5s55cpy - 5 - 49999999950000
2019/10/15 20:51:42 jnode_ta1s5t5u3kpdkp6v9muzw276m4uva2eugvv8szzp0ynv4amu7mvhdwsuf9pzkx - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1s4p7w0689mwantdt9j779g5ajlc3eyq2778a9yn3vf292xq6xla5gp5lh06 - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1s496dy5e998p733e8h6kam7fwt35x06n9z9lqrwzm43dc0mmj38nkdpf5jx - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1she3z799ka4puvfh56z757me4e9nl8akcr6n22j6jyar885wv7ztv3y0mvv - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1s5ef5pgcf6su6t79tdnvdhzq206hfyusnl35rf2dn0hvxkhqnfxu624h5me - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1shnd447lg057slzcf64l5hr3l4j8pdv7n65cams346mpxnuw5q2q620zt48 - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1s4qtez8tup99kd9q2wh707dsa3rm9q7hvt4aak3sgx4x28ft5f77uj7r5d0 - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1sh23ttcfvekxvu6zqvlmf3hmtv8dunsf57c834h0l64tupmv7uaxjhk5xyn - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1s5v3r3wh0vmejvpjq3tx3uqduyk0wp0z3devcmjpc4fx9csmm6xu7pdjwqk - 5 - 49999999950000
2019/10/15 20:51:43 jnode_ta1skg9cnc6ns09vq7re8en2s7ypfnrtk30vvhpx8p8gfj38627dksu2q80rwy - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1s4dpw4nm6n59ars56u466say7xlvtyn755qj50mday392k6vthvz7avfd6e - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1s4yq0s08gthuw0mz5gcjmt3vtnkpnsc98hhyudzkta738ryjdlv0uepy5cm - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1s5w39hejw0hcgeghn5rdm96zea8e2e8xfgn9f4a9a5e2t2vgl7fj2t3cp6q - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1s4ha79eecqv5nj3y87xrxmqw76kn5u0hcywy2zvwerpkt7xyk4h8ymdg44w - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1skyywuduvxq3y4wsjg59ycvzy2tenldy53kksghcee8g7jywf7tpsvad553 - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1s4ukn0a5cyl9gyslk4kxraaktftt2l7l95x3zdudkleg727e88esxgr03yd - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1s4ruz5ph7uel2dcs6wyleymxwe2py53kc365xd3qp89m94mtg9wdcydrxtp - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1skx4dhkazyaskslqj2st8wy4hv5c4tnukrpwtrgqnzle503cy78t5h6uvvy - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1shvjnxhjjvn58gqhngarudxrl0pt3r2dgx695hz2uwkqexpukcapgj70hn4 - 5 - 49999999950000
2019/10/15 20:51:44 jnode_ta1sh0s7ulczf6eeytxvy9k9jp7ys76skmvmz3fzwdg3gqmvr7e4tmgzju7k6d - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1skpsaglmfgpuaz9g6ugmr5hakmyr074pkyya9f46ptclx5ql88nvj7s6cfe - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1s5kfrqnjpk6fy42ae4dj9c77hrauyves5unkkvvf23u7wwqayzj7zyrtkvr - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1s5slzkz43jxfzdgs6qlqppsf5dh5vt0fzaa44kq6eg5r97jqveaz7qy5d0t - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1sh794ctzes8s99ht88xrq8vza2slf489cty2vpc094txdyaaurtt5j970e2 - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1sk43xm28hk3khpqe2fqp54xfy5dl90tjz9a8pyhts40pewsgpuvj75nftrs - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1sktaew3xc65k8t3pt9vtyt6wdavmugt4pfj85cpjsgpwvz63xcg8cjvhrsk - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1shyldsf4mtkg60lqtzjz6tk4h5077caqx8kzwwat5c9mv9kj275mqxcfpsz - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1s4cfacejps39069h0dvv8e0ngwrwjg94hqrdhfd9azeeydxupppwu5sz3lt - 5 - 49999999950000
2019/10/15 20:51:45 jnode_ta1sh3erdxaq05a8u98ssv86ck9tvep298et7q4aeys328jtmhqnv5yvs0z6rx - 5 - 49999999950000
2019/10/15 20:51:46 jnode_ta1sklhnnlws730mmnqml236tx9cayn4rmm76sveyanl06asm0c9lef7ukgu24 - 5 - 49999999950000
2019/10/15 20:51:46 jnode_ta1shuwmr8ehdt7qw98aa5jquyv3nrlfakhxa500hc3gzp2s6v2ntj86ymhew4 - 5 - 49999999950000
2019/10/15 20:51:46 jnode_ta1s40dl0pdmj5zez85usrhzjtuy36ppc97024938g3jrjhp3armm5jxl53k5p - 5 - 49999999950000
2019/10/15 20:51:46 jnode_ta1s5frku4y676aytssae0xzh0awattfg5gxcr279payz0nt9ectxn55s38yw3 - 5 - 49999999950000
```

then sends the transactions an reports back **account address** - **fragment_id**

```log
2019/10/15 20:51:46 jnode_ta1skjryzc2x6a4dwpxmkygznf3czvpcmuequ9gqca0hcufqp87vw5a5y3lgal - 59711dcac52bbbdc5d581eb9af23651a022042a8d5d34ced11510e922633a1a2
2019/10/15 20:51:46 jnode_ta1s4pmyagmqdanyfjtdrz9jj8jrf84dyqhnf2w67man3xlspt842cquhc9xek - f955eb28278fdc40073197815b8290925418d96483fcdf2b78cd756de3a80999
2019/10/15 20:51:46 jnode_ta1sk9ydx0xy3q4pz477yfc2gx33vlgepa0zjqdujmsm83n4zx8xwn4yf5tzdd - f2dfa92696a1a9349226979db80246e958cb45ace6a41785f46b3f0568214275
2019/10/15 20:51:46 jnode_ta1s53w0sa28fw0ykreyw7nyclgl3rjprw7mkccpk084k8rmg7un474gsg862x - b2665188d423a3b383bd9bcbd3fe4ccd322d88158e9be7e5ee16bb4e74a26529
2019/10/15 20:51:46 jnode_ta1sh2xfzwxrqlxal6vnmhqtzkm3k4nvum8rlk3fd2ksudpk0las39466an2tz - c35949c830ba08d2f1bc2b5187ba1dd035581544803aa2bdf935b44425d82223
2019/10/15 20:51:46 jnode_ta1s47g6peqfzfxmhzhc07xqrlhqmq3f0q24v6nr8a42lqduegtj9kvcfkm3hn - b578c53529852b6468daa17707bd05e11239f42d4b12f7b7b23c1bb86001bfad
2019/10/15 20:51:46 jnode_ta1s4gnnljx4wa7rfxr0swxqln4wg9s4vj8kt3xsa5djmya984zea5h2t8vm3k - 9b2377b35cccacf19c873c7ee565709c372c2dba7d8d6f899c9af070996241be
2019/10/15 20:51:46 jnode_ta1s5au5mf0axuppq9qsfe90kwne0jjtjdkh93pmfplap56yue24kvz6cljd03 - 4676875a15ff346b19d40e52ef1b335787dc7b945292fdeff6fd314c6ca1bf11
2019/10/15 20:51:46 jnode_ta1s47zzz8nrwww9z4w2yl9d3yjckxap6dyd5u27pryuyk8vpje0jkzjzkjnka - e002803f8670b28b885b4fcea22cfefcb143693d7a6f94634dad7d20a40387d5
2019/10/15 20:51:46 jnode_ta1s4cca439qvt7dazppu9uuvl8xpwf45dtxqmh2q3wgxr8tx4hjg2wyhdthuu - c01e9e0bb03259ed1e16b61eaee744d87f4136ddaf17bd6dee0c5b0e0a0d2066
2019/10/15 20:51:46 jnode_ta1s5xlflm5ytzvzhr6p6m9gj2qlw03g20qwdgwxnsv2ll4kj0jthvgzgfulva - d6ec5151e8eb05b771ee93afa5b36452efb4e6e97c341154109a342d444b83c9
2019/10/15 20:51:46 jnode_ta1shhagrmnc2d2gc89me7el5d57706e8vmvm35wd277u9tq7tnqu9j79h3hvd - 67262d7238d718340dfc9576f0de9691b0dba99e076fd45fa65d47e0fba05f1d
2019/10/15 20:51:46 jnode_ta1s4dd2fkla3deusgweywzge9t0cetm064kllf9gsfvtfpc9nfs7sgcgf4e3v - ed36cc87f73d02f03685a47918434eae3b811c5ec530f8873b2049aa899725ce
2019/10/15 20:51:47 jnode_ta1skkf06ftnlexqwvx0fzgs82vqnrnm9s84adhr89wu5q0ly3at5xuxtlepl3 - 435cbfe4db2d9873557a184244e766a467d618f7da6941a4c9b5a76951a558e1
2019/10/15 20:51:47 jnode_ta1s4mtn4jmefk00qfmv7j3t6vesqunarnsmht9herk7yuf8x4mzs8cztt0n58 - e315d7728a8b50da286dd3baf67e9a140d8c18643b4de904a9246d1e57c12de7
2019/10/15 20:51:47 jnode_ta1shldwftcpng0ah3kzrc7ufzl39uafsjkrmk8yu7cum05r9fuslj65d7m8ap - ea7ee92f87db53a22e2c4cafb8b2688961203e9045e919ddbabf2d3eab7a02a9
2019/10/15 20:51:47 jnode_ta1shykqlluq8ass3kgj5zalhrmwlf4d2qd4lam5rhsgcw2urjy8rwpc6qh6f6 - 458fd8e027f61db1cd2829bf69ca96516961e4bb61d44814671fa6d34bed4ac6
2019/10/15 20:51:47 jnode_ta1s49dckh3gqtfzqrfe5t4x93j6kzhq4ssuc9zfdk5fdwfd6gat8c4j2xp2x7 - 75fa6b55eb7810f763d0d948dcebce4d12bb5c0a6e8fb312fa43707b915b61c5
2019/10/15 20:51:47 jnode_ta1sk4l6zqcc4tgsfp0akeknpkd35jrrsz97827z35mg36wrc6fumhxw5tay3z - b13587f0b94c0ab677d84242b3541c7df98c882d50a13a49319d100040024a6f
2019/10/15 20:51:47 jnode_ta1skxqlkz49rrnpwd3heapa76rnrp2adent77w4asrq4sayw3wwpp9jfcn5k8 - f49dd0843f628a58868d4841c5dd14d96b0d931bb4414f41b154302eaea353cb
2019/10/15 20:51:47 jnode_ta1sh8jfxjgtmdym04x3tr0tzugxw0a7hqsj0ydr40xt2rz69hsxy99g0m7e0s - 1b98b080a67e19e16e2007652f10af38eb70669592b8fea5f4dfc44b6c0dc85c
2019/10/15 20:51:47 jnode_ta1s4yvyn93sjmr7e7rlr8r4qsq76mfrr5mpumqq8vvzur33uergnu8uudrs99 - 35ec563cb10023d47f24a87bed411a011e136d99e85180efc0e61d8f514240d3
2019/10/15 20:51:47 jnode_ta1s5hpjd5uy5npu6d6pc7kycz3pawkzd9m6kau3kq20mq2easv5pfyj802gv4 - 36c7ae4345db7137f694e6c854076397699702705d2088174f28f4175fcddba7
2019/10/15 20:51:47 jnode_ta1skuwg6kam9wlckh7kzrlul4fq022nfr52rdzd0094550wkmr2ydk7z3dwz0 - c992244b90058e2c54c772c7f5a1e188ff0296f4f66d69caf83531baf12f3ab4
2019/10/15 20:51:47 jnode_ta1s550tzlpy3xhlezde72cwn70xg94924k04luw5nkxfgwhdkf0e9g29x06ff - 91f423939c98a5b694d5cf1c72422cd310b7903063776d370b77c5954558d9fc
2019/10/15 20:51:47 jnode_ta1skfys5mtl9w5chu7dlk7f0p3xtqrvauhw866ykg6jmevtjkdmumj69sl4nq - 75b1279deba88e2f9bdbd2eec8f239148a4ca9c248017b92054be7630743759f
2019/10/15 20:51:47 jnode_ta1s5jvr7hehcme4hlgp0jl30s0jhl0hhl3lm7h406s50u65t7c3myvkklfdj9 - 47397a9ab3d5a2689dd105529993784cb6af91bba55416f1523e4af2211dbe7b
2019/10/15 20:51:47 jnode_ta1s5gmt78jszdcjthq5snrawgyf4rycdjcfkvfkjnxx437pzkvxkxrwkhh537 - 524e66a3b4642b77ba781ae3a57aa3ae230fc964b5178604c94cfd25882679ea
2019/10/15 20:51:47 jnode_ta1sh8fjw43kxwzmu6stas4rw4fvsgwd9mtnrdhykdas6833nu05ex6swx55ja - c394bef674f190a1cba0ba548f2b7b97283ba4f06e9e83202bb7187cba6615df
2019/10/15 20:51:47 jnode_ta1skzlvder8cv6fsnuxqm6dkrgra97py2jprkxcgdzlrmu5m6yglvnxs6qqwm - 9403edf67492298c9a90a8172e39e553e0870a2b2760e5f8b6427c378923bd57
2019/10/15 20:51:47 jnode_ta1sk66alarl3xqklwee22xe07kpxxs94nh96r37ye4dl6uwh96zf0m724980u - b172bae4781b5e127391c287296fda50589840ed4aec3eccbfe706b816182f0e
2019/10/15 20:51:47 jnode_ta1sh4rm259rqpuazujn7z7l3d5hcwymgum97g2c4453cnvpmg8athk5dwxkm7 - a05c4158fdd25dc20556bc896145f6d18f1b5667707cd834d3bd193beffb164f
2019/10/15 20:51:47 jnode_ta1s4y5wz08tcfv8xdmdqkf5knxnm7qk9snft3g045txswvycc3vq2m75e7utq - 0e5ff4fb177ac685db443796825de4cc7f74e2efb0290deea16dd57593e61ef9
2019/10/15 20:51:47 jnode_ta1s55zzdkw3r95hdsym29vvlacan40we90pqu5vrempuz0x48cw7tqqswzjmt - 764486c0ef82570fd938d32470940e90252f2cf738506fe7d3302832bb500b7d
2019/10/15 20:51:47 jnode_ta1s43ymn68j9yq6gmdmt745tfdj8pee8ct02qmd9zmd2d3chc6m3dqwrxhkhx - a8aa4c721dd7e045347ff7a7382fd55651e4923b1d944975a2baf0234d60355c
2019/10/15 20:51:47 jnode_ta1skng9thtxeyxkn8hg4038q423e2wkhzym4c3t87hsk5tghtfc7hlyjehx33 - a57bcb74b1ae513d3b80b4c031d2743d1cecf0ddd7a88469992e14afbf0958f0
2019/10/15 20:51:47 jnode_ta1s5jp5xhmqqv2uyzdnx2ge4tsdkxs7gqvy956trtjl2c0fhnkudl6gp46xuf - 3d02520a6bef65047b1b35abbc68805cd38d111904fecbb715e7b3542ab3a348
2019/10/15 20:51:47 jnode_ta1sk59zz4klx6j8md78246rqhk8q0gnnp6gj4qje94yr2c24ggrswgwqm0scu - 2ee44e1fee73f37f12e2cf8b7eca448cc62e3ac8ba54a25f268a48a7851d418c
2019/10/15 20:51:47 jnode_ta1s45xm9nryk8nxdfhvznjuxdav842muqr8c0czntj5500qxkckjylkwa4xa9 - c23287c692647961949e27f1e802915c09743eb34e4cd1480c75e5af8bdf7aad
2019/10/15 20:51:47 jnode_ta1s42v0gfknf648du70rx9qu68t4zfrklnl0ajpwh4mn8u4n3txn5a25wafxc - 802f10435f38c2ee926b7a4e311ede9a5fa432a0dffb614261931aa17a9802d9
2019/10/15 20:51:47 jnode_ta1s5sxsls6warl8q62juflm9ny58jzjyq9tddparqv9xj55p6ymp4lq8vu6y0 - 0c8320d9460ecd4ddd0605f299a835fb13b9e1299e4160b1a006931e18ca11bf
2019/10/15 20:51:47 jnode_ta1sh4pygj8p3d0yasksql7yjwueltrndurcv4ngqar5qxmjmytd22w7zljjnj - 6ad71630ce98869c1d2af4b3225bda1a41122b3334373af35f7601aa283609bc
2019/10/15 20:51:47 jnode_ta1s53qm6edw5trx4zeqak8mcdx0706fnmjnxp68me8y055kq5z4p9t67xwhjf - 55b2e61a7951a0e9126aff4630cead0a9a53c30b66dc1f3cb3a69f2cd4adf5a2
2019/10/15 20:51:48 jnode_ta1skylrf9ulphq6x2hq2vmcflf8kc335uyg80ll72v7pd7md8wwna5g8dquqx - cf44a2e5c6a620a2a9dffd81f0d254000d23cbc4311110ad87c48e37a6c34fce
2019/10/15 20:51:48 jnode_ta1shjz9fnn423jchm8ukf2dtcwmh65czg05d0g3xymepanpsnugk3zxjqgvm0 - 2fd706ff1facfee7483a0f1094d1550327e6bfcd1140f57f38284b06f1074804
2019/10/15 20:51:48 jnode_ta1sk3dc8cxuvad3wg9fkeap5m4fz9v3epfrmk9w0yx0c9cv7k2pvq25zxzwsp - 6aaf6d0c092930e4b8a3ea53f6d0947ed20334535dc242f09dd47c4747acb0ee
2019/10/15 20:51:48 jnode_ta1s5aze7naejjz487l6vz5zuhqf2fn48qve8fxzv5f53mfkjlapaj9cq4fqxk - 700879a74231554bece6430fb79f0cf69f79e1121f2389c2705f84533602d06e
2019/10/15 20:51:48 jnode_ta1s4pp6gntvdg39nju3q9jhy9l20lqajf2l7f0czryfzpl7kkcc3pey3545yh - 439d5f3371c183682cde867c55366c74e6b74515e9c4134f94a11bf52ff90096
2019/10/15 20:51:48 jnode_ta1s5zssp6n2x3yy72jym8r7ypjkxw6np7dgnf6e5yyslml9ddm7qj5s4yfjdw - 92c8a971aba9ec373382c161963fc5421620d72444d64e6ecf5d5869f723f747
2019/10/15 20:51:48 jnode_ta1s40579t28us2kgq5vznjzmaudt2pa9tjq4kg2q38pqesdw9mfy9ws7xke69 - f04d14f59adfb1ccfab723a33d6dd590a41c436f21afc4600ebbb149a7f28139
2019/10/15 20:51:48 jnode_ta1s5nm64z0sk9r8jp9y20nv0eegmq6gm2hpghpnp3z9lzzdk4uerhfx55cua7 - cb9b70764ca3e2e1bfb9ba8507439f4e2fd00d551d8c7e31751dd11fd0286d47
2019/10/15 20:51:48 jnode_ta1s5ec5mymsuvdaeljcg7c7s7x9wcg4dmykf35lu0n0qk4mpmcgrnp706f8e4 - e579712c6686849cc0a75b1131d5e8dac916b004c21b966d50eec6321fb05150
2019/10/15 20:51:48 jnode_ta1s4ahxnflfda06xcdw909npk92c5synlnk7pectq7rs8l8mr6p6wqsuptjp5 - 90d6e6bcf2d9ebdc5e3bd30985aec79235d0e7fd4f835f391b18ddc23c616bc1
2019/10/15 20:51:48 jnode_ta1shv0a8dx6lpj4fr8xup9drf2a4cmpzvlwh5enwjcqaj6v3gel76rwygtyg8 - e15686e98cef80bbe4cc751e543317243846572d9da065526fe65d3dd4b44300
2019/10/15 20:51:48 jnode_ta1s59zyusjregyhhwchmdf0lly3ld53tj0s49lrcya238pdmmmrt50jcwfatx - 95397a570f4926c1a733accbcb3f434258cfb8e2dd57308253d9c92fa1a447a0
2019/10/15 20:51:48 jnode_ta1s5x78szsanpd8smym2rdjvhwlhuz8ccncvj0vyu0mm8zxkc6h6kmzrlz3mj - f2da78fb90ebd26faafe0f615f7e2349e8d9737b7095a925d223a32913567aba
2019/10/15 20:51:48 jnode_ta1skmm0369a8efkf2rntvs5ymvwwfrw9zv6n9zkd8mhv96wscj5fuqgytd2dy - 600560d8f271d1aad7fe59b49e9f7643893d5fec86ac2288ab31f34fd8799453
2019/10/15 20:51:48 jnode_ta1s43s0kp0l85lld0nutvq6xfkkcmmmw3wh46x7r4y57h97u50l5ft5ec5y2g - 4853144e3f61e5211a0c28f909fba01b2389bc18f75c298ec69231f7acf85331
2019/10/15 20:51:48 jnode_ta1s4u3c0sflmltp9lnvy6s76qgvxjgtn0aw3sysyep4p5nkg2h3djywt64zay - c42d7d7a770076f1c7b7cfb68eb822f6d3d63227f4e0315a734650f70b1ef8ea
2019/10/15 20:51:48 jnode_ta1s4e9jas42rcqs3ufcnlp3mvp7nyt3mrf9qee0hca83pdxewlu63nu9qdh5u - 0a626e01cc3fabbb09f03b17633267b9d804d5e89bc4d14ec9f53bec50ab917a
2019/10/15 20:51:48 jnode_ta1s420wwu2hfgwj5tvwava78tugq06evkfw8ngu8wyn2xuwttt78qqg6ayxr5 - f9e2298ba47b000f5676f95262c9072e5f90a8a2641ca2959740d8444c177d37
2019/10/15 20:51:48 jnode_ta1s4267em4razxtdl83538kl5f350h5f353dygp0kgvmk5a82p726v5jt8tex - 4a00bf5a489d3b5306d1be84eab359724ab25c54b2d5329318bb0aa3258cb60d
2019/10/15 20:51:48 jnode_ta1shn9axum2cphz2f66r2sr6uqwjq4mzh2xgkxnyw9ksnt3u47klzy7m9zggk - d222abe74845a4a8ee8d16b8a452e56ac7f4e997b16eec93f888ea65e5e9d695
2019/10/15 20:51:48 jnode_ta1s5tq3s7r3ezmuhh9v9ha5050gltzf7mllx2aejqhck5txhmqzhqqs9y52wf - 6097d16e1240c96f0fe4c996150c5443091f6ce16134e6451508f303dc154118
2019/10/15 20:51:48 jnode_ta1skx3tk6upatfg4s063yzp84jd5g6dp4sdlkzu47gjrfct9nfugvlsgg3hqu - 7078cada20e9938483a3f84164470077f42e43d8f5d947d48d551de6673d95ce
2019/10/15 20:51:48 jnode_ta1skezrvu4m2hklr2dlfthswl2sc649ua8yhrfnz0uhfus9xsufpjryfj6sqw - 7f92c66b4233dd7366720947750cac8b9bdad9ae5817e213c08fa7eec0ad6d68
2019/10/15 20:51:48 jnode_ta1s55g3dunvppflvnu2f6a08akmhwnpyzcpxras0qvnhdwnjk0unzm5s55cpy - 221c43d0840a4d64ed72cbc9ecd4e1362115565deb5efc8ace0649ef9dbc63a6
2019/10/15 20:51:48 jnode_ta1s5t5u3kpdkp6v9muzw276m4uva2eugvv8szzp0ynv4amu7mvhdwsuf9pzkx - 23793609c97ffc82903207b58aaba89e6d31eff5ffbd1af79b045031692f122c
2019/10/15 20:51:48 jnode_ta1s4p7w0689mwantdt9j779g5ajlc3eyq2778a9yn3vf292xq6xla5gp5lh06 - 921244836c5041f9937be631e05911eb93e160f19ecade879b1127ebcc12889b
2019/10/15 20:51:48 jnode_ta1s496dy5e998p733e8h6kam7fwt35x06n9z9lqrwzm43dc0mmj38nkdpf5jx - 994206c4cee07cf5d3f2dbc53860ae198f67a2f9c51be32975c082ebeba07e41
2019/10/15 20:51:48 jnode_ta1she3z799ka4puvfh56z757me4e9nl8akcr6n22j6jyar885wv7ztv3y0mvv - 1d1e0971423107a87f951cf6fbf5734492081b1f9ff2837fb2c1d14ae6ebdefe
2019/10/15 20:51:48 jnode_ta1s5ef5pgcf6su6t79tdnvdhzq206hfyusnl35rf2dn0hvxkhqnfxu624h5me - daa6f5a2ab017740c8810a273248bb5d53456450c49ba1f41baeae482ec3379a
2019/10/15 20:51:48 jnode_ta1shnd447lg057slzcf64l5hr3l4j8pdv7n65cams346mpxnuw5q2q620zt48 - e4146eade2cda879922e99fb39c8ec173521e888d48d9e3c7b112e58bb743f5b
2019/10/15 20:51:48 jnode_ta1s4qtez8tup99kd9q2wh707dsa3rm9q7hvt4aak3sgx4x28ft5f77uj7r5d0 - e4f077ebd1f93c4736afceb3e7bd37d0e4bc80d2b236a7bc0a4e4135c5b607f4
2019/10/15 20:51:49 jnode_ta1sh23ttcfvekxvu6zqvlmf3hmtv8dunsf57c834h0l64tupmv7uaxjhk5xyn - 33a0ecba9d96330d1cd4b5627be78bb9604a5fc3ebc35e895133c6c7ec6c1e1a
2019/10/15 20:51:49 jnode_ta1s5v3r3wh0vmejvpjq3tx3uqduyk0wp0z3devcmjpc4fx9csmm6xu7pdjwqk - 7023b91902f8df3a5b56fb31abe89983c73ce627b44b025c9e87107f52d157ab
2019/10/15 20:51:49 jnode_ta1skg9cnc6ns09vq7re8en2s7ypfnrtk30vvhpx8p8gfj38627dksu2q80rwy - 43b90b10a0d1254385475a557709b48378195d95bb2ef360e83f75dd4c1d53da
2019/10/15 20:51:49 jnode_ta1s4dpw4nm6n59ars56u466say7xlvtyn755qj50mday392k6vthvz7avfd6e - 246865de9c953d6920c5dc98752eb2632612dfd5901aea35001ef49042da40c5
2019/10/15 20:51:49 jnode_ta1s4yq0s08gthuw0mz5gcjmt3vtnkpnsc98hhyudzkta738ryjdlv0uepy5cm - 6b10ec8aa481284a09aaabf6430c3abc575aa7e3e31aaf9fbfa515907624761b
2019/10/15 20:51:49 jnode_ta1s5w39hejw0hcgeghn5rdm96zea8e2e8xfgn9f4a9a5e2t2vgl7fj2t3cp6q - 8781aaab4a49aa9e0866780027fcd733e974bf80ee873852724c3605a570abcd
2019/10/15 20:51:49 jnode_ta1s4ha79eecqv5nj3y87xrxmqw76kn5u0hcywy2zvwerpkt7xyk4h8ymdg44w - 6ace0fbf3c33a47d2f3f526744d323ead669091175b039c4890fcab81639056e
2019/10/15 20:51:49 jnode_ta1skyywuduvxq3y4wsjg59ycvzy2tenldy53kksghcee8g7jywf7tpsvad553 - b4b8f45c59a7dfb44c6abdb0ed92fe792812ffb3826997f67cd5e83883a0b67a
2019/10/15 20:51:49 jnode_ta1s4ukn0a5cyl9gyslk4kxraaktftt2l7l95x3zdudkleg727e88esxgr03yd - b8be626b7b4d9fa0e5b1d7083572bf7e98982c49c3662025f00d884859904588
2019/10/15 20:51:49 jnode_ta1s4ruz5ph7uel2dcs6wyleymxwe2py53kc365xd3qp89m94mtg9wdcydrxtp - 41e6d263b6bafc44b7b628016cf11f4d9ac78a95ce8b0e1babde01d523c00a46
2019/10/15 20:51:49 jnode_ta1skx4dhkazyaskslqj2st8wy4hv5c4tnukrpwtrgqnzle503cy78t5h6uvvy - 45a783acd16b201ba407ed3783f65399568035d291ca08d123c668d18aede56f
2019/10/15 20:51:49 jnode_ta1shvjnxhjjvn58gqhngarudxrl0pt3r2dgx695hz2uwkqexpukcapgj70hn4 - f156e2fc28cde4e07520150d2cb58b5a53f7413a266a7210a6949108ff623838
2019/10/15 20:51:49 jnode_ta1sh0s7ulczf6eeytxvy9k9jp7ys76skmvmz3fzwdg3gqmvr7e4tmgzju7k6d - 4d1255301401d0bcf0665ec8e0a1c53218d8c64055da113a95ad175a0e501a92
2019/10/15 20:51:49 jnode_ta1skpsaglmfgpuaz9g6ugmr5hakmyr074pkyya9f46ptclx5ql88nvj7s6cfe - 8b8294919936ed68e83d2f41b49bbfa1d93efdefe2aad6bf3e3a11db11982ab5
2019/10/15 20:51:49 jnode_ta1s5kfrqnjpk6fy42ae4dj9c77hrauyves5unkkvvf23u7wwqayzj7zyrtkvr - 3b7cbd86e54c21ed3e541f0abb4cc9f79d459a93b9afcace00601ba197bbe025
2019/10/15 20:51:49 jnode_ta1s5slzkz43jxfzdgs6qlqppsf5dh5vt0fzaa44kq6eg5r97jqveaz7qy5d0t - 9120d78a39ec69ac0ce6ab3fde619a5a5b40b969e6b34e47c67852d26c14da2f
2019/10/15 20:51:49 jnode_ta1sh794ctzes8s99ht88xrq8vza2slf489cty2vpc094txdyaaurtt5j970e2 - 7c4bda27014fe7ad535cc7cb2a0d37052c1fd5c01ed79167ff46f00813980b6f
2019/10/15 20:51:49 jnode_ta1sk43xm28hk3khpqe2fqp54xfy5dl90tjz9a8pyhts40pewsgpuvj75nftrs - 6ebab93ee67d03f7089ff21a4b214f79a110515aa3dec614a16662abc2f5a8fa
2019/10/15 20:51:49 jnode_ta1sktaew3xc65k8t3pt9vtyt6wdavmugt4pfj85cpjsgpwvz63xcg8cjvhrsk - 25433380e5c76f780832b2e77df677127b7e9799a38b74c773f9586b09ec859f
2019/10/15 20:51:49 jnode_ta1shyldsf4mtkg60lqtzjz6tk4h5077caqx8kzwwat5c9mv9kj275mqxcfpsz - 7e85e14381c88c76c51037c2abc48344668cf16eaa892ebeb3235fe099f917db
2019/10/15 20:51:49 jnode_ta1s4cfacejps39069h0dvv8e0ngwrwjg94hqrdhfd9azeeydxupppwu5sz3lt - 86ec49bdb187e9444f0859651b709e0472db93eb7bfab423ca55bfda448d1e1c
2019/10/15 20:51:49 jnode_ta1sh3erdxaq05a8u98ssv86ck9tvep298et7q4aeys328jtmhqnv5yvs0z6rx - 84b17011c19f7d41102ce608bf7ec5d4f6229d18c3be0e02d6b7e1c33e1afcaf
2019/10/15 20:51:49 jnode_ta1sklhnnlws730mmnqml236tx9cayn4rmm76sveyanl06asm0c9lef7ukgu24 - b0f06342f11e363b777bdbe9cad9e48c02cdda996f644b79e5062ec572af0d35
2019/10/15 20:51:49 jnode_ta1shuwmr8ehdt7qw98aa5jquyv3nrlfakhxa500hc3gzp2s6v2ntj86ymhew4 - 4995727175eb3c3d1644e7914966a78c2cf7b4ffdfc69d8aace8e4dec3c3095e
2019/10/15 20:51:49 jnode_ta1s40dl0pdmj5zez85usrhzjtuy36ppc97024938g3jrjhp3armm5jxl53k5p - ec5d6c446cd83aa5e468558fc72d22f07d106319ddc72c459aff5cca9d2de77e
2019/10/15 20:51:49 jnode_ta1s5frku4y676aytssae0xzh0awattfg5gxcr279payz0nt9ectxn55s38yw3 - f86f8f07abd5b39b82ff02707d88c5169574d091761776e70bd27c97ce840c06
```
