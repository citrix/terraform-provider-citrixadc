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

func resourceCitrixAdcAaaradiusparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaaradiusparamsFunc,
		ReadContext:   readAaaradiusparamsFunc,
		UpdateContext: updateAaaradiusparamsFunc,
		DeleteContext: deleteAaaradiusparamsFunc,
		Schema: map[string]*schema.Schema{
			"radkey": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accounting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authservretry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"callingstationid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipattributetype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipvendorid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"passencoding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pwdattributetype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pwdvendorid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radattributetype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radgroupseparator": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radgroupsprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radvendorid": {
				Type:     schema.TypeInt,
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
			"tunnelendpointclientip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaaradiusparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaaradiusparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	aaaradiusparamsName := resource.PrefixedUniqueId("tf-aaaradiusparams-")

	aaaradiusparams := aaa.Aaaradiusparams{
		Accounting:                 d.Get("accounting").(string),
		Authentication:             d.Get("authentication").(string),
		Callingstationid:           d.Get("callingstationid").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Passencoding:               d.Get("passencoding").(string),
		Radgroupseparator:          d.Get("radgroupseparator").(string),
		Radgroupsprefix:            d.Get("radgroupsprefix").(string),
		Radkey:                     d.Get("radkey").(string),
		Radnasid:                   d.Get("radnasid").(string),
		Radnasip:                   d.Get("radnasip").(string),
		Serverip:                   d.Get("serverip").(string),
		Tunnelendpointclientip:     d.Get("tunnelendpointclientip").(string),
	}

	if raw := d.GetRawConfig().GetAttr("authservretry"); !raw.IsNull() {
		aaaradiusparams.Authservretry = intPtr(d.Get("authservretry").(int))
	}
	if raw := d.GetRawConfig().GetAttr("authtimeout"); !raw.IsNull() {
		aaaradiusparams.Authtimeout = intPtr(d.Get("authtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ipattributetype"); !raw.IsNull() {
		aaaradiusparams.Ipattributetype = intPtr(d.Get("ipattributetype").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ipvendorid"); !raw.IsNull() {
		aaaradiusparams.Ipvendorid = intPtr(d.Get("ipvendorid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("pwdattributetype"); !raw.IsNull() {
		aaaradiusparams.Pwdattributetype = intPtr(d.Get("pwdattributetype").(int))
	}
	if raw := d.GetRawConfig().GetAttr("pwdvendorid"); !raw.IsNull() {
		aaaradiusparams.Pwdvendorid = intPtr(d.Get("pwdvendorid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("radattributetype"); !raw.IsNull() {
		aaaradiusparams.Radattributetype = intPtr(d.Get("radattributetype").(int))
	}
	if raw := d.GetRawConfig().GetAttr("radvendorid"); !raw.IsNull() {
		aaaradiusparams.Radvendorid = intPtr(d.Get("radvendorid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("serverport"); !raw.IsNull() {
		aaaradiusparams.Serverport = intPtr(d.Get("serverport").(int))
	}

	err := client.UpdateUnnamedResource(service.Aaaradiusparams.Type(), &aaaradiusparams)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaaradiusparamsName)

	return readAaaradiusparamsFunc(ctx, d, meta)
}

func readAaaradiusparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaaradiusparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaaradiusparams state")
	data, err := client.FindResource(service.Aaaradiusparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaaradiusparams state")
		d.SetId("")
		return nil
	}
	d.Set("accounting", data["accounting"])
	d.Set("authentication", data["authentication"])
	setToInt("authservretry", d, data["authservretry"])
	setToInt("authtimeout", d, data["authtimeout"])
	d.Set("callingstationid", data["callingstationid"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	setToInt("ipattributetype", d, data["ipattributetype"])
	setToInt("ipvendorid", d, data["ipvendorid"])
	d.Set("passencoding", data["passencoding"])
	setToInt("pwdattributetype", d, data["pwdattributetype"])
	setToInt("pwdvendorid", d, data["pwdvendorid"])
	setToInt("radattributetype", d, data["radattributetype"])
	d.Set("radgroupseparator", data["radgroupseparator"])
	d.Set("radgroupsprefix", data["radgroupsprefix"])
	//d.Set("radkey", data["radkey"])
	d.Set("radnasid", data["radnasid"])
	d.Set("radnasip", data["radnasip"])
	setToInt("radvendorid", d, data["radvendorid"])
	d.Set("serverip", data["serverip"])
	setToInt("serverport", d, data["serverport"])
	d.Set("tunnelendpointclientip", data["tunnelendpointclientip"])

	return nil

}

func updateAaaradiusparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaaradiusparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	aaaradiusparams := aaa.Aaaradiusparams{
		Radkey: d.Get("radkey").(string),
	}
	hasChange := false
	if d.HasChange("accounting") {
		log.Printf("[DEBUG]  citrixadc-provider: Accounting has changed for aaaradiusparams, starting update")
		aaaradiusparams.Accounting = d.Get("accounting").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for aaaradiusparams, starting update")
		aaaradiusparams.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authservretry") {
		log.Printf("[DEBUG]  citrixadc-provider: Authservretry has changed for aaaradiusparams, starting update")
		aaaradiusparams.Authservretry = intPtr(d.Get("authservretry").(int))
		hasChange = true
	}
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for aaaradiusparams, starting update")
		aaaradiusparams.Authtimeout = intPtr(d.Get("authtimeout").(int))
		hasChange = true
	}
	if d.HasChange("callingstationid") {
		log.Printf("[DEBUG]  citrixadc-provider: Callingstationid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Callingstationid = d.Get("callingstationid").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for aaaradiusparams, starting update")
		aaaradiusparams.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("ipattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipattributetype has changed for aaaradiusparams, starting update")
		aaaradiusparams.Ipattributetype = intPtr(d.Get("ipattributetype").(int))
		hasChange = true
	}
	if d.HasChange("ipvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipvendorid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Ipvendorid = intPtr(d.Get("ipvendorid").(int))
		hasChange = true
	}
	if d.HasChange("passencoding") {
		log.Printf("[DEBUG]  citrixadc-provider: Passencoding has changed for aaaradiusparams, starting update")
		aaaradiusparams.Passencoding = d.Get("passencoding").(string)
		hasChange = true
	}
	if d.HasChange("pwdattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Pwdattributetype has changed for aaaradiusparams, starting update")
		aaaradiusparams.Pwdattributetype = intPtr(d.Get("pwdattributetype").(int))
		hasChange = true
	}
	if d.HasChange("pwdvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Pwdvendorid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Pwdvendorid = intPtr(d.Get("pwdvendorid").(int))
		hasChange = true
	}
	if d.HasChange("radattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Radattributetype has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radattributetype = intPtr(d.Get("radattributetype").(int))
		hasChange = true
	}
	if d.HasChange("radgroupseparator") {
		log.Printf("[DEBUG]  citrixadc-provider: Radgroupseparator has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radgroupseparator = d.Get("radgroupseparator").(string)
		hasChange = true
	}
	if d.HasChange("radgroupsprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Radgroupsprefix has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radgroupsprefix = d.Get("radgroupsprefix").(string)
		hasChange = true
	}
	if d.HasChange("radnasid") {
		log.Printf("[DEBUG]  citrixadc-provider: Radnasid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radnasid = d.Get("radnasid").(string)
		hasChange = true
	}
	if d.HasChange("radnasip") {
		log.Printf("[DEBUG]  citrixadc-provider: Radnasip has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radnasip = d.Get("radnasip").(string)
		hasChange = true
	}
	if d.HasChange("radvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Radvendorid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radvendorid = intPtr(d.Get("radvendorid").(int))
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for aaaradiusparams, starting update")
		aaaradiusparams.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for aaaradiusparams, starting update")
		aaaradiusparams.Serverport = intPtr(d.Get("serverport").(int))
		hasChange = true
	}
	if d.HasChange("tunnelendpointclientip") {
		log.Printf("[DEBUG]  citrixadc-provider: Tunnelendpointclientip has changed for aaaradiusparams, starting update")
		aaaradiusparams.Tunnelendpointclientip = d.Get("tunnelendpointclientip").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaaradiusparams.Type(), &aaaradiusparams)
		if err != nil {
			return diag.Errorf("Error updating aaaradiusparams")
		}
	}
	return readAaaradiusparamsFunc(ctx, d, meta)
}

func deleteAaaradiusparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaaradiusparamsFunc")
	// aaaradiusparams does not support delete operation
	d.SetId("")

	return nil
}
