package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSnmpmib() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpmibFunc,
		Read:          readSnmpmibFunc,
		Update:        updateSnmpmibFunc,
		Delete:        deleteSnmpmibFunc, // Thought snmpmib resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"contact": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSnmpmibFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpmibFunc")
	client := meta.(*NetScalerNitroClient).client

	// there is no primary key in snmpmib resource. Hence generate one for terraform state maintenance
	snmpmibName := resource.PrefixedUniqueId("tf-snmpmib-")

	snmpmib := snmp.Snmpmib{
		Contact:   d.Get("contact").(string),
		Customid:  d.Get("customid").(string),
		Location:  d.Get("location").(string),
		Name:      d.Get("name").(string),
		Ownernode: d.Get("ownernode").(int),
	}

	err := client.UpdateUnnamedResource(service.Snmpmib.Type(), &snmpmib)
	if err != nil {
		return err
	}

	d.SetId(snmpmibName)

	err = readSnmpmibFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpmib but we can't read it ??")
		return nil
	}
	return nil
}

func readSnmpmibFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpmibFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpmib state")
	data, err := client.FindResource(service.Snmpmib.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpmib state")
		d.SetId("")
		return nil
	}
	d.Set("contact", data["contact"])
	d.Set("customid", data["customid"])
	d.Set("location", data["location"])
	d.Set("name", data["name"])
	d.Set("ownernode", data["ownernode"])

	return nil

}

func updateSnmpmibFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpmibFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmib := snmp.Snmpmib{}

	hasChange := false
	if d.HasChange("contact") {
		log.Printf("[DEBUG]  citrixadc-provider: Contact has changed for snmpmib, starting update")
		snmpmib.Contact = d.Get("contact").(string)
		hasChange = true
	}
	if d.HasChange("customid") {
		log.Printf("[DEBUG]  citrixadc-provider: Customid has changed for snmpmib, starting update")
		snmpmib.Customid = d.Get("customid").(string)
		hasChange = true
	}
	if d.HasChange("location") {
		log.Printf("[DEBUG]  citrixadc-provider: Location has changed for snmpmib, starting update")
		snmpmib.Location = d.Get("location").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for snmpmib, starting update")
		snmpmib.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for snmpmib, starting update")
		snmpmib.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpmib.Type(), &snmpmib)
		if err != nil {
			return fmt.Errorf("Error updating snmpmib")
		}
	}
	return readSnmpmibFunc(d, meta)
}

func deleteSnmpmibFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpmibFunc")
	// snmpmib do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
