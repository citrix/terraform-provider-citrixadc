package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNshostname() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNshostnameFunc,
		Read:          readNshostnameFunc,
		Update:        updateNshostnameFunc,
		Delete:        deleteNshostnameFunc,
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNshostnameFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNshostnameFunc")
	client := meta.(*NetScalerNitroClient).client
	nshostnameName := resource.PrefixedUniqueId("tf-nshostname-")
	nshostname := ns.Nshostname{
		Hostname:  d.Get("hostname").(string),
		Ownernode: d.Get("ownernode").(int),
	}

	err := client.UpdateUnnamedResource(service.Nshostname.Type(), &nshostname)
	if err != nil {
		return err
	}

	d.SetId(nshostnameName)

	err = readNshostnameFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nshostname but we can't read it ??")
		return nil
	}
	return nil
}

func readNshostnameFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNshostnameFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nshostname state")
	data, err := client.FindResource(service.Nshostname.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nshostname state")
		d.SetId("")
		return nil
	}
	d.Set("hostname", data["hostname"])
	d.Set("ownernode", data["ownernode"])

	return nil

}

func updateNshostnameFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNshostnameFunc")
	client := meta.(*NetScalerNitroClient).client

	nshostname := ns.Nshostname{
		Hostname: d.Get("hostname").(string),
	}

	hasChange := false
	if d.HasChange("hostname") {
		log.Printf("[DEBUG]  citrixadc-provider: Hostname has changed for nshostname, starting update")
		nshostname.Hostname = d.Get("hostname").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for nshostname, starting update")
		nshostname.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nshostname.Type(), &nshostname)
		if err != nil {
			return fmt.Errorf("Error updating nshostname")
		}
	}
	return readNshostnameFunc(d, meta)
}

func deleteNshostnameFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNshostnameFunc")
	//nshostname does not support delete operation
	d.SetId("")

	return nil
}
