package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"
	
	"log"
)

func resourceCitrixAdcLocationfile6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLocationfile6Func,
		Read:          readLocationfile6Func,
		Delete:        deleteLocationfile6Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLocationfile6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	locationfileName := d.Get("locationfile").(string)
	locationfile := basic.Locationfile6{
		Format:       d.Get("format").(string),
		Locationfile: d.Get("locationfile").(string),
	}

	_, err := client.AddResource(service.Locationfile6.Type(), "", &locationfile)
	if err != nil {
		return err
	}

	d.SetId(locationfileName)

	err = readLocationfile6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this locationfile but we can't read it ?? %s", locationfileName)
		return nil
	}
	return nil
}

func readLocationfile6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	locationfileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading locationfile state %s", locationfileName)
	data, err := client.FindResource(service.Locationfile6.Type(), "")
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

func deleteLocationfile6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLocationfileFunc")
	client := meta.(*NetScalerNitroClient).client
	err := client.DeleteResource(service.Locationfile6.Type(), "")
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
