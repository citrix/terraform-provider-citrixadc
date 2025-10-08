package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIcaaccessprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIcaaccessprofileFunc,
		ReadContext:   readIcaaccessprofileFunc,
		UpdateContext: updateIcaaccessprofileFunc,
		DeleteContext: deleteIcaaccessprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"clientaudioredirection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientclipboardredirection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientcomportredirection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientdriveredirection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientprinterredirection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientusbdriveredirection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connectclientlptports": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localremotedatasharing": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multistream": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIcaaccessprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcaaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icaaccessprofileName := d.Get("name").(string)
	icaaccessprofile := ica.Icaaccessprofile{
		Clientaudioredirection:     d.Get("clientaudioredirection").(string),
		Clientclipboardredirection: d.Get("clientclipboardredirection").(string),
		Clientcomportredirection:   d.Get("clientcomportredirection").(string),
		Clientdriveredirection:     d.Get("clientdriveredirection").(string),
		Clientprinterredirection:   d.Get("clientprinterredirection").(string),
		Clientusbdriveredirection:  d.Get("clientusbdriveredirection").(string),
		Connectclientlptports:      d.Get("connectclientlptports").(string),
		Localremotedatasharing:     d.Get("localremotedatasharing").(string),
		Multistream:                d.Get("multistream").(string),
		Name:                       d.Get("name").(string),
	}

	_, err := client.AddResource("icaaccessprofile", icaaccessprofileName, &icaaccessprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(icaaccessprofileName)

	return readIcaaccessprofileFunc(ctx, d, meta)
}

func readIcaaccessprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcaaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icaaccessprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading icaaccessprofile state %s", icaaccessprofileName)
	data, err := client.FindResource("icaaccessprofile", icaaccessprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icaaccessprofile state %s", icaaccessprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("clientaudioredirection", data["clientaudioredirection"])
	d.Set("clientclipboardredirection", data["clientclipboardredirection"])
	d.Set("clientcomportredirection", data["clientcomportredirection"])
	d.Set("clientdriveredirection", data["clientdriveredirection"])
	d.Set("clientprinterredirection", data["clientprinterredirection"])
	d.Set("clientusbdriveredirection", data["clientusbdriveredirection"])
	d.Set("connectclientlptports", data["connectclientlptports"])
	d.Set("localremotedatasharing", data["localremotedatasharing"])
	d.Set("multistream", data["multistream"])

	return nil

}

func updateIcaaccessprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcaaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icaaccessprofileName := d.Get("name").(string)

	icaaccessprofile := ica.Icaaccessprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("clientaudioredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientaudioredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientaudioredirection = d.Get("clientaudioredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientclipboardredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientclipboardredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientclipboardredirection = d.Get("clientclipboardredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientcomportredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientcomportredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientcomportredirection = d.Get("clientcomportredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientdriveredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientdriveredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientdriveredirection = d.Get("clientdriveredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientprinterredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientprinterredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientprinterredirection = d.Get("clientprinterredirection").(string)
		hasChange = true
	}
	if d.HasChange("clientusbdriveredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientusbdriveredirection has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Clientusbdriveredirection = d.Get("clientusbdriveredirection").(string)
		hasChange = true
	}
	if d.HasChange("connectclientlptports") {
		log.Printf("[DEBUG]  citrixadc-provider: Connectclientlptports has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Connectclientlptports = d.Get("connectclientlptports").(string)
		hasChange = true
	}
	if d.HasChange("localremotedatasharing") {
		log.Printf("[DEBUG]  citrixadc-provider: Localremotedatasharing has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Localremotedatasharing = d.Get("localremotedatasharing").(string)
		hasChange = true
	}
	if d.HasChange("multistream") {
		log.Printf("[DEBUG]  citrixadc-provider: Multistream has changed for icaaccessprofile %s, starting update", icaaccessprofileName)
		icaaccessprofile.Multistream = d.Get("multistream").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icaaccessprofile", &icaaccessprofile)
		if err != nil {
			return diag.Errorf("Error updating icaaccessprofile %s", icaaccessprofileName)
		}
	}
	return readIcaaccessprofileFunc(ctx, d, meta)
}

func deleteIcaaccessprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcaaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icaaccessprofileName := d.Id()
	err := client.DeleteResource("icaaccessprofile", icaaccessprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
