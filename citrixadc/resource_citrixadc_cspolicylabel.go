package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCspolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCspolicylabelFunc,
		ReadContext:   readCspolicylabelFunc,
		DeleteContext: deleteCspolicylabelFunc,
		Schema: map[string]*schema.Schema{
			"cspolicylabeltype": {
				Type:     schema.TypeString,
				Required: true,
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

func createCspolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	var cspolicylabelName string
	if v, ok := d.GetOk("labelname"); ok {
		cspolicylabelName = v.(string)
	} else {
		cspolicylabelName = resource.PrefixedUniqueId("tf-cspolicylabel-")
		d.Set("labelname", cspolicylabelName)
	}
	cspolicylabel := cs.Cspolicylabel{
		Cspolicylabeltype: d.Get("cspolicylabeltype").(string),
		Labelname:         d.Get("labelname").(string),
	}

	_, err := client.AddResource(service.Cspolicylabel.Type(), cspolicylabelName, &cspolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cspolicylabelName)

	return readCspolicylabelFunc(ctx, d, meta)
}

func readCspolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cspolicylabel state %s", cspolicylabelName)
	data, err := client.FindResource(service.Cspolicylabel.Type(), cspolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cspolicylabel state %s", cspolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("cspolicylabeltype", data["cspolicylabeltype"])

	return nil

}

func deleteCspolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicylabelName := d.Id()
	err := client.DeleteResource(service.Cspolicylabel.Type(), cspolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
