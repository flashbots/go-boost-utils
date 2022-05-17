# Builder API types & signing

[![Goreport status](https://goreportcard.com/badge/github.com/flashbots/builder)](https://goreportcard.com/report/github.com/flashbots/builder)
[![Test status](https://github.com/flashbots/builder/workflows/Checks/badge.svg)](https://github.com/flashbots/builder/actions?query=workflow%3A%22Checks%22)

Tested types and signing routines of the [Eth2 Builder API](https://ethereum.github.io/builder-specs/).

This is useful for:

* validators and beacon nodes can use it to communicate with [mev-boost](https://github.com/flashbots/mev-boost), relays and builders after the merge
* [mev-boost](https://github.com/flashbots/mev-boost), which uses the types in this repository to comminicate with relays and builders
* for future builders, to communicate with relays
* to build testing, monitoring and verification utilities

See also:

* [Builder API specification](https://ethereum.github.io/builder-specs/) ([Github](https://github.com/ethereum/builder-specs))
* https://github.com/flashbots/mev-boost
* https://github.com/protolambda/mergemock

---

## Contents

* [`types/common.go`](https://github.com/flashbots/builder/blob/main/types/common.go): various common basic datatypes (Signature, PublicKey, Hash, etc), with tested SSZ encoding
* [`types/builder.go`](https://github.com/flashbots/builder/blob/main/types/builder.go): builder-API specific datatypes, with tested SSZ encoding
* ...

---

## Contributors

Maintainers:

* [@metachris](https://twitter.com/metachris)
* [@Ruteri](https://twitter.com/mmrosum)

Special thanks:

* [@lightclient](https://twitter.com/lightclients)
* [@terencechain](https://twitter.com/terencechain)
