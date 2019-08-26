# jorcli
[Jormungandr](https://github.com/input-output-hk/jormungandr) binary wrapper in [Go](https://golang.org/) - (**experimental**)

Right now this is just a *Proof Of Concept* and may or may not become a *Prototype*.

The idea is to build:
 - [ ] a simple and small wrapper around *JCLI* binary.
 - [ ] a simple and small wrapper around *JORMUNGANDR* binary. (*maybe*)
 - [ ] *JORMUNGANDR* rest API. (*maybe*)

**DONE** *jcli* :
- [x] **address** - Address tooling and helper
- [x] **certificate** - Certificate generation tool
- [x] **debug** - Debug tools for developers
- [x] **genesis** - Block tooling and helper
- [x] **key** - Key Generation
- [x] **transaction** - Build and view offline transaction
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
  - [x] **utxo** - UTXO information

**TODO** *jormungandr* :

**TODO** *jormungandr* rest API: