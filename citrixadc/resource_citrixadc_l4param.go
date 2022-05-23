package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcL4param() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createL4paramFunc,
		Read:          readL4paramFunc,
		Update:        updateL4paramFunc,
		Delete:        deleteL4paramFunc,
		Schema: map[string]*schema.Schema{
			"l2connmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l4switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createL4paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createL4paramFunc")
	client := meta.(*NetScalerNitroClient).client
	var l4paramName string
	// there is no primary key in l4param resource. Hence generate one for terraform state maintenance
	l4paramName = resource.PrefixedUniqueId("tf-l4param-")
	l4param := network.L4param{
		L2connmethod: d.Get("l2connmethod").(string),
		L4switch:     d.Get("l4switch").(string),
	}

	err := client.UpdateUnnamedResource(service.L4param.Type(), &l4param)
	if err != nil {
		return err
	}

	d.SetId(l4paramName)

	err = readL4paramFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this l4param but we can't read it ??")
		return nil
	}
	return nil
}

func readL4paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readL4paramFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading l4param state")
	data, err := client.FindResource(service.L4param.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing l4param state")
		d.SetId("")
		return nil
	}
	d.Set("l2connmethod", data["l2connmethod"])
	d.Set("l4switch", data["l4switch"])

	return nil

}

func updateL4paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateL4paramFunc")
	client := meta.(*NetScalerNitroClient).client

	l4param := network.L4param{}
	hasChange := false
	if d.HasChange("l2connmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: L2connmethod has changed for l4param, starting update")
		l4param.L2connmethod = d.Get("l2connmethod").(string)
		hasChange = true
	}
	if d.HasChange("l4switch") {
		log.Printf("[DEBUG]  citrixadc-provider: L4switch has changed for l4param, starting update")
		l4param.L4switch = d.Get("l4switch").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.L4param.Type(), &l4param)
		if err != nil {
			return fmt.Errorf("Error updating l4param")
		}
	}
	return readL4paramFunc(d, meta)
}

func deleteL4paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteL4paramFunc")

	d.SetId("")

	return nil
}
