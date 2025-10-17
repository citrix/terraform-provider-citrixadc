package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ipsec"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIpsecparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIpsecparameterFunc,
		ReadContext:   readIpsecparameterFunc,
		UpdateContext: updateIpsecparameterFunc,
		DeleteContext: deleteIpsecparameterFunc,
		Schema: map[string]*schema.Schema{
			"encalgo": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hashalgo": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ikeretryinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ikeversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"livenesscheckinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"perfectforwardsecrecy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replaywindowsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retransmissiontime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIpsecparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpsecparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecparameterName := resource.PrefixedUniqueId("tf-ipsecparameter-")

	ipsecparameter := ipsec.Ipsecparameter{
		Encalgo:               toStringList(d.Get("encalgo").([]interface{})),
		Hashalgo:              toStringList(d.Get("hashalgo").([]interface{})),
		Ikeversion:            d.Get("ikeversion").(string),
		Perfectforwardsecrecy: d.Get("perfectforwardsecrecy").(string),
	}

	if raw := d.GetRawConfig().GetAttr("ikeretryinterval"); !raw.IsNull() {
		ipsecparameter.Ikeretryinterval = intPtr(d.Get("ikeretryinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("lifetime"); !raw.IsNull() {
		ipsecparameter.Lifetime = intPtr(d.Get("lifetime").(int))
	}
	if raw := d.GetRawConfig().GetAttr("livenesscheckinterval"); !raw.IsNull() {
		ipsecparameter.Livenesscheckinterval = intPtr(d.Get("livenesscheckinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("replaywindowsize"); !raw.IsNull() {
		ipsecparameter.Replaywindowsize = intPtr(d.Get("replaywindowsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("retransmissiontime"); !raw.IsNull() {
		ipsecparameter.Retransmissiontime = intPtr(d.Get("retransmissiontime").(int))
	}

	err := client.UpdateUnnamedResource(service.Ipsecparameter.Type(), &ipsecparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ipsecparameterName)

	return readIpsecparameterFunc(ctx, d, meta)
}

func readIpsecparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpsecparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ipsecparameter state")
	data, err := client.FindResource(service.Ipsecparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ipsecparameter state")
		d.SetId("")
		return nil
	}
	d.Set("encalgo", data["encalgo"])
	d.Set("hashalgo", data["hashalgo"])
	setToInt("ikeretryinterval", d, data["ikeretryinterval"])
	d.Set("ikeversion", data["ikeversion"])
	setToInt("lifetime", d, data["lifetime"])
	setToInt("livenesscheckinterval", d, data["livenesscheckinterval"])
	d.Set("perfectforwardsecrecy", data["perfectforwardsecrecy"])
	setToInt("replaywindowsize", d, data["replaywindowsize"])
	setToInt("retransmissiontime", d, data["retransmissiontime"])

	return nil

}

func updateIpsecparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIpsecparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	ipsecparameter := ipsec.Ipsecparameter{}
	hasChange := false
	if d.HasChange("encalgo") {
		log.Printf("[DEBUG]  citrixadc-provider: Encalgo has changed for ipsecparameter, starting update")
		ipsecparameter.Encalgo = toStringList(d.Get("encalgo").([]interface{}))
		hasChange = true
	}
	if d.HasChange("hashalgo") {
		log.Printf("[DEBUG]  citrixadc-provider: Hashalgo has changed for ipsecparameter, starting update")
		ipsecparameter.Hashalgo = toStringList(d.Get("hashalgo").([]interface{}))
		hasChange = true
	}
	if d.HasChange("ikeretryinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Ikeretryinterval has changed for ipsecparameter, starting update")
		ipsecparameter.Ikeretryinterval = intPtr(d.Get("ikeretryinterval").(int))
		hasChange = true
	}
	if d.HasChange("ikeversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Ikeversion has changed for ipsecparameter, starting update")
		ipsecparameter.Ikeversion = d.Get("ikeversion").(string)
		hasChange = true
	}
	if d.HasChange("lifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Lifetime has changed for ipsecparameter, starting update")
		ipsecparameter.Lifetime = intPtr(d.Get("lifetime").(int))
		hasChange = true
	}
	if d.HasChange("livenesscheckinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Livenesscheckinterval has changed for ipsecparameter, starting update")
		ipsecparameter.Livenesscheckinterval = intPtr(d.Get("livenesscheckinterval").(int))
		hasChange = true
	}
	if d.HasChange("perfectforwardsecrecy") {
		log.Printf("[DEBUG]  citrixadc-provider: Perfectforwardsecrecy has changed for ipsecparameter, starting update")
		ipsecparameter.Perfectforwardsecrecy = d.Get("perfectforwardsecrecy").(string)
		hasChange = true
	}
	if d.HasChange("replaywindowsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Replaywindowsize has changed for ipsecparameter, starting update")
		ipsecparameter.Replaywindowsize = intPtr(d.Get("replaywindowsize").(int))
		hasChange = true
	}
	if d.HasChange("retransmissiontime") {
		log.Printf("[DEBUG]  citrixadc-provider: Retransmissiontime has changed for ipsecparameter, starting update")
		ipsecparameter.Retransmissiontime = intPtr(d.Get("retransmissiontime").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ipsecparameter.Type(), &ipsecparameter)
		if err != nil {
			return diag.Errorf("Error updating ipsecparameter")
		}
	}
	return readIpsecparameterFunc(ctx, d, meta)
}

func deleteIpsecparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpsecparameterFunc")
	//ipsecparameter does not support DELETE operation
	d.SetId("")

	return nil
}
