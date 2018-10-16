---
layout: "secret"
page_title: "Provider: Secret"
sidebar_current: "docs-secret-index"
description: |-
  The secret provider providers a simple resource that can be used to store
  secrets.
---

# Secret Provider

The `secret` provider has one mission: store secrets in the Terraform state.

Please be careful about your security stance before adopting this!

The main goal of this provider is that a lot of time, terraform contains
secrets in it's state file anyways. Instead of putting them in the repo and
the loading them with `"${file("./secret")}"` why not import them directly
into the state file?

When using a remote state file, the state is automatically distributed with
the new secret which makes key rotation easier.

This is only a better solution than storing secrets in Git. Look at adopting
Hashicorp Vault in the longer term.

