package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/filter"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcFilteraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createFilteractionFunc,
		ReadContext:   readFilteractionFunc,
		UpdateContext: updateFilteractionFunc,
		DeleteContext: deleteFilteractionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"qual": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"respcode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createFilteractionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createFilteractionFunc")
	client := meta.(*NetScalerNitroClient).client
	filteractionName := d.Get("name").(string)
	filteraction := filter.Filteraction{
		Name:        d.Get("name").(string),
		Page:        d.Get("page").(string),
		Qual:        d.Get("qual").(string),
		Servicename: d.Get("servicename").(string),
		Value:       d.Get("value").(string),
	}

	if raw := d.GetRawConfig().GetAttr("respcode"); !raw.IsNull() {
		filteraction.Respcode = intPtr(d.Get("respcode").(int))
	}

	_, err := client.AddResource(service.Filteraction.Type(), filteractionName, &filteraction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(filteractionName)

	return readFilteractionFunc(ctx, d, meta)
}

func readFilteractionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readFilteractionFunc")
	client := meta.(*NetScalerNitroClient).client
	filteractionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading filteraction state %s", filteractionName)
	data, err := client.FindResource(service.Filteraction.Type(), filteractionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing filteraction state %s", filteractionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("page", data["page"])
	d.Set("qual", data["qual"])
	setToInt("respcode", d, data["respcode"])
	d.Set("servicename", data["servicename"])
	d.Set("value", data["value"])

	return nil

}

func updateFilteractionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateFilteractionFunc")
	client := meta.(*NetScalerNitroClient).client
	filteractionName := d.Get("name").(string)

	filteraction := filter.Filteraction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("page") {
		log.Printf("[DEBUG]  citrixadc-provider: Page has changed for filteraction %s, starting update", filteractionName)
		filteraction.Page = d.Get("page").(string)
		hasChange = true
	}
	if d.HasChange("qual") {
		log.Printf("[DEBUG]  citrixadc-provider: Qual has changed for filteraction %s, starting update", filteractionName)
		filteraction.Qual = d.Get("qual").(string)
		hasChange = true
	}
	if d.HasChange("respcode") {
		log.Printf("[DEBUG]  citrixadc-provider: Respcode has changed for filteraction %s, starting update", filteractionName)
		filteraction.Respcode = intPtr(d.Get("respcode").(int))
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicename has changed for filteraction %s, starting update", filteractionName)
		filteraction.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("value") {
		log.Printf("[DEBUG]  citrixadc-provider: Value has changed for filteraction %s, starting update", filteractionName)
		filteraction.Value = d.Get("value").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Filteraction.Type(), filteractionName, &filteraction)
		if err != nil {
			return diag.Errorf("Error updating filteraction %s", filteractionName)
		}
	}
	return readFilteractionFunc(ctx, d, meta)
}

func deleteFilteractionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFilteractionFunc")
	client := meta.(*NetScalerNitroClient).client
	filteractionName := d.Id()
	err := client.DeleteResource(service.Filteraction.Type(), filteractionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
