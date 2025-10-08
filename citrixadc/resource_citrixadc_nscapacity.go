package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNscapacity() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNscapacityFunc,
		ReadContext:   readNscapacityFunc,
		UpdateContext: createNscapacityFunc,
		DeleteContext: deleteNscapacityFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"bandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"edition": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcpu": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNscapacityFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNscapacityFunc")
	client := meta.(*NetScalerNitroClient).client

	nscapacityId := resource.PrefixedUniqueId("tf-nscapacity-")
	nscapacity := ns.Nscapacity{
		Bandwidth: d.Get("bandwidth").(int),
		Edition:   d.Get("edition").(string),
		Nodeid:    d.Get("nodeid").(int),
		Platform:  d.Get("platform").(string),
		Unit:      d.Get("unit").(string),
		Vcpu:      d.Get("vcpu").(bool),
	}

	err := client.UpdateUnnamedResource("nscapacity", &nscapacity)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nscapacityId)

	warm := true
	if err = rebootNetScaler(d, meta, warm); err != nil {
		return diag.Errorf("Error warm rebooting ADC. %s", err.Error())
	}

	return readNscapacityFunc(ctx, d, meta)
}

func deleteNscapacityFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNscapacityFunc")
	client := meta.(*NetScalerNitroClient).client

	type nscapacityRemove struct {
		Bandwidth bool `json:"bandwidth,omitempty"`
		Platform  bool `json:"platform,omitempty"`
		Vcpu      bool `json:"vcpu,omitempty"`
	}
	nscapacity := nscapacityRemove{}
	log.Printf("nscapacitydelete struct %v", nscapacity)

	if _, ok := d.GetOk("bandwidth"); ok {
		nscapacity.Bandwidth = true
	}

	if _, ok := d.GetOk("platform"); ok {
		nscapacity.Platform = true
	}

	if _, ok := d.GetOk("vcpu"); ok {
		nscapacity.Vcpu = true
	}

	err := client.ActOnResource("nscapacity", &nscapacity, "unset")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readNscapacityFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNscapacityFunc")
	client := meta.(*NetScalerNitroClient).client

	var err error

	log.Printf("[DEBUG] citrixadc-provider: Reading nscapacity state")
	data, err := client.FindResource("nscapacity", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nscapacity state")
		d.SetId("")
		return nil
	}

	// CICO
	if value, ok := data["platform"]; ok {
		d.Set("platform", value)
	} else {
		d.Set("platform", "")
	}

	// VCPU
	if _, ok := data["vcpucount"]; ok {
		d.Set("vcpu", true)
	} else {
		d.Set("vcpu", false)
	}

	// Pooled
	if value, ok := data["bandwidth"]; ok {
		setToInt("bandwidth", d, value)
		d.Set("edition", data["edition"])
		d.Set("unit", data["unit"])
	} else {
		d.Set("bandwidth", 0)
		d.Set("edition", "")
		d.Set("unit", "")
	}

	return nil

}
