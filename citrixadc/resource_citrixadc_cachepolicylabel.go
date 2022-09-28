package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcCachepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCachepolicylabelFunc,
		Read:          readCachepolicylabelFunc,
		Delete:        deleteCachepolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"evaluates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCachepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCachepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicylabelName := d.Get("labelname").(string)
	cachepolicylabel := cache.Cachepolicylabel{
		Evaluates: d.Get("evaluates").(string),
		Labelname: d.Get("labelname").(string),
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
	d.Set("evaluates", data["evaluates"])
	d.Set("labelname", data["labelname"])

	return nil

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