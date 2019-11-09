# Examples

1) [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
2) [node_genesis_stakepool_run](#genesis-stakepool-node-run)
3) [node_stakepool_create_and_run](#stakepool-node-create-and-run)
4) [node_passive_run](#passive-node-run)
5) [jcli_rest_v0](#jcli-rest-v0)
6) [jcli_bulk_send_sequential](#jcli-bulk-send-sequential)
7) [jcli_bulk_send_concurrent](#jcli-bulk-send-concurrent) - TBD

## Genesis Node Bootstrap And Run

Bootstrap configuration and start a **leader** node with and active local stake pool.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info Genesis Node

```log
2019/11/09 10:33:01 Using: jcli 0.7.0-rc7 (master-cf5fcaea, debug, linux [x86_64]) - [rustc 1.39.0 (4560ea788 2019-11-04)]
2019/11/09 10:33:01 Using: jormungandr 0.7.0-rc7 (master-63192d6e, debug, linux [x86_64]) - [rustc 1.39.0 (4560ea788 2019-11-04)]
2019/11/09 10:33:01
2019/11/09 10:33:01 Working Directory: /tmp/jnode_genesis_223862143
2019/11/09 10:33:07
2019/11/09 10:33:07 Genesis Hash: 116f3e765a825a68dc1ac0a3f8993447dccef5641b0450e31dbe0a2cf1c79cad
2019/11/09 10:33:07
2019/11/09 10:33:07 LOCAL StakePool ID       : a93cf67dac50f84f2b74f3cccad1c21a2df2c364037e5dc1dd8017c1d320fc9d
2019/11/09 10:33:07 LOCAL StakePool Owner    : jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/11/09 10:33:07 LOCAL StakePool Owner    : jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/11/09 10:33:07 LOCAL StakePool Delegator: jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/11/09 10:33:07 LOCAL StakePool Delegator: jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/11/09 10:33:07
2019/11/09 10:33:07 EXTRA StakePool ID       : 2f3471d99a42e3c75362ea3a217f190143675300ca9893edd86519eff118c9aa
2019/11/09 10:33:07 EXTRA StakePool Owner    : jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/11/09 10:33:07 EXTRA StakePool Delegator: jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/11/09 10:33:07
2019/11/09 10:33:07 NodePublicID for trusted: 111111111111111111111111111111111111111111111111
2019/11/09 10:33:07
2019/11/09 10:33:07 Genesis Node - Running...
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
2019/11/09 10:33:56 Using: jcli 0.7.0-rc7 (master-cf5fcaea, debug, linux [x86_64]) - [rustc 1.39.0 (4560ea788 2019-11-04)]
2019/11/09 10:33:56 Using: jormungandr 0.7.0-rc7 (master-63192d6e, debug, linux [x86_64]) - [rustc 1.39.0 (4560ea788 2019-11-04)]
2019/11/09 10:33:56
2019/11/09 10:33:56 Working Directory: /tmp/jnode_gepstake_023000790
2019/11/09 10:33:58
2019/11/09 10:33:58 Genesis Hash: 116f3e765a825a68dc1ac0a3f8993447dccef5641b0450e31dbe0a2cf1c79cad
2019/11/09 10:33:58
2019/11/09 10:33:58 StakePool ID       : 2f3471d99a42e3c75362ea3a217f190143675300ca9893edd86519eff118c9aa
2019/11/09 10:33:58 StakePool Owner    : jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/11/09 10:33:58 StakePool Delegator: jnode_ta1s5rkmsfsra5chatzcjmdmh5nsu8rpar6x2ly3gr85q332ckr4quqvy8sthh
2019/11/09 10:33:58
2019/11/09 10:33:58 NodePublicID for trusted: 222222222222222222222222222222222222222222222222
2019/11/09 10:33:58
2019/11/09 10:33:58 Genesis StakePool Node - Running...
```

## StakePool Node Create And Run

Creates a StakePool node with appropriate configuration,
register it to the network, delegate to it and run:

- `--genesis-block-hash` from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--secret` self generate pool secrets
- `--trusted-peer` the node from [node_genesis_stakepool_run](#genesis-stakepool-node-run)

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)
- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info StakePool Node

```log
2019/11/09 10:34:30 Using: jcli 0.7.0-rc7 (master-cf5fcaea, debug, linux [x86_64]) - [rustc 1.39.0 (4560ea788 2019-11-04)]
2019/11/09 10:34:30 Using: jormungandr 0.7.0-rc7 (master-63192d6e, debug, linux [x86_64]) - [rustc 1.39.0 (4560ea788 2019-11-04)]
2019/11/09 10:34:30
2019/11/09 10:34:30 Working Directory: /tmp/jnode_dynstake_937920717
2019/11/09 10:34:30
2019/11/09 10:34:32 Waiting for rest interface...
2019/11/09 10:34:33 ...Node state [Running]
2019/11/09 10:34:34 SelfTip: 94678ddf09e18f1e2e68f48dfc11feb36b54acd2386d83afa6ff238bdaeec4c6
2019/11/09 10:34:34 Wait for pool registration certificate transaction [0c0edd00ae38fde833f719a5aca299d954e6ee3fe96597fe954855c033870ec0] status change...
2019/11/09 10:35:45 FragmentID: 0c0edd00ae38fde833f719a5aca299d954e6ee3fe96597fe954855c033870ec0 - InABlock [222175.22 (98d90b9e1d1f36a44f1a43c8b21185cbf8807ff1547296389b4ab297095274bf)]
2019/11/09 10:35:45 Wait for delegation certificate transaction [8d6ff3a979e51017b0a083a100337516ee0c70da89334182ab12530df767b1ec] status change...
2019/11/09 10:35:50 FragmentID: 8d6ff3a979e51017b0a083a100337516ee0c70da89334182ab12530df767b1ec - InABlock [222175.24 (98f54e539eb0596fd4f2ae699688753128894d51c562890d0a1511c046190aba)]
2019/11/09 10:35:50
2019/11/09 10:35:50 Genesis Hash: 116f3e765a825a68dc1ac0a3f8993447dccef5641b0450e31dbe0a2cf1c79cad
2019/11/09 10:35:50
2019/11/09 10:35:50 StakePool ID       : 20ba61c4d0b044962ada2536ab703eeaf95ddf7b90d9900a737988b80abb9415
2019/11/09 10:35:50 StakePool Owner    : jnode_ta1shz8a85d3xhu76n0k9s99ss8v69nf8dnqagly4ljndzr9pqyg6ktu9syl8c
2019/11/09 10:35:50 StakePool Owner    : jnode_ta1sk6pmqy3lfrr7kq4afmywn5hl9prurwy7xfqejjgazlg9r5nnmk26vjfs3z
2019/11/09 10:35:50 StakePool Delegator: jnode_ta1s5a8e4qye5rwttc9qrek0e30htttmpvvuf967mdp35pcx80t6e2psskthdh
2019/11/09 10:35:50
2019/11/09 10:35:50 NodePublicID for trusted: 333333333333333333333333333333333333333333333333
2019/11/09 10:35:50
2019/11/09 10:35:50 Delegator StakePool Node - Running...
```

## Passive Node Run

Start a passive node with:

- `--genesis-block-hash` from [node_genesis_bootstrap_and_run](#genesis-node-bootstrap-and-run)
- `--trusted-peer` the node from [node_stakepool_create_and_run](#stakepool-node-create-and-run)

It shows the usage of:

- [jnode](https://godoc.org/github.com/rinor/jorcli/jnode)

### Info Passive Node

```log
2019/11/09 10:34:52 Using: jcli 0.7.0-rc7 (master-cf5fcaea, debug, linux [x86_64]) - [rustc 1.39.0 (4560ea788 2019-11-04)]
2019/11/09 10:34:52 Using: jormungandr 0.7.0-rc7 (master-63192d6e, debug, linux [x86_64]) - [rustc 1.39.0 (4560ea788 2019-11-04)]
2019/11/09 10:34:52
2019/11/09 10:34:52 Working Directory: /tmp/jnode_passive_946502118
2019/11/09 10:34:52
2019/11/09 10:34:52 Genesis Hash: 116f3e765a825a68dc1ac0a3f8993447dccef5641b0450e31dbe0a2cf1c79cad
2019/11/09 10:34:52
2019/11/09 10:34:52 NodePublicID for trusted: 444444444444444444444444444444444444444444444444
2019/11/09 10:34:52
2019/11/09 10:34:52 Passive/Explorer Node - Running...
```

## JCLI Rest v0

Query the node using rest capabilies of `jcli`.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)

## JCLI Bulk send sequential

Bulk send 100 transaction to 1 address (delegator address).
This is a sequential version in order not to stress too much the rest interface.

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)

Example log showing first the **account address** - **spending counter** - **account balance**

```log
2019/11/09 10:49:10 jnode_ta1skjryzc2x6a4dwpxmkygznf3czvpcmuequ9gqca0hcufqp87vw5a5y3lgal - 0 - 50000000000000
2019/11/09 10:49:10 jnode_ta1s4pmyagmqdanyfjtdrz9jj8jrf84dyqhnf2w67man3xlspt842cquhc9xek - 0 - 50000000000000
2019/11/09 10:49:10 jnode_ta1sk9ydx0xy3q4pz477yfc2gx33vlgepa0zjqdujmsm83n4zx8xwn4yf5tzdd - 0 - 50000000000000
2019/11/09 10:49:10 jnode_ta1s53w0sa28fw0ykreyw7nyclgl3rjprw7mkccpk084k8rmg7un474gsg862x - 0 - 50000000000000
2019/11/09 10:49:10 jnode_ta1sh2xfzwxrqlxal6vnmhqtzkm3k4nvum8rlk3fd2ksudpk0las39466an2tz - 0 - 50000000000000
2019/11/09 10:49:10 jnode_ta1s47g6peqfzfxmhzhc07xqrlhqmq3f0q24v6nr8a42lqduegtj9kvcfkm3hn - 0 - 50000000000000
2019/11/09 10:49:11 jnode_ta1s4gnnljx4wa7rfxr0swxqln4wg9s4vj8kt3xsa5djmya984zea5h2t8vm3k - 0 - 50000000000000
2019/11/09 10:49:11 jnode_ta1s5au5mf0axuppq9qsfe90kwne0jjtjdkh93pmfplap56yue24kvz6cljd03 - 0 - 50000000000000
2019/11/09 10:49:11 jnode_ta1s47zzz8nrwww9z4w2yl9d3yjckxap6dyd5u27pryuyk8vpje0jkzjzkjnka - 0 - 50000000000000
2019/11/09 10:49:11 jnode_ta1s4cca439qvt7dazppu9uuvl8xpwf45dtxqmh2q3wgxr8tx4hjg2wyhdthuu - 0 - 50000000000000
2019/11/09 10:49:11 jnode_ta1s5xlflm5ytzvzhr6p6m9gj2qlw03g20qwdgwxnsv2ll4kj0jthvgzgfulva - 0 - 50000000000000
2019/11/09 10:49:11 jnode_ta1shhagrmnc2d2gc89me7el5d57706e8vmvm35wd277u9tq7tnqu9j79h3hvd - 0 - 50000000000000
2019/11/09 10:49:11 jnode_ta1s4dd2fkla3deusgweywzge9t0cetm064kllf9gsfvtfpc9nfs7sgcgf4e3v - 0 - 50000000000000
2019/11/09 10:49:11 jnode_ta1skkf06ftnlexqwvx0fzgs82vqnrnm9s84adhr89wu5q0ly3at5xuxtlepl3 - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1s4mtn4jmefk00qfmv7j3t6vesqunarnsmht9herk7yuf8x4mzs8cztt0n58 - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1shldwftcpng0ah3kzrc7ufzl39uafsjkrmk8yu7cum05r9fuslj65d7m8ap - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1shykqlluq8ass3kgj5zalhrmwlf4d2qd4lam5rhsgcw2urjy8rwpc6qh6f6 - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1s49dckh3gqtfzqrfe5t4x93j6kzhq4ssuc9zfdk5fdwfd6gat8c4j2xp2x7 - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1sk4l6zqcc4tgsfp0akeknpkd35jrrsz97827z35mg36wrc6fumhxw5tay3z - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1skxqlkz49rrnpwd3heapa76rnrp2adent77w4asrq4sayw3wwpp9jfcn5k8 - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1sh8jfxjgtmdym04x3tr0tzugxw0a7hqsj0ydr40xt2rz69hsxy99g0m7e0s - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1s4yvyn93sjmr7e7rlr8r4qsq76mfrr5mpumqq8vvzur33uergnu8uudrs99 - 0 - 50000000000000
2019/11/09 10:49:12 jnode_ta1s5hpjd5uy5npu6d6pc7kycz3pawkzd9m6kau3kq20mq2easv5pfyj802gv4 - 0 - 50000000000000
2019/11/09 10:49:13 jnode_ta1skuwg6kam9wlckh7kzrlul4fq022nfr52rdzd0094550wkmr2ydk7z3dwz0 - 0 - 50000000000000
2019/11/09 10:49:13 jnode_ta1s550tzlpy3xhlezde72cwn70xg94924k04luw5nkxfgwhdkf0e9g29x06ff - 0 - 50000000000000
2019/11/09 10:49:13 jnode_ta1skfys5mtl9w5chu7dlk7f0p3xtqrvauhw866ykg6jmevtjkdmumj69sl4nq - 0 - 50000000000000
2019/11/09 10:49:13 jnode_ta1s5jvr7hehcme4hlgp0jl30s0jhl0hhl3lm7h406s50u65t7c3myvkklfdj9 - 0 - 50000000000000
2019/11/09 10:49:13 jnode_ta1s5gmt78jszdcjthq5snrawgyf4rycdjcfkvfkjnxx437pzkvxkxrwkhh537 - 0 - 50000000000000
2019/11/09 10:49:13 jnode_ta1sh8fjw43kxwzmu6stas4rw4fvsgwd9mtnrdhykdas6833nu05ex6swx55ja - 0 - 50000000000000
2019/11/09 10:49:13 jnode_ta1skzlvder8cv6fsnuxqm6dkrgra97py2jprkxcgdzlrmu5m6yglvnxs6qqwm - 0 - 50000000000000
2019/11/09 10:49:13 jnode_ta1sk66alarl3xqklwee22xe07kpxxs94nh96r37ye4dl6uwh96zf0m724980u - 0 - 50000000000000
2019/11/09 10:49:14 jnode_ta1sh4rm259rqpuazujn7z7l3d5hcwymgum97g2c4453cnvpmg8athk5dwxkm7 - 0 - 50000000000000
2019/11/09 10:49:14 jnode_ta1s4y5wz08tcfv8xdmdqkf5knxnm7qk9snft3g045txswvycc3vq2m75e7utq - 0 - 50000000000000
2019/11/09 10:49:14 jnode_ta1s55zzdkw3r95hdsym29vvlacan40we90pqu5vrempuz0x48cw7tqqswzjmt - 0 - 50000000000000
2019/11/09 10:49:14 jnode_ta1s43ymn68j9yq6gmdmt745tfdj8pee8ct02qmd9zmd2d3chc6m3dqwrxhkhx - 0 - 50000000000000
2019/11/09 10:49:14 jnode_ta1skng9thtxeyxkn8hg4038q423e2wkhzym4c3t87hsk5tghtfc7hlyjehx33 - 0 - 50000000000000
2019/11/09 10:49:14 jnode_ta1s5jp5xhmqqv2uyzdnx2ge4tsdkxs7gqvy956trtjl2c0fhnkudl6gp46xuf - 0 - 50000000000000
2019/11/09 10:49:14 jnode_ta1sk59zz4klx6j8md78246rqhk8q0gnnp6gj4qje94yr2c24ggrswgwqm0scu - 0 - 50000000000000
2019/11/09 10:49:14 jnode_ta1s45xm9nryk8nxdfhvznjuxdav842muqr8c0czntj5500qxkckjylkwa4xa9 - 0 - 50000000000000
2019/11/09 10:49:15 jnode_ta1s42v0gfknf648du70rx9qu68t4zfrklnl0ajpwh4mn8u4n3txn5a25wafxc - 0 - 50000000000000
2019/11/09 10:49:15 jnode_ta1s5sxsls6warl8q62juflm9ny58jzjyq9tddparqv9xj55p6ymp4lq8vu6y0 - 0 - 50000000000000
2019/11/09 10:49:15 jnode_ta1sh4pygj8p3d0yasksql7yjwueltrndurcv4ngqar5qxmjmytd22w7zljjnj - 0 - 50000000000000
2019/11/09 10:49:15 jnode_ta1s53qm6edw5trx4zeqak8mcdx0706fnmjnxp68me8y055kq5z4p9t67xwhjf - 0 - 50000000000000
2019/11/09 10:49:15 jnode_ta1skylrf9ulphq6x2hq2vmcflf8kc335uyg80ll72v7pd7md8wwna5g8dquqx - 0 - 50000000000000
2019/11/09 10:49:15 jnode_ta1shjz9fnn423jchm8ukf2dtcwmh65czg05d0g3xymepanpsnugk3zxjqgvm0 - 0 - 50000000000000
2019/11/09 10:49:15 jnode_ta1sk3dc8cxuvad3wg9fkeap5m4fz9v3epfrmk9w0yx0c9cv7k2pvq25zxzwsp - 0 - 50000000000000
2019/11/09 10:49:15 jnode_ta1s5aze7naejjz487l6vz5zuhqf2fn48qve8fxzv5f53mfkjlapaj9cq4fqxk - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1s4pp6gntvdg39nju3q9jhy9l20lqajf2l7f0czryfzpl7kkcc3pey3545yh - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1s5zssp6n2x3yy72jym8r7ypjkxw6np7dgnf6e5yyslml9ddm7qj5s4yfjdw - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1s40579t28us2kgq5vznjzmaudt2pa9tjq4kg2q38pqesdw9mfy9ws7xke69 - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1s5nm64z0sk9r8jp9y20nv0eegmq6gm2hpghpnp3z9lzzdk4uerhfx55cua7 - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1s5ec5mymsuvdaeljcg7c7s7x9wcg4dmykf35lu0n0qk4mpmcgrnp706f8e4 - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1s4ahxnflfda06xcdw909npk92c5synlnk7pectq7rs8l8mr6p6wqsuptjp5 - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1shv0a8dx6lpj4fr8xup9drf2a4cmpzvlwh5enwjcqaj6v3gel76rwygtyg8 - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1s59zyusjregyhhwchmdf0lly3ld53tj0s49lrcya238pdmmmrt50jcwfatx - 0 - 50000000000000
2019/11/09 10:49:16 jnode_ta1s5x78szsanpd8smym2rdjvhwlhuz8ccncvj0vyu0mm8zxkc6h6kmzrlz3mj - 0 - 50000000000000
2019/11/09 10:49:17 jnode_ta1skmm0369a8efkf2rntvs5ymvwwfrw9zv6n9zkd8mhv96wscj5fuqgytd2dy - 0 - 50000000000000
2019/11/09 10:49:17 jnode_ta1s43s0kp0l85lld0nutvq6xfkkcmmmw3wh46x7r4y57h97u50l5ft5ec5y2g - 0 - 50000000000000
2019/11/09 10:49:17 jnode_ta1s4u3c0sflmltp9lnvy6s76qgvxjgtn0aw3sysyep4p5nkg2h3djywt64zay - 0 - 50000000000000
2019/11/09 10:49:17 jnode_ta1s4e9jas42rcqs3ufcnlp3mvp7nyt3mrf9qee0hca83pdxewlu63nu9qdh5u - 0 - 50000000000000
2019/11/09 10:49:17 jnode_ta1s420wwu2hfgwj5tvwava78tugq06evkfw8ngu8wyn2xuwttt78qqg6ayxr5 - 0 - 50000000000000
2019/11/09 10:49:17 jnode_ta1s4267em4razxtdl83538kl5f350h5f353dygp0kgvmk5a82p726v5jt8tex - 0 - 50000000000000
2019/11/09 10:49:17 jnode_ta1shn9axum2cphz2f66r2sr6uqwjq4mzh2xgkxnyw9ksnt3u47klzy7m9zggk - 0 - 50000000000000
2019/11/09 10:49:17 jnode_ta1s5tq3s7r3ezmuhh9v9ha5050gltzf7mllx2aejqhck5txhmqzhqqs9y52wf - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1skx3tk6upatfg4s063yzp84jd5g6dp4sdlkzu47gjrfct9nfugvlsgg3hqu - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1skezrvu4m2hklr2dlfthswl2sc649ua8yhrfnz0uhfus9xsufpjryfj6sqw - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1s55g3dunvppflvnu2f6a08akmhwnpyzcpxras0qvnhdwnjk0unzm5s55cpy - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1s5t5u3kpdkp6v9muzw276m4uva2eugvv8szzp0ynv4amu7mvhdwsuf9pzkx - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1s4p7w0689mwantdt9j779g5ajlc3eyq2778a9yn3vf292xq6xla5gp5lh06 - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1s496dy5e998p733e8h6kam7fwt35x06n9z9lqrwzm43dc0mmj38nkdpf5jx - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1she3z799ka4puvfh56z757me4e9nl8akcr6n22j6jyar885wv7ztv3y0mvv - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1s5ef5pgcf6su6t79tdnvdhzq206hfyusnl35rf2dn0hvxkhqnfxu624h5me - 0 - 50000000000000
2019/11/09 10:49:18 jnode_ta1shnd447lg057slzcf64l5hr3l4j8pdv7n65cams346mpxnuw5q2q620zt48 - 0 - 50000000000000
2019/11/09 10:49:19 jnode_ta1s4qtez8tup99kd9q2wh707dsa3rm9q7hvt4aak3sgx4x28ft5f77uj7r5d0 - 0 - 50000000000000
2019/11/09 10:49:19 jnode_ta1sh23ttcfvekxvu6zqvlmf3hmtv8dunsf57c834h0l64tupmv7uaxjhk5xyn - 0 - 50000000000000
2019/11/09 10:49:19 jnode_ta1s5v3r3wh0vmejvpjq3tx3uqduyk0wp0z3devcmjpc4fx9csmm6xu7pdjwqk - 0 - 50000000000000
2019/11/09 10:49:19 jnode_ta1skg9cnc6ns09vq7re8en2s7ypfnrtk30vvhpx8p8gfj38627dksu2q80rwy - 0 - 50000000000000
2019/11/09 10:49:19 jnode_ta1s4dpw4nm6n59ars56u466say7xlvtyn755qj50mday392k6vthvz7avfd6e - 0 - 50000000000000
2019/11/09 10:49:19 jnode_ta1s4yq0s08gthuw0mz5gcjmt3vtnkpnsc98hhyudzkta738ryjdlv0uepy5cm - 0 - 50000000000000
2019/11/09 10:49:19 jnode_ta1s5w39hejw0hcgeghn5rdm96zea8e2e8xfgn9f4a9a5e2t2vgl7fj2t3cp6q - 0 - 50000000000000
2019/11/09 10:49:19 jnode_ta1s4ha79eecqv5nj3y87xrxmqw76kn5u0hcywy2zvwerpkt7xyk4h8ymdg44w - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1skyywuduvxq3y4wsjg59ycvzy2tenldy53kksghcee8g7jywf7tpsvad553 - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1s4ukn0a5cyl9gyslk4kxraaktftt2l7l95x3zdudkleg727e88esxgr03yd - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1s4ruz5ph7uel2dcs6wyleymxwe2py53kc365xd3qp89m94mtg9wdcydrxtp - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1skx4dhkazyaskslqj2st8wy4hv5c4tnukrpwtrgqnzle503cy78t5h6uvvy - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1shvjnxhjjvn58gqhngarudxrl0pt3r2dgx695hz2uwkqexpukcapgj70hn4 - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1sh0s7ulczf6eeytxvy9k9jp7ys76skmvmz3fzwdg3gqmvr7e4tmgzju7k6d - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1skpsaglmfgpuaz9g6ugmr5hakmyr074pkyya9f46ptclx5ql88nvj7s6cfe - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1s5kfrqnjpk6fy42ae4dj9c77hrauyves5unkkvvf23u7wwqayzj7zyrtkvr - 0 - 50000000000000
2019/11/09 10:49:20 jnode_ta1s5slzkz43jxfzdgs6qlqppsf5dh5vt0fzaa44kq6eg5r97jqveaz7qy5d0t - 0 - 50000000000000
2019/11/09 10:49:21 jnode_ta1sh794ctzes8s99ht88xrq8vza2slf489cty2vpc094txdyaaurtt5j970e2 - 0 - 50000000000000
2019/11/09 10:49:21 jnode_ta1sk43xm28hk3khpqe2fqp54xfy5dl90tjz9a8pyhts40pewsgpuvj75nftrs - 0 - 50000000000000
2019/11/09 10:49:21 jnode_ta1sktaew3xc65k8t3pt9vtyt6wdavmugt4pfj85cpjsgpwvz63xcg8cjvhrsk - 0 - 50000000000000
2019/11/09 10:49:21 jnode_ta1shyldsf4mtkg60lqtzjz6tk4h5077caqx8kzwwat5c9mv9kj275mqxcfpsz - 0 - 50000000000000
2019/11/09 10:49:21 jnode_ta1s4cfacejps39069h0dvv8e0ngwrwjg94hqrdhfd9azeeydxupppwu5sz3lt - 0 - 50000000000000
2019/11/09 10:49:21 jnode_ta1sh3erdxaq05a8u98ssv86ck9tvep298et7q4aeys328jtmhqnv5yvs0z6rx - 0 - 50000000000000
2019/11/09 10:49:21 jnode_ta1sklhnnlws730mmnqml236tx9cayn4rmm76sveyanl06asm0c9lef7ukgu24 - 0 - 50000000000000
2019/11/09 10:49:21 jnode_ta1shuwmr8ehdt7qw98aa5jquyv3nrlfakhxa500hc3gzp2s6v2ntj86ymhew4 - 0 - 50000000000000
2019/11/09 10:49:22 jnode_ta1s40dl0pdmj5zez85usrhzjtuy36ppc97024938g3jrjhp3armm5jxl53k5p - 0 - 50000000000000
2019/11/09 10:49:22 jnode_ta1s5frku4y676aytssae0xzh0awattfg5gxcr279payz0nt9ectxn55s38yw3 - 0 - 50000000000000
```

then sends the transactions an reports back **account address** - **fragment_id**

```log
2019/11/09 10:49:22 jnode_ta1skjryzc2x6a4dwpxmkygznf3czvpcmuequ9gqca0hcufqp87vw5a5y3lgal - 35e273e06b262f04f1737511075f90440fab1a13c89a5f1fb49d6bbc5c220123
2019/11/09 10:49:22 jnode_ta1s4pmyagmqdanyfjtdrz9jj8jrf84dyqhnf2w67man3xlspt842cquhc9xek - 37f02767e33b6889f65d8db42161dd1f4b8df8a61beb3bdae555d4b173302b5d
2019/11/09 10:49:22 jnode_ta1sk9ydx0xy3q4pz477yfc2gx33vlgepa0zjqdujmsm83n4zx8xwn4yf5tzdd - fce7d352e29adb41553f761583af4f4e3a818d8c99b58641b849323a5284f052
2019/11/09 10:49:22 jnode_ta1s53w0sa28fw0ykreyw7nyclgl3rjprw7mkccpk084k8rmg7un474gsg862x - 73c26cc36f69e3732619d1f2cb0ad2a943929d847dfd7346ae6128ad926cb64f
2019/11/09 10:49:22 jnode_ta1sh2xfzwxrqlxal6vnmhqtzkm3k4nvum8rlk3fd2ksudpk0las39466an2tz - 9b549541839a93712f86e54cfd2464199deacacf23f1626008cb9ebba62c6bc8
2019/11/09 10:49:22 jnode_ta1s47g6peqfzfxmhzhc07xqrlhqmq3f0q24v6nr8a42lqduegtj9kvcfkm3hn - 1988777978a6c9a92ee6b1925b5722fca43f7fe50572c839593443fbb3d257eb
2019/11/09 10:49:22 jnode_ta1s4gnnljx4wa7rfxr0swxqln4wg9s4vj8kt3xsa5djmya984zea5h2t8vm3k - 9567d8f7826dd387eaa0556af79b312a33409de5c2ddba7c0e22f086ed56a351
2019/11/09 10:49:22 jnode_ta1s5au5mf0axuppq9qsfe90kwne0jjtjdkh93pmfplap56yue24kvz6cljd03 - 7ae56411dcd701118a6e024830fcb38365bc0f44b4e987f1f25e2441806a7b27
2019/11/09 10:49:22 jnode_ta1s47zzz8nrwww9z4w2yl9d3yjckxap6dyd5u27pryuyk8vpje0jkzjzkjnka - e99c4dbdec25cc782814696ed5f8190a1c304afae30b010ac8e3b48391f12104
2019/11/09 10:49:22 jnode_ta1s4cca439qvt7dazppu9uuvl8xpwf45dtxqmh2q3wgxr8tx4hjg2wyhdthuu - 83bb7ccdaec554f44245f169bad4c9a13e8114c86d7485d372cceaf34ed70dfb
2019/11/09 10:49:22 jnode_ta1s5xlflm5ytzvzhr6p6m9gj2qlw03g20qwdgwxnsv2ll4kj0jthvgzgfulva - 8b323c888d89b0b341499ae422f82d42f81a81c7b40616d3543586781461019f
2019/11/09 10:49:22 jnode_ta1shhagrmnc2d2gc89me7el5d57706e8vmvm35wd277u9tq7tnqu9j79h3hvd - c8b1015a2968cc1a5c471d6fc5df0e662cc3fbe57bca445a289d7c3f14e877ce
2019/11/09 10:49:22 jnode_ta1s4dd2fkla3deusgweywzge9t0cetm064kllf9gsfvtfpc9nfs7sgcgf4e3v - 14a106af6f126588883924c4aaa73fceb42761d9ea6a43e42ff904076314ab6f
2019/11/09 10:49:22 jnode_ta1skkf06ftnlexqwvx0fzgs82vqnrnm9s84adhr89wu5q0ly3at5xuxtlepl3 - b38bcf59496bb0e8fcb188cdb9e749c4098eeec6d9cf7a9aad292a712a648470
2019/11/09 10:49:22 jnode_ta1s4mtn4jmefk00qfmv7j3t6vesqunarnsmht9herk7yuf8x4mzs8cztt0n58 - e4ff29023edbd1676a685c635a67ceb5095024dd3f939a68cc01dbb47017bebc
2019/11/09 10:49:22 jnode_ta1shldwftcpng0ah3kzrc7ufzl39uafsjkrmk8yu7cum05r9fuslj65d7m8ap - 8f667bc3c72b0f141c60842f418e1bbddebd80f3b6b87df497aa8b61213d8c13
2019/11/09 10:49:22 jnode_ta1shykqlluq8ass3kgj5zalhrmwlf4d2qd4lam5rhsgcw2urjy8rwpc6qh6f6 - 629a46f7ae86a0d56b145191c4428a02dd814f79c0ca19836e4563e03f2df0ed
2019/11/09 10:49:22 jnode_ta1s49dckh3gqtfzqrfe5t4x93j6kzhq4ssuc9zfdk5fdwfd6gat8c4j2xp2x7 - 7c6e01db2e3b0c443b82caf89c4bc804a1bda37bdd2c50663ad50e88237182e4
2019/11/09 10:49:22 jnode_ta1sk4l6zqcc4tgsfp0akeknpkd35jrrsz97827z35mg36wrc6fumhxw5tay3z - 598c95733533551e04179bd3dfda018910a999a4917070db536cac923f82b65f
2019/11/09 10:49:22 jnode_ta1skxqlkz49rrnpwd3heapa76rnrp2adent77w4asrq4sayw3wwpp9jfcn5k8 - 11f615ffd18bce1939b014d84b4fca9314251ab416fb54eeb77627ab187641d3
2019/11/09 10:49:23 jnode_ta1sh8jfxjgtmdym04x3tr0tzugxw0a7hqsj0ydr40xt2rz69hsxy99g0m7e0s - 34c033178945e60dbdc5d963ff830561e6a98a57201b24d5ae8911276590d58c
2019/11/09 10:49:23 jnode_ta1s4yvyn93sjmr7e7rlr8r4qsq76mfrr5mpumqq8vvzur33uergnu8uudrs99 - e155b16c8eebf0c75c160e5b5bb517a5600151e23d72b2ca02bbcf5a83c89cc6
2019/11/09 10:49:23 jnode_ta1s5hpjd5uy5npu6d6pc7kycz3pawkzd9m6kau3kq20mq2easv5pfyj802gv4 - c4e43c970c07a68c1133d9225b589861edd2c076c53adf5a9daeca4ca905479f
2019/11/09 10:49:23 jnode_ta1skuwg6kam9wlckh7kzrlul4fq022nfr52rdzd0094550wkmr2ydk7z3dwz0 - 2e4f6575ba2fcdbf62f867b292ee89b7e43f3d7762c3ec4b2c3622519d9b42e1
2019/11/09 10:49:23 jnode_ta1s550tzlpy3xhlezde72cwn70xg94924k04luw5nkxfgwhdkf0e9g29x06ff - 67c9e53d24ed391515e78bae07afae9f5017e972b451b4db44e6adc8563da9ef
2019/11/09 10:49:23 jnode_ta1skfys5mtl9w5chu7dlk7f0p3xtqrvauhw866ykg6jmevtjkdmumj69sl4nq - 3abf166aa86b1e904945e1e51a1e813992b4d8e9bb5a1b2e7ad31d66a1dc6bf8
2019/11/09 10:49:23 jnode_ta1s5jvr7hehcme4hlgp0jl30s0jhl0hhl3lm7h406s50u65t7c3myvkklfdj9 - 4e44101a60c4c18e413bf352adc1313df6a1ebac29bec620b592d6c14dd5f002
2019/11/09 10:49:23 jnode_ta1s5gmt78jszdcjthq5snrawgyf4rycdjcfkvfkjnxx437pzkvxkxrwkhh537 - 68f831869c7c3354e69fd713d55a51272838bed92c22dace34b0719b974fbd8d
2019/11/09 10:49:23 jnode_ta1sh8fjw43kxwzmu6stas4rw4fvsgwd9mtnrdhykdas6833nu05ex6swx55ja - b0ecccb1b07e1c7fb42103479a9a757c5364ebac6895960da1579a97cc921c49
2019/11/09 10:49:23 jnode_ta1skzlvder8cv6fsnuxqm6dkrgra97py2jprkxcgdzlrmu5m6yglvnxs6qqwm - 28a0e6f8664dc472fb716502b6477626fb28b456809c06425ce280f431d83068
2019/11/09 10:49:23 jnode_ta1sk66alarl3xqklwee22xe07kpxxs94nh96r37ye4dl6uwh96zf0m724980u - eb6958858ddaebf3254ab4231ee3d4a45d9da65c423ae101a3817931077f0aa8
2019/11/09 10:49:23 jnode_ta1sh4rm259rqpuazujn7z7l3d5hcwymgum97g2c4453cnvpmg8athk5dwxkm7 - 2b44ef3250e8bf2e81f0f41c8b02c54a8b205a6c6ef86ef5be83b777a19ed599
2019/11/09 10:49:23 jnode_ta1s4y5wz08tcfv8xdmdqkf5knxnm7qk9snft3g045txswvycc3vq2m75e7utq - ead91b06c4f3a6cb2fa5527dd2aba8681944289fdf9309fbf25ca6a36b889792
2019/11/09 10:49:23 jnode_ta1s55zzdkw3r95hdsym29vvlacan40we90pqu5vrempuz0x48cw7tqqswzjmt - d4b659325baabb26ff5c242eeb05be5972b5692a9b6845cfa255de7de060126c
2019/11/09 10:49:23 jnode_ta1s43ymn68j9yq6gmdmt745tfdj8pee8ct02qmd9zmd2d3chc6m3dqwrxhkhx - 5a0db8b5ff1f4c7030facc74ee2cc2c758e4bfcd7fca5e5a4b2635b387ba6252
2019/11/09 10:49:23 jnode_ta1skng9thtxeyxkn8hg4038q423e2wkhzym4c3t87hsk5tghtfc7hlyjehx33 - d43e976119ae324986ad7f966d2b7a98a41d06d45c710cbd185ece3aff4c8964
2019/11/09 10:49:23 jnode_ta1s5jp5xhmqqv2uyzdnx2ge4tsdkxs7gqvy956trtjl2c0fhnkudl6gp46xuf - 0f6797f38356ab603132b7bba80ecffd5744d526757b251aac2cab3c857f91b4
2019/11/09 10:49:23 jnode_ta1sk59zz4klx6j8md78246rqhk8q0gnnp6gj4qje94yr2c24ggrswgwqm0scu - d70d4063240dcfabf2389ccab4383c60316f447338c61958055603ecea654c77
2019/11/09 10:49:23 jnode_ta1s45xm9nryk8nxdfhvznjuxdav842muqr8c0czntj5500qxkckjylkwa4xa9 - 3248d765a34cc2525afd6a51c5cb4e60aa90f816dc1de3eb747b8eb29dbdec17
2019/11/09 10:49:23 jnode_ta1s42v0gfknf648du70rx9qu68t4zfrklnl0ajpwh4mn8u4n3txn5a25wafxc - 2beda4db8bd3c6e68c5b3d1f423a0ea05a8971428e03c1dbc4fc3338ded0f84e
2019/11/09 10:49:23 jnode_ta1s5sxsls6warl8q62juflm9ny58jzjyq9tddparqv9xj55p6ymp4lq8vu6y0 - 97b48744ab4ebbb8cbc458c5cae23dda2a81bd0c38d32bdb1ca9ae0fdea895d2
2019/11/09 10:49:23 jnode_ta1sh4pygj8p3d0yasksql7yjwueltrndurcv4ngqar5qxmjmytd22w7zljjnj - 14df428319719f76935e13d1e89f64e841d002aca650a0f1f03690f8b3d84e4a
2019/11/09 10:49:23 jnode_ta1s53qm6edw5trx4zeqak8mcdx0706fnmjnxp68me8y055kq5z4p9t67xwhjf - a7bfd50f5dc7616a15405d60125c1e0980d98db7cad26948feacd8dd4f2caf98
2019/11/09 10:49:23 jnode_ta1skylrf9ulphq6x2hq2vmcflf8kc335uyg80ll72v7pd7md8wwna5g8dquqx - fa1f2572bc160cd5b51a4bf374efcaa276f21ff6aa5df8df8ebf619f4107692b
2019/11/09 10:49:23 jnode_ta1shjz9fnn423jchm8ukf2dtcwmh65czg05d0g3xymepanpsnugk3zxjqgvm0 - 3b50d72cc40d1435cc873873877d607bff46c3675f5dd49ceaa0836a011813a7
2019/11/09 10:49:23 jnode_ta1sk3dc8cxuvad3wg9fkeap5m4fz9v3epfrmk9w0yx0c9cv7k2pvq25zxzwsp - 39c83e1f5d1de03c37d7eb812b07b8abc89666f6518915afbf824f7b5c900aef
2019/11/09 10:49:23 jnode_ta1s5aze7naejjz487l6vz5zuhqf2fn48qve8fxzv5f53mfkjlapaj9cq4fqxk - 124f72b118fa16f064be105b1eb54ff8841ea777c49be10c1b9771e59a3d4ff7
2019/11/09 10:49:23 jnode_ta1s4pp6gntvdg39nju3q9jhy9l20lqajf2l7f0czryfzpl7kkcc3pey3545yh - 1cf2b607e2a77de7af212bbca9cbe7dfbb20ccb1e23aae4c4882baf0900a33b6
2019/11/09 10:49:23 jnode_ta1s5zssp6n2x3yy72jym8r7ypjkxw6np7dgnf6e5yyslml9ddm7qj5s4yfjdw - 5e65fb1b03ccacfe47e8bf7f84997f9cb4a64b88d77e5b2a92b285a91663333e
2019/11/09 10:49:23 jnode_ta1s40579t28us2kgq5vznjzmaudt2pa9tjq4kg2q38pqesdw9mfy9ws7xke69 - 6d42b1735e5b0d0a4c3e9154831b254c8210925bf479b18199497ccd96219df0
2019/11/09 10:49:24 jnode_ta1s5nm64z0sk9r8jp9y20nv0eegmq6gm2hpghpnp3z9lzzdk4uerhfx55cua7 - 8e5f94edba815a1c4761f0b94af8821f44a2ca42ace8be7b80dc1ee360d61513
2019/11/09 10:49:24 jnode_ta1s5ec5mymsuvdaeljcg7c7s7x9wcg4dmykf35lu0n0qk4mpmcgrnp706f8e4 - f82831d095b23e62d30a27b1640810c39229e66f2b6cc751c4bce0c44b8a2666
2019/11/09 10:49:24 jnode_ta1s4ahxnflfda06xcdw909npk92c5synlnk7pectq7rs8l8mr6p6wqsuptjp5 - a44fe0be51b886aa0b54bc15052306b7f7f07e810dfbb6e1fb1ffb2e9ef3d3cd
2019/11/09 10:49:24 jnode_ta1shv0a8dx6lpj4fr8xup9drf2a4cmpzvlwh5enwjcqaj6v3gel76rwygtyg8 - f5032b4c022fcd3da0a1aedf8b44523d5ce30ee89182295a9c68630e5f9ae3c6
2019/11/09 10:49:24 jnode_ta1s59zyusjregyhhwchmdf0lly3ld53tj0s49lrcya238pdmmmrt50jcwfatx - a5befe5335bb86a8426ed36f815f9adaf6e4a09e0f3d949229cba9ae4f6cb466
2019/11/09 10:49:24 jnode_ta1s5x78szsanpd8smym2rdjvhwlhuz8ccncvj0vyu0mm8zxkc6h6kmzrlz3mj - 9736e0aaafaa65ba19cf9ab0a83d6fdb38b199eeb07036112cef6b0f9a4514ef
2019/11/09 10:49:24 jnode_ta1skmm0369a8efkf2rntvs5ymvwwfrw9zv6n9zkd8mhv96wscj5fuqgytd2dy - f1d9ab9319957ed0d49f411cf21fc5adb7990195adb94cec534a96272404048b
2019/11/09 10:49:24 jnode_ta1s43s0kp0l85lld0nutvq6xfkkcmmmw3wh46x7r4y57h97u50l5ft5ec5y2g - aa125b97439b28dabc2648986bc8442c21ae49917ab8e6e6e65c1157109c2a72
2019/11/09 10:49:24 jnode_ta1s4u3c0sflmltp9lnvy6s76qgvxjgtn0aw3sysyep4p5nkg2h3djywt64zay - 73c673f412c0830149eb097a5195c84311871215f59c735b62f1ce9653d3573d
2019/11/09 10:49:24 jnode_ta1s4e9jas42rcqs3ufcnlp3mvp7nyt3mrf9qee0hca83pdxewlu63nu9qdh5u - 85627ab660d05488523ee4670dcb8eb37c0978dea4abe591ef2bb38367f78270
2019/11/09 10:49:24 jnode_ta1s420wwu2hfgwj5tvwava78tugq06evkfw8ngu8wyn2xuwttt78qqg6ayxr5 - 705077a3493a77841580879c42893ad2d2e5918064d639b92687fc43986c0e30
2019/11/09 10:49:24 jnode_ta1s4267em4razxtdl83538kl5f350h5f353dygp0kgvmk5a82p726v5jt8tex - a8bbed38164a3a21518b45383929c0c6ddb3bf534e875edb7ce5ef4a72eacced
2019/11/09 10:49:24 jnode_ta1shn9axum2cphz2f66r2sr6uqwjq4mzh2xgkxnyw9ksnt3u47klzy7m9zggk - bc0e43f03c1db2572694f70ba4550c9cbc2d9f0909eb452691af8c1459f6a77b
2019/11/09 10:49:24 jnode_ta1s5tq3s7r3ezmuhh9v9ha5050gltzf7mllx2aejqhck5txhmqzhqqs9y52wf - 9fbcaa039deb311b77a9f301a29314703e79bd472f76f39ba01ed0dfbe8920db
2019/11/09 10:49:24 jnode_ta1skx3tk6upatfg4s063yzp84jd5g6dp4sdlkzu47gjrfct9nfugvlsgg3hqu - df1a4034bc698f3973d08311358e8363a11205b43d80fca7a7c2635d6388ea5c
2019/11/09 10:49:24 jnode_ta1skezrvu4m2hklr2dlfthswl2sc649ua8yhrfnz0uhfus9xsufpjryfj6sqw - fdc877162a3a286233b5b82eaf83680d5029613eab6c6e0d9d5542c939f62a8d
2019/11/09 10:49:24 jnode_ta1s55g3dunvppflvnu2f6a08akmhwnpyzcpxras0qvnhdwnjk0unzm5s55cpy - 958fd556ed614cd2a78a795bd41cdc696997cb3a325addfc6aeb24115bb3e334
2019/11/09 10:49:24 jnode_ta1s5t5u3kpdkp6v9muzw276m4uva2eugvv8szzp0ynv4amu7mvhdwsuf9pzkx - ed225d8e31aede592ccae1ef7c0763a1efd5da233d0ec52b8685760fedb5dffe
2019/11/09 10:49:24 jnode_ta1s4p7w0689mwantdt9j779g5ajlc3eyq2778a9yn3vf292xq6xla5gp5lh06 - 213a0b598223e4eff7d5045cdf526ff343efbe0630221334b4919de1cd0d92c1
2019/11/09 10:49:24 jnode_ta1s496dy5e998p733e8h6kam7fwt35x06n9z9lqrwzm43dc0mmj38nkdpf5jx - 0c89b5a3a50eaf906b645d1144f6d32ef1bda6c2e5f70085814e5c3ddd960c1a
2019/11/09 10:49:24 jnode_ta1she3z799ka4puvfh56z757me4e9nl8akcr6n22j6jyar885wv7ztv3y0mvv - 74bcdb0e7507a40bd0426776dd435d932407c1376036a1c95b13b428fd7f1cbd
2019/11/09 10:49:24 jnode_ta1s5ef5pgcf6su6t79tdnvdhzq206hfyusnl35rf2dn0hvxkhqnfxu624h5me - 91ae316f6a69451698145d05bd2a5a3ff11f8ae8bd6ea0982871d7721eb178b8
2019/11/09 10:49:24 jnode_ta1shnd447lg057slzcf64l5hr3l4j8pdv7n65cams346mpxnuw5q2q620zt48 - 12b0bcd754f096244fc58a87d50b44ada5e1fe8406bd99119cd5908199c591ae
2019/11/09 10:49:24 jnode_ta1s4qtez8tup99kd9q2wh707dsa3rm9q7hvt4aak3sgx4x28ft5f77uj7r5d0 - 98b29d1f3d405b91dd3822eb6741661eeaecd20bca0873022332efc34122b1ac
2019/11/09 10:49:24 jnode_ta1sh23ttcfvekxvu6zqvlmf3hmtv8dunsf57c834h0l64tupmv7uaxjhk5xyn - 554bac8be094710206a1c22b77bb086a3b08fc5ff4e6b5b482023edb22d5b85c
2019/11/09 10:49:24 jnode_ta1s5v3r3wh0vmejvpjq3tx3uqduyk0wp0z3devcmjpc4fx9csmm6xu7pdjwqk - 4047b58be70a539f64ddee26800af705c20a0e938e298e74ab3b9617f4df7be4
2019/11/09 10:49:24 jnode_ta1skg9cnc6ns09vq7re8en2s7ypfnrtk30vvhpx8p8gfj38627dksu2q80rwy - faf3ca0ae50a66fa420d816f74a16cf5b4b8a114ffa0d3e6ba8da4d895f4d350
2019/11/09 10:49:24 jnode_ta1s4dpw4nm6n59ars56u466say7xlvtyn755qj50mday392k6vthvz7avfd6e - 3df006a865646733c21bf2bbc6fac635660af9d6825f1ead9b9661c8d4abe479
2019/11/09 10:49:24 jnode_ta1s4yq0s08gthuw0mz5gcjmt3vtnkpnsc98hhyudzkta738ryjdlv0uepy5cm - 1162585c455571677def3cdfaab15b59f7fc65c7601c893b9e1bb7e078bf0bf1
2019/11/09 10:49:24 jnode_ta1s5w39hejw0hcgeghn5rdm96zea8e2e8xfgn9f4a9a5e2t2vgl7fj2t3cp6q - 94fb2222e9ded801c1b29e452eb9ff4e53ec331bd8f3ff8928d6affc55f16ca6
2019/11/09 10:49:25 jnode_ta1s4ha79eecqv5nj3y87xrxmqw76kn5u0hcywy2zvwerpkt7xyk4h8ymdg44w - a757164325c3a16ffcb586d8a6792058216352c006ae85133777d34c15fbd407
2019/11/09 10:49:25 jnode_ta1skyywuduvxq3y4wsjg59ycvzy2tenldy53kksghcee8g7jywf7tpsvad553 - d010b2847b7278b0b58a0eef6a3f57304dbf9de9fbe3ce027c3b6e592b381ea0
2019/11/09 10:49:25 jnode_ta1s4ukn0a5cyl9gyslk4kxraaktftt2l7l95x3zdudkleg727e88esxgr03yd - 7665fbb613192481fec382d61f776ef5db351c388f7e35661845910335f6bc59
2019/11/09 10:49:25 jnode_ta1s4ruz5ph7uel2dcs6wyleymxwe2py53kc365xd3qp89m94mtg9wdcydrxtp - 70b535caa4929257de5cb7e936a8975f9f6f41bc401094879d49570227b3e3ba
2019/11/09 10:49:25 jnode_ta1skx4dhkazyaskslqj2st8wy4hv5c4tnukrpwtrgqnzle503cy78t5h6uvvy - ba4652fc0f1ecb7215ef6a386c151eb3c035271e35007790e7d9537cc1c76dbd
2019/11/09 10:49:25 jnode_ta1shvjnxhjjvn58gqhngarudxrl0pt3r2dgx695hz2uwkqexpukcapgj70hn4 - 0e1c56308e159a69fb5a97de50270a8163f7f887e1bc8d9f08f3e5f847e096c4
2019/11/09 10:49:25 jnode_ta1sh0s7ulczf6eeytxvy9k9jp7ys76skmvmz3fzwdg3gqmvr7e4tmgzju7k6d - bc1e2347a619d67466e4431757df8171e23ae4bfeb267fb1e8954182daa5a179
2019/11/09 10:49:25 jnode_ta1skpsaglmfgpuaz9g6ugmr5hakmyr074pkyya9f46ptclx5ql88nvj7s6cfe - ca4523b41624428555cdba5dbeaf04a562c9085f3f1277ba1a152f504285b96a
2019/11/09 10:49:25 jnode_ta1s5kfrqnjpk6fy42ae4dj9c77hrauyves5unkkvvf23u7wwqayzj7zyrtkvr - 1d9c6721dfd9a36e2cd3f7c6aeab5aeba7cbe2ee60588fbd8070d8f46949c327
2019/11/09 10:49:25 jnode_ta1s5slzkz43jxfzdgs6qlqppsf5dh5vt0fzaa44kq6eg5r97jqveaz7qy5d0t - 27c5849a66429626c2468ead7f4c61eef062a4c85eb227c15b39aa6fa114e5a7
2019/11/09 10:49:25 jnode_ta1sh794ctzes8s99ht88xrq8vza2slf489cty2vpc094txdyaaurtt5j970e2 - 44d4f0e4a4eca9a9b412958e00ec710aa304edd45d22ae0567a11fca10eba751
2019/11/09 10:49:25 jnode_ta1sk43xm28hk3khpqe2fqp54xfy5dl90tjz9a8pyhts40pewsgpuvj75nftrs - f5b2a29f4a98ae0e08958603667427af9dcc115a4df81d1a6bd9fc6415339b2a
2019/11/09 10:49:25 jnode_ta1sktaew3xc65k8t3pt9vtyt6wdavmugt4pfj85cpjsgpwvz63xcg8cjvhrsk - aacff766cd6f48f166426476bb0a446c4e5a3474bfb4efbf05668003c41f365a
2019/11/09 10:49:25 jnode_ta1shyldsf4mtkg60lqtzjz6tk4h5077caqx8kzwwat5c9mv9kj275mqxcfpsz - b14a1e47568a88cfefeb267f8bd05dd14938f3138b76bd11213ccf0c25c4312d
2019/11/09 10:49:25 jnode_ta1s4cfacejps39069h0dvv8e0ngwrwjg94hqrdhfd9azeeydxupppwu5sz3lt - 78af07073a91b29c5f10570910dda105386bca12ca3defa7d3eeedf889a7b64d
2019/11/09 10:49:25 jnode_ta1sh3erdxaq05a8u98ssv86ck9tvep298et7q4aeys328jtmhqnv5yvs0z6rx - 614ec319080391c27033d5481300ddfea8d3e7c9b58dfe8175af529d8e9ef827
2019/11/09 10:49:25 jnode_ta1sklhnnlws730mmnqml236tx9cayn4rmm76sveyanl06asm0c9lef7ukgu24 - 30d23d0ebe796a2d7a09a3e7dcfa2b9e7f5b2fb7d68b858fc73dd537392515c2
2019/11/09 10:49:25 jnode_ta1shuwmr8ehdt7qw98aa5jquyv3nrlfakhxa500hc3gzp2s6v2ntj86ymhew4 - f0d3a612feb442eec4a675244e296d363f3da2f7b3e1da31a906f95d22a2b0c9
2019/11/09 10:49:25 jnode_ta1s40dl0pdmj5zez85usrhzjtuy36ppc97024938g3jrjhp3armm5jxl53k5p - a88a24c77a319a15cc5a46170eb629d565237a340b6d5851fa4a4c08d3e8029e
2019/11/09 10:49:25 jnode_ta1s5frku4y676aytssae0xzh0awattfg5gxcr279payz0nt9ectxn55s38yw3 - 4e32a383d0c3cbcf6dc789637edb7fa14bf532e3e54b83f994bb3605f9d708b1
```

## JCLI Bulk send concurrent

It shows the usage of:

- [jcli](https://godoc.org/github.com/rinor/jorcli/jcli)

TBD
