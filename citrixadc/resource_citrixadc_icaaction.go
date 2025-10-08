package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIcaaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIcaactionFunc,
		ReadContext:   readIcaactionFunc,
		UpdateContext: updateIcaactionFunc,
		DeleteContext: deleteIcaactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accessprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"latencyprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIcaactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	icaactionName := d.Get("name").(string)
	icaaction := ica.Icaaction{
		Accessprofilename:  d.Get("accessprofilename").(string),
		Latencyprofilename: d.Get("latencyprofilename").(string),
		Name:               d.Get("name").(string),
	}

	_, err := client.AddResource("icaaction", icaactionName, &icaaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(icaactionName)

	return readIcaactionFunc(ctx, d, meta)
}

func readIcaactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	icaactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading icaaction state %s", icaactionName)
	data, err := client.FindResource("icaaction", icaactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icaaction state %s", icaactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("accessprofilename", data["accessprofilename"])
	d.Set("latencyprofilename", data["latencyprofilename"])

	return nil

}

func updateIcaactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	icaactionName := d.Get("name").(string)

	icaaction := ica.Icaaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("accessprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Accessprofilename has changed for icaaction %s, starting update", icaactionName)
		icaaction.Accessprofilename = d.Get("accessprofilename").(string)
		hasChange = true
	}
	if d.HasChange("latencyprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Latencyprofilename has changed for icaaction %s, starting update", icaactionName)
		icaaction.Latencyprofilename = d.Get("latencyprofilename").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icaaction", &icaaction)
		if err != nil {
			return diag.Errorf("Error updating icaaction %s", icaactionName)
		}
	}
	return readIcaactionFunc(ctx, d, meta)
}

func deleteIcaactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	icaactionName := d.Id()
	err := client.DeleteResource("icaaction", icaactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
