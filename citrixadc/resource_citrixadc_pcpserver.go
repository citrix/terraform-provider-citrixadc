package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/pcp"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcPcpserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPcpserverFunc,
		Read:          readPcpserverFunc,
		Update:        updatePcpserverFunc,
		Delete:        deletePcpserverFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pcpprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createPcpserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPcpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpserverName := d.Get("name").(string)
	pcpserver := pcp.Pcpserver{
		Ipaddress:  d.Get("ipaddress").(string),
		Name:       d.Get("name").(string),
		Pcpprofile: d.Get("pcpprofile").(string),
		Port:       d.Get("port").(int),
	}

	_, err := client.AddResource("pcpserver", pcpserverName, &pcpserver)
	if err != nil {
		return err
	}

	d.SetId(pcpserverName)

	err = readPcpserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this pcpserver but we can't read it ?? %s", pcpserverName)
		return nil
	}
	return nil
}

func readPcpserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPcpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading pcpserver state %s", pcpserverName)
	data, err := client.FindResource("pcpserver", pcpserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing pcpserver state %s", pcpserverName)
		d.SetId("")
		return nil
	}
	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])
	d.Set("pcpprofile", data["pcpprofile"])
	d.Set("port", data["port"])

	return nil

}

func updatePcpserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePcpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpserverName := d.Get("name").(string)

	pcpserver := pcp.Pcpserver{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("pcpprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Pcpprofile has changed for pcpserver %s, starting update", pcpserverName)
		pcpserver.Pcpprofile = d.Get("pcpprofile").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for pcpserver %s, starting update", pcpserverName)
		pcpserver.Port = d.Get("port").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("pcpserver", &pcpserver)
		if err != nil {
			return fmt.Errorf("Error updating pcpserver %s", pcpserverName)
		}
	}
	return readPcpserverFunc(d, meta)
}

func deletePcpserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePcpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpserverName := d.Id()
	err := client.DeleteResource("pcpserver", pcpserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
