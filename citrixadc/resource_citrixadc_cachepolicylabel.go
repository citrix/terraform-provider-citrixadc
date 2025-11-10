package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcCachepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCachepolicylabelFunc,
		ReadContext:   readCachepolicylabelFunc,
		DeleteContext: deleteCachepolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"evaluates": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCachepolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCachepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicylabelName := d.Get("labelname").(string)
	cachepolicylabel := cache.Cachepolicylabel{
		Evaluates: d.Get("evaluates").(string),
		Labelname: d.Get("labelname").(string),
	}

	_, err := client.AddResource(service.Cachepolicylabel.Type(), cachepolicylabelName, &cachepolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cachepolicylabelName)

	return readCachepolicylabelFunc(ctx, d, meta)
}

func readCachepolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func deleteCachepolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCachepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicylabelName := d.Id()
	err := client.DeleteResource(service.Cachepolicylabel.Type(), cachepolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
