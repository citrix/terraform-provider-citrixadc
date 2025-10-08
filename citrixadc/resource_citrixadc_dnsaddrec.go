package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnsaddrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsaddrecFunc,
		ReadContext:   readDnsaddrecFunc,
		DeleteContext: deleteDnsaddrecFunc,
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
			"ipaddress": {
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

func createDnsaddrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsaddrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsaddrecId := d.Get("hostname").(string) + "," + d.Get("ipaddress").(string)

	dnsaddrec := dns.Dnsaddrec{
		Ecssubnet: d.Get("ecssubnet").(string),
		Hostname:  d.Get("hostname").(string),
		Ipaddress: d.Get("ipaddress").(string),
		Nodeid:    d.Get("nodeid").(int),
		Ttl:       d.Get("ttl").(int),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(service.Dnsaddrec.Type(), dnsaddrecId, &dnsaddrec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsaddrecId)

	return readDnsaddrecFunc(ctx, d, meta)
}

func readDnsaddrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsaddrecFunc")
	client := meta.(*NetScalerNitroClient).client
	PrimaryId := d.Id()

	// To make the resource backward compatible, in the prev state file user will have ID with 1 value, but in release v1.27.0 we have updated Id. So here we are changing the code to make it backward compatible
	// here we are checking for id, if it has 1 elements then we are appending the 2rd attribute to the old Id.
	oldIdSlice := strings.Split(PrimaryId, ",")

	if len(oldIdSlice) == 1 {
		if val, ok := d.GetOk("ipaddress"); ok {
			PrimaryId = PrimaryId + "," + val.(string)
		}
		d.SetId(PrimaryId)
	}

	idSlice := strings.SplitN(PrimaryId, ",", 2)
	if len(idSlice) != 2 {
		log.Printf("[DEBUG] citrixadc-provider:  In readDnsaddrecFunc: PrimaryId is not in the correct format")
		return diag.Errorf("citrixadc-provider:  In readDnsaddrecFunc: PrimaryId is not in the correct format for dnsaddrec %v", PrimaryId)
	}

	hostname := idSlice[0]
	ipaddress := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading dnsaddrec state %s", PrimaryId)
	dataArr, err := client.FindAllResources(service.Dnsaddrec.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsaddrec state %s", PrimaryId)
		d.SetId("")
		return nil
	}
	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: dnsaddrec does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["hostname"].(string) == hostname && v["ipaddress"].(string) == ipaddress {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams dnsaddrec not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing dnsaddrec state %s", PrimaryId)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("hostname", data["hostname"])
	d.Set("ipaddress", data["ipaddress"])
	setToInt("nodeid", d, data["nodeid"])
	setToInt("ttl", d, data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func deleteDnsaddrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsaddrecFunc")
	client := meta.(*NetScalerNitroClient).client
	argsMap := make(map[string]string)
	if ecs, ok := d.GetOk("ecssubnet"); ok {
		argsMap["ecssubnet"] = url.QueryEscape(ecs.(string))
	}
	argsMap["ipaddress"] = url.QueryEscape(d.Get("ipaddress").(string))

	err := client.DeleteResourceWithArgsMap(service.Dnsaddrec.Type(), d.Get("hostname").(string), argsMap)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return nil
}
