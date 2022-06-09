package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRsskeytype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRsskeytypeFunc,
		Read:          readRsskeytypeFunc,
		Update:        updateRsskeytypeFunc,
		Delete:        deleteRsskeytypeFunc,
		Schema: map[string]*schema.Schema{
			"rsstype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createRsskeytypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRsskeytypeFunc")
	client := meta.(*NetScalerNitroClient).client
	var rsskeytypeName string
	// there is no primary key in rsskeytype resource. Hence generate one for terraform state maintenance
	rsskeytypeName = resource.PrefixedUniqueId("tf-rsskeytype-")

	rsskeytype := network.Rsskeytype{
		Rsstype: d.Get("rsstype").(string),
	}

	err := client.UpdateUnnamedResource(service.Rsskeytype.Type(), &rsskeytype)
	if err != nil {
		return err
	}

	d.SetId(rsskeytypeName)

	err = readRsskeytypeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rsskeytype but we can't read it ?? %s", rsskeytypeName)
		return nil
	}
	return nil
}

func readRsskeytypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRsskeytypeFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading rsskeytype state")
	data, err := client.FindResource(service.Rsskeytype.Type(),"")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rsskeytype state")
		d.SetId("")
		return nil
	}
	d.Set("rsstype", data["rsstype"])

	return nil

}

func updateRsskeytypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRsskeytypeFunc")
	client := meta.(*NetScalerNitroClient).client

	rsskeytype := network.Rsskeytype{}
	hasChange := false
	if d.HasChange("rsstype") {
		log.Printf("[DEBUG]  citrixadc-provider: Rsstype has changed for rsskeytype, starting update")
		rsskeytype.Rsstype = d.Get("rsstype").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Rsskeytype.Type(), &rsskeytype)
		if err != nil {
			return fmt.Errorf("Error updating rsskeytype")
		}
	}
	return readRsskeytypeFunc(d, meta)
}

func deleteRsskeytypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRsskeytypeFunc")
	

	d.SetId("")

	return nil
}
