package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNspbr6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNspbr6Func,
		Read:          readNspbr6Func,
		Update:        updateNspbr6Func,
		Delete:        deleteNspbr6Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"destipop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destipv6val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destport": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destportop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destportval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detail": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iptunnel": {
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
			"nexthop": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"nexthopval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nexthopvlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocolnumber": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcipop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcipv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcipv6val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmacmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcport": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcportop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcportval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
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
			"vxlanvlanmap": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNspbr6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNspbr6Func")
	client := meta.(*NetScalerNitroClient).client
	nspbr6Name := d.Get("name").(string)
	nspbr6 := ns.Nspbr6{
		Action:         d.Get("action").(string),
		Destipop:       d.Get("destipop").(string),
		Destipv6:       d.Get("destipv6").(bool),
		Destipv6val:    d.Get("destipv6val").(string),
		Destport:       d.Get("destport").(bool),
		Destportop:     d.Get("destportop").(string),
		Destportval:    d.Get("destportval").(string),
		Detail:         d.Get("detail").(bool),
		Interface:      d.Get("interface").(string),
		Iptunnel:       d.Get("iptunnel").(string),
		Monitor:        d.Get("monitor").(string),
		Msr:            d.Get("msr").(string),
		Name:           d.Get("name").(string),
		Nexthop:        d.Get("nexthop").(bool),
		Nexthopval:     d.Get("nexthopval").(string),
		Nexthopvlan:    d.Get("nexthopvlan").(int),
		Ownergroup:     d.Get("ownergroup").(string),
		Priority:       d.Get("priority").(int),
		Protocol:       d.Get("protocol").(string),
		Protocolnumber: d.Get("protocolnumber").(int),
		Srcipop:        d.Get("srcipop").(string),
		Srcipv6:        d.Get("srcipv6").(bool),
		Srcipv6val:     d.Get("srcipv6val").(string),
		Srcmac:         d.Get("srcmac").(string),
		Srcmacmask:     d.Get("srcmacmask").(string),
		Srcport:        d.Get("srcport").(bool),
		Srcportop:      d.Get("srcportop").(string),
		Srcportval:     d.Get("srcportval").(string),
		State:          d.Get("state").(string),
		Td:             d.Get("td").(int),
		Vlan:           d.Get("vlan").(int),
		Vxlan:          d.Get("vxlan").(int),
		Vxlanvlanmap:   d.Get("vxlanvlanmap").(string),
	}

	_, err := client.AddResource(service.Nspbr6.Type(), nspbr6Name, &nspbr6)
	if err != nil {
		return err
	}

	d.SetId(nspbr6Name)

	err = readNspbr6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nspbr6 but we can't read it ?? %s", nspbr6Name)
		return nil
	}
	return nil
}

func readNspbr6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNspbr6Func")
	client := meta.(*NetScalerNitroClient).client
	nspbr6Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nspbr6 state %s", nspbr6Name)
	data, err := client.FindResource(service.Nspbr6.Type(), nspbr6Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nspbr6 state %s", nspbr6Name)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("destipop", data["destipop"])
	d.Set("destipv6", data["destipv6"])
	d.Set("destipv6val", data["destipv6val"])
	d.Set("destport", data["destport"])
	d.Set("destportop", data["destportop"])
	d.Set("destportval", data["destportval"])
	d.Set("detail", data["detail"])
	d.Set("interface", data["interface"])
	d.Set("iptunnel", data["iptunnel"])
	d.Set("monitor", data["monitor"])
	d.Set("msr", data["msr"])
	d.Set("name", data["name"])
	d.Set("nexthop", data["nexthop"])
	d.Set("nexthopval", data["nexthopval"])
	d.Set("nexthopvlan", data["nexthopvlan"])
	d.Set("ownergroup", data["ownergroup"])
	setToInt("priority", d, data["priority"])
	d.Set("protocol", data["protocol"])
	d.Set("protocolnumber", data["protocolnumber"])
	d.Set("srcipop", data["srcipop"])
	d.Set("srcipv6", data["srcipv6"])
	d.Set("srcipv6val", data["srcipv6val"])
	d.Set("srcmac", data["srcmac"])
	d.Set("srcmacmask", data["srcmacmask"])
	d.Set("srcport", data["srcport"])
	d.Set("srcportop", data["srcportop"])
	d.Set("srcportval", data["srcportval"])
	d.Set("state", data["state"])
	d.Set("td", data["td"])
	d.Set("vlan", data["vlan"])
	d.Set("vxlan", data["vxlan"])
	d.Set("vxlanvlanmap", data["vxlanvlanmap"])

	return nil

}

func updateNspbr6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNspbr6Func")
	client := meta.(*NetScalerNitroClient).client
	nspbr6Name := d.Get("name").(string)

	nspbr6 := ns.Nspbr6{
		Name: d.Get("name").(string),
	}
	hasChange := false
	stateChange := false

	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("destipop") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipop has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Destipop = d.Get("destipop").(string)
		hasChange = true
	}
	if d.HasChange("destipv6") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipv6 has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Destipv6 = d.Get("destipv6").(bool)
		hasChange = true
	}
	if d.HasChange("destipv6val") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipv6val has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Destipv6val = d.Get("destipv6val").(string)
		hasChange = true
	}
	if d.HasChange("destport") {
		log.Printf("[DEBUG]  citrixadc-provider: Destport has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Destport = d.Get("destport").(bool)
		hasChange = true
	}
	if d.HasChange("destportop") {
		log.Printf("[DEBUG]  citrixadc-provider: Destportop has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Destportop = d.Get("destportop").(string)
		hasChange = true
	}
	if d.HasChange("destportval") {
		log.Printf("[DEBUG]  citrixadc-provider: Destportval has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Destportval = d.Get("destportval").(string)
		hasChange = true
	}
	if d.HasChange("detail") {
		log.Printf("[DEBUG]  citrixadc-provider: Detail has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Detail = d.Get("detail").(bool)
		hasChange = true
	}
	if d.HasChange("interface") {
		log.Printf("[DEBUG]  citrixadc-provider: Interface has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Interface = d.Get("interface").(string)
		hasChange = true
	}
	if d.HasChange("iptunnel") {
		log.Printf("[DEBUG]  citrixadc-provider: Iptunnel has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Iptunnel = d.Get("iptunnel").(string)
		hasChange = true
	}
	if d.HasChange("monitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitor has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Monitor = d.Get("monitor").(string)
		hasChange = true
	}
	if d.HasChange("msr") {
		log.Printf("[DEBUG]  citrixadc-provider: Msr has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Msr = d.Get("msr").(string)
		hasChange = true
	}
	if d.HasChange("nexthop") {
		log.Printf("[DEBUG]  citrixadc-provider: Nexthop has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Nexthop = d.Get("nexthop").(bool)
		hasChange = true
	}
	if d.HasChange("nexthopval") {
		log.Printf("[DEBUG]  citrixadc-provider: Nexthopval has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Nexthopval = d.Get("nexthopval").(string)
		hasChange = true
	}
	if d.HasChange("nexthopvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Nexthopvlan has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Nexthopvlan = d.Get("nexthopvlan").(int)
		hasChange = true
	}
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Ownergroup = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocol has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("protocolnumber") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocolnumber has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Protocolnumber = d.Get("protocolnumber").(int)
		hasChange = true
	}
	if d.HasChange("srcipop") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipop has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Srcipop = d.Get("srcipop").(string)
		hasChange = true
	}
	if d.HasChange("srcipv6") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipv6 has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Srcipv6 = d.Get("srcipv6").(bool)
		hasChange = true
	}
	if d.HasChange("srcipv6val") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipv6val has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Srcipv6val = d.Get("srcipv6val").(string)
		hasChange = true
	}
	if d.HasChange("srcmac") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcmac has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Srcmac = d.Get("srcmac").(string)
		hasChange = true
	}
	if d.HasChange("srcmacmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcmacmask has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Srcmacmask = d.Get("srcmacmask").(string)
		hasChange = true
	}
	if d.HasChange("srcport") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcport has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Srcport = d.Get("srcport").(bool)
		hasChange = true
	}
	if d.HasChange("srcportop") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcportop has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Srcportop = d.Get("srcportop").(string)
		hasChange = true
	}
	if d.HasChange("srcportval") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcportval has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Srcportval = d.Get("srcportval").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for nspbr6 %s, starting update", nspbr6Name)
		stateChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Vlan = d.Get("vlan").(int)
		hasChange = true
	}
	if d.HasChange("vxlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlan has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Vxlan = d.Get("vxlan").(int)
		hasChange = true
	}
	if d.HasChange("vxlanvlanmap") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlanvlanmap has changed for nspbr6 %s, starting update", nspbr6Name)
		nspbr6.Vxlanvlanmap = d.Get("vxlanvlanmap").(string)
		hasChange = true
	}

	if stateChange {
		err := doNspbr6StateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling nspbr6 %s", nspbr6Name)
		}
	}
	if hasChange {
		_, err := client.UpdateResource(service.Nspbr6.Type(), nspbr6Name, &nspbr6)
		if err != nil {
			return fmt.Errorf("Error updating nspbr6 %s", nspbr6Name)
		}
	}
	return readNspbr6Func(d, meta)
}

func deleteNspbr6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNspbr6Func")
	client := meta.(*NetScalerNitroClient).client
	nspbr6Name := d.Id()
	err := client.DeleteResource(service.Nspbr6.Type(), nspbr6Name)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func doNspbr6StateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doNspbr6StateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes

	nspbr6 := ns.Nspbr6{
		Name: d.Get("name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Nspbr6.Type(), nspbr6, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Nspbr6.Type(), nspbr6, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
