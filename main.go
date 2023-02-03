package main

import (
	"github.com/cnanaaron/terraform-provider-icinga2api/icinga2api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: icinga2api.Provider})
}
