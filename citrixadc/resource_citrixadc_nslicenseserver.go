package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNslicenseserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNslicenseserverFunc,
		Read:          readNslicenseserverFunc,
		Delete:        deleteNslicenseserverFunc,
		Schema: map[string]*schema.Schema{
			"forceupdateip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createNslicenseserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNslicenseserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseserverId := d.Get("servername").(string)

	nslicenseserver := ns.Nslicenseserver{
		Forceupdateip: d.Get("forceupdateip").(bool),
		Nodeid:        uint32(d.Get("nodeid").(int)),
		Port:          uint32(d.Get("port").(int)),
		Servername:    d.Get("servername").(string),
	}

	_, err := client.AddResource("nslicenseserver", "", &nslicenseserver)
	if err != nil {
		return err
	}

	d.SetId(nslicenseserverId)

	err = readNslicenseserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nslicenseserver but we can't read it ?? %s", nslicenseserverId)
		return nil
	}
	return nil
}

func readNslicenseserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNslicenseserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseserverId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nslicenseserver state %s", nslicenseserverId)

	findParams := service.FindParams{
		ResourceType: "nslicenseserver",
	}

	licenseServers, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nslicenseserver state %s", nslicenseserverId)
		d.SetId("")
		return nil
	}
	if len(licenseServers) == 0 {
		// There is no license server configured
		d.SetId("")
	} else {
		// License server will return at most 1 element
		data := licenseServers[0]

		d.Set("forceupdateip", data["forceupdateip"])
		d.Set("nodeid", data["nodeid"])
		d.Set("port", data["port"])
		d.Set("servername", data["servername"])
	}

	return nil

}

func deleteNslicenseserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNslicenseserverFunc")
	client := meta.(*NetScalerNitroClient).client
	args := make([]string, 0, 1)
	args = append(args, fmt.Sprintf("servername:%s", d.Get("servername").(string)))
	err := client.DeleteResourceWithArgs("nslicenseserver", "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
