# Simple Suite Sample

Here is an example application to show https://assay.it in actions.


## Inspiration

Microservices have become a design style to evolve systems architecture in parallel,
implement stable and consistent interfaces. This architecture style brings additional
complexity and new problems. One of them is the assessment of system behavior while its
components communicate over the network - like integration testing but for distributed
environment. We need an ability to quantitatively evaluate and trade-off architecture
to ensure quality of the solutions.

https://assay.it is designed performs a formal (objective) proof of the quality using
Behavior as a Code contracts. It connects cause-and-effect (Given/When/Then) to the networking
concepts (Input/Process/Output). The expected behavior of each network component is declared
using simple Golang program.

This repository delivers a minimal example about development of Behavior as a Code contracts.


## Getting Started

You can fork this repo and play with it in own account... Let's have a look on the repository:
* [suite.go](suite.go) implements a minimal quality assessment contract of the service.
* [.assay.json](.assay.json) configuration file, it declares what suites shall be executed.

Are you ready to on-board with this example? 

1. Sign Up to https://assay.it
2. Fork the repository to your GitHub account
3. Go To Account > Setting and Integrate the repository with https://assay.it
4. Run the contract assessment


## Next Steps

1. Look into [advanced example](https://github.com/assay-it/example.assay.it). It shows more complex examples, talks about integration with CI\CD and depict the typical workflow.

2. Study ["Behavior as a Code"](https://assay.it/doc/core) syntax for contract development.


## License

[![See LICENSE](https://img.shields.io/github/license/assay-it/sample.assay.it.svg?style=for-the-badge)](LICENSE)
