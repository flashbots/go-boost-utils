# Builder API types & signing

[![Goreport status](https://goreportcard.com/badge/github.com/flashbots/go-boost-utils)](https://goreportcard.com/report/github.com/flashbots/go-boost-utils)
[![Test status](https://github.com/flashbots/go-boost-utils/workflows/Checks/badge.svg)](https://github.com/flashbots/go-boost-utils/actions?query=workflow%3A%22Checks%22)

Tested types and SSZ/BLS signing routines for the [Eth2 Builder API](https://ethereum.github.io/builder-specs/).

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


## Contributing

Useful commands:

```bash
make lint
make test

# Install sszgen command
go install github.com/ferranbt/fastssz/sszgen@v0.1.1

# Create SSZ encoding methods
make generate-ssz
```

---

## Contributors

Maintainers:

* [@metachris](https://twitter.com/metachris)
* [@Ruteri](https://twitter.com/mmrosum)

Special thanks to:

* [@lightclient](https://twitter.com/lightclients)
* [@terencechain](https://twitter.com/terencechain)
* [@JustinTraglia](https://twitter.com/JustinTraglia)
* https://github.com/flashbots/go-boost-utils/graphs/contributors
