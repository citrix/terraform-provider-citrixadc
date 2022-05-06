package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"net/url"
)

func resourceCitrixAdcLocation() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLocationFunc,
		Read:          readLocationFunc,
		Delete:        deleteLocationFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ipfrom": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ipto": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"preferredlocation": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"latitude": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"longitude": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLocationFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLocationFunc")
	client := meta.(*NetScalerNitroClient).client
	locationName := d.Get("ipfrom").(string)
	location := basic.Location{
		Ipfrom:            d.Get("ipfrom").(string),
		Ipto:              d.Get("ipto").(string),
		Latitude:          d.Get("latitude").(int),
		Longitude:         d.Get("longitude").(int),
		Preferredlocation: d.Get("preferredlocation").(string),
	}

	_, err := client.AddResource(service.Location.Type(), locationName, &location)
	if err != nil {
		return err
	}

	d.SetId(locationName)

	err = readLocationFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this location but we can't read it ?? %s", locationName)
		return nil
	}
	return nil
}

func readLocationFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLocationFunc")
	client := meta.(*NetScalerNitroClient).client
	locationName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading location state %s", locationName)
	data, err := client.FindResource(service.Location.Type(), locationName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing location state %s", locationName)
		d.SetId("")
		return nil
	}
	d.Set("ipfrom", data["ipfrom"])
	d.Set("ipto", data["ipto"])
	d.Set("latitude", data["latitude"])
	d.Set("longitude", data["longitude"])
	//d.Set("preferredlocation", data["preferredlocation"])

	return nil

}

func deleteLocationFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLocationFunc")
	client := meta.(*NetScalerNitroClient).client
	
	argsMap := make(map[string]string)
	// Only the ipfrom and ipto properties are required for deletion
	argsMap["ipfrom"] = url.QueryEscape(d.Get("ipfrom").(string))
	argsMap["ipto"] = url.QueryEscape(d.Get("ipto").(string))
	err := client.DeleteResourceWithArgsMap(service.Location.Type(), "",argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
