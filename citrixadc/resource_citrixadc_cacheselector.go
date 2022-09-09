package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCacheselector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCacheselectorFunc,
		Read:          readCacheselectorFunc,
		Update:        updateCacheselectorFunc,
		Delete:        deleteCacheselectorFunc,
		Schema: map[string]*schema.Schema{
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				Computed: false,
			},
			"selectorname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createCacheselectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCacheselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	var cacheselectorName string
	cacheselectorName = d.Get("selectorname").(string)
	cacheselector := cache.Cacheselector{
		Rule:         toStringList(d.Get("rule").([]interface{})),
		Selectorname: d.Get("selectorname").(string),
	}

	_, err := client.AddResource(service.Cacheselector.Type(), cacheselectorName, &cacheselector)
	if err != nil {
		return err
	}

	d.SetId(cacheselectorName)

	err = readCacheselectorFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cacheselector but we can't read it ?? %s", cacheselectorName)
		return nil
	}
	return nil
}

func readCacheselectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCacheselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheselectorName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cacheselector state %s", cacheselectorName)
	data, err := client.FindResource(service.Cacheselector.Type(), cacheselectorName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cacheselector state %s", cacheselectorName)
		d.SetId("")
		return nil
	}
	d.Set("selectorname", data["selectorname"])
	d.Set("rule", data["rule"])

	return nil

}

func updateCacheselectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCacheselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheselectorName := d.Get("selectorname").(string)

	cacheselector := cache.Cacheselector{
		Selectorname: d.Get("selectorname").(string),
	}
	hasChange := false
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for cacheselector %s, starting update", cacheselectorName)
		cacheselector.Rule = toStringList(d.Get("rule").([]interface{}))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Cacheselector.Type(), cacheselectorName, &cacheselector)
		if err != nil {
			return fmt.Errorf("Error updating cacheselector %s", cacheselectorName)
		}
	}
	return readCacheselectorFunc(d, meta)
}

func deleteCacheselectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCacheselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheselectorName := d.Id()
	err := client.DeleteResource(service.Cacheselector.Type(), cacheselectorName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
