package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRsskeytype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRsskeytypeFunc,
		ReadContext:   readRsskeytypeFunc,
		UpdateContext: updateRsskeytypeFunc,
		DeleteContext: deleteRsskeytypeFunc,
		Schema: map[string]*schema.Schema{
			"rsstype": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createRsskeytypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(rsskeytypeName)

	return readRsskeytypeFunc(ctx, d, meta)
}

func readRsskeytypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRsskeytypeFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading rsskeytype state")
	data, err := client.FindResource(service.Rsskeytype.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rsskeytype state")
		d.SetId("")
		return nil
	}
	d.Set("rsstype", data["rsstype"])

	return nil

}

func updateRsskeytypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("Error updating rsskeytype")
		}
	}
	return readRsskeytypeFunc(ctx, d, meta)
}

func deleteRsskeytypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRsskeytypeFunc")

	d.SetId("")

	return nil
}
