package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strconv"
)

func resourceCitrixAdcRoute6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRoute6Func,
		ReadContext:   readRoute6Func,
		UpdateContext: updateRoute6Func,
		DeleteContext: deleteRoute6Func,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"mgmt": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"advertise": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cost": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"detail": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"distance": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"msr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vxlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRoute6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRoute6Func")
	client := meta.(*NetScalerNitroClient).client
	route6Network := d.Get("network").(string)
	route6 := network.Route6{
		Advertise:  d.Get("advertise").(string),
		Detail:     d.Get("detail").(bool),
		Gateway:    d.Get("gateway").(string),
		Monitor:    d.Get("monitor").(string),
		Msr:        d.Get("msr").(string),
		Network:    d.Get("network").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Routetype:  d.Get("routetype").(string),
		Mgmt:       d.Get("mgmt").(bool),
	}

	if raw := d.GetRawConfig().GetAttr("cost"); !raw.IsNull() {
		route6.Cost = intPtr(d.Get("cost").(int))
	}
	if raw := d.GetRawConfig().GetAttr("distance"); !raw.IsNull() {
		route6.Distance = intPtr(d.Get("distance").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		route6.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		route6.Vlan = intPtr(d.Get("vlan").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vxlan"); !raw.IsNull() {
		route6.Vxlan = intPtr(d.Get("vxlan").(int))
	}
	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		route6.Weight = intPtr(d.Get("weight").(int))
	}

	_, err := client.AddResource(service.Route6.Type(), route6Network, &route6)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(route6Network)

	return readRoute6Func(ctx, d, meta)
}

func readRoute6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRoute6Func")
	client := meta.(*NetScalerNitroClient).client
	route6Network := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading route6 state %s", route6Network)
	dataArr, err := client.FindAllResources(service.Route6.Type())
	if err != nil {
		return diag.FromErr(err)
	}
	foundIndex := -1
	for i, v := range dataArr {
		if v["network"] == route6Network &&
			v["vlan"] == strconv.Itoa(d.Get("vlan").(int)) {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindAllresources network and vlan not found in array")
		log.Printf("network:%s value:%v ", route6Network, intPtr(d.Get("vlan").(int)))
		log.Printf("[WARN] citrixadc-provider: Clearing route6 %s", route6Network)
		d.SetId("")
		return nil
	}
	data := dataArr[foundIndex]
	d.Set("advertise", data["advertise"])
	d.Set("mgmt", data["mgmt"])
	setToInt("cost", d, data["cost"])
	d.Set("detail", data["detail"])
	setToInt("distance", d, data["distance"])
	d.Set("gateway", data["gateway"])
	d.Set("monitor", data["monitor"])
	d.Set("msr", data["msr"])
	d.Set("network", data["network"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("routetype", data["routetype"])
	setToInt("td", d, data["td"])
	setToInt("vlan", d, data["vlan"])
	setToInt("vxlan", d, data["vxlan"])
	setToInt("weight", d, data["weight"])

	return nil

}

func updateRoute6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRoute6Func")
	client := meta.(*NetScalerNitroClient).client
	route6Network := d.Get("network").(string)

	route6 := network.Route6{
		Network: route6Network,
	}
	hasChange := false
	if d.HasChange("advertise") {
		log.Printf("[DEBUG]  citrixadc-provider: Advertise has changed for route6 %s, starting update", route6Network)
		route6.Advertise = d.Get("advertise").(string)
		hasChange = true
	}
	if d.HasChange("cost") {
		log.Printf("[DEBUG]  citrixadc-provider: Cost has changed for route6 %s, starting update", route6Network)
		route6.Cost = intPtr(d.Get("cost").(int))
		hasChange = true
	}
	if d.HasChange("detail") {
		log.Printf("[DEBUG]  citrixadc-provider: Detail has changed for route6 %s, starting update", route6Network)
		route6.Detail = d.Get("detail").(bool)
		hasChange = true
	}
	if d.HasChange("distance") {
		log.Printf("[DEBUG]  citrixadc-provider: Distance has changed for route6 %s, starting update", route6Network)
		route6.Distance = intPtr(d.Get("distance").(int))
		hasChange = true
	}
	if d.HasChange("gateway") {
		log.Printf("[DEBUG]  citrixadc-provider: Gateway has changed for route6 %s, starting update", route6Network)
		route6.Gateway = d.Get("gateway").(string)
		hasChange = true
	}
	if d.HasChange("monitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitor has changed for route6 %s, starting update", route6Network)
		route6.Monitor = d.Get("monitor").(string)
		hasChange = true
	}
	if d.HasChange("msr") {
		log.Printf("[DEBUG]  citrixadc-provider: Msr has changed for route6 %s, starting update", route6Network)
		route6.Msr = d.Get("msr").(string)
		hasChange = true
	}
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for route6 %s, starting update", route6Network)
		route6.Ownergroup = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("routetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Routetype has changed for route6 %s, starting update", route6Network)
		route6.Routetype = d.Get("routetype").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for route6 %s, starting update", route6Network)
		route6.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for route6 %s, starting update", route6Network)
		route6.Vlan = intPtr(d.Get("vlan").(int))
		hasChange = true
	}
	if d.HasChange("vxlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlan has changed for route6 %s, starting update", route6Network)
		route6.Vxlan = intPtr(d.Get("vxlan").(int))
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for route6 %s, starting update", route6Network)
		route6.Weight = intPtr(d.Get("weight").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Route6.Type(), route6Network, &route6)
		if err != nil {
			return diag.Errorf("Error updating route6 %s", route6Network)
		}
	}
	return readRoute6Func(ctx, d, meta)
}

func deleteRoute6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRoute6Func")
	client := meta.(*NetScalerNitroClient).client
	//route6Name := d.Id()
	args := make([]string, 0)
	if v, ok := d.GetOk("network"); ok {
		gateway := v.(string)
		args = append(args, fmt.Sprintf("network:%s", url.QueryEscape(gateway)))
	}
	if v, ok := d.GetOk("gateway"); ok {
		gateway := v.(string)
		args = append(args, fmt.Sprintf("gateway:%s", url.QueryEscape(gateway)))
	}
	if v, ok := d.GetOk("vlan"); ok {
		vlan := v.(int)
		args = append(args, fmt.Sprintf("vlan:%v", vlan))

	}
	if v, ok := d.GetOk("vxlan"); ok {
		vxlan := v.(string)
		args = append(args, fmt.Sprintf("vxlan:%v", vxlan))

	}
	if v, ok := d.GetOk("ownergroup"); ok {
		ownergroup := v.(string)
		args = append(args, fmt.Sprintf("ownergroup:%s", url.QueryEscape(ownergroup)))

	}
	err := client.DeleteResourceWithArgs(service.Route6.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
