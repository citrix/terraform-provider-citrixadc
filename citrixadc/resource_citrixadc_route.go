package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"log"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRoute() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRouteFunc,
		ReadContext:   readRouteFunc,
		UpdateContext: updateRouteFunc,
		DeleteContext: deleteRouteFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gateway": {
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
			"cost1": {
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
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRouteFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRouteFunc")
	client := meta.(*NetScalerNitroClient).client
	// setting ID as `network__netmask__gateway` for better logging
	routeName := d.Get("network").(string) + "__" + d.Get("netmask").(string) + "__" + d.Get("gateway").(string)

	route := network.Route{
		Advertise:  d.Get("advertise").(string),
		Detail:     d.Get("detail").(bool),
		Gateway:    d.Get("gateway").(string),
		Monitor:    d.Get("monitor").(string),
		Msr:        d.Get("msr").(string),
		Netmask:    d.Get("netmask").(string),
		Network:    d.Get("network").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Routetype:  d.Get("routetype").(string),
	}

	if raw := d.GetRawConfig().GetAttr("cost"); !raw.IsNull() {
		route.Cost = intPtr(d.Get("cost").(int))
	}
	if raw := d.GetRawConfig().GetAttr("cost1"); !raw.IsNull() {
		route.Cost1 = intPtr(d.Get("cost1").(int))
	}
	if raw := d.GetRawConfig().GetAttr("distance"); !raw.IsNull() {
		route.Distance = intPtr(d.Get("distance").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		route.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		route.Vlan = intPtr(d.Get("vlan").(int))
	}
	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		route.Weight = intPtr(d.Get("weight").(int))
	}

	_, err := client.AddResource(service.Route.Type(), route.Network, &route)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(routeName)

	return readRouteFunc(ctx, d, meta)
}

func readRouteFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRouteFunc")
	client := meta.(*NetScalerNitroClient).client
	routeName := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading route state %s", routeName)
	findParams := service.FindParams{
		ResourceType: service.Route.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing route state %s", routeName)
		d.SetId("")
		return nil
	}
	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: route does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	idSlice := strings.SplitN(routeName, "__", 3)

	network := idSlice[0]
	netmask := idSlice[1]
	gateway := idSlice[2]

	foundIndex := -1
	for i, route := range dataArray {
		match := true
		if route["network"] != network {
			match = false
		}
		if route["netmask"] != netmask {
			match = false
		}
		if route["gateway"] != gateway {
			match = false
		}
		if val, ok := d.GetOk("ownergroup"); ok {
			if route["ownergroup"] != val.(string) {
				match = false
			}
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams route not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing route state %s", routeName)
		d.SetId("")
		return nil
	}

	data := dataArray[foundIndex]

	d.Set("advertise", data["advertise"])
	setToInt("cost", d, data["cost"])
	setToInt("cost1", d, data["cost1"])
	d.Set("detail", data["detail"])
	setToInt("distance", d, data["distance"])
	d.Set("gateway", data["gateway"])
	d.Set("monitor", data["monitor"])
	d.Set("msr", data["msr"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("routetype", data["routetype"])
	setToInt("td", d, data["td"])
	setToInt("vlan", d, data["vlan"])
	setToInt("weight", d, data["weight"])

	return nil

}

func updateRouteFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRouteFunc")
	client := meta.(*NetScalerNitroClient).client
	routeName := d.Id()

	route := network.Route{}
	hasChange := false
	if d.HasChange("advertise") {
		log.Printf("[DEBUG]  citrixadc-provider: Advertise has changed for route %s, starting update", routeName)
		route.Advertise = d.Get("advertise").(string)
		hasChange = true
	}
	if d.HasChange("cost") {
		log.Printf("[DEBUG]  citrixadc-provider: Cost has changed for route %s, starting update", routeName)
		route.Cost = intPtr(d.Get("cost").(int))
		hasChange = true
	}
	if d.HasChange("cost1") {
		log.Printf("[DEBUG]  citrixadc-provider: Cost1 has changed for route %s, starting update", routeName)
		route.Cost1 = intPtr(d.Get("cost1").(int))
		hasChange = true
	}
	if d.HasChange("detail") {
		log.Printf("[DEBUG]  citrixadc-provider: Detail has changed for route %s, starting update", routeName)
		route.Detail = d.Get("detail").(bool)
		hasChange = true
	}
	if d.HasChange("distance") {
		log.Printf("[DEBUG]  citrixadc-provider: Distance has changed for route %s, starting update", routeName)
		route.Distance = intPtr(d.Get("distance").(int))
		hasChange = true
	}
	if d.HasChange("gateway") {
		log.Printf("[DEBUG]  citrixadc-provider: Gateway has changed for route %s, starting update", routeName)
		route.Gateway = d.Get("gateway").(string)
		hasChange = true
	}
	if d.HasChange("monitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitor has changed for route %s, starting update", routeName)
		route.Monitor = d.Get("monitor").(string)
		hasChange = true
	}
	if d.HasChange("msr") {
		log.Printf("[DEBUG]  citrixadc-provider: Msr has changed for route %s, starting update", routeName)
		route.Msr = d.Get("msr").(string)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Netmask has changed for route %s, starting update", routeName)
		route.Netmask = d.Get("netmask").(string)
		hasChange = true
	}
	if d.HasChange("network") {
		log.Printf("[DEBUG]  citrixadc-provider: Network has changed for route %s, starting update", routeName)
		route.Network = d.Get("network").(string)
		hasChange = true
	}
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for route %s, starting update", routeName)
		route.Ownergroup = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("routetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Routetype has changed for route %s, starting update", routeName)
		route.Routetype = d.Get("routetype").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for route %s, starting update", routeName)
		route.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for route %s, starting update", routeName)
		route.Vlan = intPtr(d.Get("vlan").(int))
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for route %s, starting update", routeName)
		route.Weight = intPtr(d.Get("weight").(int))
		hasChange = true
	}

	if hasChange {
		// network, netmask, gateway attributes are mandatory in UPDATE data
		route.Network = d.Get("network").(string)
		route.Netmask = d.Get("netmask").(string)
		route.Gateway = d.Get("gateway").(string)
		err := client.UpdateUnnamedResource(service.Route.Type(), &route)
		if err != nil {
			return diag.Errorf("Error updating route %s", routeName)
		}
	}
	return readRouteFunc(ctx, d, meta)
}

func deleteRouteFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRouteFunc")
	client := meta.(*NetScalerNitroClient).client

	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(d.Get("network").(string))
	argsMap["netmask"] = url.QueryEscape(d.Get("netmask").(string))
	argsMap["gateway"] = url.QueryEscape(d.Get("gateway").(string))
	if val, ok := d.GetOk("ownergroup"); ok {
		argsMap["ownergroup"] = url.QueryEscape(val.(string))
	}

	err := client.DeleteResourceWithArgsMap(service.Route.Type(), "", argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
