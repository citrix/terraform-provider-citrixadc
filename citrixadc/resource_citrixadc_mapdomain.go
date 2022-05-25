package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcMapdomain() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createMapdomainFunc,
		Read:          readMapdomainFunc,
		Delete:        deleteMapdomainFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"mapdmrname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createMapdomainFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createMapdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdomainName := d.Get("name").(string)
	mapdomain := network.Mapdomain{
		Mapdmrname: d.Get("mapdmrname").(string),
		Name:       d.Get("name").(string),
	}

	_, err := client.AddResource("mapdomain", mapdomainName, &mapdomain)
	if err != nil {
		return err
	}

	d.SetId(mapdomainName)

	err = readMapdomainFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this mapdomain but we can't read it ?? %s", mapdomainName)
		return nil
	}
	return nil
}

func readMapdomainFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readMapdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdomainName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading mapdomain state %s", mapdomainName)
	data, err := client.FindResource("mapdomain", mapdomainName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing mapdomain state %s", mapdomainName)
		d.SetId("")
		return nil
	}
	d.Set("mapdmrname", data["mapdmrname"])
	d.Set("name", data["name"])

	return nil

}

func deleteMapdomainFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteMapdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdomainName := d.Id()
	err := client.DeleteResource("mapdomain", mapdomainName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
