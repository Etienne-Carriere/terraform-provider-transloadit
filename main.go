package main

import (
	"github.com/Etienne-Carriere/terraform-provider-transloadit/transloadit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return transloadit.Provider()
		},
	})
}
