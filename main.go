package main

import (
	"github.com/cloudfirst-dev/terraform-provider-kea/kea"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return kea.Provider()
		},
	})
}
