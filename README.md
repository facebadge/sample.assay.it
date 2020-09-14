# Minimal Quality Assessment Suite

Learn about quality assessments of Serverless applications with https://assay.it. This example implements a minimal quality assessment suite using Behaviour as a Code.


## Inspiration

Quality assurance of serverless applications is more complex than doing it for other runtimes. Engineering teams spend twice as much time maintaining testing environments and mocks of cloud dependencies instead of building a loyal relationship with their customers, assay.it has you covered.

[**https://assay.it**](https://assay.it) is a Software as a Service for developers to perform formal proofs of quality using type safe Behavior as a Code. It automates validation of cause-and-effect in loosely coupled topologies such as serverless applications, microservices and other systems that rely on interface syntaxes and its behaviors. It emphasizes deployment and quality assessment as a key feature along the development pipelines. Continuous proofs of the quality helps to eliminate defects at earlier phases of the feature lifecycle. It impacts on engineering teams philosophy and commitments, ensuring that your microservice(s) are always in a release-ready state.


## Getting Started


1. **Sign up for [assay.it](https://assay.it)** with your GitHub developer account. Initially, the service requires only access to your public profile, public repositories and access to commit status of connected repositories. Later, you can enable quality assessments of private repositories. 

2. **Fork [assay-it/sample.assay.it](https://github.com/assay-it/sample.assay.it)** to your own GitHub account and then add to the service workspace. The example implements a minimal quality assessment job using [category pattern](https://assay.it/doc/core/category) to connect cause-and-effect (Given/When/Then) with the networking concepts (Input/Process/Output). Just write [pure functional code](https://assay.it/doc/core) instead of clicking through UI or maintaining endless XML, YAML or JSON documents.
```go
func TestOk() gurl.Arrow {
	return gurl.HTTP(
		ø.GET("https://assay.it"),
		ƒ.Code(gurl.StatusCodeOK),
		ƒ.Header("Content-Type").Is("text/html"),
	)
}
```

3. **Launch the quality assessment** through the user interface. The service schedules the job and returns results of assessments in a few seconds. Here, a manual job trigger is used for ad-hoc and illustration purposes. [assay.it](https://assay.it) supports integration with CI/CD so that [continuous quality evaluation](https://assay.it/doc/case-study/everything-is-continuous) is a part of the development culture. 
![](https://assay.it/doc/assets/images/screen.png)


Let's have a look on the content of repository:
* [suite.go](suite.go) implements a minimal quality assessment contract of the service.
* [.assay.json](.assay.json) configuration file, it declares what suites shall be executed.


## Further Reading

Please continue to [the core](https://assay.it/doc/core) sections for details about Behavior as a Code development and see [our advanced example on GitHub](https://github.com/assay-it/example.assay.it).


## License

[![See LICENSE](https://img.shields.io/github/license/assay-it/sample.assay.it.svg?style=for-the-badge)](LICENSE)
