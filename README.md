Terraform `secret` Provider &#x1F49C;
===================================

The `secret` provider has one mission: store secrets in the Terraform state.

Please be careful about your security stance before adopting this!

The main goal of this provider is that a lot of time, terraform contains
secrets in it's state file anyways. Instead of putting them in the repo and
the loading them with `"${file("./secret")}"` why not import them directly
into the state file?

When using a remote state file, the state is automatically distributed with
the new secret which makes key rotation easier.

This is a better solution than storing secrets in Git. Look at adopting
Hashicorp Vault in the longer term.

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

## Installation

### Install via `go get`

1. Follow these [instructions](https://golang.org/doc/install) to setup a Golang development environment.
2. Use `go get` to pull down this repository and compile the binary:

```
go get -u -v github.com/tweag/terraform-provider-secret
```

The binary will be placed in `$GOPATH/bin` or `$HOME/go/bin` if `$GOPATH` is not set.

### Install via Nix

If you are lucky enough to use [Nix](https://builtwithnix.org), it's
already part of the full terraform distribution:

```sh
nix-env -iA nixpkgs.terraform-full
```

### Compile from source

Clone the repository:

```sh
$ git clone git@github.com:tweag/terraform-provider-secret
```

Enter the provider directory and build the provider

```sh
$ cd terraform-provider-secret
$ GO111MODULE=on go build
```

## Usage

### Provider installation

* Copy the `terraform-provider-secret` binary to `~/.terraform.d/plugins` (recommended) or any location specified by [Terraform documentation](https://www.terraform.io/docs/extend/how-terraform-works.html#plugin-locations).

* Add the line `provider "secret" {}` line to `main.tf`
To prevent warnings, you may optionally add a version lock to the provider entry in the form of `provider "secret" { version = "~> X.Y"}` where `X.Y` is the version you wish to pin. Note that when the binary is built no version suffix is specified; you will need to manually add `_vX.Y` to the provider binary unless you directly use release from Github.

* Run `terraform init`.

### Using `secret_resource`

**Schema**:

* `value`, string: Returns the value of the secret

### Example

Here we declare a new resource that will contain the secret.

```tf
resource "secret_resource" "datadog_api_key" {
  lifecycle {
    # avoid accidentally loosing the secret
    prevent_destroy = true
  }
}
```

To populate the secret, run
```sh
terraform import secret_resource.datadog_api_key TOKEN
```
where `TOKEN` is the value of the token.

Or to import from a file:
```sh
terraform import secret_resource.datadog_api_key "$(< ./datadog-api-key)"
```

Once imported, the secret can be accessed using
`secret_resource.datadog_api_key.value`

### Rotating secrets

```sh
terraform state rm secret_resource.datadog_api_key
terraform import secret_resource.datadog_api_key NEW_TOKEN
```

### Importing binary secrets

The secret values can only contain UTF-8 encoded strings. If the secret is a
binary key, a workaround it to encode it first as base64, then use the
terraform `base64decode()` function on usage.

Eg:

```sh
terraform import secret_resource.my_binary_key "$(base64 ./binary-key)"
```

Then on usage:

```tf
resource "other_resource" "xxx" {
  secret = base64decode(secret_resource.my_binary_key.value)
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-secret
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Related projects

* https://github.com/carlpett/terraform-provider-sops - allows to decode
  in-repo secrets on the fly.

## License

This work is licensed under the Mozilla Public License 2.0. See
[LICENSE](LICENSE) for more details.

## Sponsors

This work has been sponsored by [Digital Asset](https://digitalasset.com) and [Tweag I/O](https://tweag.io).

[![Digital Asset](https://avatars1.githubusercontent.com/u/9829909?s=200&v=4)](http://digitalasset.com)
[![Tweag I/O](https://avatars1.githubusercontent.com/u/6057932?s=200&v=4)](https://tweag.io)

This repository is maintained by [Tweag I/O](http://tweag.io)

Have questions? Need help? Tweet at
[@tweagio](http://twitter.com/tweagio).
