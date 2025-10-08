package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationpushservice() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationpushserviceFunc,
		ReadContext:   readAuthenticationpushserviceFunc,
		UpdateContext: updateAuthenticationpushserviceFunc,
		DeleteContext: deleteAuthenticationpushserviceFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"clientid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientsecret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customerid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refreshinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationpushserviceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationpushserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpushserviceName := d.Get("name").(string)
	authenticationpushservice := authentication.Authenticationpushservice{
		Clientid:        d.Get("clientid").(string),
		Clientsecret:    d.Get("clientsecret").(string),
		Customerid:      d.Get("customerid").(string),
		Name:            d.Get("name").(string),
		Refreshinterval: d.Get("refreshinterval").(int),
	}

	_, err := client.AddResource("authenticationpushservice", authenticationpushserviceName, &authenticationpushservice)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationpushserviceName)

	return readAuthenticationpushserviceFunc(ctx, d, meta)
}

func readAuthenticationpushserviceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationpushserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpushserviceName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationpushservice state %s", authenticationpushserviceName)
	data, err := client.FindResource("authenticationpushservice", authenticationpushserviceName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationpushservice state %s", authenticationpushserviceName)
		d.SetId("")
		return nil
	}
	d.Set("clientid", data["clientid"])
	//d.Set("clientsecret", data["clientsecret"]) different value is received each time
	d.Set("customerid", data["customerid"])
	d.Set("name", data["name"])
	setToInt("refreshinterval", d, data["refreshinterval"])

	return nil

}

func updateAuthenticationpushserviceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationpushserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpushserviceName := d.Get("name").(string)

	authenticationpushservice := authentication.Authenticationpushservice{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("clientid") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientid has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Clientid = d.Get("clientid").(string)
		hasChange = true
	}
	if d.HasChange("clientsecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecret has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Clientsecret = d.Get("clientsecret").(string)
		hasChange = true
	}
	if d.HasChange("customerid") {
		log.Printf("[DEBUG]  citrixadc-provider: Customerid has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Customerid = d.Get("customerid").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("refreshinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Refreshinterval has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Refreshinterval = d.Get("refreshinterval").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationpushservice", authenticationpushserviceName, &authenticationpushservice)
		if err != nil {
			return diag.Errorf("Error updating authenticationpushservice %s", authenticationpushserviceName)
		}
	}
	return readAuthenticationpushserviceFunc(ctx, d, meta)
}

func deleteAuthenticationpushserviceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationpushserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpushserviceName := d.Id()
	err := client.DeleteResource("authenticationpushservice", authenticationpushserviceName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
