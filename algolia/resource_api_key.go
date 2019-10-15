package algolia

import (
	"errors"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApiKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiKeyCreate,
		Read:   resourceApiKeyRead,
		Update: resourceApiKeyUpdate,
		Delete: resourceApiKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"acl": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"indexes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"max_queries_per_ip_peer_hour": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15000,
			},
			"max_hits_per_query": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"referers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}
func resourceApiKeyCreate(d *schema.ResourceData, m interface{}) error {
	res, err := m.(*search.Client).AddAPIKey(getAlgoliaKey(d))
	if err != nil {
		return errors.New("Error in create: " + err.Error())
	}

	err = res.Wait()
	if err != nil {
		return err
	}

	d.SetId(res.Key)

	return resourceApiKeyRead(d, m)
}

func resourceApiKeyRead(d *schema.ResourceData, m interface{}) error {
	key, err := m.(*search.Client).GetAPIKey(d.Id())
	if err != nil {
		d.SetId("")
		return errors.New("Error in read: " + err.Error())
	}

	d.SetId(key.Value)
	_ = d.Set("acl", key.ACL)
	_ = d.Set("description", key.Description)
	_ = d.Set("indexes", key.Indexes)
	_ = d.Set("max_queries_per_ip_peer_hour", key.MaxQueriesPerIPPerHour)
	_ = d.Set("max_hits_per_query", key.MaxHitsPerQuery)
	_ = d.Set("referers", key.Referers)

	return nil
}

func resourceApiKeyUpdate(d *schema.ResourceData, m interface{}) error {
	res, err := m.(*search.Client).UpdateAPIKey(getAlgoliaKey(d))
	if err != nil {
		return errors.New("Error in update: " + err.Error())
	}

	err = res.Wait()
	if err != nil {
		return err
	}

	return resourceApiKeyRead(d, m)
}

func resourceApiKeyDelete(d *schema.ResourceData, m interface{}) error {
	res, err := m.(*search.Client).DeleteAPIKey(d.Id())
	if err != nil {
		return errors.New("Error in delete: " + err.Error())
	}

	return res.Wait()
}

func getAlgoliaKey(d *schema.ResourceData) search.Key {
	var acl []string
	if value := d.Get("acl"); value != nil {
		for _, v := range value.(*schema.Set).List() {
			acl = append(acl, v.(string))
		}
	}

	var indexes []string
	if value := d.Get("indexes"); value != nil {
		for _, v := range value.(*schema.Set).List() {
			indexes = append(indexes, v.(string))
		}
	}

	var referers []string
	if value := d.Get("referers"); value != nil {
		for _, v := range value.(*schema.Set).List() {
			referers = append(referers, v.(string))
		}
	}

	return search.Key{
		Value:                  d.Id(),
		ACL:                    acl,
		Description:            d.Get("description").(string),
		Indexes:                indexes,
		MaxQueriesPerIPPerHour: d.Get("max_queries_per_ip_peer_hour").(int),
		MaxHitsPerQuery:        d.Get("max_hits_per_query").(int),
		Referers:               referers,
	}
}
