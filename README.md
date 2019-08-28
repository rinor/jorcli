# jorcli

[![Go Report Card](https://goreportcard.com/badge/github.com/rinor/jorcli)](https://goreportcard.com/report/github.com/rinor/jorcli)
[![GoDoc](https://godoc.org/github.com/rinor/jorcli?status.svg)](https://godoc.org/github.com/rinor/jorcli)
[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/emersion/stability-badges#experimental)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

[Jormungandr](https://github.com/input-output-hk/jormungandr) binary wrapper in [Go](https://golang.org/) - (**experimental**)

Right now this is just a *Proof Of Concept* and may or may not become a *Prototype*.

The idea is to build:
 - [ ] a simple and small wrapper around *jcli* binary. (*WIP*)
 - [ ] a simple and small wrapper around *jormungandr* binary. (*maybe*)
 - [ ] *jormungandr* rest API. (*maybe*)
 - [ ] *jormungandr explorer node* graphql API. (*maybe*)

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

**TODO** *jormungandr* :

**TODO** *jormungandr* rest API:

**TODO** *jormungandr explorer node* graphql API:
