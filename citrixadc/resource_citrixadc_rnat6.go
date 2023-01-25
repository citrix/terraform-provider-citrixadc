package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRnat6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRnat6Func,
		Read:          readRnat6Func,
		Update:        updateRnat6Func,
		Delete:        deleteRnat6Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"acl6name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ownergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcippersistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createRnat6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRnat6Func")
	client := meta.(*NetScalerNitroClient).client
	rnat6Name := d.Get("name").(string)
	rnat6 := network.Rnat6{
		Acl6name:         d.Get("acl6name").(string),
		Name:             d.Get("name").(string),
		Network:          d.Get("network").(string),
		Ownergroup:       d.Get("ownergroup").(string),
		Redirectport:     d.Get("redirectport").(int),
		Srcippersistency: d.Get("srcippersistency").(string),
		Td:               d.Get("td").(int),
	}

	_, err := client.AddResource(service.Rnat6.Type(), rnat6Name, &rnat6)
	if err != nil {
		return err
	}

	d.SetId(rnat6Name)

	err = readRnat6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rnat6 but we can't read it ?? %s", rnat6Name)
		return nil
	}
	return nil
}

func readRnat6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRnat6Func")
	client := meta.(*NetScalerNitroClient).client
	rnat6Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rnat6 state %s", rnat6Name)
	data, err := client.FindResource(service.Rnat6.Type(), rnat6Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rnat6 state %s", rnat6Name)
		d.SetId("")
		return nil
	}
	d.Set("acl6name", data["acl6name"])
	d.Set("name", data["name"])
	d.Set("network", data["network"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("redirectport", data["redirectport"])
	d.Set("srcippersistency", data["srcippersistency"])
	d.Set("td", data["td"])

	return nil

}

func updateRnat6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRnat6Func")
	client := meta.(*NetScalerNitroClient).client
	rnat6Name := d.Get("name").(string)

	rnat6 := make(map[string]interface{})
	rnat6["name"] = d.Get("name").(string)
	hasChange := false
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for rnat6 %s, starting update", rnat6Name)
		rnat6["ownergroup"] = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("redirectport") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectport has changed for rnat6 %s, starting update", rnat6Name)
		rnat6["redirectport"] = d.Get("redirectport").(int)
		hasChange = true
	}
	if d.HasChange("srcippersistency") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcippersistency has changed for rnat6 %s, starting update", rnat6Name)
		rnat6["srcippersistency"] = d.Get("srcippersistency").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Rnat6.Type(), &rnat6)
		if err != nil {
			return fmt.Errorf("Error updating rnat6 %s", rnat6Name)
		}
	}
	return readRnat6Func(d, meta)
}

func deleteRnat6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRnat6Func")
	// rnat6 does not support DELETE operation
	d.SetId("")

	return nil
}
