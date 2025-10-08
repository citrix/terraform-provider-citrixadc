package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthorizationpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthorizationpolicyFunc,
		ReadContext:   readAuthorizationpolicyFunc,
		UpdateContext: updateAuthorizationpolicyFunc,
		DeleteContext: deleteAuthorizationpolicyFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Computed: false,
			},
			// "newname": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAuthorizationpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthorizationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicyName := d.Get("name").(string)
	authorizationpolicy := authorization.Authorizationpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		// Newname: d.Get("newname").(string),
		Rule: d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authorizationpolicy.Type(), authorizationpolicyName, &authorizationpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authorizationpolicyName)

	return readAuthorizationpolicyFunc(ctx, d, meta)
}

func readAuthorizationpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthorizationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authorizationpolicy state %s", authorizationpolicyName)
	data, err := client.FindResource(service.Authorizationpolicy.Type(), authorizationpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authorizationpolicy state %s", authorizationpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	// d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthorizationpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthorizationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicyName := d.Get("name").(string)

	authorizationpolicy := authorization.Authorizationpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authorizationpolicy %s, starting update", authorizationpolicyName)
		authorizationpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	// if d.HasChange("name") {
	// 	log.Printf("[DEBUG]  citrixadc-provider: Name has changed for authorizationpolicy %s, starting update", authorizationpolicyName)
	// 	authorizationpolicy.Name = d.Get("name").(string)
	// 	hasChange = true
	// }
	// if d.HasChange("newname") {
	// 	log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for authorizationpolicy %s, starting update", authorizationpolicyName)
	// 	authorizationpolicy.Newname = d.Get("newname").(string)
	// 	hasChange = true
	// }
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authorizationpolicy %s, starting update", authorizationpolicyName)
		authorizationpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authorizationpolicy.Type(), authorizationpolicyName, &authorizationpolicy)
		if err != nil {
			return diag.Errorf("error updating authorizationpolicy %s", authorizationpolicyName)
		}
	}
	return readAuthorizationpolicyFunc(ctx, d, meta)
}

func deleteAuthorizationpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthorizationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicyName := d.Id()
	err := client.DeleteResource(service.Authorizationpolicy.Type(), authorizationpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
