package citrixadc

import (
	"context"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnssrvrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnssrvrecFunc,
		ReadContext:   readDnssrvrecFunc,
		UpdateContext: updateDnssrvrecFunc,
		DeleteContext: deleteDnssrvrecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"target": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnssrvrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnssrvrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssrvrecName := d.Get("domain").(string) + "," + d.Get("target").(string)
	dnssrvrec := dns.Dnssrvrec{
		Domain:    d.Get("domain").(string),
		Ecssubnet: d.Get("ecssubnet").(string),
		Target:    d.Get("target").(string),
		Type:      d.Get("type").(string),
	}

	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		dnssrvrec.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		dnssrvrec.Port = intPtr(d.Get("port").(int))
	}
	if raw := d.GetRawConfig().GetAttr("priority"); !raw.IsNull() {
		dnssrvrec.Priority = intPtr(d.Get("priority").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		dnssrvrec.Ttl = intPtr(d.Get("ttl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		dnssrvrec.Weight = intPtr(d.Get("weight").(int))
	}

	_, err := client.AddResource(service.Dnssrvrec.Type(), dnssrvrecName, &dnssrvrec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnssrvrecName)

	return readDnssrvrecFunc(ctx, d, meta)
}

func readDnssrvrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnssrvrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssrvrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnssrvrec state %s", dnssrvrecName)
	findParams := service.FindParams{
		ResourceType: service.Dnssrvrec.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnssrvrec state %s", dnssrvrecName)
		d.SetId("")
		return nil
	}

	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: dns srvrec does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	idSlice := strings.SplitN(dnssrvrecName, ",", 2)

	domain := idSlice[0]
	target := idSlice[1]

	foundIndex := -1
	for i, dnssrvrec := range dataArray {
		match := true
		if dnssrvrec["domain"] != domain {
			match = false
		}
		if dnssrvrec["target"] != target {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams dnssrvrec not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing dnssrvrec state %s", dnssrvrecName)
		d.SetId("")
		return nil
	}
	data := dataArray[foundIndex]
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	setToInt("nodeid", d, data["nodeid"])
	setToInt("port", d, data["port"])
	setToInt("priority", d, data["priority"])
	d.Set("target", data["target"])
	setToInt("ttl", d, data["ttl"])
	d.Set("type", data["type"])
	setToInt("weight", d, data["weight"])

	return nil

}

func updateDnssrvrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnssrvrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssrvrecName := d.Get("domain").(string)
	dnssrvrec := dns.Dnssrvrec{}
	log.Println(dnssrvrec)
	hasChange := false

	if d.HasChange("ecssubnet") {
		log.Printf("[DEBUG]  citrixadc-provider: Ecssubnet has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Ecssubnet = d.Get("ecssubnet").(string)
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Nodeid = intPtr(d.Get("nodeid").(int))
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Port = intPtr(d.Get("port").(int))
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Priority = intPtr(d.Get("priority").(int))
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Ttl = intPtr(d.Get("ttl").(int))
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Weight = intPtr(d.Get("weight").(int))
		hasChange = true
	}

	if hasChange {
		dnssrvrec.Domain = d.Get("domain").(string)
		dnssrvrec.Target = d.Get("target").(string)
		err := client.UpdateUnnamedResource(service.Dnssrvrec.Type(), &dnssrvrec)
		if err != nil {
			return diag.Errorf("Error updating dnssrvrec %s", dnssrvrecName)
		}
	}
	return readDnssrvrecFunc(ctx, d, meta)
}

func deleteDnssrvrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnssrvrecFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	domain := idSlice[0]
	target := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["target"] = url.QueryEscape(target)
	if ecs, ok := d.GetOk("ecssubnet"); ok {
		argsMap["ecssubnet"] = url.QueryEscape(ecs.(string))
	}
	err := client.DeleteResourceWithArgsMap(service.Dnssrvrec.Type(), domain, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
