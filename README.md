# jorcli
[Jormungandr](https://github.com/input-output-hk/jormungandr) jcli wrapper in [Go](https://golang.org/) - (**experimental**)

Right now this is just a *Proof Of Concept* and may or may not become a *Prototype*.

The idea is to build a simple and small wrapper around *JCLI* binary.

TODO:
- [x] **address** - Address tooling and helper
- [x] **certificate** - Certificate generation tool
- [x] **debug** - Debug tools for developers
- [x] **genesis** - Block tooling and helper
- [x] **key** - Key Generation
- [x] **transaction** - Build and view offline transaction
- [x] **utils** - Utilities that perform specialized tasks
- [ ] **rest** - Send request to node REST API
  - [x] **account** - Account operations
  - [x] **block** - Block operations
  - [x] **leaders** - Node leaders operations
  - [ ] **message** - Message sending
  - [ ] **node** - Node information
  - [ ] **settings** - Node settings
  - [ ] **shutdown** - Shutdown node
  - [ ] **stake-pools** - Stake pools operations
  - [ ] **tip** - Blockchain tip information
  - [ ] **utxo** - UTXO information