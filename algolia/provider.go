package algolia

import (
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Algolia application ID",
				DefaultFunc: schema.EnvDefaultFunc("ALGOLIA_APPLICATION_ID", nil),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API key",
				DefaultFunc: schema.EnvDefaultFunc("ALGOLIA_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"algolia_api_key": resourceApiKey(),
		},
	}

	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		return providerConfigure(d)
	}

	return p
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := search.NewClient(
		d.Get("application_id").(string),
		d.Get("api_key").(string),
	)

	return client, nil
}
