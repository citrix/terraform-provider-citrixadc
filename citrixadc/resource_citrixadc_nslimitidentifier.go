package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNslimitidentifier() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNslimitidentifierFunc,
		ReadContext:   readNslimitidentifierFunc,
		UpdateContext: updateNslimitidentifierFunc,
		DeleteContext: deleteNslimitidentifierFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"limitidentifier": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"limittype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"selectorname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeslice": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trapsintimeslice": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNslimitidentifierFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNslimitidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	nslimitidentifierName := d.Get("limitidentifier").(string)

	nslimitidentifier := ns.Nslimitidentifier{
		Limitidentifier: d.Get("limitidentifier").(string),
		Limittype:       d.Get("limittype").(string),
		Mode:            d.Get("mode").(string),
		Selectorname:    d.Get("selectorname").(string),
	}

	if raw := d.GetRawConfig().GetAttr("maxbandwidth"); !raw.IsNull() {
		nslimitidentifier.Maxbandwidth = intPtr(d.Get("maxbandwidth").(int))
	}
	if raw := d.GetRawConfig().GetAttr("threshold"); !raw.IsNull() {
		nslimitidentifier.Threshold = intPtr(d.Get("threshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("timeslice"); !raw.IsNull() {
		nslimitidentifier.Timeslice = intPtr(d.Get("timeslice").(int))
	}
	if raw := d.GetRawConfig().GetAttr("trapsintimeslice"); !raw.IsNull() {
		nslimitidentifier.Trapsintimeslice = intPtr(d.Get("trapsintimeslice").(int))
	}

	_, err := client.AddResource(service.Nslimitidentifier.Type(), nslimitidentifierName, &nslimitidentifier)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nslimitidentifierName)

	return readNslimitidentifierFunc(ctx, d, meta)
}

func readNslimitidentifierFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNslimitidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	nslimitidentifierName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nslimitidentifier state %s", nslimitidentifierName)
	data, err := client.FindResource(service.Nslimitidentifier.Type(), nslimitidentifierName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nslimitidentifier state %s", nslimitidentifierName)
		d.SetId("")
		return nil
	}
	d.Set("limitidentifier", data["limitidentifier"])
	d.Set("limittype", data["limittype"])
	setToInt("maxbandwidth", d, data["maxbandwidth"])
	d.Set("mode", data["mode"])
	d.Set("selectorname", data["selectorname"])
	setToInt("threshold", d, data["threshold"])
	setToInt("timeslice", d, data["timeslice"])
	setToInt("trapsintimeslice", d, data["trapsintimeslice"])

	return nil

}

func updateNslimitidentifierFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNslimitidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	nslimitidentifierName := d.Get("limitidentifier").(string)

	nslimitidentifier := ns.Nslimitidentifier{
		Limitidentifier: d.Get("limitidentifier").(string),
	}
	hasChange := false
	if d.HasChange("limittype") {
		log.Printf("[DEBUG]  citrixadc-provider: Limittype has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Limittype = d.Get("limittype").(string)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbandwidth has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Maxbandwidth = intPtr(d.Get("maxbandwidth").(int))
		hasChange = true
	}
	if d.HasChange("mode") {
		log.Printf("[DEBUG]  citrixadc-provider: Mode has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Mode = d.Get("mode").(string)
		hasChange = true
	}
	if d.HasChange("selectorname") {
		log.Printf("[DEBUG]  citrixadc-provider: Selectorname has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Selectorname = d.Get("selectorname").(string)
		hasChange = true
	}
	if d.HasChange("threshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Threshold has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Threshold = intPtr(d.Get("threshold").(int))
		hasChange = true
	}
	if d.HasChange("timeslice") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeslice has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Timeslice = intPtr(d.Get("timeslice").(int))
		hasChange = true
	}
	if d.HasChange("trapsintimeslice") {
		log.Printf("[DEBUG]  citrixadc-provider: Trapsintimeslice has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Trapsintimeslice = intPtr(d.Get("trapsintimeslice").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nslimitidentifier.Type(), nslimitidentifierName, &nslimitidentifier)
		if err != nil {
			return diag.Errorf("Error updating nslimitidentifier %s", nslimitidentifierName)
		}
	}
	return readNslimitidentifierFunc(ctx, d, meta)
}

func deleteNslimitidentifierFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNslimitidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	nslimitidentifierName := d.Id()
	err := client.DeleteResource(service.Nslimitidentifier.Type(), nslimitidentifierName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
