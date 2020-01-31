package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ns"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsacls() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsaclsFunc,
		Read:          readNsaclsFunc,
		Update:        updateNsaclsFunc,
		Delete:        deleteNsaclsFunc,
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsaclsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsaclsName string
	if v, ok := d.GetOk("name"); ok {
		nsaclsName = v.(string)
	} else {
		nsaclsName = resource.PrefixedUniqueId("tf-nsacls-")
		d.Set("name", nsaclsName)
	}
	nsacls := ns.Nsacls{
		Type: d.Get("type").(string),
	}

	_, err := client.AddResource(netscaler.Nsacls.Type(), nsaclsName, &nsacls)
	if err != nil {
		return err
	}

	d.SetId(nsaclsName)

	err = readNsaclsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsacls but we can't read it ?? %s", nsaclsName)
		return nil
	}
	return nil
}

func readNsaclsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclsName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsacls state %s", nsaclsName)
	data, err := client.FindResource(netscaler.Nsacls.Type(), nsaclsName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsacls state %s", nsaclsName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("type", data["type"])

	return nil

}

func updateNsaclsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclsName := d.Get("name").(string)

	nsacls := ns.Nsacls{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for nsacls %s, starting update", nsaclsName)
		nsacls.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Nsacls.Type(), nsaclsName, &nsacls)
		if err != nil {
			return fmt.Errorf("Error updating nsacls %s", nsaclsName)
		}
	}
	return readNsaclsFunc(d, meta)
}

func deleteNsaclsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsaclsFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclsName := d.Id()
	err := client.DeleteResource(netscaler.Nsacls.Type(), nsaclsName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
