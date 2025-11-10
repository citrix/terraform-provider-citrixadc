package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRnat6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRnat6Func,
		ReadContext:   readRnat6Func,
		UpdateContext: updateRnat6Func,
		DeleteContext: deleteRnat6Func,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"acl6name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcippersistency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createRnat6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRnat6Func")
	client := meta.(*NetScalerNitroClient).client
	rnat6Name := d.Get("name").(string)
	rnat6 := network.Rnat6{
		Acl6name:         d.Get("acl6name").(string),
		Name:             d.Get("name").(string),
		Network:          d.Get("network").(string),
		Ownergroup:       d.Get("ownergroup").(string),
		Srcippersistency: d.Get("srcippersistency").(string),
	}

	if raw := d.GetRawConfig().GetAttr("redirectport"); !raw.IsNull() {
		rnat6.Redirectport = intPtr(d.Get("redirectport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		rnat6.Td = intPtr(d.Get("td").(int))
	}

	_, err := client.AddResource(service.Rnat6.Type(), rnat6Name, &rnat6)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rnat6Name)

	return readRnat6Func(ctx, d, meta)
}

func readRnat6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRnat6Func")
	client := meta.(*NetScalerNitroClient).client
	rnat6Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rnat6 state %s", rnat6Name)
	data, err := client.FindResource(service.Rnat6.Type(), rnat6Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rnat6 state %s", rnat6Name)
		d.SetId("")
		return nil
	}
	d.Set("acl6name", data["acl6name"])
	d.Set("name", data["name"])
	d.Set("network", data["network"])
	d.Set("ownergroup", data["ownergroup"])
	setToInt("redirectport", d, data["redirectport"])
	d.Set("srcippersistency", data["srcippersistency"])
	setToInt("td", d, data["td"])

	return nil

}

func updateRnat6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRnat6Func")
	client := meta.(*NetScalerNitroClient).client
	rnat6Name := d.Get("name").(string)

	rnat6 := make(map[string]interface{})
	rnat6["name"] = d.Get("name").(string)
	hasChange := false
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for rnat6 %s, starting update", rnat6Name)
		rnat6["ownergroup"] = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("redirectport") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectport has changed for rnat6 %s, starting update", rnat6Name)
		rnat6["redirectport"] = intPtr(d.Get("redirectport").(int))
		hasChange = true
	}
	if d.HasChange("srcippersistency") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcippersistency has changed for rnat6 %s, starting update", rnat6Name)
		rnat6["srcippersistency"] = d.Get("srcippersistency").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Rnat6.Type(), &rnat6)
		if err != nil {
			return diag.Errorf("Error updating rnat6 %s", rnat6Name)
		}
	}
	return readRnat6Func(ctx, d, meta)
}

func deleteRnat6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRnat6Func")
	// rnat6 does not support DELETE operation
	d.SetId("")

	return nil
}
