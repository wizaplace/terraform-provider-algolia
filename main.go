package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"
    "github.com/wizaplace/terraform-provider-algolia/algolia"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return algolia.Provider()
        },
    })
}
