package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAaatacacsparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaatacacsparamsFunc,
		ReadContext:   readAaatacacsparamsFunc,
		UpdateContext: updateAaatacacsparamsFunc,
		DeleteContext: deleteAaatacacsparamsFunc,
		Schema: map[string]*schema.Schema{
			"accounting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auditfailedcmds": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupattrname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tacacssecret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaatacacsparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaatacacsparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	aaatacacsparamsName := resource.PrefixedUniqueId("tf-aaatacacsparams-")

	aaatacacsparams := aaa.Aaatacacsparams{
		Accounting:                 d.Get("accounting").(string),
		Auditfailedcmds:            d.Get("auditfailedcmds").(string),
		Authorization:              d.Get("authorization").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Groupattrname:              d.Get("groupattrname").(string),
		Serverip:                   d.Get("serverip").(string),
		Tacacssecret:               d.Get("tacacssecret").(string),
	}

	if raw := d.GetRawConfig().GetAttr("authtimeout"); !raw.IsNull() {
		aaatacacsparams.Authtimeout = intPtr(d.Get("authtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("serverport"); !raw.IsNull() {
		aaatacacsparams.Serverport = intPtr(d.Get("serverport").(int))
	}

	err := client.UpdateUnnamedResource(service.Aaatacacsparams.Type(), &aaatacacsparams)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaatacacsparamsName)

	return readAaatacacsparamsFunc(ctx, d, meta)
}

func readAaatacacsparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaatacacsparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaatacacsparams state")
	data, err := client.FindResource(service.Aaatacacsparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaatacacsparams state")
		d.SetId("")
		return nil
	}
	d.Set("accounting", data["accounting"])
	d.Set("auditfailedcmds", data["auditfailedcmds"])
	d.Set("authorization", data["authorization"])
	setToInt("authtimeout", d, data["authtimeout"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("groupattrname", data["groupattrname"])
	d.Set("serverip", data["serverip"])
	setToInt("serverport", d, data["serverport"])
	d.Set("tacacssecret", data["tacacssecret"])

	return nil

}

func updateAaatacacsparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaatacacsparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	aaatacacsparams := aaa.Aaatacacsparams{}
	hasChange := false
	if d.HasChange("accounting") {
		log.Printf("[DEBUG]  citrixadc-provider: Accounting has changed for aaatacacsparams, starting update")
		aaatacacsparams.Accounting = d.Get("accounting").(string)
		hasChange = true
	}
	if d.HasChange("auditfailedcmds") {
		log.Printf("[DEBUG]  citrixadc-provider: Auditfailedcmds has changed for aaatacacsparams, starting update")
		aaatacacsparams.Auditfailedcmds = d.Get("auditfailedcmds").(string)
		hasChange = true
	}
	if d.HasChange("authorization") {
		log.Printf("[DEBUG]  citrixadc-provider: Authorization has changed for aaatacacsparams, starting update")
		aaatacacsparams.Authorization = d.Get("authorization").(string)
		hasChange = true
	}
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for aaatacacsparams, starting update")
		aaatacacsparams.Authtimeout = intPtr(d.Get("authtimeout").(int))
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for aaatacacsparams, starting update")
		aaatacacsparams.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("groupattrname") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupattrname has changed for aaatacacsparams, starting update")
		aaatacacsparams.Groupattrname = d.Get("groupattrname").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for aaatacacsparams, starting update")
		aaatacacsparams.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for aaatacacsparams, starting update")
		aaatacacsparams.Serverport = intPtr(d.Get("serverport").(int))
		hasChange = true
	}
	if d.HasChange("tacacssecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Tacacssecret has changed for aaatacacsparams, starting update")
		aaatacacsparams.Tacacssecret = d.Get("tacacssecret").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaatacacsparams.Type(), &aaatacacsparams)
		if err != nil {
			return diag.Errorf("Error updating aaatacacsparams")
		}
	}
	return readAaatacacsparamsFunc(ctx, d, meta)
}

func deleteAaatacacsparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaatacacsparamsFunc")
	// aaatacacsparams does not support DELETE operation
	d.SetId("")

	return nil
}
