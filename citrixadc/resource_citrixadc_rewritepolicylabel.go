package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRewritepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRewritepolicylabelFunc,
		ReadContext:   readRewritepolicylabelFunc,
		DeleteContext: deleteRewritepolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createRewritepolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewritepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicylabelName := d.Get("labelname").(string)

	rewritepolicylabel := rewrite.Rewritepolicylabel{
		Comment:   d.Get("comment").(string),
		Labelname: d.Get("labelname").(string),
		Transform: d.Get("transform").(string),
	}

	_, err := client.AddResource(service.Rewritepolicylabel.Type(), rewritepolicylabelName, &rewritepolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rewritepolicylabelName)

	return readRewritepolicylabelFunc(ctx, d, meta)
}

func readRewritepolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRewritepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rewritepolicylabel state %s", rewritepolicylabelName)
	data, err := client.FindResource(service.Rewritepolicylabel.Type(), rewritepolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rewritepolicylabel state %s", rewritepolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("labelname", data["labelname"])
	d.Set("transform", data["transform"])

	return nil

}

func deleteRewritepolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewritepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicylabelName := d.Id()
	err := client.DeleteResource(service.Rewritepolicylabel.Type(), rewritepolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
