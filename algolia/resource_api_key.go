package algolia

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceApiKey() *schema.Resource {
    return &schema.Resource{
        Create: resourceApiKeyCreate,
        Read:   resourceApiKeyRead,
        Update: resourceApiKeyUpdate,
        Delete: resourceApiKeyDelete,

        Schema: map[string]*schema.Schema{
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceApiKeyCreate(d *schema.ResourceData, m interface{}) error {
     address := d.Get("address").(string)
     d.SetId(address)

     return resourceApiKeyRead(d, m)
}

func resourceApiKeyRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceApiKeyUpdate(d *schema.ResourceData, m interface{}) error {
    d.Partial(true)

    if d.HasChange("address") {
        // Try updating the address
//        if err := updateAddress(d, m); err != nil {
//            return err
//        }

        d.SetPartial("address")
    }

    return resourceApiKeyRead(d, m)
}

func resourceApiKeyDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}
