package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAutoscaleprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAutoscaleprofileFunc,
		ReadContext:   readAutoscaleprofileFunc,
		UpdateContext: updateAutoscaleprofileFunc,
		DeleteContext: deleteAutoscaleprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apikey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sharedsecret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAutoscaleprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAutoscaleprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleprofileName := d.Get("name").(string)
	autoscaleprofile := autoscale.Autoscaleprofile{
		Apikey:       d.Get("apikey").(string),
		Name:         d.Get("name").(string),
		Sharedsecret: d.Get("sharedsecret").(string),
		Type:         d.Get("type").(string),
		Url:          d.Get("url").(string),
	}

	_, err := client.AddResource(service.Autoscaleprofile.Type(), autoscaleprofileName, &autoscaleprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(autoscaleprofileName)

	return readAutoscaleprofileFunc(ctx, d, meta)
}

func readAutoscaleprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAutoscaleprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading autoscaleprofile state %s", autoscaleprofileName)
	data, err := client.FindResource(service.Autoscaleprofile.Type(), autoscaleprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing autoscaleprofile state %s", autoscaleprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("type", data["type"])
	d.Set("url", data["url"])

	return nil

}

func updateAutoscaleprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAutoscaleprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleprofileName := d.Get("name").(string)

	autoscaleprofile := autoscale.Autoscaleprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("apikey") {
		log.Printf("[DEBUG]  citrixadc-provider: Apikey has changed for autoscaleprofile %s, starting update", autoscaleprofileName)
		autoscaleprofile.Apikey = d.Get("apikey").(string)
		hasChange = true
	}
	if d.HasChange("sharedsecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Sharedsecret has changed for autoscaleprofile %s, starting update", autoscaleprofileName)
		autoscaleprofile.Sharedsecret = d.Get("sharedsecret").(string)
		hasChange = true
	}
	if d.HasChange("url") {
		log.Printf("[DEBUG]  citrixadc-provider: Url has changed for autoscaleprofile %s, starting update", autoscaleprofileName)
		autoscaleprofile.Url = d.Get("url").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Autoscaleprofile.Type(), &autoscaleprofile)
		if err != nil {
			return diag.Errorf("Error updating autoscaleprofile %s", autoscaleprofileName)
		}
	}
	return readAutoscaleprofileFunc(ctx, d, meta)
}

func deleteAutoscaleprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAutoscaleprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleprofileName := d.Id()
	err := client.DeleteResource(service.Autoscaleprofile.Type(), autoscaleprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
