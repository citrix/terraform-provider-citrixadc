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
	"strings"
)

func resourceCitrixAdcDnsnsrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsnsrecFunc,
		ReadContext:   readDnsnsrecFunc,
		DeleteContext: deleteDnsnsrecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nameserver": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createDnsnsrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsnsrecFunc")
	client := meta.(*NetScalerNitroClient).client
	domain := d.Get("domain").(string)
	nameserver := d.Get("nameserver").(string)

	dnsnsrecId := fmt.Sprintf("%s,%s", domain, nameserver)
	dnsnsrec := dns.Dnsnsrec{
		Domain:     d.Get("domain").(string),
		Nameserver: d.Get("nameserver").(string),
		Ttl:        d.Get("ttl").(int),
	}

	_, err := client.AddResource(service.Dnsnsrec.Type(), "", &dnsnsrec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsnsrecId)

	return readDnsnsrecFunc(ctx, d, meta)
}

func readDnsnsrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsnsrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnsrecId := d.Id()

	idSlice := strings.SplitN(dnsnsrecId, ",", 2)
	domain := idSlice[0]
	nameserver := idSlice[1]
	findParams := service.FindParams{
		ResourceType: "dnsnsrec",
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return diag.FromErr(err)
	}
	foundIndex := -1
	for i, v := range dataArr {
		if v["domain"] == domain && v["nameserver"] == nameserver {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams dnsnsrec not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing dnsnsrec state %s", dnsnsrecId)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]

	d.Set("domain", data["domain"])
	d.Set("nameserver", data["nameserver"])
	setToInt("ttl", d, data["ttl"])

	return nil

}

func deleteDnsnsrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsnsrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnsrecId := d.Id()

	idSlice := strings.SplitN(dnsnsrecId, ",", 2)

	domain := idSlice[0]
	nameserver := idSlice[1]
	argsMap := make(map[string]string)
	argsMap["nameserver"] = url.QueryEscape(nameserver)
	err := client.DeleteResourceWithArgsMap(service.Dnsnsrec.Type(), domain, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
