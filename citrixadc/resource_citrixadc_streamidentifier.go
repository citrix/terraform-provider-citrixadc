package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/stream"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcStreamidentifier() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createStreamidentifierFunc,
		ReadContext:   readStreamidentifierFunc,
		UpdateContext: updateStreamidentifierFunc,
		DeleteContext: deleteStreamidentifierFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"acceptancethreshold": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"breachthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxtransactionthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mintransactionthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"samplecount": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"selectorname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmptrap": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sort": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trackackonlypackets": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tracktransactions": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createStreamidentifierFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createStreamidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	streamidentifierName := d.Get("name").(string)
	streamidentifier := stream.Streamidentifier{
		Acceptancethreshold: d.Get("acceptancethreshold").(string),
		Appflowlog:          d.Get("appflowlog").(string),
		Name:                d.Get("name").(string),
		Selectorname:        d.Get("selectorname").(string),
		Snmptrap:            d.Get("snmptrap").(string),
		Sort:                d.Get("sort").(string),
		Trackackonlypackets: d.Get("trackackonlypackets").(string),
		Tracktransactions:   d.Get("tracktransactions").(string),
	}

	if raw := d.GetRawConfig().GetAttr("breachthreshold"); !raw.IsNull() {
		streamidentifier.Breachthreshold = intPtr(d.Get("breachthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("interval"); !raw.IsNull() {
		streamidentifier.Interval = intPtr(d.Get("interval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxtransactionthreshold"); !raw.IsNull() {
		streamidentifier.Maxtransactionthreshold = intPtr(d.Get("maxtransactionthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mintransactionthreshold"); !raw.IsNull() {
		streamidentifier.Mintransactionthreshold = intPtr(d.Get("mintransactionthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("samplecount"); !raw.IsNull() {
		streamidentifier.Samplecount = intPtr(d.Get("samplecount").(int))
	}

	_, err := client.AddResource(service.Streamidentifier.Type(), streamidentifierName, &streamidentifier)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(streamidentifierName)

	return readStreamidentifierFunc(ctx, d, meta)
}

func readStreamidentifierFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readStreamidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	streamidentifierName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading streamidentifier state %s", streamidentifierName)
	data, err := client.FindResource(service.Streamidentifier.Type(), streamidentifierName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing streamidentifier state %s", streamidentifierName)
		d.SetId("")
		return nil
	}
	d.Set("acceptancethreshold", data["acceptancethreshold"])
	d.Set("appflowlog", data["appflowlog"])
	setToInt("breachthreshold", d, data["breachthreshold"])
	setToInt("interval", d, data["interval"])
	setToInt("maxtransactionthreshold", d, data["maxtransactionthreshold"])
	setToInt("mintransactionthreshold", d, data["mintransactionthreshold"])
	d.Set("name", data["name"])
	setToInt("samplecount", d, data["samplecount"])
	d.Set("selectorname", data["selectorname"])
	d.Set("snmptrap", data["snmptrap"])
	d.Set("sort", data["sort"])
	d.Set("trackackonlypackets", data["trackackonlypackets"])
	d.Set("tracktransactions", data["tracktransactions"])

	return nil

}

func updateStreamidentifierFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateStreamidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	streamidentifierName := d.Get("name").(string)

	streamidentifier := stream.Streamidentifier{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("acceptancethreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Acceptancethreshold has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Acceptancethreshold = d.Get("acceptancethreshold").(string)
		hasChange = true
	}
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("breachthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Breachthreshold has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Breachthreshold = intPtr(d.Get("breachthreshold").(int))
		hasChange = true
	}
	if d.HasChange("interval") {
		log.Printf("[DEBUG]  citrixadc-provider: Interval has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Interval = intPtr(d.Get("interval").(int))
		hasChange = true
	}
	if d.HasChange("maxtransactionthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxtransactionthreshold has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Maxtransactionthreshold = intPtr(d.Get("maxtransactionthreshold").(int))
		hasChange = true
	}
	if d.HasChange("mintransactionthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Mintransactionthreshold has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Mintransactionthreshold = intPtr(d.Get("mintransactionthreshold").(int))
		hasChange = true
	}
	if d.HasChange("samplecount") {
		log.Printf("[DEBUG]  citrixadc-provider: Samplecount has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Samplecount = intPtr(d.Get("samplecount").(int))
		hasChange = true
	}
	if d.HasChange("selectorname") {
		log.Printf("[DEBUG]  citrixadc-provider: Selectorname has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Selectorname = d.Get("selectorname").(string)
		hasChange = true
	}
	if d.HasChange("snmptrap") {
		log.Printf("[DEBUG]  citrixadc-provider: Snmptrap has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Snmptrap = d.Get("snmptrap").(string)
		hasChange = true
	}
	if d.HasChange("sort") {
		log.Printf("[DEBUG]  citrixadc-provider: Sort has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Sort = d.Get("sort").(string)
		hasChange = true
	}
	if d.HasChange("trackackonlypackets") {
		log.Printf("[DEBUG]  citrixadc-provider: Trackackonlypackets has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Trackackonlypackets = d.Get("trackackonlypackets").(string)
		hasChange = true
	}
	if d.HasChange("tracktransactions") {
		log.Printf("[DEBUG]  citrixadc-provider: Tracktransactions has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Tracktransactions = d.Get("tracktransactions").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Streamidentifier.Type(), &streamidentifier)
		if err != nil {
			return diag.Errorf("Error updating streamidentifier %s", streamidentifierName)
		}
	}
	return readStreamidentifierFunc(ctx, d, meta)
}

func deleteStreamidentifierFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteStreamidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	streamidentifierName := d.Id()
	err := client.DeleteResource(service.Streamidentifier.Type(), streamidentifierName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
