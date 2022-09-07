package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCachepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCachepolicylabelFunc,
		Read:          readCachepolicylabelFunc,
		Update:        updateCachepolicylabelFunc,
		Delete:        deleteCachepolicylabelFunc,
		Schema: map[string]*schema.Schema{
			"evaluates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCachepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCachepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	var cachepolicylabelName string
	if v, ok := d.GetOk("labelname"); ok {
		cachepolicylabelName = v.(string)
	} else {
		cachepolicylabelName = resource.PrefixedUniqueId("tf-cachepolicylabel-")
		d.Set("labelname", cachepolicylabelName)
	}
	cachepolicylabel := cache.Cachepolicylabel{
		Evaluates: d.Get("evaluates").(string),
		Labelname: d.Get("labelname").(string),
		Newname:   d.Get("newname").(string),
	}

	_, err := client.AddResource(service.Cachepolicylabel.Type(), cachepolicylabelName, &cachepolicylabel)
	if err != nil {
		return err
	}

	d.SetId(cachepolicylabelName)

	err = readCachepolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cachepolicylabel but we can't read it ?? %s", cachepolicylabelName)
		return nil
	}
	return nil
}

func readCachepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCachepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cachepolicylabel state %s", cachepolicylabelName)
	data, err := client.FindResource(service.Cachepolicylabel.Type(), cachepolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cachepolicylabel state %s", cachepolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("evaluates", data["evaluates"])
	d.Set("labelname", data["labelname"])
	d.Set("newname", data["newname"])

	return nil

}

func updateCachepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCachepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicylabelName := d.Get("labelname").(string)

	cachepolicylabel := cache.Cachepolicylabel{
		Labelname: d.Get("labelname").(string),
	}
	hasChange := false
	if d.HasChange("evaluates") {
		log.Printf("[DEBUG]  citrixadc-provider: Evaluates has changed for cachepolicylabel %s, starting update", cachepolicylabelName)
		cachepolicylabel.Evaluates = d.Get("evaluates").(string)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for cachepolicylabel %s, starting update", cachepolicylabelName)
		cachepolicylabel.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for cachepolicylabel %s, starting update", cachepolicylabelName)
		cachepolicylabel.Newname = d.Get("newname").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Cachepolicylabel.Type(), cachepolicylabelName, &cachepolicylabel)
		if err != nil {
			return fmt.Errorf("Error updating cachepolicylabel %s", cachepolicylabelName)
		}
	}
	return readCachepolicylabelFunc(d, meta)
}

func deleteCachepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCachepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicylabelName := d.Id()
	err := client.DeleteResource(service.Cachepolicylabel.Type(), cachepolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
