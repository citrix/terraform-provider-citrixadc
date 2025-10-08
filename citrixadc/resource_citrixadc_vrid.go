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

func resourceCitrixAdcVrid() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVridFunc,
		ReadContext:   readVridFunc,
		UpdateContext: updateVridFunc,
		DeleteContext: deleteVridFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vrid_id": {
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

func createVridFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVridFunc")
	client := meta.(*NetScalerNitroClient).client
	vridId := d.Get("vrid_id").(int)
	vrid := network.Vrid{
		All:                  d.Get("all").(bool),
		Id:                   d.Get("vrid_id").(int),
		Ownernode:            d.Get("ownernode").(int),
		Preemption:           d.Get("preemption").(string),
		Preemptiondelaytimer: d.Get("preemptiondelaytimer").(int),
		Priority:             d.Get("priority").(int),
		Sharing:              d.Get("sharing").(string),
		Trackifnumpriority:   d.Get("trackifnumpriority").(int),
		Tracking:             d.Get("tracking").(string),
	}
	vridIdStr := strconv.Itoa(vridId)
	_, err := client.AddResource(service.Vrid.Type(), vridIdStr, &vrid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vridIdStr)

	return readVridFunc(ctx, d, meta)
}

func readVridFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVridFunc")
	client := meta.(*NetScalerNitroClient).client
	vridIdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vrid state %s", vridIdStr)
	data, err := client.FindResource(service.Vrid.Type(), vridIdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vrid state %s", vridIdStr)
		d.SetId("")
		return nil
	}
	d.Set("all", data["all"])
	setToInt("vrid_id", d, data["id"])
	setToInt("ownernode", d, data["ownernode"])
	d.Set("preemption", data["preemption"])
	setToInt("preemptiondelaytimer", d, data["preemptiondelaytimer"])
	setToInt("priority", d, data["priority"])
	d.Set("sharing", data["sharing"])
	setToInt("trackifnumpriority", d, data["trackifnumpriority"])
	d.Set("tracking", data["tracking"])

	return nil

}

func updateVridFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVridFunc")
	client := meta.(*NetScalerNitroClient).client
	vridId := d.Get("vrid_id").(int)

	vrid := network.Vrid{
		Id: d.Get("vrid_id").(int),
	}
	hasChange := false
	if d.HasChange("all") {
		log.Printf("[DEBUG]  citrixadc-provider: All has changed for vrid %d, starting update", vridId)
		vrid.All = d.Get("all").(bool)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for vrid %d, starting update", vridId)
		vrid.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}
	if d.HasChange("preemption") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemption has changed for vrid %d, starting update", vridId)
		vrid.Preemption = d.Get("preemption").(string)
		hasChange = true
	}
	if d.HasChange("preemptiondelaytimer") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemptiondelaytimer has changed for vrid %d, starting update", vridId)
		vrid.Preemptiondelaytimer = d.Get("preemptiondelaytimer").(int)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for vrid %d, starting update", vridId)
		vrid.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("sharing") {
		log.Printf("[DEBUG]  citrixadc-provider: Sharing has changed for vrid %d, starting update", vridId)
		vrid.Sharing = d.Get("sharing").(string)
		hasChange = true
	}
	if d.HasChange("trackifnumpriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Trackifnumpriority has changed for vrid %d, starting update", vridId)
		vrid.Trackifnumpriority = d.Get("trackifnumpriority").(int)
		hasChange = true
	}
	if d.HasChange("tracking") {
		log.Printf("[DEBUG]  citrixadc-provider: Tracking has changed for vrid %d, starting update", vridId)
		vrid.Tracking = d.Get("tracking").(string)
		hasChange = true
	}
	vridIdStr := strconv.Itoa(vridId)
	if hasChange {
		_, err := client.UpdateResource(service.Vrid.Type(), vridIdStr, &vrid)
		if err != nil {
			return diag.Errorf("Error updating vrid %s", vridIdStr)
		}
	}
	return readVridFunc(ctx, d, meta)
}

func deleteVridFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVridFunc")
	client := meta.(*NetScalerNitroClient).client
	vridName := d.Id()
	err := client.DeleteResource(service.Vrid.Type(), vridName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
