package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVxlan() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVxlanFunc,
		ReadContext:   readVxlanFunc,
		UpdateContext: updateVxlanFunc,
		DeleteContext: deleteVxlanFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vxlanid": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"innervlantagging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVxlanFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVxlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanId := d.Get("vxlanid").(int)
	vxlan := network.Vxlan{
		Dynamicrouting:     d.Get("dynamicrouting").(string),
		Innervlantagging:   d.Get("innervlantagging").(string),
		Ipv6dynamicrouting: d.Get("ipv6dynamicrouting").(string),
		Protocol:           d.Get("protocol").(string),
		Type:               d.Get("type").(string),
	}

	if raw := d.GetRawConfig().GetAttr("vxlanid"); !raw.IsNull() {
		vxlan.Id = intPtr(d.Get("vxlanid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		vxlan.Port = intPtr(d.Get("port").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		vxlan.Vlan = intPtr(d.Get("vlan").(int))
	}
	vxlanIdStr := strconv.Itoa(vxlanId)
	_, err := client.AddResource(service.Vxlan.Type(), vxlanIdStr, &vxlan)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vxlanIdStr)

	return readVxlanFunc(ctx, d, meta)
}

func readVxlanFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVxlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanIdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vxlan state %s", vxlanIdStr)
	data, err := client.FindResource(service.Vxlan.Type(), vxlanIdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vxlan state %s", vxlanIdStr)
		d.SetId("")
		return nil
	}
	d.Set("dynamicrouting", data["dynamicrouting"])
	setToInt("vxlanid", d, data["id"])
	d.Set("innervlantagging", data["innervlantagging"])
	d.Set("ipv6dynamicrouting", data["ipv6dynamicrouting"])
	setToInt("port", d, data["port"])
	d.Set("protocol", data["protocol"])
	d.Set("type", data["type"])
	setToInt("vlan", d, data["vlan"])

	return nil

}

func updateVxlanFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVxlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanId := d.Get("vxlanid").(int)

	vxlan := network.Vxlan{}

	if raw := d.GetRawConfig().GetAttr("vxlanid"); !raw.IsNull() {
		vxlan.Id = intPtr(d.Get("vxlanid").(int))
	}
	hasChange := false
	if d.HasChange("dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicrouting has changed for vxlan %d, starting update", vxlanId)
		vxlan.Dynamicrouting = d.Get("dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("innervlantagging") {
		log.Printf("[DEBUG]  citrixadc-provider: Innervlantagging has changed for vxlan %d, starting update", vxlanId)
		vxlan.Innervlantagging = d.Get("innervlantagging").(string)
		hasChange = true
	}
	if d.HasChange("ipv6dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv6dynamicrouting has changed for vxlan %d, starting update", vxlanId)
		vxlan.Ipv6dynamicrouting = d.Get("ipv6dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for vxlan %d, starting update", vxlanId)
		vxlan.Port = intPtr(d.Get("port").(int))
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocol has changed for vxlan %d, starting update", vxlanId)
		vxlan.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for vxlan %d, starting update", vxlanId)
		vxlan.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for vxlan %d, starting update", vxlanId)
		vxlan.Vlan = intPtr(d.Get("vlan").(int))
		hasChange = true
	}
	vxlanIdStr := strconv.Itoa(vxlanId)
	if hasChange {
		_, err := client.UpdateResource(service.Vxlan.Type(), vxlanIdStr, &vxlan)
		if err != nil {
			return diag.Errorf("Error updating vxlan %d", vxlanId)
		}
	}
	return readVxlanFunc(ctx, d, meta)
}

func deleteVxlanFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVxlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanName := d.Id()
	err := client.DeleteResource(service.Vxlan.Type(), vxlanName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
