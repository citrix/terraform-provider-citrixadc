package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcExtendedmemoryparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createExtendedmemoryparamFunc,
		Read:          readExtendedmemoryparamFunc,
		Update:        updateExtendedmemoryparamFunc,
		Delete:        deleteExtendedmemoryparamFunc,
		Schema: map[string]*schema.Schema{
			"memlimit": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createExtendedmemoryparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createExtendedmemoryparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var extendedmemoryparamName string
	extendedmemoryparamName = resource.PrefixedUniqueId("tf-extendedmemoryparam-")
	extendedmemoryparam := basic.Extendedmemoryparam{
		Memlimit: d.Get("memlimit").(int),
	}

	err := client.UpdateUnnamedResource(service.Extendedmemoryparam.Type(), &extendedmemoryparam)
	if err != nil {
		return fmt.Errorf("Error updating extendedmemoryparam")
	}

	d.SetId(extendedmemoryparamName)

	err = readExtendedmemoryparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this extendedmemoryparam but we can't read it ??")
		return nil
	}
	return nil
}

func readExtendedmemoryparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readExtendedmemoryparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading extendedmemoryparam state")
	data, err := client.FindResource(service.Extendedmemoryparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing extendedmemoryparam state")
		d.SetId("")
		return nil
	}
	d.Set("memlimit", data["memlimit"])

	return nil

}

func updateExtendedmemoryparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateExtendedmemoryparamFunc")
	client := meta.(*NetScalerNitroClient).client

	extendedmemoryparam := basic.Extendedmemoryparam{}
	hasChange := false
	if d.HasChange("memlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Memlimit has changed for extendedmemoryparam , starting update")
		extendedmemoryparam.Memlimit = d.Get("memlimit").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Extendedmemoryparam.Type(), &extendedmemoryparam)
		if err != nil {
			return fmt.Errorf("Error updating extendedmemoryparam %s", err.Error())
		}
	}
	return readExtendedmemoryparamFunc(d, meta)
}

func deleteExtendedmemoryparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteExtendedmemoryparamFunc")

	d.SetId("")

	return nil
}
