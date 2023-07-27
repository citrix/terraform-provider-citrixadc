package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcLocationfile6Import() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLocationfile6ImportFunc,
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

func createLocationfile6ImportFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	locationfileName := resource.PrefixedUniqueId("tf-locationfile6-")
	locationfile := basic.Locationfile6{
		Src:          d.Get("src").(string),
		Locationfile: d.Get("locationfile").(string),
	}

	err := client.ActOnResource(service.Locationfile6.Type(), locationfile, "import")
	if err != nil {
		return err
	}

	d.SetId(locationfileName)
	return nil
}
