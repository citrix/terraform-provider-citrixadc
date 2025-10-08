package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcDnsview() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsviewFunc,
		ReadContext:   readDnsviewFunc,
		DeleteContext: deleteDnsviewFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"viewname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDnsviewFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsviewFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsviewName := d.Get("viewname").(string)
	dnsview := dns.Dnsview{
		Viewname: d.Get("viewname").(string),
	}

	_, err := client.AddResource(service.Dnsview.Type(), dnsviewName, &dnsview)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsviewName)

	return readDnsviewFunc(ctx, d, meta)
}

func readDnsviewFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsviewFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsviewName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsview state %s", dnsviewName)
	data, err := client.FindResource(service.Dnsview.Type(), dnsviewName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsview state %s", dnsviewName)
		d.SetId("")
		return nil
	}
	d.Set("viewname", data["viewname"])

	return nil

}

func deleteDnsviewFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsviewFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsviewName := d.Id()
	err := client.DeleteResource(service.Dnsview.Type(), dnsviewName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
