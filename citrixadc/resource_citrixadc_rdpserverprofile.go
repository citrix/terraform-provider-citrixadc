package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/rdp"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRdpserverprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRdpserverprofileFunc,
		ReadContext:   readRdpserverprofileFunc,
		UpdateContext: updateRdpserverprofileFunc,
		DeleteContext: deleteRdpserverprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"psk": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdpip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdpport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rdpredirection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRdpserverprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRdpserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpserverprofileName := d.Get("name").(string)
	rdpserverprofile := rdp.Rdpserverprofile{
		Name:           d.Get("name").(string),
		Psk:            d.Get("psk").(string),
		Rdpip:          d.Get("rdpip").(string),
		Rdpredirection: d.Get("rdpredirection").(string),
	}

	if raw := d.GetRawConfig().GetAttr("rdpport"); !raw.IsNull() {
		rdpserverprofile.Rdpport = intPtr(d.Get("rdpport").(int))
	}

	_, err := client.AddResource("rdpserverprofile", rdpserverprofileName, &rdpserverprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rdpserverprofileName)

	return readRdpserverprofileFunc(ctx, d, meta)
}

func readRdpserverprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRdpserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpserverprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rdpserverprofile state %s", rdpserverprofileName)
	data, err := client.FindResource("rdpserverprofile", rdpserverprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rdpserverprofile state %s", rdpserverprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("rdpip", data["rdpip"])
	setToInt("rdpport", d, data["rdpport"])
	d.Set("rdpredirection", data["rdpredirection"])

	return nil

}

func updateRdpserverprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRdpserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpserverprofileName := d.Get("name").(string)

	rdpserverprofile := rdp.Rdpserverprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("psk") {
		log.Printf("[DEBUG]  citrixadc-provider: Psk has changed for rdpserverprofile %s, starting update", rdpserverprofileName)
		rdpserverprofile.Psk = d.Get("psk").(string)
		hasChange = true
	}
	if d.HasChange("rdpip") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpip has changed for rdpserverprofile %s, starting update", rdpserverprofileName)
		rdpserverprofile.Rdpip = d.Get("rdpip").(string)
		hasChange = true
	}
	if d.HasChange("rdpport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpport has changed for rdpserverprofile %s, starting update", rdpserverprofileName)
		rdpserverprofile.Rdpport = intPtr(d.Get("rdpport").(int))
		hasChange = true
	}
	if d.HasChange("rdpredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpredirection has changed for rdpserverprofile %s, starting update", rdpserverprofileName)
		rdpserverprofile.Rdpredirection = d.Get("rdpredirection").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("rdpserverprofile", &rdpserverprofile)
		if err != nil {
			return diag.Errorf("Error updating rdpserverprofile %s", rdpserverprofileName)
		}
	}
	return readRdpserverprofileFunc(ctx, d, meta)
}

func deleteRdpserverprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRdpserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	rdpserverprofileName := d.Id()
	err := client.DeleteResource("rdpserverprofile", rdpserverprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
