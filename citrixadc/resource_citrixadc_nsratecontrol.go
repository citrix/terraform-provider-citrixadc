package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcNsratecontrol() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsratecontrolFunc,
		Read:          readNsratecontrolFunc,
		Update:        updateNsratecontrolFunc,
		Delete:        deleteNsratecontrolFunc,
		Schema: map[string]*schema.Schema{
			"icmpthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcprstthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"udpthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsratecontrolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsratecontrolFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsratecontrolName string
	nsratecontrolName = resource.PrefixedUniqueId("tf-nsratecontrol-")
	nsratecontrol := ns.Nsratecontrol{
		Icmpthreshold:   d.Get("icmpthreshold").(int),
		Tcprstthreshold: d.Get("tcprstthreshold").(int),
		Tcpthreshold:    d.Get("tcpthreshold").(int),
		Udpthreshold:    d.Get("udpthreshold").(int),
	}

	err := client.UpdateUnnamedResource(service.Nsratecontrol.Type(), &nsratecontrol)
	if err != nil {
		return err
	}

	d.SetId(nsratecontrolName)

	err = readNsratecontrolFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsratecontrol but we can't read it ??")
		return nil
	}
	return nil
}

func readNsratecontrolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsratecontrolFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsratecontrol state")
	data, err := client.FindResource(service.Nsratecontrol.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsratecontrol state")
		d.SetId("")
		return nil
	}
	value, _ := strconv.Atoi(data["icmpthreshold"].(string))
	d.Set("icmpthreshold", value)
	value, _ = strconv.Atoi(data["tcprstthreshold"].(string))
	d.Set("tcprstthreshold", value)
	value, _ = strconv.Atoi(data["tcpthreshold"].(string))
	d.Set("tcpthreshold", value)
	value, _ = strconv.Atoi(data["udpthreshold"].(string))
	d.Set("udpthreshold", value)

	return nil

}

func updateNsratecontrolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsratecontrolFunc")
	client := meta.(*NetScalerNitroClient).client

	nsratecontrol := ns.Nsratecontrol{
		Icmpthreshold:   d.Get("icmpthreshold").(int),
		Tcprstthreshold: d.Get("tcprstthreshold").(int),
		Tcpthreshold:    d.Get("tcpthreshold").(int),
		Udpthreshold:    d.Get("udpthreshold").(int),
	}

	err := client.UpdateUnnamedResource(service.Nsratecontrol.Type(), &nsratecontrol)
	if err != nil {
		return fmt.Errorf("Error updating nsratecontrol")
	}

	return readNsratecontrolFunc(d, meta)
}

func deleteNsratecontrolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsratecontrolFunc")
	// nsratecontrol do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
