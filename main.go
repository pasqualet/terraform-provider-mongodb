package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pasqualet/terraform-provider-mongodb/mongodb"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mongodb.Provider})
}
