package main

import (
	"github.com/cnanaaron/terraform-provider-icinga2/icinga2api"
	"github.com/cnanaaron/terraform-provider-icinga2/iapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: icinga2api.Provider})
}
