# Vultr Resource Provider

The Vultr Resource Provider lets you manage [Vultr](https://www.vultr.com) resources.

## Installing

This package is currently not available for most languages/platforms.

### Python

To use from Python, install using `pip`:

```bash
pip install git+https://github.com/vincentbernat/pulumi-vultr.git#egg=pulumi_vultr&subdirectory=sdk/python
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/vincentbernat/pulumi-vultr/sdk/go/...
```

## Configuration

The following configuration points are available for the `vultr` provider:

- `vultr:apiKey` (environment: `VULTR_API_KEY`) - the API key for `vultr`

## Reference

For detailed reference documentation, please visit [the Pulumi
registry](https://www.pulumi.com/registry/packages/vultr/api-docs/).
