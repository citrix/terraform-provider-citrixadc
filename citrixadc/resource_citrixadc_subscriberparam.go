package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSubscriberparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSubscriberparamFunc,
		ReadContext:   readSubscriberparamFunc,
		UpdateContext: updateSubscriberparamFunc,
		DeleteContext: deleteSubscriberparamFunc,
		Schema: map[string]*schema.Schema{
			"idleaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idlettl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interfacetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6prefixlookuplist": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"keytype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSubscriberparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSubscriberparamFunc")
	client := meta.(*NetScalerNitroClient).client
	subscriberparamName := resource.PrefixedUniqueId("tf-subscriberparam-")

	subscriberparam := subscriber.Subscriberparam{
		Idleaction:           d.Get("idleaction").(string),
		Interfacetype:        d.Get("interfacetype").(string),
		Ipv6prefixlookuplist: toIntegerList(d.Get("ipv6prefixlookuplist").([]interface{})),
		Keytype:              d.Get("keytype").(string),
	}

	if raw := d.GetRawConfig().GetAttr("idlettl"); !raw.IsNull() {
		subscriberparam.Idlettl = intPtr(d.Get("idlettl").(int))
	}

	err := client.UpdateUnnamedResource("subscriberparam", &subscriberparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(subscriberparamName)

	return readSubscriberparamFunc(ctx, d, meta)
}

func readSubscriberparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSubscriberparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading subscriberparam state")
	data, err := client.FindResource("subscriberparam", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing subscriberparam state")
		d.SetId("")
		return nil
	}
	d.Set("idleaction", data["idleaction"])
	setToInt("idlettl", d, data["idlettl"])
	d.Set("interfacetype", data["interfacetype"])
	//d.Set("ipv6prefixlookuplist", data["ipv6prefixlookuplist"])
	d.Set("keytype", data["keytype"])

	return nil

}

func updateSubscriberparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSubscriberparamFunc")
	client := meta.(*NetScalerNitroClient).client

	subscriberparam := subscriber.Subscriberparam{}
	hasChange := false
	if d.HasChange("idleaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Idleaction has changed for subscriberparam, starting update")
		subscriberparam.Idleaction = d.Get("idleaction").(string)
		hasChange = true
	}
	if d.HasChange("idlettl") {
		log.Printf("[DEBUG]  citrixadc-provider: Idlettl has changed for subscriberparam, starting update")
		subscriberparam.Idlettl = intPtr(d.Get("idlettl").(int))
		hasChange = true
	}
	if d.HasChange("interfacetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Interfacetype has changed for subscriberparam, starting update")
		subscriberparam.Interfacetype = d.Get("interfacetype").(string)
		hasChange = true
	}
	if d.HasChange("ipv6prefixlookuplist") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv6prefixlookuplist has changed for subscriberparam, starting update")
		subscriberparam.Ipv6prefixlookuplist = toIntegerList(d.Get("ipv6prefixlookuplist").([]interface{}))
		hasChange = true
	}
	if d.HasChange("keytype") {
		log.Printf("[DEBUG]  citrixadc-provider: Keytype has changed for subscriberparam, starting update")
		subscriberparam.Keytype = d.Get("keytype").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("subscriberparam", &subscriberparam)
		if err != nil {
			return diag.Errorf("Error updating subscriberparam")
		}
	}
	return readSubscriberparamFunc(ctx, d, meta)
}

func deleteSubscriberparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSubscriberparamFunc")
	//subscriberparam does not support DELETE operation
	d.SetId("")

	return nil
}
