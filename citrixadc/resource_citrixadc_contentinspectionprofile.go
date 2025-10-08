package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcContentinspectionprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createContentinspectionprofileFunc,
		ReadContext:   readContentinspectionprofileFunc,
		UpdateContext: updateContentinspectionprofileFunc,
		DeleteContext: deleteContentinspectionprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"egressinterface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egressvlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ingressinterface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ingressvlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iptunnel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createContentinspectionprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionprofileName := d.Get("name").(string)
	contentinspectionprofile := contentinspection.Contentinspectionprofile{
		Egressinterface:  d.Get("egressinterface").(string),
		Egressvlan:       d.Get("egressvlan").(int),
		Ingressinterface: d.Get("ingressinterface").(string),
		Ingressvlan:      d.Get("ingressvlan").(int),
		Iptunnel:         d.Get("iptunnel").(string),
		Name:             d.Get("name").(string),
		Type:             d.Get("type").(string),
	}

	_, err := client.AddResource("contentinspectionprofile", contentinspectionprofileName, &contentinspectionprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(contentinspectionprofileName)

	return readContentinspectionprofileFunc(ctx, d, meta)
}

func readContentinspectionprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionprofile state %s", contentinspectionprofileName)
	data, err := client.FindResource("contentinspectionprofile", contentinspectionprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionprofile state %s", contentinspectionprofileName)
		d.SetId("")
		return nil
	}
	d.Set("egressinterface", data["egressinterface"])
	setToInt("egressvlan", d, data["egressvlan"])
	d.Set("ingressinterface", data["ingressinterface"])
	setToInt("ingressvlan", d, data["ingressvlan"])
	d.Set("iptunnel", data["iptunnel"])
	d.Set("name", data["name"])
	d.Set("type", data["type"])

	return nil

}

func updateContentinspectionprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectionprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionprofileName := d.Get("name").(string)

	contentinspectionprofile := contentinspection.Contentinspectionprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("egressinterface") {
		log.Printf("[DEBUG]  citrixadc-provider: Egressinterface has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Egressinterface = d.Get("egressinterface").(string)
		hasChange = true
	}
	if d.HasChange("egressvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Egressvlan has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Egressvlan = d.Get("egressvlan").(int)
		hasChange = true
	}
	if d.HasChange("ingressinterface") {
		log.Printf("[DEBUG]  citrixadc-provider: Ingressinterface has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Ingressinterface = d.Get("ingressinterface").(string)
		hasChange = true
	}
	if d.HasChange("ingressvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Ingressvlan has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Ingressvlan = d.Get("ingressvlan").(int)
		hasChange = true
	}
	if d.HasChange("iptunnel") {
		log.Printf("[DEBUG]  citrixadc-provider: Iptunnel has changed for contentinspectionprofile %s, starting update", contentinspectionprofileName)
		contentinspectionprofile.Iptunnel = d.Get("iptunnel").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectionprofile", &contentinspectionprofile)
		if err != nil {
			return diag.Errorf("Error updating contentinspectionprofile %s", contentinspectionprofileName)
		}
	}
	return readContentinspectionprofileFunc(ctx, d, meta)
}

func deleteContentinspectionprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionprofileName := d.Id()
	err := client.DeleteResource("contentinspectionprofile", contentinspectionprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
