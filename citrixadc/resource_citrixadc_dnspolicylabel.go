package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcDnspolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnspolicylabelFunc,
		ReadContext:   readDnspolicylabelFunc,
		DeleteContext: deleteDnspolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transform": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDnspolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicylabelName := d.Get("labelname").(string)
	dnspolicylabel := dns.Dnspolicylabel{
		Labelname: d.Get("labelname").(string),
		Transform: d.Get("transform").(string),
	}

	_, err := client.AddResource(service.Dnspolicylabel.Type(), dnspolicylabelName, &dnspolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnspolicylabelName)

	return readDnspolicylabelFunc(ctx, d, meta)
}

func readDnspolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnspolicylabel state %s", dnspolicylabelName)
	data, err := client.FindResource(service.Dnspolicylabel.Type(), dnspolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnspolicylabel state %s", dnspolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("transform", data["transform"])

	return nil

}

func deleteDnspolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicylabelName := d.Id()
	err := client.DeleteResource(service.Dnspolicylabel.Type(), dnspolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
