package algolia

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "application_id": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The Algolia application ID",
            },
            "api_key": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The application ID",
            },
        },
        ResourcesMap: map[string]*schema.Resource{
            "algolia_api_key": resourceApiKey(),
        },
    }
}
