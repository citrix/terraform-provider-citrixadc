package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSsllogprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSsllogprofileFunc,
		ReadContext:   readSsllogprofileFunc,
		UpdateContext: updateSsllogprofileFunc,
		DeleteContext: deleteSsllogprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssllogclauth": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssllogclauthfailures": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslloghs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslloghsfailures": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSsllogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSsllogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	var ssllogprofileName string
	ssllogprofileName = d.Get("name").(string)

	ssllogprofile := ssl.Ssllogprofile{
		Name:                 ssllogprofileName,
		Ssllogclauth:         d.Get("ssllogclauth").(string),
		Ssllogclauthfailures: d.Get("ssllogclauthfailures").(string),
		Sslloghs:             d.Get("sslloghs").(string),
		Sslloghsfailures:     d.Get("sslloghsfailures").(string),
	}

	_, err := client.AddResource("ssllogprofile", ssllogprofileName, &ssllogprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ssllogprofileName)

	return readSsllogprofileFunc(ctx, d, meta)
}

func readSsllogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSsllogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ssllogprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ssllogprofile state %s", ssllogprofileName)
	data, err := client.FindResource("ssllogprofile", ssllogprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ssllogprofile state %s", ssllogprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("name", data["name"])
	d.Set("ssllogclauth", data["ssllogclauth"])
	d.Set("ssllogclauthfailures", data["ssllogclauthfailures"])
	d.Set("sslloghs", data["sslloghs"])
	d.Set("sslloghsfailures", data["sslloghsfailures"])

	return nil

}

func updateSsllogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSsllogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ssllogprofileName := d.Get("name").(string)

	ssllogprofile := ssl.Ssllogprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for ssllogprofile %s, starting update", ssllogprofileName)
		ssllogprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("ssllogclauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssllogclauth has changed for ssllogprofile %s, starting update", ssllogprofileName)
		ssllogprofile.Ssllogclauth = d.Get("ssllogclauth").(string)
		hasChange = true
	}
	if d.HasChange("ssllogclauthfailures") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssllogclauthfailures has changed for ssllogprofile %s, starting update", ssllogprofileName)
		ssllogprofile.Ssllogclauthfailures = d.Get("ssllogclauthfailures").(string)
		hasChange = true
	}
	if d.HasChange("sslloghs") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslloghs has changed for ssllogprofile %s, starting update", ssllogprofileName)
		ssllogprofile.Sslloghs = d.Get("sslloghs").(string)
		hasChange = true
	}
	if d.HasChange("sslloghsfailures") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslloghsfailures has changed for ssllogprofile %s, starting update", ssllogprofileName)
		ssllogprofile.Sslloghsfailures = d.Get("sslloghsfailures").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("ssllogprofile", ssllogprofileName, &ssllogprofile)
		if err != nil {
			return diag.Errorf("Error updating ssllogprofile %s", ssllogprofileName)
		}
	}
	return readSsllogprofileFunc(ctx, d, meta)
}

func deleteSsllogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSsllogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ssllogprofileName := d.Id()
	err := client.DeleteResource("ssllogprofile", ssllogprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
