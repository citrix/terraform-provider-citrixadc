package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcLocationfile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLocationfileFunc,
		ReadContext:   readLocationfileFunc,
		DeleteContext: deleteLocationfileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"locationfile": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"src": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLocationfileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	locationfileName := d.Get("locationfile").(string)
	locationfile := basic.Locationfile{
		Format:       d.Get("format").(string),
		Locationfile: d.Get("locationfile").(string),
	}

	_, err := client.AddResource(service.Locationfile.Type(), "", &locationfile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(locationfileName)

	return readLocationfileFunc(ctx, d, meta)
}

func readLocationfileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	locationfileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading locationfile state %s", locationfileName)
	data, err := client.FindResource(service.Locationfile.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing locationfile state %s", locationfileName)
		d.SetId("")
		return nil
	}
	d.Set("format", data["format"])
	d.Set("locationfile", data["Locationfile"])
	d.Set("src", data["src"])

	return nil

}

func deleteLocationfileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	err := client.DeleteResource(service.Locationfile.Type(), "")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
