package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strconv"
	"log"
)

func resourceCitrixAdcLacp() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLacpFunc,
		Read:          readLacpFunc,
		Update:        updateLacpFunc,
		Delete:        deleteLacpFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"syspriority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ownernode": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func createLacpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLacpFunc")
	client := meta.(*NetScalerNitroClient).client
	lacpId := strconv.Itoa(d.Get("ownernode").(int))
	lacp := network.Lacp{
		Ownernode:   d.Get("ownernode").(int),
		Syspriority: d.Get("syspriority").(int),
	}

	err := client.UpdateUnnamedResource(service.Lacp.Type(), &lacp)
	if err != nil {
		return err
	}

	d.SetId(lacpId)

	err = readLacpFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lacp but we can't read it %s??", lacpId)
		return nil
	}
	return nil
}

func readLacpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLacpFunc")
	client := meta.(*NetScalerNitroClient).client
	lacpId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lacp state")
	data, err := client.FindResource(service.Lacp.Type(), strconv.Itoa(d.Get("ownernode").(int)))
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lacp state %s", lacpId)
		d.SetId("")
		return nil
	}
	
	d.Set("ownernode", data["ownernode"])
	d.Set("syspriority", data["syspriority"])

	return nil

}

func updateLacpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLacpFunc")
	client := meta.(*NetScalerNitroClient).client
	lacpId := d.Id()

	lacp := network.Lacp{
		Ownernode: d.Get("ownernode").(int),
	}
	hasChange := false

	if d.HasChange("syspriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Syspriority has changed for lacp, starting update %s", lacpId)
		lacp.Syspriority = d.Get("syspriority").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Lacp.Type(), &lacp)
		if err != nil {
			return fmt.Errorf("Error updating lacp %s", lacpId)
		}
	}
	return readLacpFunc(d, meta)
}

func deleteLacpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLacpFunc")
	//lacp does not support delete operation
	d.SetId("")

	return nil
}
