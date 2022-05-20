package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsdhcpparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsdhcpparamsFunc,
		Read:          readNsdhcpparamsFunc,
		Update:        updateNsdhcpparamsFunc,
		Delete:        deleteNsdhcpparamsFunc,
		Schema: map[string]*schema.Schema{
			"dhcpclient": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saveroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsdhcpparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsdhcpparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsdhcpparamsName string
	// there is no primary key in nsdhcpparams resource. Hence generate one for terraform state maintenance
	nsdhcpparamsName = resource.PrefixedUniqueId("tf-nsdhcpparams-")
	nsdhcpparams := ns.Nsdhcpparams{
		Dhcpclient: d.Get("dhcpclient").(string),
		Saveroute:  d.Get("saveroute").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsdhcpparams.Type(), &nsdhcpparams)
	if err != nil {
		return err
	}

	d.SetId(nsdhcpparamsName)

	err = readNsdhcpparamsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsdhcpparams but we can't read it ??")
		return nil
	}
	return nil
}

func readNsdhcpparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsdhcpparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsdhcpparams state")
	data, err := client.FindResource(service.Nsdhcpparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsdhcpparams state")
		d.SetId("")
		return nil
	}
	d.Set("dhcpclient", data["dhcpclient"])
	d.Set("saveroute", data["saveroute"])

	return nil

}

func updateNsdhcpparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsdhcpparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	nsdhcpparams := ns.Nsdhcpparams{}
	hasChange := false
	if d.HasChange("dhcpclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhcpclient has changed for nsdhcpparams , starting update")
		nsdhcpparams.Dhcpclient = d.Get("dhcpclient").(string)
		hasChange = true
	}
	if d.HasChange("saveroute") {
		log.Printf("[DEBUG]  citrixadc-provider: Saveroute has changed for nsdhcpparams , starting update")
		nsdhcpparams.Saveroute = d.Get("saveroute").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsdhcpparams.Type(), &nsdhcpparams)
		if err != nil {
			return fmt.Errorf("Error updating nsdhcpparams")
		}
	}
	return readNsdhcpparamsFunc(d, meta)
}

func deleteNsdhcpparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsdhcpparamsFunc")

	d.SetId("")

	return nil
}
