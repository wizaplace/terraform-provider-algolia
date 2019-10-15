# Terraform Provider Algolia

The Algolia provider is used to interact with the many resources supported by Algolia.
The provider needs to be configured with the proper credentials before it can be used.

## Installation

### Goland

require
* golang 1.13.1

clone the repository

````bash
git clone https://github.com/wizaplace/terraform-provider-algolia
````

compile terraform provider algolia
https://www.terraform.io/docs/configuration/providers.html#third-party-plugins

````bash
go build -o ~/.terraform.d/plugins/terraform-provider-algolia;
````

### Docker

require
* docker

````bash
./install.sh
````

## Authentication

### Static credentials

Warning: Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system.

Static credentials can be provided by adding an application_id and api_key in-line in the Algolia provider block:

usage:

````hcl-terraform
provider "algolia" {
    application_id = "{your application id}"
    api_key        = "{your api key}"
}
````

### Environment variables 

You can provide your credentials via the ALGOLIA_APPLICATION_ID and ALGOLIA_API_KEY, environment variables, representing your Algolia Application Id and Algolia Api Key, respectively.

usage:

````hcl-terraform
provider "algolia" {}
````

````bash
export ALGOLIA_APPLICATION_ID="{your application id}"
export ALGOLIA_API_KEY="{your api key}"
terraform plan
````

## Resource algolia_api_key

Provides an Algolia Api Key.

Example Usage:

````hcl-terraform
resource "algolia_api_key" "example" {
  acl         = list("search")
  description = "example"
  indexes     = list("example*")
}
````

### Argument Reference

The following arguments are supported:
* acl - (Required) Specify the list of permissions associated to the key. The possible acls are:
  * search: Allows search.
  * browse: Allows retrieval of all index contents via the browse API.
  * addObject: Allows adding/updating an object in the index. (Copying/moving indices are also allowed with this permission.)
  * deleteObject: Allows deleting an existing object.
  * deleteIndex: Allows deleting index content.
  * settings: allows getting index settings.
  * editSettings: Allows changing index settings.
  * analytics: Allows retrieval of analytics through the analytics API.
  * recommendation: Allows usage of the Personalization dashboard and the Recommendation API.
  * listIndexes: Allows listing all accessible indices.
  * logs: Allows getting the logs.
  * seeUnretrievableAttributes: Disables the unretrievableAttributes feature for all operations returning records.
* description - (Required) Specify a description to describe where the key is used.
* indexes - (Optional) Specify the list of targeted indices. You can target all indices starting with a prefix or ending with a suffix using the ‘*’ character. For example, “dev_*” matches all indices starting with “dev_” and “*_dev” matches all indices ending with “_dev”.  
* max_queries_per_ip_peer_hour - (Optional) Specify the maximum number of API calls allowed from an IP address per hour. Each time an API call is performed with this key, a check is performed. If the IP at the source of the call did more than this number of calls in the last hour, a 429 code is returned. This parameter can be used to protect you from attempts at retrieving your entire index contents by massively querying the index.
* max_hits_per_query - (Optional) Specify the maximum number of hits this API key can retrieve in one call. This parameter can be used to protect you from attempts at retrieving your entire index contents by massively querying the index.
* referers - (Optional) Specify the list of query parameters. You can force the query parameters for a query using the url string format. Example: “typoTolerance=strict&ignorePlurals=false”.

### Attributes Reference

In addition to all arguments above, the following attributes are exported:
* id - The API Key
