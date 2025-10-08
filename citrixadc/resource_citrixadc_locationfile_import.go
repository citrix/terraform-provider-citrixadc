package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLocationImportfile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLocationfileImportFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
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
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLocationfileImportFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	locationfileName := resource.PrefixedUniqueId("tf-locationfile-")
	locationfile := basic.Locationfile{
		Src:          d.Get("src").(string),
		Locationfile: d.Get("locationfile").(string),
	}

	err := client.ActOnResource(service.Locationfile.Type(), locationfile, "import")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(locationfileName)
	return nil
}
