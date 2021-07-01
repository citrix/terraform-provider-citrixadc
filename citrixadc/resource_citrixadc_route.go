package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcRoute() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRouteFunc,
		Read:          readRouteFunc,
		Update:        updateRouteFunc,
		Delete:        deleteRouteFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"advertise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cost1": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"detail": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"msr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRouteFunc")
	client := meta.(*NetScalerNitroClient).client
	// setting ID as `network__netmask__gateway` for better logging
	routeName := d.Get("network").(string) + "__" + d.Get("netmask").(string) + "__" + d.Get("gateway").(string)

	route := network.Route{
		Advertise:  d.Get("advertise").(string),
		Cost:       uint32(d.Get("cost").(int)),
		Cost1:      uint32(d.Get("cost1").(int)),
		Detail:     d.Get("detail").(bool),
		Distance:   uint32(d.Get("distance").(int)),
		Gateway:    d.Get("gateway").(string),
		Monitor:    d.Get("monitor").(string),
		Msr:        d.Get("msr").(string),
		Netmask:    d.Get("netmask").(string),
		Network:    d.Get("network").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Routetype:  d.Get("routetype").(string),
		Td:         uint32(d.Get("td").(int)),
		Vlan:       uint32(d.Get("vlan").(int)),
		Weight:     uint32(d.Get("weight").(int)),
	}

	_, err := client.AddResource(service.Route.Type(), route.Network, &route)
	if err != nil {
		return err
	}

	d.SetId(routeName)

	err = readRouteFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this route but we can't read it ?? %s", routeName)
		return nil
	}
	return nil
}

func readRouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRouteFunc")
	client := meta.(*NetScalerNitroClient).client
	routeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading route state %s", routeName)
	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(d.Get("network").(string))
	argsMap["netmask"] = url.QueryEscape(d.Get("netmask").(string))
	argsMap["gateway"] = url.QueryEscape(d.Get("gateway").(string))
	findParams := service.FindParams{
		ResourceType: service.Route.Type(),
		ArgsMap:      argsMap,
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

	if len(dataArray) > 1 {
		return fmt.Errorf("multiple entries found for route")
	}

	data := dataArray[0]

	d.Set("advertise", data["advertise"])
	d.Set("cost", data["cost"])
	d.Set("cost1", data["cost1"])
	d.Set("detail", data["detail"])
	d.Set("distance", data["distance"])
	d.Set("gateway", data["gateway"])
	d.Set("monitor", data["monitor"])
	d.Set("msr", data["msr"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("routetype", data["routetype"])
	d.Set("td", data["td"])
	d.Set("vlan", data["vlan"])
	d.Set("weight", data["weight"])

	return nil

}

func updateRouteFunc(d *schema.ResourceData, meta interface{}) error {
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
		route.Cost = uint32(d.Get("cost").(int))
		hasChange = true
	}
	if d.HasChange("cost1") {
		log.Printf("[DEBUG]  citrixadc-provider: Cost1 has changed for route %s, starting update", routeName)
		route.Cost1 = uint32(d.Get("cost1").(int))
		hasChange = true
	}
	if d.HasChange("detail") {
		log.Printf("[DEBUG]  citrixadc-provider: Detail has changed for route %s, starting update", routeName)
		route.Detail = d.Get("detail").(bool)
		hasChange = true
	}
	if d.HasChange("distance") {
		log.Printf("[DEBUG]  citrixadc-provider: Distance has changed for route %s, starting update", routeName)
		route.Distance = uint32(d.Get("distance").(int))
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
		route.Td = uint32(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for route %s, starting update", routeName)
		route.Vlan = uint32(d.Get("vlan").(int))
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for route %s, starting update", routeName)
		route.Weight = uint32(d.Get("weight").(int))
		hasChange = true
	}

	if hasChange {
		// network, netmask, gateway attributes are mandatory in UPDATE data
		route.Network = d.Get("network").(string)
		route.Netmask = d.Get("netmask").(string)
		route.Gateway = d.Get("gateway").(string)
		err := client.UpdateUnnamedResource(service.Route.Type(), &route)
		if err != nil {
			return fmt.Errorf("Error updating route %s", routeName)
		}
	}
	return readRouteFunc(d, meta)
}

func deleteRouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRouteFunc")
	client := meta.(*NetScalerNitroClient).client

	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(d.Get("network").(string))
	argsMap["netmask"] = url.QueryEscape(d.Get("netmask").(string))
	argsMap["gateway"] = url.QueryEscape(d.Get("gateway").(string))

	err := client.DeleteResourceWithArgsMap(service.Route.Type(), d.Get("network").(string), argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
