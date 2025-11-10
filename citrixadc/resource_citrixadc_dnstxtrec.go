package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcDnstxtrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnstxtrecFunc,
		ReadContext:   readDnstxtrecFunc,
		DeleteContext: deleteDnstxtrecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"string": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createDnstxtrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnstxtrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnstxtrecName := d.Get("domain").(string)
	dnstxtrec := dns.Dnstxtrec{
		Domain: dnstxtrecName,
		String: toStringList(d.Get("string").([]interface{})),
	}

	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		dnstxtrec.Ttl = intPtr(d.Get("ttl").(int))
	}

	_, err := client.AddResource(service.Dnstxtrec.Type(), dnstxtrecName, &dnstxtrec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnstxtrecName)

	return readDnstxtrecFunc(ctx, d, meta)
}

func readDnstxtrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnstxtrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnstxtrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnstxtrec state %s", dnstxtrecName)
	data, err := client.FindResource(service.Dnstxtrec.Type(), dnstxtrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnstxtrec state %s", dnstxtrecName)
		d.SetId("")
		return nil
	}

	d.Set("domain", data["domain"])
	d.Set("string", data["String"])
	setToInt("ttl", d, data["ttl"])

	return nil

}

func deleteDnstxtrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnstxtrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnstxtrecName := d.Id()
	data, err := client.FindResource(service.Dnstxtrec.Type(), dnstxtrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnstxtrec state %s", dnstxtrecName)
		d.SetId("")
		return nil
	}
	argsMap := make(map[string]string)
	argsMap["recordid"] = fmt.Sprintf("%v", data["recordid"])
	argsMap["domain"] = url.QueryEscape(d.Id())

	err = client.DeleteResourceWithArgsMap(service.Dnstxtrec.Type(), dnstxtrecName, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
