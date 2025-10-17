package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAutoscaleaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAutoscaleactionFunc,
		ReadContext:   readAutoscaleactionFunc,
		UpdateContext: updateAutoscaleactionFunc,
		DeleteContext: deleteAutoscaleactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quiettime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vmdestroygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAutoscaleactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAutoscaleactionFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleactionName := d.Get("name").(string)
	autoscaleaction := autoscale.Autoscaleaction{
		Name:        d.Get("name").(string),
		Parameters:  d.Get("parameters").(string),
		Profilename: d.Get("profilename").(string),
		Type:        d.Get("type").(string),
		Vserver:     d.Get("vserver").(string),
	}

	if raw := d.GetRawConfig().GetAttr("quiettime"); !raw.IsNull() {
		autoscaleaction.Quiettime = intPtr(d.Get("quiettime").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vmdestroygraceperiod"); !raw.IsNull() {
		autoscaleaction.Vmdestroygraceperiod = intPtr(d.Get("vmdestroygraceperiod").(int))
	}

	_, err := client.AddResource(service.Autoscaleaction.Type(), autoscaleactionName, &autoscaleaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(autoscaleactionName)

	return readAutoscaleactionFunc(ctx, d, meta)
}

func readAutoscaleactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAutoscaleactionFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading autoscaleaction state %s", autoscaleactionName)
	data, err := client.FindResource(service.Autoscaleaction.Type(), autoscaleactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing autoscaleaction state %s", autoscaleactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("parameters", data["parameters"])
	d.Set("profilename", data["profilename"])
	setToInt("quiettime", d, data["quiettime"])
	d.Set("type", data["type"])
	setToInt("vmdestroygraceperiod", d, data["vmdestroygraceperiod"])
	d.Set("vserver", data["vserver"])

	return nil

}

func updateAutoscaleactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAutoscaleactionFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleactionName := d.Get("name").(string)

	autoscaleaction := autoscale.Autoscaleaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("parameters") {
		log.Printf("[DEBUG]  citrixadc-provider: Parameters has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Parameters = d.Get("parameters").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("quiettime") {
		log.Printf("[DEBUG]  citrixadc-provider: Quiettime has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Quiettime = intPtr(d.Get("quiettime").(int))
		hasChange = true
	}
	if d.HasChange("vmdestroygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Vmdestroygraceperiod has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Vmdestroygraceperiod = intPtr(d.Get("vmdestroygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("vserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserver has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Vserver = d.Get("vserver").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Autoscaleaction.Type(), &autoscaleaction)
		if err != nil {
			return diag.Errorf("Error updating autoscaleaction %s", autoscaleactionName)
		}
	}
	return readAutoscaleactionFunc(ctx, d, meta)
}

func deleteAutoscaleactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAutoscaleactionFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleactionName := d.Id()
	err := client.DeleteResource(service.Autoscaleaction.Type(), autoscaleactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
