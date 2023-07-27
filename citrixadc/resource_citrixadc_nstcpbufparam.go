package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNstcpbufparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNstcpbufparamFunc,
		Read:          readNstcpbufparamFunc,
		Update:        updateNstcpbufparamFunc,
		Delete:        deleteNstcpbufparamFunc,
		Schema: map[string]*schema.Schema{
			"memlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNstcpbufparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstcpbufparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var nstcpbufparamName string
	// there is no primary key in nstcpbufparam resource. Hence generate one for terraform state maintenance
	nstcpbufparamName = resource.PrefixedUniqueId("tf-nstcpbufparam-")
	nstcpbufparam := ns.Nstcpbufparam{
		Memlimit: d.Get("memlimit").(int),
		Size:     d.Get("size").(int),
	}

	err := client.UpdateUnnamedResource(service.Nstcpbufparam.Type(), &nstcpbufparam)
	if err != nil {
		return err
	}

	d.SetId(nstcpbufparamName)

	err = readNstcpbufparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nstcpbufparam but we can't read it ?? %s", nstcpbufparamName)
		return nil
	}
	return nil
}

func readNstcpbufparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstcpbufparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nstcpbufparam state")
	data, err := client.FindResource(service.Nstcpbufparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nstcpbufparam state")
		d.SetId("")
		return nil
	}
	d.Set("memlimit", data["memlimit"])
	d.Set("size", data["size"])

	return nil

}

func updateNstcpbufparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNstcpbufparamFunc")
	client := meta.(*NetScalerNitroClient).client

	nstcpbufparam := ns.Nstcpbufparam{
		Memlimit: d.Get("memlimit").(int),
	}
	hasChange := false
	if d.HasChange("memlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Memlimit has changed for nstcpbufparam, starting update")
		hasChange = true
	}
	if d.HasChange("size") {
		log.Printf("[DEBUG]  citrixadc-provider: Size has changed for nstcpbufparam, starting update")
		nstcpbufparam.Size = d.Get("size").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nstcpbufparam.Type(), &nstcpbufparam)
		if err != nil {
			return fmt.Errorf("Error updating nstcpbufparam")
		}
	}
	return readNstcpbufparamFunc(d, meta)
}

func deleteNstcpbufparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstcpbufparamFunc")

	d.SetId("")

	return nil
}
