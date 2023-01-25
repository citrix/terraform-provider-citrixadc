package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	
	"log"
)

func resourceCitrixAdcLocationImportfile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLocationfileImportFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"locationfile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLocationfileImportFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	locationfileName := resource.PrefixedUniqueId("tf-locationfile-")
	locationfile := basic.Locationfile{
		Src:       d.Get("src").(string),
		Locationfile: d.Get("locationfile").(string),
	}

	err := client.ActOnResource(service.Locationfile.Type(), locationfile, "import")
	if err != nil {
		return err
	}

	d.SetId(locationfileName)
	return nil
}