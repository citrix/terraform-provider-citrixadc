package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcCmppolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCmppolicylabelFunc,
		ReadContext:   readCmppolicylabelFunc,
		DeleteContext: deleteCmppolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCmppolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCmppolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicylabelName := d.Get("labelname").(string)
	cmppolicylabel := cmp.Cmppolicylabel{
		Labelname: d.Get("labelname").(string),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(service.Cmppolicylabel.Type(), cmppolicylabelName, &cmppolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cmppolicylabelName)

	return readCmppolicylabelFunc(ctx, d, meta)
}

func readCmppolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCmppolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cmppolicylabel state %s", cmppolicylabelName)
	data, err := client.FindResource(service.Cmppolicylabel.Type(), cmppolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cmppolicylabel state %s", cmppolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("type", data["type"])

	return nil

}

func deleteCmppolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCmppolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicylabelName := d.Id()
	err := client.DeleteResource(service.Cmppolicylabel.Type(), cmppolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
