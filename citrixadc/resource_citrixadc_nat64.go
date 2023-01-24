package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNat64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNat64Func,
		Read:          readNat64Func,
		Update:        updateNat64Func,
		Delete:        deleteNat64Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"acl6name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNat64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNat64Func")
	client := meta.(*NetScalerNitroClient).client
	nat64Name := d.Get("name").(string)
	nat64 := network.Nat64{
		Acl6name:   d.Get("acl6name").(string),
		Name:       d.Get("name").(string),
		Netprofile: d.Get("netprofile").(string),
	}

	_, err := client.AddResource(service.Nat64.Type(), nat64Name, &nat64)
	if err != nil {
		return err
	}

	d.SetId(nat64Name)

	err = readNat64Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nat64 but we can't read it ?? %s", nat64Name)
		return nil
	}
	return nil
}

func readNat64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNat64Func")
	client := meta.(*NetScalerNitroClient).client
	nat64Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nat64 state %s", nat64Name)
	data, err := client.FindResource(service.Nat64.Type(), nat64Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nat64 state %s", nat64Name)
		d.SetId("")
		return nil
	}
	d.Set("acl6name", data["acl6name"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])

	return nil

}

func updateNat64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNat64Func")
	client := meta.(*NetScalerNitroClient).client
	nat64Name := d.Get("name").(string)

	nat64 := network.Nat64{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("acl6name") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl6name has changed for nat64 %s, starting update", nat64Name)
		nat64.Acl6name = d.Get("acl6name").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for nat64 %s, starting update", nat64Name)
		nat64.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nat64.Type(), nat64Name, &nat64)
		if err != nil {
			return fmt.Errorf("Error updating nat64 %s", nat64Name)
		}
	}
	return readNat64Func(d, meta)
}

func deleteNat64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNat64Func")
	client := meta.(*NetScalerNitroClient).client
	nat64Name := d.Id()
	err := client.DeleteResource(service.Nat64.Type(), nat64Name)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
