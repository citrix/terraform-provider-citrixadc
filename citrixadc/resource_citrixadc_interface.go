package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcInterface() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createInterfaceFunc,
		ReadContext:   readInterfaceFunc,
		UpdateContext: updateInterfaceFunc,
		DeleteContext: deleteInterfaceFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"autoneg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidthhigh": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bandwidthnormal": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"duplex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flowctl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"haheartbeat": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hamonitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ifalias": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lacpkey": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lacpmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lacppriority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lacptimeout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lagtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linkredundancy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldpmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lrsetpriority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ringsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ringtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"speed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tagall": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"throughput": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trunk": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trunkmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createInterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createInterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	interfaceId := d.Get("interface_id").(string)

	Interface := network.Interface{
		Autoneg:         d.Get("autoneg").(string),
		Bandwidthhigh:   d.Get("bandwidthhigh").(int),
		Bandwidthnormal: d.Get("bandwidthnormal").(int),
		Duplex:          d.Get("duplex").(string),
		Flowctl:         d.Get("flowctl").(string),
		Haheartbeat:     d.Get("haheartbeat").(string),
		Hamonitor:       d.Get("hamonitor").(string),
		Id:              d.Get("interface_id").(string),
		Ifalias:         d.Get("ifalias").(string),
		Lacpkey:         d.Get("lacpkey").(int),
		Lacpmode:        d.Get("lacpmode").(string),
		Lacppriority:    d.Get("lacppriority").(int),
		Lacptimeout:     d.Get("lacptimeout").(string),
		Lagtype:         d.Get("lagtype").(string),
		Linkredundancy:  d.Get("linkredundancy").(string),
		Lldpmode:        d.Get("lldpmode").(string),
		Lrsetpriority:   d.Get("lrsetpriority").(int),
		Mtu:             d.Get("mtu").(int),
		Ringsize:        d.Get("ringsize").(int),
		Ringtype:        d.Get("ringtype").(string),
		Speed:           d.Get("speed").(string),
		Tagall:          d.Get("tagall").(string),
		Throughput:      d.Get("throughput").(int),
		Trunk:           d.Get("trunk").(string),
		Trunkmode:       d.Get("trunkmode").(string),
	}

	_, err := client.UpdateResource(service.Interface.Type(), "", &Interface)
	if err != nil {
		return diag.Errorf("Error creating Interface %s. %s", interfaceId, err.Error())
	}

	d.SetId(interfaceId)

	return readInterfaceFunc(ctx, d, meta)
}

func readInterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readInterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	interfaceId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading Interface state %s", interfaceId)

	array, _ := client.FindAllResources(service.Interface.Type())

	// Iterate over the retrieved addresses to find the particular interface id
	foundInterface := false
	foundIndex := -1
	for i, item := range array {
		if item["id"] == interfaceId {
			foundInterface = true
			foundIndex = i
			break
		}
	}
	if !foundInterface {
		log.Printf("[WARN] citrixadc-provider: Clearing interface state %s", interfaceId)
		d.SetId("")
		return diag.Errorf("Could not read interface %v", interfaceId)
	}
	// Fallthrough

	data := array[foundIndex]

	d.Set("autoneg", data["autoneg"])
	setToInt("bandwidthhigh", d, data["bandwidthhigh"])
	setToInt("bandwidthnormal", d, data["bandwidthnormal"])
	d.Set("duplex", data["actduplex"])
	d.Set("flowctl", data["actflowctl"])
	d.Set("haheartbeat", data["haheartbeat"])
	d.Set("hamonitor", data["hamonitor"])
	d.Set("interface_id", interfaceId)
	d.Set("ifalias", data["ifalias"])
	setToInt("lacpkey", d, data["lacpkey"])
	d.Set("lacpmode", data["lacpmode"])
	setToInt("lacppriority", d, data["lacppriority"])
	d.Set("lacptimeout", data["lacptimeout"])
	d.Set("lagtype", data["lagtype"])
	d.Set("linkredundancy", data["linkredundancy"])
	d.Set("lldpmode", data["lldpmode"])
	setToInt("lrsetpriority", d, data["lrsetpriority"])
	setToInt("mtu", d, data["mtu"])
	setToInt("ringsize", d, data["ringsize"])
	d.Set("ringtype", data["ringtype"])
	d.Set("speed", data["actspeed"])
	d.Set("tagall", data["tagall"])
	setToInt("throughput", d, data["actthroughput"])
	d.Set("trunk", data["trunk"])
	d.Set("trunkmode", data["trunkmode"])

	return nil

}

func updateInterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateInterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	interfaceId := d.Get("interface_id").(string)

	Interface := network.Interface{
		Id: d.Get("interface_id").(string),
	}
	hasChange := false
	stateChange := false
	if d.HasChange("autoneg") {
		log.Printf("[DEBUG]  citrixadc-provider: Autoneg has changed for Interface %s, starting update", interfaceId)
		Interface.Autoneg = d.Get("autoneg").(string)
		hasChange = true
	}
	if d.HasChange("bandwidthhigh") {
		log.Printf("[DEBUG]  citrixadc-provider: Bandwidthhigh has changed for Interface %s, starting update", interfaceId)
		Interface.Bandwidthhigh = d.Get("bandwidthhigh").(int)
		hasChange = true
	}
	if d.HasChange("bandwidthnormal") {
		log.Printf("[DEBUG]  citrixadc-provider: Bandwidthnormal has changed for Interface %s, starting update", interfaceId)
		Interface.Bandwidthnormal = d.Get("bandwidthnormal").(int)
		hasChange = true
	}
	if d.HasChange("duplex") {
		log.Printf("[DEBUG]  citrixadc-provider: Duplex has changed for Interface %s, starting update", interfaceId)
		Interface.Duplex = d.Get("duplex").(string)
		hasChange = true
	}
	if d.HasChange("flowctl") {
		log.Printf("[DEBUG]  citrixadc-provider: Flowctl has changed for Interface %s, starting update", interfaceId)
		Interface.Flowctl = d.Get("flowctl").(string)
		hasChange = true
	}
	if d.HasChange("haheartbeat") {
		log.Printf("[DEBUG]  citrixadc-provider: Haheartbeat has changed for Interface %s, starting update", interfaceId)
		Interface.Haheartbeat = d.Get("haheartbeat").(string)
		hasChange = true
	}
	if d.HasChange("hamonitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Hamonitor has changed for Interface %s, starting update", interfaceId)
		Interface.Hamonitor = d.Get("hamonitor").(string)
		hasChange = true
	}
	if d.HasChange("interface_id") {
		log.Printf("[DEBUG]  citrixadc-provider: Id has changed for Interface %s, starting update", interfaceId)
		Interface.Id = d.Get("interface_id").(string)
		hasChange = true
	}
	if d.HasChange("ifalias") {
		log.Printf("[DEBUG]  citrixadc-provider: Ifalias has changed for Interface %s, starting update", interfaceId)
		Interface.Ifalias = d.Get("ifalias").(string)
		hasChange = true
	}
	if d.HasChange("lacpkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Lacpkey has changed for Interface %s, starting update", interfaceId)
		Interface.Lacpkey = d.Get("lacpkey").(int)
		hasChange = true
	}
	if d.HasChange("lacpmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Lacpmode has changed for Interface %s, starting update", interfaceId)
		Interface.Lacpmode = d.Get("lacpmode").(string)
		hasChange = true
	}
	if d.HasChange("lacppriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Lacppriority has changed for Interface %s, starting update", interfaceId)
		Interface.Lacppriority = d.Get("lacppriority").(int)
		hasChange = true
	}
	if d.HasChange("lacptimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Lacptimeout has changed for Interface %s, starting update", interfaceId)
		Interface.Lacptimeout = d.Get("lacptimeout").(string)
		hasChange = true
	}
	if d.HasChange("lagtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Lagtype has changed for Interface %s, starting update", interfaceId)
		Interface.Lagtype = d.Get("lagtype").(string)
		hasChange = true
	}
	if d.HasChange("linkredundancy") {
		log.Printf("[DEBUG]  citrixadc-provider: Linkredundancy has changed for Interface %s, starting update", interfaceId)
		Interface.Linkredundancy = d.Get("linkredundancy").(string)
		hasChange = true
	}
	if d.HasChange("lldpmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Lldpmode has changed for Interface %s, starting update", interfaceId)
		Interface.Lldpmode = d.Get("lldpmode").(string)
		hasChange = true
	}
	if d.HasChange("lrsetpriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Lrsetpriority has changed for Interface %s, starting update", interfaceId)
		Interface.Lrsetpriority = d.Get("lrsetpriority").(int)
		hasChange = true
	}
	if d.HasChange("mtu") {
		log.Printf("[DEBUG]  citrixadc-provider: Mtu has changed for Interface %s, starting update", interfaceId)
		Interface.Mtu = d.Get("mtu").(int)
		hasChange = true
	}
	if d.HasChange("ringsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Ringsize has changed for Interface %s, starting update", interfaceId)
		Interface.Ringsize = d.Get("ringsize").(int)
		hasChange = true
	}
	if d.HasChange("ringtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Ringtype has changed for Interface %s, starting update", interfaceId)
		Interface.Ringtype = d.Get("ringtype").(string)
		hasChange = true
	}
	if d.HasChange("speed") {
		log.Printf("[DEBUG]  citrixadc-provider: Speed has changed for Interface %s, starting update", interfaceId)
		Interface.Speed = d.Get("speed").(string)
		hasChange = true
	}
	if d.HasChange("tagall") {
		log.Printf("[DEBUG]  citrixadc-provider: Tagall has changed for Interface %s, starting update", interfaceId)
		Interface.Tagall = d.Get("tagall").(string)
		hasChange = true
	}
	if d.HasChange("throughput") {
		log.Printf("[DEBUG]  citrixadc-provider: Throughput has changed for Interface %s, starting update", interfaceId)
		Interface.Throughput = d.Get("throughput").(int)
		hasChange = true
	}
	if d.HasChange("trunk") {
		log.Printf("[DEBUG]  citrixadc-provider: Trunk has changed for Interface %s, starting update", interfaceId)
		Interface.Trunk = d.Get("trunk").(string)
		hasChange = true
	}
	if d.HasChange("trunkmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Trunkmode has changed for Interface %s, starting update", interfaceId)
		Interface.Trunkmode = d.Get("trunkmode").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for Interface %s, starting update", interfaceId)
		stateChange = true
	}
	if stateChange {
		err := doInterfaceStateChange(d, client)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if hasChange {
		_, err := client.UpdateResource(service.Interface.Type(), "", &Interface)
		if err != nil {
			return diag.Errorf("Error updating Interface %s. %s", interfaceId, err.Error())
		}
	}
	return readInterfaceFunc(ctx, d, meta)
}

func doInterfaceStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doLbvserverStateChange")

	Interface := network.Interface{
		Id: d.Get("interface_id").(string),
	}

	newstate := d.Get("state").(string)

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Interface.Type(), Interface, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Interface.Type(), Interface, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
func deleteInterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteInterfaceFunc")
	// We cannot really delete the interface.
	// Hardware changes can only delete interfaces
	// We just delete it from the local state
	d.SetId("")

	return nil
}
