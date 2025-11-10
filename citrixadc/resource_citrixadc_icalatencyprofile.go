package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIcalatencyprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIcalatencyprofileFunc,
		ReadContext:   readIcalatencyprofileFunc,
		UpdateContext: updateIcalatencyprofileFunc,
		DeleteContext: deleteIcalatencyprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"l7latencymaxnotifycount": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l7latencymonitoring": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l7latencynotifyinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l7latencythresholdfactor": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l7latencywaittime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIcalatencyprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcalatencyprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icalatencyprofileName := d.Get("name").(string)
	icalatencyprofile := ica.Icalatencyprofile{
		L7latencymonitoring: d.Get("l7latencymonitoring").(string),
		Name:                d.Get("name").(string),
	}

	if raw := d.GetRawConfig().GetAttr("l7latencymaxnotifycount"); !raw.IsNull() {
		icalatencyprofile.L7latencymaxnotifycount = intPtr(d.Get("l7latencymaxnotifycount").(int))
	}
	if raw := d.GetRawConfig().GetAttr("l7latencynotifyinterval"); !raw.IsNull() {
		icalatencyprofile.L7latencynotifyinterval = intPtr(d.Get("l7latencynotifyinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("l7latencythresholdfactor"); !raw.IsNull() {
		icalatencyprofile.L7latencythresholdfactor = intPtr(d.Get("l7latencythresholdfactor").(int))
	}
	if raw := d.GetRawConfig().GetAttr("l7latencywaittime"); !raw.IsNull() {
		icalatencyprofile.L7latencywaittime = intPtr(d.Get("l7latencywaittime").(int))
	}

	_, err := client.AddResource("icalatencyprofile", icalatencyprofileName, &icalatencyprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(icalatencyprofileName)

	return readIcalatencyprofileFunc(ctx, d, meta)
}

func readIcalatencyprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcalatencyprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icalatencyprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading icalatencyprofile state %s", icalatencyprofileName)
	data, err := client.FindResource("icalatencyprofile", icalatencyprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icalatencyprofile state %s", icalatencyprofileName)
		d.SetId("")
		return nil
	}
	setToInt("l7latencymaxnotifycount", d, data["l7latencymaxnotifycount"])
	d.Set("l7latencymonitoring", data["l7latencymonitoring"])
	setToInt("l7latencynotifyinterval", d, data["l7latencynotifyinterval"])
	setToInt("l7latencythresholdfactor", d, data["l7latencythresholdfactor"])
	setToInt("l7latencywaittime", d, data["l7latencywaittime"])
	d.Set("name", data["name"])

	return nil

}

func updateIcalatencyprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcalatencyprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icalatencyprofileName := d.Get("name").(string)

	icalatencyprofile := ica.Icalatencyprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("l7latencymaxnotifycount") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencymaxnotifycount has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencymaxnotifycount = intPtr(d.Get("l7latencymaxnotifycount").(int))
		hasChange = true
	}
	if d.HasChange("l7latencymonitoring") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencymonitoring has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencymonitoring = d.Get("l7latencymonitoring").(string)
		hasChange = true
	}
	if d.HasChange("l7latencynotifyinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencynotifyinterval has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencynotifyinterval = intPtr(d.Get("l7latencynotifyinterval").(int))
		hasChange = true
	}
	if d.HasChange("l7latencythresholdfactor") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencythresholdfactor has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencythresholdfactor = intPtr(d.Get("l7latencythresholdfactor").(int))
		hasChange = true
	}
	if d.HasChange("l7latencywaittime") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencywaittime has changed for icalatencyprofile %s, starting update", icalatencyprofileName)
		icalatencyprofile.L7latencywaittime = intPtr(d.Get("l7latencywaittime").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icalatencyprofile", &icalatencyprofile)
		if err != nil {
			return diag.Errorf("Error updating icalatencyprofile %s", icalatencyprofileName)
		}
	}
	return readIcalatencyprofileFunc(ctx, d, meta)
}

func deleteIcalatencyprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcalatencyprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	icalatencyprofileName := d.Id()
	err := client.DeleteResource("icalatencyprofile", icalatencyprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
