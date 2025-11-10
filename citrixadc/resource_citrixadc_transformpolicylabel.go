package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/transform"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcTransformpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTransformpolicylabelFunc,
		ReadContext:   readTransformpolicylabelFunc,
		DeleteContext: deleteTransformpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policylabeltype": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createTransformpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTransformpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicylabelName := d.Get("labelname").(string)
	transformpolicylabel := transform.Transformpolicylabel{
		Labelname:       transformpolicylabelName,
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(service.Transformpolicylabel.Type(), transformpolicylabelName, &transformpolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(transformpolicylabelName)

	return readTransformpolicylabelFunc(ctx, d, meta)
}

func readTransformpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTransformpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading transformpolicylabel state %s", transformpolicylabelName)
	data, err := client.FindResource(service.Transformpolicylabel.Type(), transformpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing transformpolicylabel state %s", transformpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}

func deleteTransformpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTransformpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicylabelName := d.Id()
	err := client.DeleteResource(service.Transformpolicylabel.Type(), transformpolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
