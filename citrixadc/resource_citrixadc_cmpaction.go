package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCmpaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCmpactionFunc,
		ReadContext:   readCmpactionFunc,
		UpdateContext: updateCmpactionFunc,
		DeleteContext: deleteCmpactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cmptype": {
				Type:     schema.TypeString,
				Required: true,
			},
			"addvaryheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deltatype": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"varyheadervalue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCmpactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCmpactionFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpactionName := d.Get("name").(string)
	cmpaction := cmp.Cmpaction{
		Addvaryheader:   d.Get("addvaryheader").(string),
		Cmptype:         d.Get("cmptype").(string),
		Deltatype:       d.Get("deltatype").(string),
		Name:            d.Get("name").(string),
		Varyheadervalue: d.Get("varyheadervalue").(string),
	}

	_, err := client.AddResource(service.Cmpaction.Type(), cmpactionName, &cmpaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cmpactionName)

	return readCmpactionFunc(ctx, d, meta)
}

func readCmpactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCmpactionFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cmpaction state %s", cmpactionName)
	data, err := client.FindResource(service.Cmpaction.Type(), cmpactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cmpaction state %s", cmpactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	//d.Set("addvaryheader", data["addvaryheader"])
	d.Set("cmptype", data["cmptype"])
	//d.Set("deltatype", data["deltatype"])
	d.Set("varyheadervalue", data["varyheadervalue"])

	return nil

}

func updateCmpactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCmpactionFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpactionName := d.Get("name").(string)

	cmpaction := cmp.Cmpaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("addvaryheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Addvaryheader has changed for cmpaction %s, starting update", cmpactionName)
		cmpaction.Addvaryheader = d.Get("addvaryheader").(string)
		hasChange = true
	}
	if d.HasChange("cmptype") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmptype has changed for cmpaction %s, starting update", cmpactionName)
		cmpaction.Cmptype = d.Get("cmptype").(string)
		hasChange = true
	}
	if d.HasChange("varyheadervalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Varyheadervalue has changed for cmpaction %s, starting update", cmpactionName)
		cmpaction.Varyheadervalue = d.Get("varyheadervalue").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Cmpaction.Type(), &cmpaction)
		if err != nil {
			return diag.Errorf("Error updating cmpaction %s", cmpactionName)
		}
	}
	return readCmpactionFunc(ctx, d, meta)
}

func deleteCmpactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCmpactionFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpactionName := d.Id()
	err := client.DeleteResource(service.Cmpaction.Type(), cmpactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
