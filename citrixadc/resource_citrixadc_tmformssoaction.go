package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcTmformssoaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTmformssoactionFunc,
		ReadContext:   readTmformssoactionFunc,
		UpdateContext: updateTmformssoactionFunc,
		DeleteContext: deleteTmformssoactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"actionurl": {
				Type:     schema.TypeString,
				Required: true,
			},
			"userfield": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"passwdfield": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ssosuccessrule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namevaluepair": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nvtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"responsesize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"submitmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createTmformssoactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmformssoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmformssoactionName := d.Get("name").(string)

	tmformssoaction := tm.Tmformssoaction{
		Actionurl:      d.Get("actionurl").(string),
		Name:           d.Get("name").(string),
		Namevaluepair:  d.Get("namevaluepair").(string),
		Nvtype:         d.Get("nvtype").(string),
		Passwdfield:    d.Get("passwdfield").(string),
		Ssosuccessrule: d.Get("ssosuccessrule").(string),
		Submitmethod:   d.Get("submitmethod").(string),
		Userfield:      d.Get("userfield").(string),
	}

	if raw := d.GetRawConfig().GetAttr("responsesize"); !raw.IsNull() {
		tmformssoaction.Responsesize = intPtr(d.Get("responsesize").(int))
	}

	_, err := client.AddResource(service.Tmformssoaction.Type(), tmformssoactionName, &tmformssoaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(tmformssoactionName)

	return readTmformssoactionFunc(ctx, d, meta)
}

func readTmformssoactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmformssoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmformssoactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading tmformssoaction state %s", tmformssoactionName)
	data, err := client.FindResource(service.Tmformssoaction.Type(), tmformssoactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tmformssoaction state %s", tmformssoactionName)
		d.SetId("")
		return nil
	}
	d.Set("actionurl", data["actionurl"])
	d.Set("name", data["name"])
	d.Set("namevaluepair", data["namevaluepair"])
	d.Set("nvtype", data["nvtype"])
	d.Set("passwdfield", data["passwdfield"])
	setToInt("responsesize", d, data["responsesize"])
	d.Set("ssosuccessrule", data["ssosuccessrule"])
	d.Set("submitmethod", data["submitmethod"])
	d.Set("userfield", data["userfield"])

	return nil

}

func updateTmformssoactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTmformssoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmformssoactionName := d.Get("name").(string)

	tmformssoaction := tm.Tmformssoaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("actionurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Actionurl has changed for tmformssoaction %s, starting update", tmformssoactionName)
		tmformssoaction.Actionurl = d.Get("actionurl").(string)
		hasChange = true
	}
	if d.HasChange("namevaluepair") {
		log.Printf("[DEBUG]  citrixadc-provider: Namevaluepair has changed for tmformssoaction %s, starting update", tmformssoactionName)
		tmformssoaction.Namevaluepair = d.Get("namevaluepair").(string)
		hasChange = true
	}
	if d.HasChange("nvtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Nvtype has changed for tmformssoaction %s, starting update", tmformssoactionName)
		tmformssoaction.Nvtype = d.Get("nvtype").(string)
		hasChange = true
	}
	if d.HasChange("passwdfield") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwdfield has changed for tmformssoaction %s, starting update", tmformssoactionName)
		tmformssoaction.Passwdfield = d.Get("passwdfield").(string)
		hasChange = true
	}
	if d.HasChange("responsesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Responsesize has changed for tmformssoaction %s, starting update", tmformssoactionName)
		tmformssoaction.Responsesize = intPtr(d.Get("responsesize").(int))
		hasChange = true
	}
	if d.HasChange("ssosuccessrule") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssosuccessrule has changed for tmformssoaction %s, starting update", tmformssoactionName)
		tmformssoaction.Ssosuccessrule = d.Get("ssosuccessrule").(string)
		hasChange = true
	}
	if d.HasChange("submitmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Submitmethod has changed for tmformssoaction %s, starting update", tmformssoactionName)
		tmformssoaction.Submitmethod = d.Get("submitmethod").(string)
		hasChange = true
	}
	if d.HasChange("userfield") {
		log.Printf("[DEBUG]  citrixadc-provider: Userfield has changed for tmformssoaction %s, starting update", tmformssoactionName)
		tmformssoaction.Userfield = d.Get("userfield").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tmformssoaction.Type(), &tmformssoaction)
		if err != nil {
			return diag.Errorf("Error updating tmformssoaction %s", tmformssoactionName)
		}
	}
	return readTmformssoactionFunc(ctx, d, meta)
}

func deleteTmformssoactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmformssoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmformssoactionName := d.Id()
	err := client.DeleteResource(service.Tmformssoaction.Type(), tmformssoactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
