package citrixadc

import (
	"context"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnscnamerec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnscnamerecFunc,
		ReadContext:   readDnscnamerecFunc,
		DeleteContext: deleteDnscnamerecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"aliasname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"canonicalname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ecssubnet": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
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

func createDnscnamerecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnscnamerecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnscnamerecName := d.Get("aliasname").(string)
	dnscnamerec := dns.Dnscnamerec{
		Aliasname:     dnscnamerecName,
		Canonicalname: d.Get("canonicalname").(string),
		Ecssubnet:     d.Get("ecssubnet").(string),
	}

	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		dnscnamerec.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		dnscnamerec.Ttl = intPtr(d.Get("ttl").(int))
	}

	_, err := client.AddResource(service.Dnscnamerec.Type(), dnscnamerecName, &dnscnamerec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnscnamerecName)

	return readDnscnamerecFunc(ctx, d, meta)
}

func readDnscnamerecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnscnamerecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnscnamerecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnscnamerec state %s", dnscnamerecName)
	data, err := client.FindResource(service.Dnscnamerec.Type(), dnscnamerecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnscnamerec state %s", dnscnamerecName)
		d.SetId("")
		return nil
	}
	d.Set("aliasname", data["aliasname"])
	d.Set("canonicalname", data["canonicalname"])
	d.Set("ecssubnet", data["ecssubnet"])
	setToInt("nodeid", d, data["nodeid"])
	setToInt("ttl", d, data["ttl"])

	return nil

}

func deleteDnscnamerecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnscnamerecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnscnamerecName := d.Id()
	var err error
	argsMap := make(map[string]string)
	if ecs, ok := d.GetOk("ecssubnet"); ok {
		argsMap["ecssubnet"] = url.QueryEscape(ecs.(string))
		err = client.DeleteResourceWithArgsMap(service.Dnscnamerec.Type(), dnscnamerecName, argsMap)
	} else {
		err = client.DeleteResource(service.Dnscnamerec.Type(), dnscnamerecName)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
