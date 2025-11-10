package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLsnhttphdrlogprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnhttphdrlogprofileFunc,
		ReadContext:   readLsnhttphdrlogprofileFunc,
		UpdateContext: updateLsnhttphdrlogprofileFunc,
		DeleteContext: deleteLsnhttphdrlogprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"httphdrlogprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"loghost": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnhttphdrlogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnhttphdrlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnhttphdrlogprofileName := d.Get("httphdrlogprofilename").(string)
	lsnhttphdrlogprofile := lsn.Lsnhttphdrlogprofile{
		Httphdrlogprofilename: d.Get("httphdrlogprofilename").(string),
		Loghost:               d.Get("loghost").(string),
		Logmethod:             d.Get("logmethod").(string),
		Logurl:                d.Get("logurl").(string),
		Logversion:            d.Get("logversion").(string),
	}

	_, err := client.AddResource("lsnhttphdrlogprofile", lsnhttphdrlogprofileName, &lsnhttphdrlogprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnhttphdrlogprofileName)

	return readLsnhttphdrlogprofileFunc(ctx, d, meta)
}

func readLsnhttphdrlogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnhttphdrlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnhttphdrlogprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnhttphdrlogprofile state %s", lsnhttphdrlogprofileName)
	data, err := client.FindResource("lsnhttphdrlogprofile", lsnhttphdrlogprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnhttphdrlogprofile state %s", lsnhttphdrlogprofileName)
		d.SetId("")
		return nil
	}
	d.Set("httphdrlogprofilename", data["httphdrlogprofilename"])
	d.Set("loghost", data["loghost"])
	d.Set("logmethod", data["logmethod"])
	d.Set("logurl", data["logurl"])
	d.Set("logversion", data["logversion"])

	return nil

}

func updateLsnhttphdrlogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnhttphdrlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnhttphdrlogprofileName := d.Get("httphdrlogprofilename").(string)

	lsnhttphdrlogprofile := lsn.Lsnhttphdrlogprofile{
		Httphdrlogprofilename: d.Get("httphdrlogprofilename").(string),
	}
	hasChange := false
	if d.HasChange("loghost") {
		log.Printf("[DEBUG]  citrixadc-provider: Loghost has changed for lsnhttphdrlogprofile %s, starting update", lsnhttphdrlogprofileName)
		lsnhttphdrlogprofile.Loghost = d.Get("loghost").(string)
		hasChange = true
	}
	if d.HasChange("logmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Logmethod has changed for lsnhttphdrlogprofile %s, starting update", lsnhttphdrlogprofileName)
		lsnhttphdrlogprofile.Logmethod = d.Get("logmethod").(string)
		hasChange = true
	}
	if d.HasChange("logurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Logurl has changed for lsnhttphdrlogprofile %s, starting update", lsnhttphdrlogprofileName)
		lsnhttphdrlogprofile.Logurl = d.Get("logurl").(string)
		hasChange = true
	}
	if d.HasChange("logversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Logversion has changed for lsnhttphdrlogprofile %s, starting update", lsnhttphdrlogprofileName)
		lsnhttphdrlogprofile.Logversion = d.Get("logversion").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnhttphdrlogprofile", &lsnhttphdrlogprofile)
		if err != nil {
			return diag.Errorf("Error updating lsnhttphdrlogprofile %s", lsnhttphdrlogprofileName)
		}
	}
	return readLsnhttphdrlogprofileFunc(ctx, d, meta)
}

func deleteLsnhttphdrlogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnhttphdrlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnhttphdrlogprofileName := d.Id()
	err := client.DeleteResource("lsnhttphdrlogprofile", lsnhttphdrlogprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
