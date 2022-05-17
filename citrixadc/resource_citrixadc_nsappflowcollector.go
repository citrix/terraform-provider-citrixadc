package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcNsappflowcollector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsappflowcollectorFunc,
		Read:          readNsappflowcollectorFunc,
		Delete:        deleteNsappflowcollectorFunc,
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
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNsappflowcollectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsappflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	nsappflowcollectorName := d.Get("name").(string)
	nsappflowcollector := ns.Nsappflowcollector{
		Ipaddress: d.Get("ipaddress").(string),
		Name:      d.Get("name").(string),
		Port:      d.Get("port").(int),
	}

	_, err := client.AddResource(service.Nsappflowcollector.Type(), nsappflowcollectorName, &nsappflowcollector)
	if err != nil {
		return err
	}

	d.SetId(nsappflowcollectorName)

	err = readNsappflowcollectorFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsappflowcollector but we can't read it ?? %s", nsappflowcollectorName)
		return nil
	}
	return nil
}

func readNsappflowcollectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsappflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	nsappflowcollectorName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsappflowcollector state %s", nsappflowcollectorName)
	data, err := client.FindResource(service.Nsappflowcollector.Type(), nsappflowcollectorName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsappflowcollector state %s", nsappflowcollectorName)
		d.SetId("")
		return nil
	}
	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])
	d.Set("port", data["port"])

	return nil

}

func deleteNsappflowcollectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsappflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	nsappflowcollectorName := d.Id()
	err := client.DeleteResource(service.Nsappflowcollector.Type(), nsappflowcollectorName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
