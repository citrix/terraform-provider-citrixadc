package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsdiameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsdiameterFunc,
		Read:          readNsdiameterFunc,
		Update:        updateNsdiameterFunc,
		Delete:        deleteNsdiameterFunc,
		Schema: map[string]*schema.Schema{
			"identity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"realm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverclosepropagation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsdiameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsdiameterFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsdiameterName string
	// there is no primary key in nsdiameter resource. Hence generate one for terraform state maintenance
	nsdiameterName = resource.PrefixedUniqueId("tf-nsdiameter-")
	nsdiameter := ns.Nsdiameter{
		Identity:               d.Get("identity").(string),
		Ownernode:              d.Get("ownernode").(int),
		Realm:                  d.Get("realm").(string),
		Serverclosepropagation: d.Get("serverclosepropagation").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsdiameter.Type(), &nsdiameter)
	if err != nil {
		return err
	}

	d.SetId(nsdiameterName)

	err = readNsdiameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsdiameter but we can't read it ??")
		return nil
	}
	return nil
}

func readNsdiameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsdiameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsdiameter state")
	data, err := client.FindResource(service.Nsdiameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsdiameter state")
		d.SetId("")
		return nil
	}
	d.Set("identity", data["identity"])
	d.Set("ownernode", data["ownernode"])
	d.Set("realm", data["realm"])
	// d.Set("serverclosepropagation", data["serverclosepropagation"])

	return nil

}

func updateNsdiameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsdiameterFunc")
	client := meta.(*NetScalerNitroClient).client

	nsdiameter := ns.Nsdiameter{}

	hasChange := false
	if d.HasChange("identity") {
		log.Printf("[DEBUG]  citrixadc-provider: Identity has changed for nsdiameter ,starting update")
		nsdiameter.Identity = d.Get("identity").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for nsdiameter ,starting update")
		nsdiameter.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}
	if d.HasChange("realm") {
		log.Printf("[DEBUG]  citrixadc-provider: Realm has changed for nsdiameter ,starting update")
		nsdiameter.Realm = d.Get("realm").(string)
		hasChange = true
	}
	if d.HasChange("serverclosepropagation") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverclosepropagation has changed for nsdiameter ,starting update")
		nsdiameter.Serverclosepropagation = d.Get("serverclosepropagation").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsdiameter.Type(), &nsdiameter)
		if err != nil {
			return fmt.Errorf("Error updating nsdiameter")
		}
	}
	return readNsdiameterFunc(d, meta)
}

func deleteNsdiameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsdiameterFunc")
	// nsdiameter do not have DELETE operation, but this function is required to set the ID to ""

	d.SetId("")

	return nil
}
