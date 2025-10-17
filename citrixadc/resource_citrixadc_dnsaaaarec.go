package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"log"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnsaaaarec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsaaaarecFunc,
		ReadContext:   readDnsaaaarecFunc,
		DeleteContext: deleteDnsaaaarecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ecssubnet": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipv6address": {
				Type:     schema.TypeString,
				Required: true,
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createDnsaaaarecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsaaaarecFunc")
	client := meta.(*NetScalerNitroClient).client

	dnsaaaarec := dns.Dnsaaaarec{
		Ecssubnet:   d.Get("ecssubnet").(string),
		Hostname:    d.Get("hostname").(string),
		Ipv6address: d.Get("ipv6address").(string),
		Type:        d.Get("type").(string),
	}

	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		dnsaaaarec.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		dnsaaaarec.Ttl = intPtr(d.Get("ttl").(int))
	}
	var dnsaaaarecName string
	if Hostname, ok := d.GetOk("hostname"); ok {
		dnsaaaarecName = Hostname.(string)
	}

	_, err := client.AddResource(service.Dnsaaaarec.Type(), dnsaaaarecName, &dnsaaaarec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsaaaarecName)

	return readDnsaaaarecFunc(ctx, d, meta)
}

func readDnsaaaarecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsaaaarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsaaaarecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsaaaarec state %s", dnsaaaarecName)
	findParams := service.FindParams{
		ResourceType: service.Dnsaaaarec.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsaaaarec state %s", dnsaaaarecName)
		d.SetId("")
		return nil
	}

	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: dns aaaarec does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, dnsaaaarec := range dataArray {
		if dnsaaaarec["hostname"] == d.Get("hostname").(string) {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams dnsaaaarec not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing dnsaaaarec state %s", dnsaaaarecName)
		d.SetId("")
		return nil
	}
	data := dataArray[foundIndex]
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("hostname", data["hostname"])
	d.Set("ipv6address", data["ipv6address"])
	setToInt("nodeid", d, data["nodeid"])
	setToInt("ttl", d, data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func deleteDnsaaaarecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsaaaarecFunc")
	client := meta.(*NetScalerNitroClient).client
	argsMap := make(map[string]string)
	if ecs, ok := d.GetOk("ecssubnet"); ok {
		argsMap["ecssubnet"] = url.QueryEscape(ecs.(string))
	}
	argsMap["ipv6address"] = url.QueryEscape(d.Get("ipv6address").(string))

	err := client.DeleteResourceWithArgsMap(service.Dnsaaaarec.Type(), d.Id(), argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
