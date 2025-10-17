package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	//"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcBotsettings() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBotsettingsFunc,
		ReadContext:   readBotsettingsFunc,
		UpdateContext: updateBotsettingsFunc,
		DeleteContext: deleteBotsettingsFunc, // Thought botsettings resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"defaultprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultnonintrusiveprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"javascriptname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessiontimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessioncookiename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dfprequestlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signatureautoupdate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signatureurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxyserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxyport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trapurlautogenerate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trapurlinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trapurllength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createBotsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	var botsettingsName string

	// there is no primary key in BOTSETTINGS resource. Hence generate one for terraform state maintenance
	botsettingsName = resource.PrefixedUniqueId("tf-botsettings")

	botsettings := bot.Botsettings{
		Defaultprofile:             d.Get("defaultprofile").(string),
		Defaultnonintrusiveprofile: d.Get("defaultnonintrusiveprofile").(string),
		Javascriptname:             d.Get("javascriptname").(string),
		Sessioncookiename:          d.Get("sessioncookiename").(string),
		Signatureautoupdate:        d.Get("signatureautoupdate").(string),
		Signatureurl:               d.Get("signatureurl").(string),
		Proxyserver:                d.Get("proxyserver").(string),
		Trapurlautogenerate:        d.Get("trapurlautogenerate").(string),
	}

	if raw := d.GetRawConfig().GetAttr("sessiontimeout"); !raw.IsNull() {
		botsettings.Sessiontimeout = intPtr(d.Get("sessiontimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("dfprequestlimit"); !raw.IsNull() {
		botsettings.Dfprequestlimit = intPtr(d.Get("dfprequestlimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("proxyport"); !raw.IsNull() {
		botsettings.Proxyport = intPtr(d.Get("proxyport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("trapurlinterval"); !raw.IsNull() {
		botsettings.Trapurlinterval = intPtr(d.Get("trapurlinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("trapurllength"); !raw.IsNull() {
		botsettings.Trapurllength = intPtr(d.Get("trapurllength").(int))
	}

	err := client.UpdateUnnamedResource("botsettings", &botsettings)
	if err != nil {
		return diag.Errorf("Error updating botsettings")
	}

	d.SetId(botsettingsName)

	return readBotsettingsFunc(ctx, d, meta)
}

func readBotsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading Botsettings state")
	data, err := client.FindResource("botsettings", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing botsettings state")
		d.SetId("")
		return nil
	}
	d.Set("defaultprofile", data["defaultprofile"])
	d.Set("defaultnonintrusiveprofile", data["defaultnonintrusiveprofile"])
	d.Set("javascriptname", data["javascriptname"])
	setToInt("sessiontimeout", d, data["sessiontimeout"])
	d.Set("sessioncookiename", data["sessioncookiename"])
	setToInt("dfprequestlimit", d, data["dfprequestlimit"])
	d.Set("signatureautoupdate", data["signatureautoupdate"])
	d.Set("signatureurl", data["signatureurl"])
	d.Set("proxyserver", data["proxyserver"])
	setToInt("proxyport", d, data["proxyport"])
	d.Set("trapurlautogenerate", data["trapurlautogenerate"])
	setToInt("trapurlinterval", d, data["trapurlinterval"])
	setToInt("trapurllength", d, data["trapurllength"])

	return nil
}

func updateBotsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateBotsettingsFunc")
	client := meta.(*NetScalerNitroClient).client

	botsettings := bot.Botsettings{}

	hasChange := false
	if d.HasChange("defaultprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultprofile has changed for botsettings, starting update")
		botsettings.Defaultprofile = d.Get("defaultprofile").(string)
		hasChange = true
	}
	if d.HasChange("defaultnonintrusiveprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultprofile has changed for botsettings, starting update")
		botsettings.Defaultnonintrusiveprofile = d.Get("defaultnonintrusiveprofile").(string)
		hasChange = true
	}
	if d.HasChange("javascriptname") {
		log.Printf("[DEBUG]  citrixadc-provider: Javascriptname  has changed for botsettings, starting update")
		botsettings.Javascriptname = d.Get("javascriptname").(string)
		hasChange = true
	}
	if d.HasChange("sessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessiontimeout has changed for botsettings, starting update")
		botsettings.Sessiontimeout = intPtr(d.Get("sessiontimeout").(int))
		hasChange = true
	}
	if d.HasChange("sessioncookiename") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessioncookiename has changed for botsettings, starting update")
		botsettings.Sessioncookiename = d.Get("sessioncookiename").(string)
		hasChange = true
	}
	if d.HasChange("dfprequestlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Dfprequestlimit has changed for botsettings, starting update")
		botsettings.Dfprequestlimit = intPtr(d.Get("dfprequestlimit").(int))
		hasChange = true
	}
	if d.HasChange("signatureautoupdate") {
		log.Printf("[DEBUG]  citrixadc-provider: Signatureautoupdate has changed for botsettings, starting update")
		botsettings.Signatureautoupdate = d.Get("signatureautoupdate").(string)
		hasChange = true
	}
	if d.HasChange("signatureurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Signatureurl has changed for botsettings, starting update")
		botsettings.Signatureurl = d.Get("signatureurl").(string)
		hasChange = true
	}
	if d.HasChange("proxyserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyserver has changed for botsettings, starting update")
		botsettings.Proxyserver = d.Get("proxyserver").(string)
		hasChange = true
	}
	if d.HasChange("proxyport") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyport has changed for botsettings, starting update")
		botsettings.Proxyport = intPtr(d.Get("proxyport").(int))
		hasChange = true
	}
	if d.HasChange("trapurlautogenerate") {
		log.Printf("[DEBUG]  citrixadc-provider: Trapurlautogenerate has changed for botsettings, starting update")
		botsettings.Trapurlautogenerate = d.Get("trapurlautogenerate").(string)
		hasChange = true
	}
	if d.HasChange("trapurlinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Trapurlinterval has changed for botsettings, starting update")
		botsettings.Trapurlinterval = intPtr(d.Get("trapurlinterval").(int))
		hasChange = true
	}
	if d.HasChange("trapurllength") {
		log.Printf("[DEBUG]  citrixadc-provider: Trapurllength has changed for botsettings, starting update")
		botsettings.Trapurllength = intPtr(d.Get("trapurllength").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("botsettings", &botsettings)
		if err != nil {
			return diag.Errorf("Error updating botsettings: %s", err.Error())
		}
	}
	return readBotsettingsFunc(ctx, d, meta)
}

func deleteBotsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotsettingsFunc")
	// botsettings do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
