package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcTmsessionaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTmsessionactionFunc,
		ReadContext:   readTmsessionactionFunc,
		UpdateContext: updateTmsessionactionFunc,
		DeleteContext: deleteTmsessionactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"defaultauthorizationaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"homepage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httponlycookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookievalidity": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sesstimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sso": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssocredential": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssodomain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createTmsessionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionactionName := d.Get("name").(string)

	tmsessionaction := tm.Tmsessionaction{
		Defaultauthorizationaction: d.Get("defaultauthorizationaction").(string),
		Homepage:                   d.Get("homepage").(string),
		Httponlycookie:             d.Get("httponlycookie").(string),
		Kcdaccount:                 d.Get("kcdaccount").(string),
		Name:                       d.Get("name").(string),
		Persistentcookie:           d.Get("persistentcookie").(string),
		Sso:                        d.Get("sso").(string),
		Ssocredential:              d.Get("ssocredential").(string),
		Ssodomain:                  d.Get("ssodomain").(string),
	}

	if raw := d.GetRawConfig().GetAttr("persistentcookievalidity"); !raw.IsNull() {
		tmsessionaction.Persistentcookievalidity = intPtr(d.Get("persistentcookievalidity").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sesstimeout"); !raw.IsNull() {
		tmsessionaction.Sesstimeout = intPtr(d.Get("sesstimeout").(int))
	}

	_, err := client.AddResource(service.Tmsessionaction.Type(), tmsessionactionName, &tmsessionaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(tmsessionactionName)

	return readTmsessionactionFunc(ctx, d, meta)
}

func readTmsessionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading tmsessionaction state %s", tmsessionactionName)
	data, err := client.FindResource(service.Tmsessionaction.Type(), tmsessionactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tmsessionaction state %s", tmsessionactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthorizationaction", data["defaultauthorizationaction"])
	d.Set("homepage", data["homepage"])
	d.Set("httponlycookie", data["httponlycookie"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("name", data["name"])
	d.Set("persistentcookie", data["persistentcookie"])
	setToInt("persistentcookievalidity", d, data["persistentcookievalidity"])
	setToInt("sesstimeout", d, data["sesstimeout"])
	d.Set("sso", data["sso"])
	d.Set("ssocredential", data["ssocredential"])
	d.Set("ssodomain", data["ssodomain"])

	return nil

}

func updateTmsessionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTmsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionactionName := d.Get("name").(string)

	tmsessionaction := tm.Tmsessionaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthorizationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthorizationaction has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Defaultauthorizationaction = d.Get("defaultauthorizationaction").(string)
		hasChange = true
	}
	if d.HasChange("homepage") {
		log.Printf("[DEBUG]  citrixadc-provider: Homepage has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Homepage = d.Get("homepage").(string)
		hasChange = true
	}
	if d.HasChange("httponlycookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httponlycookie has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Httponlycookie = d.Get("httponlycookie").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookie has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Persistentcookie = d.Get("persistentcookie").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookievalidity") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookievalidity has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Persistentcookievalidity = intPtr(d.Get("persistentcookievalidity").(int))
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Sesstimeout = intPtr(d.Get("sesstimeout").(int))
		hasChange = true
	}
	if d.HasChange("sso") {
		log.Printf("[DEBUG]  citrixadc-provider: Sso has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Sso = d.Get("sso").(string)
		hasChange = true
	}
	if d.HasChange("ssocredential") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssocredential has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Ssocredential = d.Get("ssocredential").(string)
		hasChange = true
	}
	if d.HasChange("ssodomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssodomain has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Ssodomain = d.Get("ssodomain").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tmsessionaction.Type(), &tmsessionaction)
		if err != nil {
			return diag.Errorf("Error updating tmsessionaction %s", tmsessionactionName)
		}
	}
	return readTmsessionactionFunc(ctx, d, meta)
}

func deleteTmsessionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionactionName := d.Id()
	err := client.DeleteResource(service.Tmsessionaction.Type(), tmsessionactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
