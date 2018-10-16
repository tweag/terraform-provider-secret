---
layout: "secret"
page_title: "Secret Resource"
sidebar_current: "docs-secret-resource"
description: |-
  A resource that stores a secret
---

# Secret Resource

~> **Warning** Make sure that the Terraform state is stored in a remote state
before using this providers. This includes using proper ACL and at rest encryption.

The `secret_resource` resource does only one thing: store an arbitrary secret
string as a state variable.

## Example Usage

The primary use-case for the secret resource is to store arbitrary secret in
the Terraform state instead of having them stored into git.
arbitrary actions taken by a provisioner, as follows:

```hcl
resource "secret_resource" "datadog_api_key" { }
resource "secret_resource" "datadog_app_key" { }

provider "datadog" {
  api_key = "${secret_resource.datadog_api_key.value}"
  app_key = "${secret_resource.datadog_app_key.value}"
}
```

In this example, the DataDog API keys are stored in the terraform state
instead of being written in a file on disk. To actually use the resource an
admin would have to first run:

```
$ terraform import secret_resource.datadog_api_key ACTUAL_API_KEY
$ terraform import secret_resource.datadog_api_key ACTUAL_APP_KEY
```

Once the state is imported the rest can be used as usual.

### Secret rotation

To rotate a secret, run:

```
$ terraform state rm secret_resource.mysecret
$ terraform import secret_resource.mysecret THENEWSECRET
```

Once the new secret is imported, run `terraform apply` to propagate the
change.

## Argument Reference

None

## Attributes Reference

The following attributes are exported:

* `value` (Sensitive) - Contains the imported secret.
