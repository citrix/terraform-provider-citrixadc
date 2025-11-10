package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnsmxrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsmxrecFunc,
		ReadContext:   readDnsmxrecFunc,
		UpdateContext: updateDnsmxrecFunc,
		DeleteContext: deleteDnsmxrecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mx": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pref": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"ecssubnet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnsmxrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsmxrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsmxrecName := d.Get("domain").(string)
	dnsmxrec := dns.Dnsmxrec{
		Domain:    d.Get("domain").(string),
		Ecssubnet: d.Get("ecssubnet").(string),
		Mx:        d.Get("mx").(string),
	}

	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		dnsmxrec.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("pref"); !raw.IsNull() {
		dnsmxrec.Pref = intPtr(d.Get("pref").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		dnsmxrec.Ttl = intPtr(d.Get("ttl").(int))
	}

	_, err := client.AddResource(service.Dnsmxrec.Type(), dnsmxrecName, &dnsmxrec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsmxrecName)

	return readDnsmxrecFunc(ctx, d, meta)
}

func readDnsmxrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsmxrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsmxrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsmxrec state %s", dnsmxrecName)
	data, err := client.FindResource(service.Dnsmxrec.Type(), dnsmxrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsmxrec state %s", dnsmxrecName)
		d.SetId("")
		return nil
	}

	pref1, _ := strconv.Atoi(data["pref"].(string))
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("mx", data["mx"])
	setToInt("nodeid", d, data["nodeid"])
	d.Set("pref", pref1)
	setToInt("ttl", d, data["ttl"])
	log.Printf("set functionality:  %v", data)
	return nil

}

func updateDnsmxrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsmxrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsmxrecName := d.Get("domain").(string)
	dnsmxrec := dns.Dnsmxrec{
		Domain: dnsmxrecName,
		Mx:     d.Get("mx").(string),
	}
	hasChange := false
	if d.HasChange("ecssubnet") {
		log.Printf("[DEBUG]  citrixadc-provider: Ecssubnet has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Ecssubnet = d.Get("ecssubnet").(string)
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Nodeid = intPtr(d.Get("nodeid").(int))
		hasChange = true
	}
	if d.HasChange("pref") {
		log.Printf("[DEBUG]  citrixadc-provider: Pref has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Pref = intPtr(d.Get("pref").(int))
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Ttl = intPtr(d.Get("ttl").(int))
		hasChange = true
	}
	if hasChange {
		_, err := client.UpdateResource(service.Dnsmxrec.Type(), dnsmxrecName, &dnsmxrec)
		if err != nil {
			return diag.Errorf("Error updating dnsmxrec %s", dnsmxrecName)
		}
	}
	return readDnsmxrecFunc(ctx, d, meta)
}

func deleteDnsmxrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsmxrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsmxrecName := d.Id()
	argsMap := make(map[string]string)
	argsMap["mx"] = url.QueryEscape(d.Get("mx").(string))

	if ecscheck, ok := d.GetOk("ecssubnet"); ok {
		argsMap["ecssubnet"] = url.QueryEscape(ecscheck.(string))
	}

	//argsMap["domain"] = dnsmxrecName
	err := client.DeleteResourceWithArgsMap(service.Dnsmxrec.Type(), dnsmxrecName, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
