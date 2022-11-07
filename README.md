# Terrajet Taikun Provider

`provider-jet-taikun` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for the
Taikun API.

## Getting Started

Install the provider by using the following command:

```
kubectl apply -f examples/install.yaml
```

In order to learn about basics of this provider, you can found [here](https://github.com/itera-io/provider-jet-taikun-workshop) a little workshop.

## Documentation

You can see the API reference [here](https://doc.crds.dev/github.com/itera-io/provider-jet-taikun).

## Contributing

To add a new resources, please follow [adding new resources](https://github.com/crossplane/terrajet/blob/main/docs/generating-a-provider.md#adding-new-resources) section.

To configure existing resources, please follow [configuring a resource document](https://github.com/crossplane/terrajet/blob/main/docs/configuring-a-resource.md).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/itera-io/provider-jet-taikun/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

provider-jet-taikun is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-taikun adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-taikun is under the Apache 2.0 license.
