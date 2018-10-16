package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tweag/terraform-provider-secret/secret"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: secret.Provider})
}
