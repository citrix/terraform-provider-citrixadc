package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationepaaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationepaactionFunc,
		ReadContext:   readAuthenticationepaactionFunc,
		UpdateContext: updateAuthenticationepaactionFunc,
		DeleteContext: deleteAuthenticationepaactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"deviceposture": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"csecexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
			},
			"defaultepagroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deletefiles": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"killprocess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantinegroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationepaactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationepaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationepaactionName := d.Get("name").(string)
	authenticationepaaction := authentication.Authenticationepaaction{
		Csecexpr:        d.Get("csecexpr").(string),
		Defaultepagroup: d.Get("defaultepagroup").(string),
		Deletefiles:     d.Get("deletefiles").(string),
		Killprocess:     d.Get("killprocess").(string),
		Name:            d.Get("name").(string),
		Quarantinegroup: d.Get("quarantinegroup").(string),
		Deviceposture:   d.Get("deviceposture").(string),
	}

	_, err := client.AddResource("authenticationepaaction", authenticationepaactionName, &authenticationepaaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationepaactionName)

	return readAuthenticationepaactionFunc(ctx, d, meta)
}

func readAuthenticationepaactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationepaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationepaactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationepaaction state %s", authenticationepaactionName)
	data, err := client.FindResource("authenticationepaaction", authenticationepaactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationepaaction state %s", authenticationepaactionName)
		d.SetId("")
		return nil
	}
	d.Set("csecexpr", data["csecexpr"])
	d.Set("deviceposture", data["deviceposture"])
	d.Set("defaultepagroup", data["defaultepagroup"])
	d.Set("deletefiles", data["deletefiles"])
	d.Set("killprocess", data["killprocess"])
	d.Set("name", data["name"])
	d.Set("quarantinegroup", data["quarantinegroup"])

	return nil

}

func updateAuthenticationepaactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationepaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationepaactionName := d.Get("name").(string)

	authenticationepaaction := authentication.Authenticationepaaction{
		Name:     d.Get("name").(string),
		Csecexpr: d.Get("csecexpr").(string),
	}
	hasChange := false
	if d.HasChange("deviceposture") {
		log.Printf("[DEBUG]  citrixadc-provider: Deviceposture has changed for authenticationepaaction, starting update")
		authenticationepaaction.Deviceposture = d.Get("deviceposture").(string)
		hasChange = true
	}
	if d.HasChange("csecexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Csecexpr has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Csecexpr = d.Get("csecexpr").(string)
		hasChange = true
	}
	if d.HasChange("defaultepagroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultepagroup has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Defaultepagroup = d.Get("defaultepagroup").(string)
		hasChange = true
	}
	if d.HasChange("deletefiles") {
		log.Printf("[DEBUG]  citrixadc-provider: Deletefiles has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Deletefiles = d.Get("deletefiles").(string)
		hasChange = true
	}
	if d.HasChange("killprocess") {
		log.Printf("[DEBUG]  citrixadc-provider: Killprocess has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Killprocess = d.Get("killprocess").(string)
		hasChange = true
	}
	if d.HasChange("quarantinegroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Quarantinegroup has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Quarantinegroup = d.Get("quarantinegroup").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationepaaction", authenticationepaactionName, &authenticationepaaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationepaaction %s", authenticationepaactionName)
		}
	}
	return readAuthenticationepaactionFunc(ctx, d, meta)
}

func deleteAuthenticationepaactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationepaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationepaactionName := d.Id()
	err := client.DeleteResource("authenticationepaaction", authenticationepaactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
