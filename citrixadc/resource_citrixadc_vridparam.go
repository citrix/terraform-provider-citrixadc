package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVridparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVridparamFunc,
		Read:          readVridparamFunc,
		Update:        updateVridparamFunc,
		Delete:        deleteVridparamFunc,
		Schema: map[string]*schema.Schema{
			"deadinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hellointerval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sendtomaster": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVridparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVridparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var vridparamName string
	// there is no primary key in vridparam resource. Hence generate one for terraform state maintenance
	vridparamName = resource.PrefixedUniqueId("tf-vridparam-")
	vridparam := network.Vridparam{
		Deadinterval:  d.Get("deadinterval").(int),
		Hellointerval: d.Get("hellointerval").(int),
		Sendtomaster:  d.Get("sendtomaster").(string),
	}

	err := client.UpdateUnnamedResource(service.Vridparam.Type(), &vridparam)
	if err != nil {
		return err
	}

	d.SetId(vridparamName)

	err = readVridparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vridparam but we can't read it ?? %s", vridparamName)
		return nil
	}
	return nil
}

func readVridparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVridparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading vridparam state")
	data, err := client.FindResource(service.Vridparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vridparam state ")
		d.SetId("")
		return nil
	}
	d.Set("deadinterval", data["deadinterval"])
	d.Set("hellointerval", data["hellointerval"])
	d.Set("sendtomaster", data["sendtomaster"])

	return nil

}

func updateVridparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVridparamFunc")
	client := meta.(*NetScalerNitroClient).client
	vridparam := network.Vridparam{}
	hasChange := false
	if d.HasChange("deadinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Deadinterval has changed for vridparam, starting update")
		vridparam.Deadinterval = d.Get("deadinterval").(int)
		hasChange = true
	}
	if d.HasChange("hellointerval") {
		log.Printf("[DEBUG]  citrixadc-provider: Hellointerval has changed for vridparam, starting update")
		vridparam.Hellointerval = d.Get("hellointerval").(int)
		hasChange = true
	}
	if d.HasChange("sendtomaster") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendtomaster has changed for vridparam, starting update")
		vridparam.Sendtomaster = d.Get("sendtomaster").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Vridparam.Type(), &vridparam)
		if err != nil {
			return fmt.Errorf("Error updating vridparam")
		}
	}
	return readVridparamFunc(d, meta)
}

func deleteVridparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVridparamFunc")

	d.SetId("")

	return nil
}
