# jorcli

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/fb345d4e21584e38b13695438c733c79)](https://www.codacy.com/app/rinor/jorcli?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=rinor/jorcli&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/rinor/jorcli)](https://goreportcard.com/report/github.com/rinor/jorcli)
[![GoDoc](https://godoc.org/github.com/rinor/jorcli?status.svg)](https://godoc.org/github.com/rinor/jorcli)
[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/emersion/stability-badges#experimental)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frinor%2Fjorcli.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Frinor%2Fjorcli?ref=badge_shield)


[Jörmungandr](https://github.com/input-output-hk/jormungandr) tools in [Go](https://golang.org/) - (**experimental**)

Right now this is just a *Proof Of Concept* and may or may not become a *Prototype*.

The idea is to build:

- [x] a simple and small wrapper around *jcli* binary. (*alpha*)
- [x] a simple and small wrapper around *Jörmungandr node* binary. (*alpha*)
- [ ] *Jörmungandr rest* API. (*wip*)
- [ ] *Jörmungandr explorer node* graphql API. (*wip*)
- [ ] *Jörmungandr* grpc. (*maybe*)

**DONE** *jcli* :

- [x] **address** - Address tooling and helper
- [x] **certificate** - Certificate generation tool
- [x] **debug** - Debug tools for developers
- [x] **genesis** - Block tooling and helper
- [x] **key** - Key Generation
- [x] **transaction** - Build and view offline transaction
  - [ ] **data-for-witness** - Sign data hash [(not yet available on jcli)](https://github.com/input-output-hk/jormungandr/issues/674)
- [x] **utils** - Utilities that perform specialized tasks
- [x] **rest** - Send request to node REST API
  - [x] **account** - Account operations
  - [x] **block** - Block operations
  - [x] **leaders** - Node leaders operations
  - [x] **message** - Message sending
  - [x] **node** - Node information
  - [x] **settings** - Node settings
  - [x] **shutdown** - Shutdown node
  - [x] **stake-pools** - Stake pools operations
  - [x] **stake** - Gets stake distribution
  - [x] **tip** - Blockchain tip information
  - [x] **utxo** - UTxO information

**DONE** *Jörmungandr node*:

- [x] node controller
  - [x] Run node
  - [x] Stop node
  - [x] PID of running node
- [x] configs
  - [x] block0 (genesis) config
  - [x] secrets config
  - [x] node config
    - [ ] implement output gelf (to be implemented)

**WIP** *Jörmungandr* rest API:

- [ ] build from [openapi](https://github.com/input-output-hk/jormungandr/blob/master/doc/openapi.yaml)

**WIP** *Jörmungandr explorer node* graphql API:

- [ ] build from [explorer](https://github.com/input-output-hk/jormungandr/tree/master/jormungandr/src/rest/explorer)

**TODO** *Jörmungandr* grpc:
