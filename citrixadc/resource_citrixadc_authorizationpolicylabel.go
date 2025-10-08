package citrixadc

import (
	"context"
	"log"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthorizationpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthorizationpolicylabelFunc,
		ReadContext:   readAuthorizationpolicylabelFunc,
		DeleteContext: deleteAuthorizationpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthorizationpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthorizationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicylabelName := d.Get("labelname").(string)
	authorizationpolicylabel := authorization.Authorizationpolicylabel{
		Labelname: authorizationpolicylabelName,
		Newname:   d.Get("newname").(string),
	}

	_, err := client.AddResource(service.Authorizationpolicylabel.Type(), authorizationpolicylabelName, &authorizationpolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authorizationpolicylabelName)

	return readAuthorizationpolicylabelFunc(ctx, d, meta)
}

func readAuthorizationpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthorizationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authorizationpolicylabel state %s", authorizationpolicylabelName)
	data, err := client.FindResource(service.Authorizationpolicylabel.Type(), authorizationpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authorizationpolicylabel state %s", authorizationpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("newname", data["newname"])

	return nil

}

func deleteAuthorizationpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthorizationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicylabelName := d.Id()
	err := client.DeleteResource(service.Authorizationpolicylabel.Type(), authorizationpolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
