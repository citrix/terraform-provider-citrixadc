package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/responder"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcResponderpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createResponderpolicylabelFunc,
		ReadContext:   readResponderpolicylabelFunc,
		DeleteContext: deleteResponderpolicylabelFunc,
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
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policylabeltype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createResponderpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Get("labelname").(string)
	responderpolicylabel := responder.Responderpolicylabel{
		Comment:         d.Get("comment").(string),
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(service.Responderpolicylabel.Type(), responderpolicylabelName, &responderpolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(responderpolicylabelName)

	return readResponderpolicylabelFunc(ctx, d, meta)
}

func readResponderpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading responderpolicylabel state %s", responderpolicylabelName)
	data, err := client.FindResource(service.Responderpolicylabel.Type(), responderpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderpolicylabel state %s", responderpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("labelname", data["labelname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}

func deleteResponderpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Id()
	err := client.DeleteResource(service.Responderpolicylabel.Type(), responderpolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
