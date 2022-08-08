package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNspbr() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNspbrFunc,
		Read:          readNspbrFunc,
		Update:        updateNspbrFunc,
		Delete:        deleteNspbrFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"destip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destipop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destport": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destportop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destportval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detail": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iptunnel": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"iptunnelname": &schema.Schema{
				Type:     schema.TypeString,
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
			"nexthop": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"nexthopval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocolnumber": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcipop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcipval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmacmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcport": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcportop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcportval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
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
			"vxlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vxlanvlanmap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNspbrFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNspbrFunc")
	client := meta.(*NetScalerNitroClient).client
	nspbrName:= d.Get("name").(string)
	nspbr := ns.Nspbr{
		Action:         d.Get("action").(string),
		Destip:         d.Get("destip").(bool),
		Destipop:       d.Get("destipop").(string),
		Destipval:      d.Get("destipval").(string),
		Destport:       d.Get("destport").(bool),
		Destportop:     d.Get("destportop").(string),
		Destportval:    d.Get("destportval").(string),
		Detail:         d.Get("detail").(bool),
		Interface:      d.Get("interface").(string),
		Iptunnel:       d.Get("iptunnel").(bool),
		Iptunnelname:   d.Get("iptunnelname").(string),
		Monitor:        d.Get("monitor").(string),
		Msr:            d.Get("msr").(string),
		Name:           d.Get("name").(string),
		Nexthop:        d.Get("nexthop").(bool),
		Nexthopval:     d.Get("nexthopval").(string),
		Ownergroup:     d.Get("ownergroup").(string),
		Priority:       d.Get("priority").(int),
		Protocol:       d.Get("protocol").(string),
		Protocolnumber: d.Get("protocolnumber").(int),
		Srcip:          d.Get("srcip").(bool),
		Srcipop:        d.Get("srcipop").(string),
		Srcipval:       d.Get("srcipval").(string),
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

	_, err := client.AddResource(service.Nspbr.Type(), nspbrName, &nspbr)
	if err != nil {
		return err
	}

	d.SetId(nspbrName)

	err = readNspbrFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nspbr but we can't read it ?? %s", nspbrName)
		return nil
	}
	return nil
}

func readNspbrFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNspbrFunc")
	client := meta.(*NetScalerNitroClient).client
	nspbrName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nspbr state %s", nspbrName)
	data, err := client.FindResource(service.Nspbr.Type(), nspbrName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nspbr state %s", nspbrName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("destip", data["destip"])
	d.Set("destipop", data["destipop"])
	d.Set("destipval", data["destipval"])
	d.Set("destport", data["destport"])
	d.Set("destportop", data["destportop"])
	d.Set("destportval", data["destportval"])
	d.Set("detail", data["detail"])
	d.Set("interface", data["interface"])
	d.Set("iptunnel", data["iptunnel"])
	d.Set("iptunnelname", data["iptunnelname"])
	d.Set("monitor", data["monitor"])
	d.Set("msr", data["msr"])
	//d.Set("nexthop", data["nexthop"])
	d.Set("nexthopval", data["nexthopval"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("priority", data["priority"])
	d.Set("protocol", data["protocol"])
	d.Set("protocolnumber", data["protocolnumber"])
	d.Set("srcip", data["srcip"])
	d.Set("srcipop", data["srcipop"])
	d.Set("srcipval", data["srcipval"])
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

func updateNspbrFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNspbrFunc")
	client := meta.(*NetScalerNitroClient).client
	nspbrName := d.Get("name").(string)

	nspbr := ns.Nspbr{
		Name: d.Get("name").(string),
	}
	hasChange := false
	stateChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for nspbr %s, starting update", nspbrName)
		nspbr.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("destip") {
		log.Printf("[DEBUG]  citrixadc-provider: Destip has changed for nspbr %s, starting update", nspbrName)
		nspbr.Destip = d.Get("destip").(bool)
		hasChange = true
	}
	if d.HasChange("destipop") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipop has changed for nspbr %s, starting update", nspbrName)
		nspbr.Destipop = d.Get("destipop").(string)
		hasChange = true
	}
	if d.HasChange("destipval") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipval has changed for nspbr %s, starting update", nspbrName)
		nspbr.Destipval = d.Get("destipval").(string)
		hasChange = true
	}
	if d.HasChange("destport") {
		log.Printf("[DEBUG]  citrixadc-provider: Destport has changed for nspbr %s, starting update", nspbrName)
		nspbr.Destport = d.Get("destport").(bool)
		hasChange = true
	}
	if d.HasChange("destportop") {
		log.Printf("[DEBUG]  citrixadc-provider: Destportop has changed for nspbr %s, starting update", nspbrName)
		nspbr.Destportop = d.Get("destportop").(string)
		hasChange = true
	}
	if d.HasChange("destportval") {
		log.Printf("[DEBUG]  citrixadc-provider: Destportval has changed for nspbr %s, starting update", nspbrName)
		nspbr.Destportval = d.Get("destportval").(string)
		hasChange = true
	}
	if d.HasChange("detail") {
		log.Printf("[DEBUG]  citrixadc-provider: Detail has changed for nspbr %s, starting update", nspbrName)
		nspbr.Detail = d.Get("detail").(bool)
		hasChange = true
	}
	if d.HasChange("interface") {
		log.Printf("[DEBUG]  citrixadc-provider: Interface has changed for nspbr %s, starting update", nspbrName)
		nspbr.Interface = d.Get("interface").(string)
		hasChange = true
	}
	if d.HasChange("iptunnel") {
		log.Printf("[DEBUG]  citrixadc-provider: Iptunnel has changed for nspbr %s, starting update", nspbrName)
		nspbr.Iptunnel = d.Get("iptunnel").(bool)
		hasChange = true
	}
	if d.HasChange("iptunnelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Iptunnelname has changed for nspbr %s, starting update", nspbrName)
		nspbr.Iptunnelname = d.Get("iptunnelname").(string)
		hasChange = true
	}
	if d.HasChange("monitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitor has changed for nspbr %s, starting update", nspbrName)
		nspbr.Monitor = d.Get("monitor").(string)
		hasChange = true
	}
	if d.HasChange("msr") {
		log.Printf("[DEBUG]  citrixadc-provider: Msr has changed for nspbr %s, starting update", nspbrName)
		nspbr.Msr = d.Get("msr").(string)
		hasChange = true
	}
	if d.HasChange("nexthop") {
		log.Printf("[DEBUG]  citrixadc-provider: Nexthop has changed for nspbr %s, starting update", nspbrName)
		nspbr.Nexthop = d.Get("nexthop").(bool)
		hasChange = true
	}
	if d.HasChange("nexthopval") {
		log.Printf("[DEBUG]  citrixadc-provider: Nexthopval has changed for nspbr %s, starting update", nspbrName)
		nspbr.Nexthopval = d.Get("nexthopval").(string)
		nspbr.Nexthop = d.Get("nexthop").(bool)
		hasChange = true
	}
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for nspbr %s, starting update", nspbrName)
		nspbr.Ownergroup = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for nspbr %s, starting update", nspbrName)
		nspbr.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocol has changed for nspbr %s, starting update", nspbrName)
		nspbr.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("protocolnumber") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocolnumber has changed for nspbr %s, starting update", nspbrName)
		nspbr.Protocolnumber = d.Get("protocolnumber").(int)
		hasChange = true
	}
	if d.HasChange("srcip") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcip has changed for nspbr %s, starting update", nspbrName)
		nspbr.Srcip = d.Get("srcip").(bool)
		hasChange = true
	}
	if d.HasChange("srcipop") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipop has changed for nspbr %s, starting update", nspbrName)
		nspbr.Srcipop = d.Get("srcipop").(string)
		hasChange = true
	}
	if d.HasChange("srcipval") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipval has changed for nspbr %s, starting update", nspbrName)
		nspbr.Srcipval = d.Get("srcipval").(string)
		hasChange = true
	}
	if d.HasChange("srcmac") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcmac has changed for nspbr %s, starting update", nspbrName)
		nspbr.Srcmac = d.Get("srcmac").(string)
		hasChange = true
	}
	if d.HasChange("srcmacmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcmacmask has changed for nspbr %s, starting update", nspbrName)
		nspbr.Srcmacmask = d.Get("srcmacmask").(string)
		hasChange = true
	}
	if d.HasChange("srcport") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcport has changed for nspbr %s, starting update", nspbrName)
		nspbr.Srcport = d.Get("srcport").(bool)
		hasChange = true
	}
	if d.HasChange("srcportop") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcportop has changed for nspbr %s, starting update", nspbrName)
		nspbr.Srcportop = d.Get("srcportop").(string)
		hasChange = true
	}
	if d.HasChange("srcportval") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcportval has changed for nspbr %s, starting update", nspbrName)
		nspbr.Srcportval = d.Get("srcportval").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for nspbr %s, starting update", nspbrName)
		nspbr.State = d.Get("state").(string)
		stateChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nspbr %s, starting update", nspbrName)
		nspbr.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for nspbr %s, starting update", nspbrName)
		nspbr.Vlan = d.Get("vlan").(int)
		hasChange = true
	}
	if d.HasChange("vxlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlan has changed for nspbr %s, starting update", nspbrName)
		nspbr.Vxlan = d.Get("vxlan").(int)
		hasChange = true
	}
	if d.HasChange("vxlanvlanmap") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlanvlanmap has changed for nspbr %s, starting update", nspbrName)
		nspbr.Vxlanvlanmap = d.Get("vxlanvlanmap").(string)
		hasChange = true
	}

	if stateChange {
		err := doNspbrStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling nspbr %s", nspbrName)
		}
	}
	if hasChange {
		err := client.UpdateUnnamedResource(service.Nspbr.Type(), &nspbr)
		if err != nil {
			return fmt.Errorf("Error updating nspbr %s", nspbrName)
		}
	}
	return readNspbrFunc(d, meta)
}

func doNspbrStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doNspbrStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	
	nspbr := ns.Nspbr{
		Name: d.Get("name").(string),
	}
	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Nspbr.Type(), nspbr, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Snmpalarm.Type(), nspbr, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
func deleteNspbrFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNspbrFunc")
	client := meta.(*NetScalerNitroClient).client
	nspbrName := d.Id()
	err := client.DeleteResource(service.Nspbr.Type(), nspbrName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
