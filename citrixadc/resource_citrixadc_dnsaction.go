package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnsaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsactionFunc,
		ReadContext:   readDnsactionFunc,
		UpdateContext: updateDnsactionFunc,
		DeleteContext: deleteDnsactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"actionname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"actiontype": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"dnsprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipaddress": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"preferredloclist": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"viewname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsactionName := d.Get("actionname").(string)
	dnsaction := dns.Dnsaction{
		Actionname:       d.Get("actionname").(string),
		Actiontype:       d.Get("actiontype").(string),
		Dnsprofilename:   d.Get("dnsprofilename").(string),
		Ipaddress:        toStringList(d.Get("ipaddress").([]interface{})),
		Preferredloclist: toStringList(d.Get("preferredloclist").([]interface{})),
		Ttl:              d.Get("ttl").(int),
		Viewname:         d.Get("viewname").(string),
	}

	_, err := client.AddResource(service.Dnsaction.Type(), dnsactionName, &dnsaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsactionName)

	return readDnsactionFunc(ctx, d, meta)
}

func readDnsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsaction state %s", dnsactionName)
	data, err := client.FindResource(service.Dnsaction.Type(), dnsactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsaction state %s", dnsactionName)
		d.SetId("")
		return nil
	}
	d.Set("actionname", data["actionname"])
	d.Set("actionname", data["actionname"])
	d.Set("actiontype", data["actiontype"])
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("preferredloclist", data["preferredloclist"])
	setToInt("ttl", d, data["ttl"])
	d.Set("viewname", data["viewname"])

	return nil

}

func updateDnsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsactionName := d.Get("actionname").(string)

	dnsaction := dns.Dnsaction{
		Actionname: d.Get("actionname").(string),
	}
	hasChange := false
	if d.HasChange("actiontype") {
		log.Printf("[DEBUG]  citrixadc-provider: Actiontype has changed for dnsaction %s, starting update", dnsactionName)
		dnsaction.Actiontype = d.Get("actiontype").(string)
		hasChange = true
	}
	if d.HasChange("dnsprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsprofilename has changed for dnsaction %s, starting update", dnsactionName)
		dnsaction.Dnsprofilename = d.Get("dnsprofilename").(string)
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipaddress has changed for dnsaction %s, starting update", dnsactionName)
		dnsaction.Ipaddress = toStringList(d.Get("ipaddress").([]interface{}))
		hasChange = true
	}
	if d.HasChange("preferredloclist") {
		log.Printf("[DEBUG]  citrixadc-provider: Preferredloclist has changed for dnsaction %s, starting update", dnsactionName)
		dnsaction.Preferredloclist = toStringList(d.Get("preferredloclist").([]interface{}))
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for dnsaction %s, starting update", dnsactionName)
		dnsaction.Ttl = d.Get("ttl").(int)
		hasChange = true
	}
	if d.HasChange("viewname") {
		log.Printf("[DEBUG]  citrixadc-provider: Viewname has changed for dnsaction %s, starting update", dnsactionName)
		dnsaction.Viewname = d.Get("viewname").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnsaction.Type(), dnsactionName, &dnsaction)
		if err != nil {
			return diag.Errorf("Error updating dnsaction %s", dnsactionName)
		}
	}
	return readDnsactionFunc(ctx, d, meta)
}

func deleteDnsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsactionName := d.Id()
	err := client.DeleteResource(service.Dnsaction.Type(), dnsactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
