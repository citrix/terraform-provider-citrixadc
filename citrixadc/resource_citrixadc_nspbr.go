package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNspbr() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNspbrFunc,
		ReadContext:   readNspbrFunc,
		UpdateContext: updateNspbrFunc,
		DeleteContext: deleteNspbrFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"targettd": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"destip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destipop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipval": {
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
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"iptunnelname": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			"srcip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcipop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcipval": {
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

func createNspbrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNspbrFunc")
	client := meta.(*NetScalerNitroClient).client
	nspbrName := d.Get("name").(string)
	nspbr := ns.Nspbr{
		Action:       d.Get("action").(string),
		Destip:       d.Get("destip").(bool),
		Destipop:     d.Get("destipop").(string),
		Destipval:    d.Get("destipval").(string),
		Destport:     d.Get("destport").(bool),
		Destportop:   d.Get("destportop").(string),
		Destportval:  d.Get("destportval").(string),
		Detail:       d.Get("detail").(bool),
		Interface:    d.Get("interface").(string),
		Iptunnel:     d.Get("iptunnel").(bool),
		Iptunnelname: d.Get("iptunnelname").(string),
		Monitor:      d.Get("monitor").(string),
		Msr:          d.Get("msr").(string),
		Name:         d.Get("name").(string),
		Nexthop:      d.Get("nexthop").(bool),
		Nexthopval:   d.Get("nexthopval").(string),
		Ownergroup:   d.Get("ownergroup").(string),
		Protocol:     d.Get("protocol").(string),
		Srcip:        d.Get("srcip").(bool),
		Srcipop:      d.Get("srcipop").(string),
		Srcipval:     d.Get("srcipval").(string),
		Srcmac:       d.Get("srcmac").(string),
		Srcmacmask:   d.Get("srcmacmask").(string),
		Srcport:      d.Get("srcport").(bool),
		Srcportop:    d.Get("srcportop").(string),
		Srcportval:   d.Get("srcportval").(string),
		State:        d.Get("state").(string),
		Vxlanvlanmap: d.Get("vxlanvlanmap").(string),
	}
	if raw := d.GetRawConfig().GetAttr("targettd"); !raw.IsNull() {
		nspbr.Targettd = intPtr(d.Get("targettd").(int))
	}
	if raw := d.GetRawConfig().GetAttr("priority"); !raw.IsNull() {
		nspbr.Priority = intPtr(d.Get("priority").(int))
	}
	if raw := d.GetRawConfig().GetAttr("protocolnumber"); !raw.IsNull() {
		nspbr.Protocolnumber = intPtr(d.Get("protocolnumber").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		nspbr.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		nspbr.Vlan = intPtr(d.Get("vlan").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vxlan"); !raw.IsNull() {
		nspbr.Vxlan = intPtr(d.Get("vxlan").(int))
	}

	_, err := client.AddResource(service.Nspbr.Type(), nspbrName, &nspbr)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nspbrName)

	return readNspbrFunc(ctx, d, meta)
}

func readNspbrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	setToInt("targettd", d, data["targettd"])
	d.Set("action", data["action"])
	// d.Set("destip", data["destip"]) // We don't recieve from the NetScaler
	d.Set("destipop", data["destipop"])
	d.Set("destipval", data["destipval"])
	// d.Set("destport", data["destport"]) // We don't recieve from the NetScaler
	d.Set("destportop", data["destportop"])
	d.Set("destportval", data["destportval"])
	d.Set("detail", data["detail"])
	d.Set("interface", data["interface"])
	// d.Set("iptunnel", data["iptunnel"]) // We don't recieve from the NetScaler
	d.Set("iptunnelname", data["iptunnelname"])
	d.Set("monitor", data["monitor"])
	d.Set("msr", data["msr"])
	//d.Set("nexthop", data["nexthop"]) // We don't recieve from the NetScaler
	d.Set("nexthopval", data["nexthopval"])
	d.Set("ownergroup", data["ownergroup"])
	setToInt("priority", d, data["priority"])
	d.Set("protocol", data["protocol"])
	setToInt("protocolnumber", d, data["protocolnumber"])
	// d.Set("srcip", data["srcip"]) // We don't recieve from the NetScaler
	d.Set("srcipop", data["srcipop"])
	d.Set("srcipval", data["srcipval"])
	d.Set("srcmac", data["srcmac"])
	d.Set("srcmacmask", data["srcmacmask"])
	// d.Set("srcport", data["srcport"]) // We don't recieve from the NetScaler
	d.Set("srcportop", data["srcportop"])
	d.Set("srcportval", data["srcportval"])
	d.Set("state", data["state"])
	setToInt("td", d, data["td"])
	setToInt("vlan", d, data["vlan"])
	setToInt("vxlan", d, data["vxlan"])
	d.Set("vxlanvlanmap", data["vxlanvlanmap"])

	return nil

}

func updateNspbrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNspbrFunc")
	client := meta.(*NetScalerNitroClient).client
	nspbrName := d.Get("name").(string)

	nspbr := ns.Nspbr{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("targettd") {
		log.Printf("[DEBUG]  citrixadc-provider: Targettd has changed for nspbr, starting update")
		nspbr.Targettd = intPtr(d.Get("targettd").(int))
		hasChange = true
	}
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
		nspbr.Destip = d.Get("destip").(bool) // whenever the `destipval` is included in the payload then `destip` should also be included
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
		nspbr.Destport = d.Get("destport").(bool) // whenever the `destportval` is included in the payload then `destport` should also be included
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
		nspbr.Nexthop = d.Get("nexthop").(bool) // whenever the `nexthopval` is included in the payload then `nexthop` should also be included
		hasChange = true
	}
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for nspbr %s, starting update", nspbrName)
		nspbr.Ownergroup = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for nspbr %s, starting update", nspbrName)
		nspbr.Priority = intPtr(d.Get("priority").(int))
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocol has changed for nspbr %s, starting update", nspbrName)
		nspbr.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("protocolnumber") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocolnumber has changed for nspbr %s, starting update", nspbrName)
		nspbr.Protocolnumber = intPtr(d.Get("protocolnumber").(int))
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
		nspbr.Srcip = d.Get("srcip").(bool) // whenever the `srcipval` is included in the payload then `srcip` should also be included
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
		nspbr.Srcport = d.Get("srcport").(bool) // whenever the `srcportval` is included in the payload then `srcport` should also be included
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for nspbr %s, starting update", nspbrName)
		nspbr.State = d.Get("state").(string)
		stateChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nspbr %s, starting update", nspbrName)
		nspbr.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for nspbr %s, starting update", nspbrName)
		nspbr.Vlan = intPtr(d.Get("vlan").(int))
		hasChange = true
	}
	if d.HasChange("vxlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlan has changed for nspbr %s, starting update", nspbrName)
		nspbr.Vxlan = intPtr(d.Get("vxlan").(int))
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
			return diag.Errorf("Error enabling/disabling nspbr %s", nspbrName)
		}
	}
	if hasChange {
		err := client.UpdateUnnamedResource(service.Nspbr.Type(), &nspbr)
		if err != nil {
			return diag.Errorf("Error updating nspbr %s", nspbrName)
		}
	}
	return readNspbrFunc(ctx, d, meta)
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
func deleteNspbrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNspbrFunc")
	client := meta.(*NetScalerNitroClient).client
	nspbrName := d.Id()
	err := client.DeleteResource(service.Nspbr.Type(), nspbrName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
