package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcDnssuffix() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnssuffixFunc,
		ReadContext:   readDnssuffixFunc,
		DeleteContext: deleteDnssuffixFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"dnssuffix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDnssuffixFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnssuffixFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssuffixName := d.Get("dnssuffix").(string)
	dnssuffix := dns.Dnssuffix{
		Dnssuffix: dnssuffixName,
	}

	_, err := client.AddResource(service.Dnssuffix.Type(), dnssuffixName, &dnssuffix)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnssuffixName)

	return readDnssuffixFunc(ctx, d, meta)
}

func readDnssuffixFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnssuffixFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssuffixName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnssuffix state %s", dnssuffixName)
	data, err := client.FindResource(service.Dnssuffix.Type(), dnssuffixName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnssuffix state %s", dnssuffixName)
		d.SetId("")
		return nil
	}
	d.Set("dnssuffix", data["Dnssuffix"])

	return nil

}

func deleteDnssuffixFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnssuffixFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssuffixName := d.Id()
	err := client.DeleteResource(service.Dnssuffix.Type(), dnssuffixName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
