package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func resourceCitrixAdcVrid6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVrid6Func,
		ReadContext:   readVrid6Func,
		UpdateContext: updateVrid6Func,
		DeleteContext: deleteVrid6Func,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vrid6_id": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"all": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"preemption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preemptiondelaytimer": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sharing": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trackifnumpriority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tracking": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVrid6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVrid6Func")
	client := meta.(*NetScalerNitroClient).client
	vrid6Id := d.Get("vrid6_id").(int)
	vrid6 := network.Vrid6{
		All:                  d.Get("all").(bool),
		Id:                   d.Get("vrid6_id").(int),
		Ownernode:            d.Get("ownernode").(int),
		Preemption:           d.Get("preemption").(string),
		Preemptiondelaytimer: d.Get("preemptiondelaytimer").(int),
		Priority:             d.Get("priority").(int),
		Sharing:              d.Get("sharing").(string),
		Trackifnumpriority:   d.Get("trackifnumpriority").(int),
		Tracking:             d.Get("tracking").(string),
	}
	vrid6IdStr := strconv.Itoa(vrid6Id)
	_, err := client.AddResource(service.Vrid6.Type(), vrid6IdStr, &vrid6)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vrid6IdStr)

	return readVrid6Func(ctx, d, meta)
}

func readVrid6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVrid6Func")
	client := meta.(*NetScalerNitroClient).client
	vrid6IdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vrid6 state %s", vrid6IdStr)
	data, err := client.FindResource(service.Vrid6.Type(), vrid6IdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vrid6 state %s", vrid6IdStr)
		d.SetId("")
		return nil
	}
	d.Set("all", data["all"])
	setToInt("vrid6_id", d, data["id"])
	setToInt("ownernode", d, data["ownernode"])
	d.Set("preemption", data["preemption"])
	setToInt("preemptiondelaytimer", d, data["preemptiondelaytimer"])
	setToInt("priority", d, data["priority"])
	d.Set("sharing", data["sharing"])
	setToInt("trackifnumpriority", d, data["trackifnumpriority"])
	d.Set("tracking", data["tracking"])

	return nil

}

func updateVrid6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVrid6Func")
	client := meta.(*NetScalerNitroClient).client
	vrid6Id := d.Get("vrid6_id").(int)

	vrid6 := network.Vrid6{
		Id: d.Get("vrid6_id").(int),
	}
	hasChange := false
	if d.HasChange("all") {
		log.Printf("[DEBUG]  citrixadc-provider: All has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.All = d.Get("all").(bool)
		hasChange = true
	}
	if d.HasChange("id") {
		log.Printf("[DEBUG]  citrixadc-provider: Id has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Id = d.Get("id").(int)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}
	if d.HasChange("preemption") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemption has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Preemption = d.Get("preemption").(string)
		hasChange = true
	}
	if d.HasChange("preemptiondelaytimer") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemptiondelaytimer has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Preemptiondelaytimer = d.Get("preemptiondelaytimer").(int)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("sharing") {
		log.Printf("[DEBUG]  citrixadc-provider: Sharing has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Sharing = d.Get("sharing").(string)
		hasChange = true
	}
	if d.HasChange("trackifnumpriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Trackifnumpriority has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Trackifnumpriority = d.Get("trackifnumpriority").(int)
		hasChange = true
	}
	if d.HasChange("tracking") {
		log.Printf("[DEBUG]  citrixadc-provider: Tracking has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Tracking = d.Get("tracking").(string)
		hasChange = true
	}
	vrid6IdStr := strconv.Itoa(vrid6Id)
	if hasChange {
		_, err := client.UpdateResource(service.Vrid6.Type(), vrid6IdStr, &vrid6)
		if err != nil {
			return diag.Errorf("Error updating vrid6 %s", vrid6IdStr)
		}
	}
	return readVrid6Func(ctx, d, meta)
}

func deleteVrid6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVrid6Func")
	client := meta.(*NetScalerNitroClient).client
	vrid6Id := d.Id()
	err := client.DeleteResource(service.Vrid6.Type(), vrid6Id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
